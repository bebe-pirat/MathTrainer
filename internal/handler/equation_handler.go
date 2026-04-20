package handler

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
)

type EquationHandler struct {
	gameService service.GameService
}

func NewEquationHandler(gameService service.GameService) *EquationHandler {
	return &EquationHandler{
		gameService: gameService,
	}
}

func (h *EquationHandler) GetEquationsSet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	var sectionId int

	if err := json.NewDecoder(r.Body).Decode(&sectionId); err != nil {
		http.Error(w, "failed to get session", http.StatusBadRequest)
		slog.Error("failed to get section_id", "error", err)
		return
	}

	equations, err := h.gameService.GenerateAdaptiveEquationSet(ctx, sectionId, sessionData.UserID)
	if err != nil {
		http.Error(w, "failed to generate set of equations", http.StatusInternalServerError)
		slog.Error("failed to create a set of equations", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(equations); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *EquationHandler) CheckEquations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	answers := make([]model.Answer, 0)
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, "failed to get answers data", http.StatusBadRequest)
		slog.Error("bad request", "error", err)
		return
	}

	feedback, err := h.gameService.CheckEquations(ctx, answers)
	if err != nil {
		http.Error(w, "failed to check equations", http.StatusInternalServerError)
		slog.Error("failed to check equations", "error", err)
		return
	}

	err = h.gameService.CreateAttempts(ctx, answers, sessionData.UserID)
	if err != nil {
		http.Error(w, "failed to save student attempts", http.StatusInternalServerError)
		slog.Error("failed to save student attempts", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(feedback); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *EquationHandler) FinishLevel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var data struct {
		Feedback   []model.EquationFeedback `json:"feedback"`
		SectionId  int                      `json:"section_id"`
		LevelOrder int                      `json:"level_order"`
	}

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "failed to get data", http.StatusBadRequest)
		slog.Error("bad request", "error", err)
		return
	}

	starsAndxp, err := h.gameService.FinishLevel(ctx, data.Feedback, data.SectionId, data.LevelOrder, sessionData.UserID)
	if err != nil {
		http.Error(w, "failed to finish level", http.StatusInternalServerError)
		slog.Error("failed to finish level", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(starsAndxp); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}
