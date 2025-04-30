package http

import (
	"1337b04rd/internal/app/common/logger"
	"1337b04rd/internal/app/common/utils"
	"1337b04rd/internal/app/services"
	"1337b04rd/internal/domain/errors"
	"encoding/json"
	"net/http"
	"strings"
)

type ThreadHandler struct {
	threadSvc *services.ThreadService
}

func NewThreadHandler(threadSvc *services.ThreadService) *ThreadHandler {
	return &ThreadHandler{
		threadSvc: threadSvc,
	}
}

func (h *ThreadHandler) CreateThread(w http.ResponseWriter, r *http.Request) {
	sess, ok := GetSessionFromContext(r.Context())
	if !ok {
		logger.Warn("session not found in CreateThread", "context_value", r.Context().Value(sessionKey))
		Respond(w, http.StatusUnauthorized, map[string]string{"error": "session not found"})
		return
	}

	if err := r.ParseMultipartForm(20 << 20); err != nil {
		logger.Error("failed to parse multipart form", "error", err)
		Respond(w, http.StatusBadRequest, map[string]string{"error": "invalid form data"})
		return
	}

	title := strings.TrimSpace(r.FormValue("title"))
	content := strings.TrimSpace(r.FormValue("content"))

	if title == "" || content == "" {
		Respond(w, http.StatusBadRequest, map[string]string{"error": "title and content are required"})
		return
	}

	files, contentTypes, err := h.threadSvc.PrepareFilesFromMultipart(r.MultipartForm)
	if err != nil {
		logger.Error("failed to process files", "error", err)
		Respond(w, http.StatusBadRequest, map[string]string{"error": "failed to process images"})
		return
	}

	thread, err := h.threadSvc.CreateThread(r.Context(), title, content, files, contentTypes, sess.ID)
	if err != nil {
		logger.Error("failed to create thread", "error", err)
		Respond(w, http.StatusInternalServerError, map[string]string{"error": "could not create thread"})
		return
	}

	Respond(w, http.StatusCreated, map[string]string{
		"thread_id": thread.ID.String(),
	})
}

func (h *ThreadHandler) GetThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Warn("invalid method", "method", r.Method, "path", r.URL.Path)
		Respond(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	ctx := r.Context()
	path := r.URL.Path
	if !strings.HasPrefix(path, "/threads/view/") {
		logger.Warn("invalid path", "path", path)
		Respond(w, http.StatusNotFound, map[string]string{"error": "Not found"})
		return
	}

	idStr := strings.TrimPrefix(path, "/threads/view/")
	id, err := utils.ParseUUID(idStr)
	if err != nil {
		logger.Error("invalid thread id", "error", err, "id", idStr)
		Respond(w, http.StatusBadRequest, map[string]string{"error": "Invalid thread ID"})
		return
	}

	thread, err := h.threadSvc.GetThreadByID(ctx, id)
	if err != nil {
		if err == errors.ErrThreadNotFound {
			logger.Warn("thread not found", "id", id)
			Respond(w, http.StatusNotFound, map[string]string{"error": "Thread not found"})
			return
		}
		logger.Error("failed to get thread", "error", err, "id", id)
		Respond(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get thread"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(thread); err != nil {
		logger.Error("failed to encode response", "error", err)
	}
}

func (h *ThreadHandler) ListActiveThreads(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Warn("invalid method", "method", r.Method, "path", r.URL.Path)
		Respond(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	if r.URL.Path != "/threads" {
		logger.Warn("invalid path", "path", r.URL.Path)
		Respond(w, http.StatusNotFound, map[string]string{"error": "Not found"})
		return
	}

	ctx := r.Context()
	threads, err := h.threadSvc.ListActiveThreads(ctx)
	if err != nil {
		logger.Error("failed to list active threads", "error", err)
		Respond(w, http.StatusInternalServerError, map[string]string{"error": "Failed to list threads"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(threads); err != nil {
		logger.Error("failed to encode response", "error", err)
	}
}

func (h *ThreadHandler) ListAllThreads(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Warn("invalid method", "method", r.Method, "path", r.URL.Path)
		Respond(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	if r.URL.Path != "/threads/all" {
		logger.Warn("invalid path", "path", r.URL.Path)
		Respond(w, http.StatusNotFound, map[string]string{"error": "Not found"})
		return
	}

	ctx := r.Context()
	threads, err := h.threadSvc.ListAllThreads(ctx)
	if err != nil {
		logger.Error("failed to list all threads", "error", err)
		Respond(w, http.StatusInternalServerError, map[string]string{"error": "Failed to list threads"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(threads); err != nil {
		logger.Error("failed to encode response", "error", err)
	}
}

func (h *ThreadHandler) LikeAdd(w http.ResponseWriter, r *http.Request) {
	sess, ok := GetSessionFromContext(r.Context())
	if !ok {
		logger.Warn("session not found in LikeAdd", "context_value", r.Context().Value(sessionKey))
		Respond(w, http.StatusUnauthorized, map[string]string{"error": "session not found"})
		return
	}

	threadIDStr := r.URL.Query().Get("id")
	if threadIDStr == "" {
		Respond(w, http.StatusBadRequest, map[string]string{"error": "missing thread ID"})
		return
	}

	threadID, err := uuidHelper.ParseUUID(threadIDStr)
	if err != nil {
		logger.Error("invalid thread ID", "error", err, "thread_id", threadIDStr)
		Respond(w, http.StatusBadRequest, map[string]string{"error": "invalid thread ID"})
		return
	}

	err = h.threadSvc.LikeAdd(r.Context(), threadID)
	if err != nil {
		logger.Error("failed to add like", "error", err, "thread_id", threadID)
		Respond(w, http.StatusInternalServerError, map[string]string{"error": "could not add like"})
		return
	}

	Respond(w, http.StatusOK, map[string]string{"status": "like added"})
}
