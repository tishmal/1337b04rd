package http

import (
	"net/http"
	"path/filepath"

	"1337B04RD/internal/domain/service"
)

func SetupRoutes(handler *Handler, sessionService *service.SessionService) http.Handler {
	router := http.NewServeMux()

	// ---------- HTML Views ----------
	router.HandleFunc("/catalog", handler.CatalogHandler)
	// router.HandleFunc("/post/", handler.SinglePostPageHandler)
	router.HandleFunc("/submit-post", handler.CreatePostHandler)

	// ---------- API: Post ----------
	// router.HandleFunc("/api/posts", handler.PostsHandler)
	// router.HandleFunc("/api/posts/", handler.PostByIDHandler)

	// ---------- API: Session / Cookies ----------
	// router.HandleFunc("/api/session", handler.SessionHandler)
	// router.HandleFunc("/api/logout", handler.LogoutHandler)

	// ---------- Static Files ----------
	fs := http.FileServer(http.Dir(filepath.Join("web", "static", "templates")))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("web/media"))))

	// ---------- Healthcheck ----------
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// ---------- Оборачиваем весь роутер ----------
	return WithConditionalSessionMiddleware(sessionService, router)
}
