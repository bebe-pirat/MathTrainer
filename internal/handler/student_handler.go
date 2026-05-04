package handler

import (
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
)

type StudentHandler struct {
	studentService service.StudentService
	statsService   service.StatsService
}

func NewStudentHandler(studentService service.StudentService, statsService service.StatsService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
		statsService:   statsService,
	}
}

// TODO: добавь в middleware для аутенфикации
// TODO: и еще что-то, что я не могу вспомниьт .................. аааааааааааааааааааааааааа

func (h *StudentHandler) GetLevelsMap(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	levelsMap, err := h.studentService.GetStudentLevelsMap(ctx, sessionData.UserID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("levels read failed", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(levelsMap); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *StudentHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.Info("hola, seniorita")
	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	profile, err := h.studentService.GetProfile(ctx, sessionData.UserID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed get profile", "user_id", sessionData.UserID, "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(profile); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *StudentHandler) GetAchievements(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	achs, err := h.studentService.GetAchievements(ctx, sessionData.UserID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed get achievements", "user_id", sessionData.UserID)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(achs); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *StudentHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	stats, err := h.statsService.GetStudentStats(ctx, sessionData.UserID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed to get students stats", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(stats); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}
