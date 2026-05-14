package handler

import (
	"MathTrainer/internal/model"
	"MathTrainer/internal/service"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (h *AdminHandler) CreateSection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type CreateSectionRequest struct {
		Name         string `json:"name"`
		Class        int    `json:"class"`
		SectionOrder int    `json:"section_order"`
	}

	var req CreateSectionRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = h.adminService.CreateSection(ctx, model.Section{Name: req.Name, Class: req.Class, Order: req.SectionOrder})
	if err != nil && errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "failed to create new section", http.StatusInternalServerError)
		slog.Error("failed to create section", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) UpdateSection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "bad request, no id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	type UpdateSectionRequest struct {
		Name         string `json:"name"`
		Class        int    `json:"class"`
		SectionOrder int    `json:"section_order"`
	}

	var req UpdateSectionRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	err = h.adminService.UpdateSection(ctx, model.Section{Id: id, Name: req.Name, Class: req.Class, Order: req.SectionOrder})
	if errors.Is(err, model.ErrNotFound) {
		http.Error(w, "failed to find section with id", http.StatusNotFound)
		slog.Error("failed to find section to update", "id", id, "error", err)
		return
	}
	if errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("bad request", "error", err)
		return
	}
	if err != nil {
		http.Error(w, "failed to update section", http.StatusInternalServerError)
		slog.Error("failed to update section", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) DeleteSection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "bad request, no id", http.StatusBadRequest)
		slog.Error("failed to get id to delete section", "URL", r.RequestURI)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		slog.Error("failed to converte id to INT", "id", idStr)
		return
	}

	err = h.adminService.DeleteSection(ctx, id)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		http.Error(w, "failed to find section with id", http.StatusNotFound)
		slog.Error("failed to find section to update", "id", id, "error", err)
		return
	}
	if err != nil {
		http.Error(w, "failed to delete section", http.StatusInternalServerError)
		slog.Error("failed to delete section", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) GetSections(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	class := 0
	var err error
	classStr := r.URL.Query().Get("class")
	if classStr != "" {
		class, err = strconv.Atoi(classStr)
		if err != nil {
			class = 0
		}
	}

	sections, err := h.adminService.GetSections(ctx, class)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		slog.Error("failed to get sections", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(sections); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.CreateAndUpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		slog.Error("invalid input to create user", "error", err, "req", req)
		return
	}

	creds, err := h.adminService.CreateUser(ctx, req)
	if err != nil && errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("invalid input to create user jj", "error", err)
		return
	}
	if err != nil {
		http.Error(w, "failed to create new user", http.StatusInternalServerError)
		slog.Error("failed to create user", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(creds); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("no id for update user")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("failed to convertd id to int update user", "error", err)
	}

	var req model.CreateAndUpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		slog.Error("invalid input to update user", "error", err)
		return
	}

	req.Id = id
	err = h.adminService.UpdateUser(ctx, req)
	if err != nil && errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("invalid input to update user", "error", err)
		return
	}
	if err != nil && errors.Is(err, model.ErrNotFound) {
		http.Error(w, "not found", http.StatusBadRequest)
		slog.Error("failed to found user for update ", "error", err, "id", id)
		return
	}
	if err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		slog.Error("failed to update user", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("no id for update user")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("failed to convertd id to int update user", "error", err)
	}

	err = h.adminService.DeleteUser(ctx, id)
	if err != nil && errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("invalid input to update user", "error", err)
		return
	}
	if err != nil && errors.Is(err, model.ErrNotFound) {
		http.Error(w, "not found", http.StatusBadRequest)
		slog.Error("failed to found user for delete ", "error", err, "id", id)
		return
	}
	if err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		slog.Error("failed to delete user", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	slog.Info("holas")

	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("no id for update user")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("failed to convertd id to int update user", "error", err)
	}

	creds, err := h.adminService.UpdatePassword(ctx, id)
	if err != nil && errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("invalid input to create user", "error", err)
		return
	}
	if err != nil {
		http.Error(w, "failed to change user password", http.StatusInternalServerError)
		slog.Error("failed to change user password", "error", err, "id", id)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(creds); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) CreateEquationType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req model.CreateEquationTypeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		slog.Error("invalid input to create equation type", "error", err, "req", req)
		return
	}

	err = h.adminService.CreateEquationType(ctx, req)
	if err != nil && errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("invalid input to create equation type", "error", err)
		return
	}
	if err != nil {
		http.Error(w, "failed to create new equation type", http.StatusInternalServerError)
		slog.Error("failed to create equation type", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) UpdateEquationType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		http.Error(w, "no id", http.StatusBadRequest)
		slog.Error("no id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		slog.Error("failed to convert id into int update equationtype", "error", err, "idStr", idStr)
		return
	}

	var req model.UpdateEquationTypeRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		slog.Error("invalid input to update equation type", "error", err, "req", req)
		return
	}

	err = h.adminService.UpdateEquationType(ctx, id, req)
	if err != nil && errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("invalid input to update equation type", "error", err)
		return
	}
	if err != nil && errors.Is(err, model.ErrNotFound) {
		http.Error(w, "not found", http.StatusBadRequest)
		slog.Error("not found equation type for update", "error", err, "id", id)
		return
	}
	if err != nil {
		http.Error(w, "failed to update new equation type", http.StatusInternalServerError)
		slog.Error("failed to update equation type", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) DeleteEquationType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		http.Error(w, "no id", http.StatusBadRequest)
		slog.Error("no id")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		slog.Error("failed to convert id into int update equationtype", "error", err, "idStr", idStr)
		return
	}

	err = h.adminService.DeleteEquationType(ctx, id)
	if err != nil && errors.Is(err, model.ErrBadRequest) {
		http.Error(w, "bad request", http.StatusBadRequest)
		slog.Error("invalid input to update equation type", "error", err)
		return
	}
	if err != nil && errors.Is(err, model.ErrNotFound) {
		http.Error(w, "not found", http.StatusBadRequest)
		slog.Error("not found equation type for update", "error", err, "id", id)
		return
	}
	if err != nil {
		http.Error(w, "failed to update new equation type", http.StatusInternalServerError)
		slog.Error("failed to update equation type", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) GetEquationTypes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	types, err := h.adminService.GetEquationTypes(ctx)

	if err != nil {
		http.Error(w, "bad request", http.StatusInternalServerError)
		slog.Error("failed to get equation types", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(types); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) GetOperandsForEquationType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	equationTypeIdStr := r.URL.Query().Get("equation_type_id")
	if equationTypeIdStr == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	equationTypeId, err := strconv.Atoi(equationTypeIdStr)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ops, err := h.adminService.GetOperandsForEquationType(ctx, equationTypeId)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		http.Error(w, "not found", http.StatusNotFound)
		slog.Error("operands for this equationTypId is not found", "id", equationTypeId, "error", err)
		return
	}
	if err != nil {
		http.Error(w, "bad request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(ops); err != nil {
		slog.Error("serializtion failed", "error", err)
	}
}

func (h *AdminHandler) JoinEquationTypeAndSection(w http.ResponseWriter, r *http.Request) {
	var req struct {
		EquationTypeID int `json:"equation_type_id"`
		SectionID      int `json:"section_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err := h.adminService.JoinEquationTypeAndSection(ctx, req.EquationTypeID, req.SectionID)
	if err != nil {
		http.Error(w, "failed to join equationType with section", http.StatusInternalServerError)
		slog.Error("failed to join equationType with section", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) UnJoinEquationTypeAndSection(w http.ResponseWriter, r *http.Request) {
	eqID, err := strconv.Atoi(r.URL.Query().Get("equation_type_id"))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	secID, err := strconv.Atoi(r.URL.Query().Get("section_id"))
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = h.adminService.UnJoinEquationTypeAndSection(ctx, eqID, secID)
	if err != nil {
		http.Error(w, "failed to unjoin equationType with section", http.StatusInternalServerError)
		slog.Error("failed to unjoin equationType with section", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AdminHandler) GetSectionsAndEquationTypes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	secAndeqs, err := h.adminService.GetSectionsAndEquationTypes(ctx)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed to get sections and eqatio n types", "error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if writer := json.NewEncoder(w); writer.Encode(secAndeqs) != nil {
		slog.Error("serialization failed", "error", err)
		return
	}
}
