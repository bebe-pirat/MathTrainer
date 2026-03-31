package handler

import (
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	adminService service.AdminService
	classService service.ClassService
}

func NewAdminHandler(adminService service.AdminService, classService service.ClassService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
		classService: classService,
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
	type UserBlockRequest struct {
		UserId  int  `json:"user_id"`
		Blocked bool `json:"blocked"`
	}

	var req UserBlockRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = h.adminService.ChangeBlockingUser(ctx, req.UserId, req.Blocked)
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

	type CreateTeacherRequest struct {
		Login    string `json:"login"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		ClassId  int    `json:"class_id"`
	}

	var req CreateTeacherRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, err := h.adminService.CreateTeacher(ctx, req.Fullname, req.Login, req.Email, req.ClassId)
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

func (h *AdminHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type CreateClassRequest struct {
		Name     string `json:"name"`
		Grade    int    `json:"grade"`
		SchoolId int    `json:"school_id"`
	}

	var req CreateClassRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	classId, err := h.classService.CreateClass(ctx, req.Name, req.Grade, req.SchoolId)
	if err != nil {
		http.Error(w, "failed to create new class", http.StatusInternalServerError)
		slog.Error("failed to create class", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	slog.Info("hero")

	if err := json.NewEncoder(w).Encode(classId); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) GetClassesBySchoolId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	schoolId := r.URL.Query().Get("school_id")

	var classes interface{}
	var err error

	if schoolId != "" {
		schoolId, err := strconv.Atoi(schoolId)
		if err != nil {
			http.Error(w, "invalid school_id", http.StatusBadRequest)
			slog.Error("invalid school_id", "error", err)
			return
		}

		classes, err = h.classService.GetClassesBySchool(ctx, schoolId)
		if err != nil {
			http.Error(w, "failed to get classes", http.StatusInternalServerError)
			slog.Error("failed to get classes", "error", err)
			return
		}
	} else {
		classes, err = h.classService.GetClasses(ctx)
		if err != nil {
			http.Error(w, "failed to get classes", http.StatusInternalServerError)
			slog.Error("failed to get classes", "error", err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(classes); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}
