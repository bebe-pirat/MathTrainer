package main

import (
	"MathTrainer/internal/database"
	"MathTrainer/internal/handler"
	"MathTrainer/internal/repository"
	"MathTrainer/internal/service"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		logger.Info("unable to load .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		logger.Error("no connection string")
		return
	}

	db, err := database.OpenDB(connectionString)
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		return
	}
	defer db.Close()

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		logger.Error("SECRET_KEY is required")
		return
	}

	handler.InitCookieStore(secretKey)

	// Repo
	userRepo := repository.NewUserRepositoryStruct(db)
	sessionRepo := repository.NewSessionRepositoryStruct(db)
	schoolRepo := repository.NewSchoolRepositoryStruct(db)
	classRepo := repository.NewClassRepositoryStruct(db)
	attemptRepo := repository.NewEquationAttemptsRepositoryStruct(db)
	progressRepo := repository.NewStudentProgressRepositoryStruct(db)
	achievRepo := repository.NewAchievementOfStudentRepositoryStruct(db)
	sectionRepo := repository.NewSectionRepositoryStruct(db)
	equationRepo := repository.NewEquationTypeRepositoryStruct(db)

	// Service
	authService := service.NewAuthServiceStruct(userRepo, sessionRepo)
	adminService := service.NewAdminServiceStruct(userRepo, schoolRepo, sectionRepo, equationRepo)
	classService := service.NewClassServiceStruct(classRepo)
	statsService := service.NewStatStatsServiceStruct(classRepo, schoolRepo, userRepo, attemptRepo, progressRepo, achievRepo)
	teacherService := service.NewTeacherServiceStruct(userRepo, attemptRepo, equationRepo)
	studentService := service.NewStudentServiceStruct(userRepo, achievRepo, sectionRepo)
	gameService := service.NewGameServiceStruct(equationRepo, attemptRepo, progressRepo, userRepo)
	directorService := service.NewDirectorService(schoolRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	adminHandler := handler.NewAdminHandler(adminService, classService)
	teacherHandler := handler.NewTeacherHandler(teacherService, statsService)
	studentHandler := handler.NewStudentHandler(studentService, statsService)
	equationHandler := handler.NewEquationHandler(gameService)
	directorHandler := handler.NewDirectorHandler(statsService, classService, directorService)

	mainRouter := mux.NewRouter()

	createAuthRouter(mainRouter, authHandler)
	createAdminRouter(mainRouter, adminHandler)
	createTeacherRouter(mainRouter, teacherHandler)
	createStudentRouter(mainRouter, studentHandler)
	createEquationRouter(mainRouter, equationHandler)
	createDirectorRouter(mainRouter, directorHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: c.Handler(mainRouter),
	}

	go func() {
		slog.Info("server started", "port", port)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server didn't started", "error", err)
			return
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	slog.Info("server is shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "error", err)
	}

	slog.Info("server exiting")
}

func createAdminRouter(router *mux.Router, adminHandler *handler.AdminHandler) *mux.Router {
	adminRouter := router.PathPrefix("/admin").Subrouter()

	adminRouter.HandleFunc("/schools", adminHandler.GetSchools).Methods("GET")
	adminRouter.HandleFunc("/users", adminHandler.GetUsers).Methods("GET")
	adminRouter.HandleFunc("/teachers", adminHandler.GetTeachers).Methods("GET")
	adminRouter.HandleFunc("/classes", adminHandler.GetClassesBySchoolId).Methods("GET")

	adminRouter.HandleFunc("/users/block", adminHandler.ChangeUserBlock).Methods("PUT")

	adminRouter.HandleFunc("/classes", adminHandler.CreateClass).Methods("POST")
	adminRouter.HandleFunc("/schools", adminHandler.CreateSchool).Methods("POST")
	adminRouter.HandleFunc("/teachers", adminHandler.CreateTeacher).Methods("POST")

	adminRouter.HandleFunc("/sections", adminHandler.CreateSection).Methods("POST")
	adminRouter.HandleFunc("/sections/{id}", adminHandler.UpdateSection).Methods("PUT")
	adminRouter.HandleFunc("/sections/{id}", adminHandler.DeleteSection).Methods("DELETE")
	adminRouter.HandleFunc("/sections", adminHandler.GetSections).Methods("GET")

	adminRouter.HandleFunc("/users", adminHandler.CreateUser).Methods("POST")
	adminRouter.HandleFunc("/users/{id}", adminHandler.UpdateUser).Methods("PUT")
	adminRouter.HandleFunc("/users/{id}", adminHandler.DeleteUser).Methods("DELETE")
	adminRouter.HandleFunc("/users/{id}", adminHandler.UpdateUserPassword).Methods("PATCH")

	adminRouter.HandleFunc("/equation-types", adminHandler.CreateEquationType).Methods("POST")
	adminRouter.HandleFunc("/equation-types/{id}", adminHandler.DeleteEquationType).Methods("DELETE")
	adminRouter.HandleFunc("/equation-types/{id}", adminHandler.UpdateEquationType).Methods("PUT")
	adminRouter.HandleFunc("/equation-types", adminHandler.GetEquationTypes).Methods("GET")
	adminRouter.HandleFunc("/operands", adminHandler.GetOperandsForEquationType).Methods("GET")

	adminRouter.HandleFunc("/sections-equation-types", adminHandler.GetSectionsAndEquationTypes).Methods("GET")
	adminRouter.HandleFunc("/sections-equation-types", adminHandler.JoinEquationTypeAndSection).Methods("POST")
	adminRouter.HandleFunc("/sections-equation-types", adminHandler.UnJoinEquationTypeAndSection).Methods("DELETE")

	return adminRouter
}

func createTeacherRouter(router *mux.Router, teacherHandler *handler.TeacherHandler) *mux.Router {
	teacherRouter := router.PathPrefix("/teacher").Subrouter()

	teacherRouter.HandleFunc("/class/stats", teacherHandler.GetClassStats).Methods("GET")
	teacherRouter.HandleFunc("/students", teacherHandler.GetStudents).Methods("GET")
	teacherRouter.HandleFunc("/students/stats/{student_id}", teacherHandler.GetStudentById).Methods("GET")
	teacherRouter.HandleFunc("/students/attempts", teacherHandler.GetStudentsAttempts).Methods("GET")
	teacherRouter.HandleFunc("/equation-types", teacherHandler.GetEquationTypesByStudentId).Methods("GET")

	teacherRouter.HandleFunc("/students", teacherHandler.CreateStudent).Methods("POST")

	teacherRouter.HandleFunc("/students", teacherHandler.UpdateStudent).Methods("PUT")

	teacherRouter.HandleFunc("/students", teacherHandler.DeleteStudent).Methods("DELETE")

	return teacherRouter
}

func createStudentRouter(router *mux.Router, studentHandler *handler.StudentHandler) *mux.Router {
	studentRouter := router.PathPrefix("/student").Subrouter()

	studentRouter.HandleFunc("/level-map", studentHandler.GetLevelsMap).Methods("GET")
	studentRouter.HandleFunc("/profile", studentHandler.GetProfile).Methods("GET")
	studentRouter.HandleFunc("/stats", studentHandler.GetStats).Methods("GET")

	return studentRouter
}

func createEquationRouter(router *mux.Router, equationHandler *handler.EquationHandler) *mux.Router {
	equationRouter := router.PathPrefix("/game").Subrouter()

	equationRouter.HandleFunc("/check", equationHandler.CheckEquations).Methods("POST")
	equationRouter.HandleFunc("/equations-set", equationHandler.GetEquationsSet).Methods("POST")
	equationRouter.HandleFunc("/finish-level", equationHandler.FinishLevel).Methods("POST")

	return equationRouter
}

func createAuthRouter(router *mux.Router, authHandler *handler.AuthHandler) *mux.Router {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", authHandler.Login).Methods("POST")
	authRouter.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/session", authHandler.CheckSession).Methods("GET")

	return authRouter
}

func createDirectorRouter(router *mux.Router, directorHandler *handler.DirectorHandler) *mux.Router {
	directorRouter := router.PathPrefix("/director").Subrouter()

	directorRouter.HandleFunc("/school-stats", directorHandler.GetSchoolStats).Methods("GET")
	directorRouter.HandleFunc("/class-stats/{class_id}", directorHandler.GetClassStats).Methods("GET")
	directorRouter.HandleFunc("/student-stats/{student_id}", directorHandler.GetStudentStats).Methods("GET")
	directorRouter.HandleFunc("/classes", directorHandler.GetClassesBySchool).Methods("GET")

	return directorRouter
}
