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

	// Инициализируем cookie store с ключом

	// Repo
	userRepo := repository.NewUserRepositoryStruct(db)
	sessionRepo := repository.NewSessionRepositoryStruct(db)
	schoolRepo := repository.NewSchoolRepositoryStruct(db)
	classRepo := repository.NewClassRepositoryStruct(db)
	attemptRepo := repository.NewEquationAttemptsRepositoryStruct(db)
	progressRepo := repository.NewStudentProgressRepositoryStruct(db)
	achievRepo := repository.NewAchievementOfStudentRepositoryStruct(db)
	levelRepo := repository.NewLevelRepositoryStruct(db)
	sectionRepo := repository.NewSectionRepositoryStruct(db)

	// Service
	authService := service.NewAuthServiceStruct(userRepo, sessionRepo)
	adminService := service.NewAdminServiceStruct(userRepo, schoolRepo)
	classService := service.NewClassServiceStruct(classRepo)
	statsService := service.NewStatStatsServiceStruct(classRepo, schoolRepo, userRepo, attemptRepo, progressRepo, achievRepo)
	teacherService := service.NewTeacherServiceStruct(userRepo, attemptRepo)
	levelService := service.NewLevelServiceStruct(levelRepo, progressRepo)
	studentService := service.NewStudentServiceStruct(userRepo, achievRepo, sectionRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	adminHandler := handler.NewAdminHandler(adminService, classService)
	teacherHandler := handler.NewTeacherHandler(teacherService, statsService)
	studentHandler := handler.NewStudentHandler(studentService, levelService, statsService)

	mainRouter := mux.NewRouter()

	createAuthRouter(mainRouter, authHandler)
	createAdminRouter(mainRouter, adminHandler)
	createTeacherRouter(mainRouter, teacherHandler)
	createStudentRouter(mainRouter, studentHandler)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
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

	return adminRouter
}

func createTeacherRouter(router *mux.Router, teacherHandler *handler.TeacherHandler) *mux.Router {
	teacherRouter := router.PathPrefix("/teacher").Subrouter()

	teacherRouter.HandleFunc("/class/stats", teacherHandler.GetClassStats).Methods("GET")
	teacherRouter.HandleFunc("/students", teacherHandler.GetStudents).Methods("GET")
	teacherRouter.HandleFunc("/students/stats", teacherHandler.GetStudentById).Methods("GET")

	teacherRouter.HandleFunc("/students", teacherHandler.CreateStudent).Methods("POST")

	teacherRouter.HandleFunc("/students", teacherHandler.UpdateStudent).Methods("PUT")

	teacherRouter.HandleFunc("/students", teacherHandler.DeleteStudent).Methods("DELETE")

	return teacherRouter
}

func createStudentRouter(router *mux.Router, studentHandler *handler.StudentHandler) *mux.Router {
	studentRouter := router.PathPrefix("/student").Subrouter()

	studentRouter.HandleFunc("/level-map", studentHandler.GetLevelsMap).Methods("GET")

	return studentRouter
}

// func createEquationRouter() *mux.Router {

// }

func createAuthRouter(router *mux.Router, authHandler *handler.AuthHandler) *mux.Router {
	authRouter := router.PathPrefix("/auth").Subrouter()

	authRouter.HandleFunc("/login", authHandler.Login).Methods("POST")
	authRouter.HandleFunc("/logout", authHandler.Logout).Methods("POST")
	authRouter.HandleFunc("/session", authHandler.CheckSession).Methods("GET")

	return authRouter
}
