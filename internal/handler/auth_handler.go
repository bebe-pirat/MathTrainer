package handler

import (
	"MathTrainer/internal/service"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	authService service.AuthService
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	login := vars["login"]
	password := vars["password"]

	sessionData, err := h.authService.Login(ctx, login, password)
	if err != nil {
		http.Error(w, "something went wrong auth", http.StatusInternalServerError)
		slog.Error("authentication failed", "error", err)
		return
	}

	if err := setSessionCookie(w, *sessionData); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		slog.Error("failed to set cookie", "error", err)
		return
	}
	slog.Info("user logged in successfully", "user_id", sessionData.UserID, "role", sessionData.Role)
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		slog.Error("failed to get session from cookie", "error", err)
		return
	}

	err = h.authService.Logout(ctx, sessionData.SessionID)
	if err != nil {
		http.Error(w, "logout failed", http.StatusInternalServerError)
		slog.Error("logout failed", "error", err)
		return
	}

	clearSessionCookie(w)

	slog.Info("user logged out successfully", "user_id", sessionData.UserID)
	w.WriteHeader(http.StatusOK)
}
