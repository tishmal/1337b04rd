package http

import (
	"1337B04RD/internal/domain/service"
	"net/http"
	"text/template"
)

// internal/adapter/http/handler.go
type Handler struct {
	postService    *service.PostService
	sessionService *service.SessionService
	templates      *template.Template
}

func NewHandler(postService *service.PostService, sessionService *service.SessionService) *Handler {
	// Загрузка шаблонов
	templates := template.Must(template.ParseGlob("web/templates/*.html"))
	return &Handler{postService, sessionService, templates}
}

// Обработчики для различных эндпоинтов
func (h *Handler) CatalogHandler(w http.ResponseWriter, r *http.Request) {
	// Логика обработки запроса к каталогу
}

// Другие обработчики
