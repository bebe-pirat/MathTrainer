package handler

import (
	"MathTrainer/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) CheckSession(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.Info("i am here")
	
	sessionData, err := getSessionFromCookie(r)
	if err != nil {
		http.Error(w, "failed to find session",http.StatusUnauthorized )
		slog.Info("no active session found")
		return
	}
	
	valid, err := h.authService.ValidateSession(ctx, sessionData.SessionID)
	if err != nil || !valid {
		clearSessionCookie(w)
		http.Error(w, "session expired or invalid",http.StatusUnauthorized )
		slog.Info("invalid session", "session_id", sessionData.SessionID)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(sessionData); err != nil {
		slog.Info("failed to serialize sessionData", "user_id", sessionData.UserID)
	}
	
	slog.Info("session checked successfully", "user_id", sessionData.UserID)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var credentials struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		slog.Error("failed to decode JSON", "error", err)
		return
	}

	login := strings.TrimSpace(credentials.Login)
	password := credentials.Password

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

	var response struct {
		UserID int `json:"user_id"`
		Role   int `json:"role"`
	}

	response.UserID = sessionData.UserID
	response.Role = sessionData.Role

	if err = json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("failed to turn to json session data")
	}
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
