package handler

import (
	"MathTrainer/internal/model"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func InitCookieStore(secretKey string) {
	s = securecookie.New([]byte(secretKey), nil)
}

func getSessionFromCookie(r *http.Request) (*model.SessionData, error) {
	cookie, err := r.Cookie("session_data")
	if err != nil {
		return nil, err
	}

	var sessionData model.SessionData
	if err := s.Decode("session_data", cookie.Value, &sessionData); err != nil {
		return nil, err
	}

	return &sessionData, nil
}

func setSessionCookie(w http.ResponseWriter, sessionData model.SessionData) error {
	encoded, err := s.Encode("session_data", sessionData)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "session_data",
		Value:    encoded,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		// Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
	return nil
}

func clearSessionCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session_data",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}
