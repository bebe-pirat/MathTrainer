package handler

import (
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) GetSchools(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	schools, err := h.adminService.GetSchools(ctx)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed to get schoools", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(schools); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) GetTeachers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	schoolId, err := strconv.Atoi(r.URL.Query().Get("school_id"))
	if err != nil {
		schoolId = 0
	}

	teachers, err := h.adminService.GetTeachersBySchoolId(ctx, schoolId)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("malchika dura, netu uchiteley v schoole", "error", err)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(teachers); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) ChangeUserBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blocked_string, ok := vars["blocked"]
	if !ok {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("failed to convert school_id into int")
		return
	}

	var blocked bool = false
	if blocked_string == "true" {
		blocked = true
	}

	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("failed to convert school_id into int", "error", err)
		return
	}

	ctx := r.Context()
	err = h.adminService.ChangeBlockingUser(ctx, userId, blocked)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed to change is user blocked or not ", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *AdminHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.adminService.GetAllUsers(ctx)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("something went wrong, i'm tired of this shit. ", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) CreateSchool(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type CreateSchoolRequest struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	var req CreateSchoolRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = h.adminService.CreateSchool(ctx, req.Name, req.Address)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("something went wrong, i'm tired of this shit. ", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	fullname, ok := vars["fullname"]
	if !ok {
		http.Error(w, "name is required", http.StatusBadRequest)
		slog.Error("name is required")
		return
	}
	email, ok := vars["email"]
	if !ok {
		http.Error(w, "email is required", http.StatusBadRequest)
		slog.Error("email is required")
		return
	}
	login, ok := vars["login"]
	if !ok {
		http.Error(w, "login is required", http.StatusBadRequest)
		slog.Error("login is required")
		return
	}
	schoolId, err := strconv.Atoi(vars["school_id"])
	if !ok {
		http.Error(w, "school_id is required", http.StatusBadRequest)
		slog.Error("school_id is required")
		return
	}
	classId, err := strconv.Atoi(vars["class_id"])
	if !ok {
		http.Error(w, "class_id is required", http.StatusBadRequest)
		slog.Error("class_id is required")
		return
	}

	user, err := h.adminService.CreateTeacher(ctx, fullname, login, email, schoolId, classId)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("something went wrong, i'm tired of this shit. ", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(user); err != nil {
		slog.Error("something went wrong, i'm tired of this shit. ", "error", err)
	}
}

// POST admin/school +
// POST admin/teacher +

// GET admin/schools/ +
// GET admin/teachers?school_id= +

// PUT admin/user/block?blocked=&user_id= +
// GET admin/users +
