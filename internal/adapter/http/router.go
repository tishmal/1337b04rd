package http

import "net/http"

func SetupRoutes(handler *Handler) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", handler.CatalogHandler)
	router.HandleFunc("/catalog", handler.CatalogHandler)
	router.HandleFunc("/create", handler.CreatePostHandler)

	return router
}
