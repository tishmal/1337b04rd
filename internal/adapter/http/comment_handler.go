package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"1337b04rd/internal/app/common/logger"
	"1337b04rd/internal/app/common/utils"
	"1337b04rd/internal/app/services"
	"1337b04rd/internal/domain/errors"
)

type CommentHandler struct {
	commentSvc *services.CommentService
}

func NewCommentHandler(commentSvc *services.CommentService, logger *slog.Logger) *CommentHandler {
	return &CommentHandler{
		commentSvc: commentSvc,
	}
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Warn("invalid method", "method", r.Method, "path", r.URL.Path)
		Respond(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	sess, ok := GetSessionFromContext(r.Context())
	if !ok {
		logger.Error("session not found in context")
		Respond(w, http.StatusUnauthorized, map[string]string{"error": "Session not found"})
		return
	}
	sessionID := sess.ID
	displayName := sess.DisplayName
	avatarURL := sess.AvatarURL

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		logger.Error("failed to parse form", "error", err)
		Respond(w, http.StatusBadRequest, map[string]string{"error": "Invalid form data"})
		return
	}

	threadIDStr := r.FormValue("thread_id")
	content := strings.TrimSpace(r.FormValue("content"))
	parentIDStr := r.FormValue("parent_id")

	if content == "" {
		Respond(w, http.StatusBadRequest, map[string]string{"error": "Content is required"})
		return
	}

	threadID, err := utils.ParseUUID(threadIDStr)
	if err != nil {
		Respond(w, http.StatusBadRequest, map[string]string{"error": "Invalid thread_id"})
		return
	}

	var parentID *utils.UUID
	if parentIDStr != "" {
		parsedID, err := utils.ParseUUID(parentIDStr)
		if err != nil {
			Respond(w, http.StatusBadRequest, map[string]string{"error": "Invalid parent_id"})
			return
		}
		parentID = &parsedID
	}

	files, contentTypes, err := h.commentSvc.PrepareFilesFromMultipart(r.MultipartForm)
	if err != nil {
		logger.Error("failed to process uploaded files", "error", err)
		Respond(w, http.StatusBadRequest, map[string]string{"error": "Invalid image upload"})
		return
	}

	comment, err := h.commentSvc.CreateComment(r.Context(), threadID, parentID, content, files, contentTypes, sessionID, displayName, avatarURL)
	if err != nil {
		if err == errors.ErrThreadNotFound {
			Respond(w, http.StatusNotFound, map[string]string{"error": "Thread not found"})
			return
		}
		logger.Error("failed to create comment", "error", err)
		Respond(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create comment"})
		return
	}

	Respond(w, http.StatusCreated, comment)
}

func (h *CommentHandler) GetCommentsByThreadID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Warn("invalid method", "method", r.Method, "path", r.URL.Path)
		Respond(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	threadIDStr := r.URL.Query().Get("thread_id")
	if threadIDStr == "" {
		logger.Warn("missing thread_id in query")
		Respond(w, http.StatusBadRequest, map[string]string{"error": "Missing thread_id"})
		return
	}

	threadID, err := utils.ParseUUID(threadIDStr)
	if err != nil {
		logger.Error("invalid thread_id", "error", err, "thread_id", threadIDStr)
		Respond(w, http.StatusBadRequest, map[string]string{"error": "Invalid thread ID"})
		return
	}

	comments, err := h.commentSvc.GetCommentsByThreadID(r.Context(), threadID)
	if err != nil {
		if err == errors.ErrThreadNotFound {
			logger.Warn("thread not found", "thread_id", threadID)
			Respond(w, http.StatusNotFound, map[string]string{"error": "Thread not found"})
			return
		}
		logger.Error("failed to get comments", "error", err, "thread_id", threadID)
		Respond(w, http.StatusInternalServerError, map[string]string{"error": "Failed to get comments"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		logger.Error("failed to encode response", "error", err)
	}
}
