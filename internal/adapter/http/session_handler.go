package http

import (
	"net/http"

	"1337B04RD/internal/domain/service"
)

func EnsureSession(sessionService *service.SessionService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		var sessionID string

		if err != nil || cookie.Value == "" {
			session, err := sessionService.CreateSession()
			if err != nil {
				http.Error(w, "Failed to create session", http.StatusInternalServerError)
				return
			}
			sessionID = session.ID
			http.SetCookie(w, &http.Cookie{
				Name:     "session_id",
				Value:    sessionID,
				Path:     "/",
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
				MaxAge:   60 * 60 * 24 * 30, // 30 дней
			})
		} else {
			sessionID = cookie.Value
		}

		// Прокидываем session_id в заголовок (или можно в context, если нужно глубже)
		r.Header.Set("X-Session-ID", sessionID)

		next(w, r)
	}
}
