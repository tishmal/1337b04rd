package http

import (
	"net/http"
	"text/template"

	"1337B04RD/internal/domain/service"
)

// internal/adapter/http/handler.go
type Handler struct {
	postService    *service.PostService
	sessionService *service.SessionService
	templates      *template.Template
}

func NewHandler(postService *service.PostService, sessionService *service.SessionService) *Handler {
	// Загрузка шаблонов
	templates := template.Must(template.ParseGlob("web/static/templates/*.html"))
	return &Handler{postService, sessionService, templates}
}

func (h *Handler) CatalogHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/catalog.html", http.StatusSeeOther)
}

// Другие обработчики
func (h *Handler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим multipart форму
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	userName := r.FormValue("name")
	title := r.FormValue("subject")
	content := r.FormValue("comment")
	file, fileHeader, err := r.FormFile("file")

	// var imageURL string
	if err == nil {
		defer file.Close()

		// Буфер для файла
		buf := make([]byte, fileHeader.Size)
		_, err = file.Read(buf)
		if err != nil {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}

		// // Загрузка изображения в S3
		// imageURL, err = h.postService.UploadImage(buf, fileHeader.Filename)
		// if err != nil {
		// 	http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		// 	return
		// }
	}

	// TODO: session ID из куки
	sessionID := "some-session"

	// Создание поста
	post, err := h.postService.CreatePost(r.Context(), userName, title, content, sessionID)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
	_ = post

	http.Redirect(w, r, "/catalog", http.StatusSeeOther)
}
