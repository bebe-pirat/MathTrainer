package handler

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
)

type StudentHandler struct {
	studentService service.StudentService
	levelService   service.LevelService
	statsService   service.StatsService
	gameService    service.GameService
}

func NewStudentHandler(studentService service.StudentService, levelService service.LevelService, statsService service.StatsService, gameService service.GameService) *StudentHandler {
	return &StudentHandler{
		studentService: studentService,
		levelService:   levelService,
		statsService:   statsService,
		gameService:    gameService,
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

func (h *StudentHandler) GetEquationsSet(w http.ResponseWriter, r *http.Request) {
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
		slog.Error("bad request", "error", err)
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

func (h *StudentHandler) CheckEquations(w http.ResponseWriter, r *http.Request) {
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

func (h *StudentHandler) FinishLevel(w http.ResponseWriter, r *http.Request) {
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

// func (h *StudentHandler) GetLevelsHandler(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	levels, err := h.levelService.GetLevels(ctx)
// 	if err != nil {
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		slog.Error("levels read failed", "error", err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(w).Encode(levels); err != nil {
// 		slog.Error("serializtion failed", "error", err)
// 	}
// }

// func (h *StudentHandler) GetLevelHandler(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	vars := mux.Vars(r)
// 	id, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "invalid json", http.StatusBadRequest)
// 		slog.Error("invalid json", "error", err)
// 		return
// 	}

// 	level, err := h.levelService.GetLevelById(ctx, id)
// 	if err != nil {
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		slog.Error("levels read failed", "error", err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(w).Encode(level); err != nil {
// 		slog.Error("serializtion failed", "error", err)
// 	}
// }

// func (h *StudentHandler) StartLevel(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	vars := mux.Vars(r)
// 	levelId, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		http.Error(w, "bad request", http.StatusBadRequest)
// 		slog.Error("failed to convert id", "error", err)
// 		return
// 	}

// 	sessionData, err := getSessionFromCookie(r)
// 	if err != nil {
// 		http.Error(w, "invalid session", http.StatusUnauthorized)
// 		slog.Error("failed to get session from cookie", "error", err)
// 		return
// 	}

// 	errResp := h.levelService.StartLevel(ctx, sessionData.UserID, levelId).(*model.ErrorResponse)
// 	if err != nil {
// 		http.Error(w, errResp.Message, errResp.Code)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// }

// func (h *StudentHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	sessionData, err := getSessionFromCookie(r)
// 	if err != nil {
// 		http.Error(w, "invalid session", http.StatusUnauthorized)
// 		slog.Error("failed to get session from cookie", "error", err)
// 		return
// 	}

// 	profile, err := h.studentService.GetProfile(ctx, sessionData.UserID)
// 	if err != nil {
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		slog.Error("failed get profile", "user_id", sessionData.UserID)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(w).Encode(profile); err != nil {
// 		slog.Error("serializtion failed", "error", err)
// 	}
// }

// func (h *StudentHandler) GetAchievements(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	sessionData, err := getSessionFromCookie(r)
// 	if err != nil {
// 		http.Error(w, "invalid session", http.StatusUnauthorized)
// 		slog.Error("failed to get session from cookie", "error", err)
// 		return
// 	}

// 	achs, err := h.studentService.GetAchievements(ctx, sessionData.UserID)
// 	if err != nil {
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		slog.Error("failed get achievements", "user_id", sessionData.UserID)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(w).Encode(achs); err != nil {
// 		slog.Error("serializtion failed", "error", err)
// 	}
// }

// func (h *StudentHandler) GetStats(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	sessionData, err := getSessionFromCookie(r)
// 	if err != nil {
// 		http.Error(w, "invalid session", http.StatusUnauthorized)
// 		slog.Error("failed to get session from cookie", "error", err)
// 		return
// 	}

// 	stats, err := h.statsService.GetStudentStats(ctx, sessionData.UserID)
// 	if err != nil {
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		slog.Error("failed to get students stats", "error", err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(w).Encode(stats); err != nil {
// 		slog.Error("serializtion failed", "error", err)
// 	}
// }
