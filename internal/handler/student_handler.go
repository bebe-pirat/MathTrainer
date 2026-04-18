package handler

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	studentService service.StudentService
	levelService   service.LevelService
	statsService   service.StatsService
}

func NewStudentHandler(studentService service.StudentService, levelService service.LevelService, statsService service.StatsService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
		levelService:   levelService,
		statsService:   statsService,
	}
}

// TODO: добавь в middleware для аутенфикации
// TODO: и еще что-то, что я не могу вспомниьт .................. аааааааааааааааааааааааааа

func (h *StudentHandler) GetLevelsMap(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.Info("hola")

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

func (h *StudentHandler) GetLevelsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	levels, err := h.levelService.GetLevels(ctx)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("levels read failed", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(levels); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *StudentHandler) GetLevelHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		slog.Error("invalid json", "error", err)
		return
	}

	level, err := h.levelService.GetLevelById(ctx, id)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("levels read failed", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(level); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *StudentHandler) StartLevel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	levelId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("failed to convert id", "error", err)
		return
	}

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	errResp := h.levelService.StartLevel(ctx, sessionData.UserID, levelId).(*model.ErrorResponse)
	if err != nil {
		http.Error(w, errResp.Message, errResp.Code)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *StudentHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	profile, err := h.studentService.GetProfile(ctx, sessionData.UserID)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed get profile", "user_id", sessionData.UserID)
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
