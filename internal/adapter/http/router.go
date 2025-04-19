package http

import (
	"net/http"
	"path/filepath"
)

func SetupRoutes(handler *Handler) *http.ServeMux {
	router := http.NewServeMux()

	// Обработка перенаправлений на catalog.html
	router.HandleFunc("/", handler.CatalogHandler)
	router.HandleFunc("/catalog", handler.CatalogHandler)

	// Обработка маршрута для создания постов
	router.HandleFunc("/submit-post", handler.CreatePostHandler)

	// Подключаем статические файлы из папки web/static/templates
	// Пример: если файл catalog.html находится в web/static/templates/catalog.html
	fs := http.FileServer(http.Dir(filepath.Join("web", "static", "templates")))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	return router
}
