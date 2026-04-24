package handler

import (
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TeacherHandler struct {
	teacherService service.TeacherService
	statsService   service.StatsService
}

func NewTeacherHandler(teacherService service.TeacherService, statsService service.StatsService) *TeacherHandler {
	return &TeacherHandler{
		teacherService: teacherService,
		statsService:   statsService,
	}
}

func (h *TeacherHandler) GetClassStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	classId, err := h.teacherService.GetClassByTeacherId(ctx, sessionData.UserID)
	if err != nil {
		http.Error(w, "failed to get teacher's class", http.StatusInternalServerError)
		slog.Error("failed to get teacher's class", "error", err)
		return
	}

	classStats, err := h.statsService.GetClassStats(ctx, classId)
	if err != nil {
		http.Error(w, "failed to get class's stats", http.StatusInternalServerError)
		slog.Error("failed to get class's stats", "error", err)
		return
	}

	slog.Info("classStats", classStats)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(classStats); err != nil {
		slog.Error("serialization faild", "error", err)
	}
}

func (h *TeacherHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	classId, err := h.teacherService.GetClassByTeacherId(ctx, sessionData.UserID)
	if err != nil {
		http.Error(w, "failed to get teacher's class", http.StatusInternalServerError)
		slog.Error("failed to get teacher's class", "error", err)
		return
	}

	students, err := h.teacherService.GetClassStudents(ctx, classId)
	if err != nil {
		http.Error(w, "failed to get students", http.StatusInternalServerError)
		slog.Error("failed to get students", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(students); err != nil {
		slog.Error("serialization faild", "error", err)
	}
}

func (h *TeacherHandler) GetStudentById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	studentId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("bad request", "error", err)
		return
	}

	studentStats, err := h.statsService.GetStudentStats(ctx, studentId)
	if err != nil {
		http.Error(w, "failed to get students stats", http.StatusBadRequest)
		slog.Error("failed to get students stats", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(studentStats); err != nil {
		slog.Error("serialization faild", "error", err)
	}
}

func (h *TeacherHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	t := struct {
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Login    string `json:"login"`
		ClassId  int    `json:"class_id"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("bad request", "error", err)
		return
	}

	student, err := h.teacherService.CreateStudent(ctx, t.ClassId, t.Fullname, t.Email, t.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("bad request", "error", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(student); err != nil {
		slog.Error("serialization faild", "error", err)
	}
	slog.Info("created")
}

func (h *TeacherHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	studentId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("bad request", "error", err)
		return
	}

	t := struct {
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("bad request", "error", err)
		return
	}

	err = h.teacherService.UpdateStudent(ctx, studentId, t.Fullname, t.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("bad request", "error", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *TeacherHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	studentId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("bad request", "error", err)
		return
	}

	err = h.teacherService.DeleteStudent(ctx, studentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		slog.Error("bad request", "error", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *TeacherHandler) GetStudentsAttempts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	studentIdStr := r.URL.Query().Get("student_id")
	if studentIdStr == "" {
		http.Error(w, "missing student_id", http.StatusBadRequest)
		slog.Error("missing student_id parameter")
		return
	}
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		http.Error(w, "invalid student_id", http.StatusBadRequest)
		slog.Error("invalid student_id", "error", err)
		return
	}

	equationTypeId := 0
	if eqTypeStr := r.URL.Query().Get("equation_type_id"); eqTypeStr != "" {
		equationTypeId, err = strconv.Atoi(eqTypeStr)
		if err != nil {
			http.Error(w, "invalid equation_type_id", http.StatusBadRequest)
			slog.Error("invalid equation_type_id", "error", err)
			return
		}
	}

	attempts, err := h.teacherService.GetStudentAttempts(ctx, studentId, equationTypeId)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("internal server error", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(attempts); err != nil {
		slog.Error("serialization failed", "error", err)
	}
}

func (h *TeacherHandler) GetEquationTypesByStudentId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	studentIdStr := r.URL.Query().Get("student_id")
	if studentIdStr == "" {
		http.Error(w, "missing student_id", http.StatusBadRequest)
		slog.Error("missing student_id parameter")
		return
	}
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil {
		http.Error(w, "invalid student_id", http.StatusBadRequest)
		slog.Error("invalid student_id", "error", err)
		return
	}

	types, err := h.teacherService.GetEquationTypesByStudentId(ctx, studentId)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("internal server error", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(types); err != nil {
		slog.Error("serialization failed", "error", err)
	}
}
