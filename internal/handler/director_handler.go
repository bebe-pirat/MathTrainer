package handler

import (
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DirectorHandler struct {
	statsService service.StatsService
	classService service.ClassServiceStruct
}

func NewDirectorHandler(statsService service.StatsService, classService service.ClassServiceStruct) *DirectorHandler {
	return &DirectorHandler{
		statsService: statsService,
		classService: classService,
	}
}

func (h *DirectorHandler) GetClasses(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	schoolId, err := strconv.Atoi(vars["school_id"])
	if err != nil {
		http.Error(w, "school_id is required", http.StatusBadRequest)
		slog.Error("failed to convert schoolId into int", "error", err)
		return
	}

	classes, err := h.classService.GetClassesBySchool(ctx, schoolId)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed to get classes of school", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(classes); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *DirectorHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//name string, grade int, schoolId int
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		http.Error(w, "name is required", http.StatusBadRequest)
		slog.Error("no name in request")
		return
	}

	grade, err := strconv.Atoi(vars["grade"])
	if err != nil {
		http.Error(w, "grade is required", http.StatusBadRequest)
		slog.Error("failed to get grade from request", "error", err)
		return
	}

	school_id, err := strconv.Atoi(vars["school_id"])
	if err != nil {
		http.Error(w, "school_id is required", http.StatusBadRequest)
		slog.Error("failed to get school_id from request", "error", err)
		return
	}

	classId, err := h.classService.CreateClass(ctx, name, grade, school_id)
	if err != nil {
		http.Error(w, "failed to create new class", http.StatusInternalServerError)
		slog.Error("failed to create class", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(classId); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *DirectorHandler) DeleteClass(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	classId, err := strconv.Atoi(vars["class_id"])
	if err != nil {
		http.Error(w, "class_id is required", http.StatusBadRequest)
		slog.Error("class_id is required", "error", err)
		return
	}

	err = h.classService.DeleteClass(ctx, classId)
	if err != nil {
		http.Error(w, "failed to delete record from classes", http.StatusBadRequest)
		slog.Error("failed to delete record", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *DirectorHandler) GetSchoolStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	schoolId, err := strconv.Atoi(vars["school_id"])
	if err != nil {
		http.Error(w, "school_id is required", http.StatusBadRequest)
		slog.Error("school_id is required", "error", err)
		return
	}

	stats, err := h.statsService.GetSchoolStats(ctx, schoolId)
	if err != nil {
		http.Error(w, "failed to get school stats", http.StatusInternalServerError)
		slog.Error("failed to get school stats", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(stats); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

// director
// GET director/classes?school_id= +
// POST director/classes +
// DELETE director/classesid +
// GET /director/stats +
