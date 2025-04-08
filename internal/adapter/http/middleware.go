package http

import (
	"1337B04RD/internal/domain/service"
	"net/http"
)

// internal/adapter/http/middleware.go
func SessionMiddleware(sessionService *service.SessionService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Проверка и управление сессией
		})
	}
}

// Другие middleware
