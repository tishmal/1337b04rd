package http

import (
	"net/http"
	"path/filepath"

	"1337B04RD/internal/domain/service"
)

func WithConditionalSessionMiddleware(sessionService *service.SessionService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Эти пути пропускаем без куки
		if r.URL.Path == "/ping" ||
			filepath.HasPrefix(r.URL.Path, "/static/") ||
			filepath.HasPrefix(r.URL.Path, "/media/") {
			next.ServeHTTP(w, r)
			return
		}

		// Логика с кукой прямо здесь (БЕЗ доп обёртки)
		cookie, err := r.Cookie("session_id")
		var sessionID string

		if err != nil || cookie.Value == "" {
			// Куки нет — создаём новую сессию
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
				MaxAge:   60 * 60 * 24 * 30, // 30 дней
			})
		} else {
			sessionID = cookie.Value
		}

		// Пробрасываем в заголовки
		r.Header.Set("X-Session-ID", sessionID)

		// Теперь вызываем основной роутер
		next.ServeHTTP(w, r)
	})
}
