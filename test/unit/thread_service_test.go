package unit

import (
	"bytes"
	"context"
	"io"
	"testing"

	"1337b04rd/internal/app/common/utils"
	"1337b04rd/internal/app/services"
	"1337b04rd/internal/domain/thread"
)

// --- Моки ---

type mockThreadRepo struct {
	created *thread.Thread
	err     error
}

func (m *mockThreadRepo) CreateThread(ctx context.Context, t *thread.Thread) error {
	m.created = t
	return m.err
}

type mockS3 struct {
	uploadedURLs []string
	err          error
}

func (m *mockS3) UploadImages(files map[string]io.Reader, contentTypes map[string]string) ([]string, error) {
	return m.uploadedURLs, m.err
}

// --- Тест ---
func TestCreateThread_Success(t *testing.T) {
	repo := &mockThreadRepo{}
	s3 := &mockS3{
		uploadedURLs: []string{"http://minio:9000/thread/image1.jpg"},
	}

	svc := services.NewThreadService(repo, s3)

	files := map[string]io.Reader{
		"file1.jpg": bytes.NewBuffer([]byte("image data")),
	}
	contentTypes := map[string]string{
		"file1.jpg": "image/jpeg",
	}

	sessionID, err := utils.NewUUID() // или какой-то фиксированный UUID
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	threadResult, err := svc.CreateThread(context.TODO(), "Test Title", "Test Content", files, contentTypes, sessionID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if threadResult == nil {
		t.Fatal("expected thread, got nil")
	}

	if threadResult.Title != "Test Title" {
		t.Errorf("expected title 'Test Title', got '%s'", threadResult.Title)
	}

	expectedURL := "http://localhost:9000/thread/image1.jpg"
	if len(threadResult.ImageURLs) != 1 || threadResult.ImageURLs[0] != expectedURL {
		t.Errorf("expected image URL '%s', got %+v", expectedURL, threadResult.ImageURLs)
	}

	if repo.created != threadResult {
		t.Error("expected thread to be saved in repo")
	}
}
