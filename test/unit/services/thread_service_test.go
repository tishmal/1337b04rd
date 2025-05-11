package services

import (
	"bytes"
	"context"
	"io"
	"testing"

	uuidHelper "1337b04rd/internal/app/common/utils"
	"1337b04rd/internal/app/services"

	"1337b04rd/internal/domain/thread"
)

type mockThreadRepo struct {
	GetThreadByIDFunc func(ctx context.Context, id uuidHelper.UUID) (*thread.Thread, error)
	created           *thread.Thread
	err               error
}

func (m *mockThreadRepo) CreateThread(ctx context.Context, t *thread.Thread) error {
	m.created = t
	return m.err
}

// Заглушки
func (m *mockThreadRepo) GetThreadByID(ctx context.Context, id uuidHelper.UUID) (*thread.Thread, error) {
	if m.GetThreadByIDFunc != nil {
		return m.GetThreadByIDFunc(ctx, id)
	}
	return nil, nil
}
func (m *mockThreadRepo) UpdateThread(ctx context.Context, t *thread.Thread) error {
	return nil
}
func (m *mockThreadRepo) ListActiveThreads(ctx context.Context) ([]*thread.Thread, error) {
	return nil, nil
}
func (m *mockThreadRepo) ListAllThreads(ctx context.Context) ([]*thread.Thread, error) {
	return nil, nil
}
func (m *mockThreadRepo) LikeAdd(ctx context.Context, threadID, sessionID uuidHelper.UUID) error {
	return nil
}
func (m *mockThreadRepo) LikeRemove(ctx context.Context, threadID, sessionID uuidHelper.UUID) error {
	return nil
}
func (m *mockThreadRepo) GetCountLikes(ctx context.Context, threadID uuidHelper.UUID) (int, error) {
	return 0, nil
}

type mockS3 struct {
	uploaded map[string]string
	err      error
}

func (s *mockS3) UploadImages(files map[string]io.Reader, contentTypes map[string]string) (map[string]string, error) {
	return s.uploaded, s.err
}
func (s *mockS3) UploadImage(file io.Reader, size int64, contentType string) (string, error) {
	return "", nil
}
func (s *mockS3) DeleteFile(fileName string) error {
	return nil
}

func TestCreateThread_Success(t *testing.T) {
	repo := &mockThreadRepo{}
	s3 := &mockS3{
		uploaded: map[string]string{
			"file1.jpg": "http://localhost:9000/thread/image1.jpg",
		},
	}

	svc := services.NewThreadService(repo, s3)

	files := map[string]io.Reader{
		"file1.jpg": bytes.NewBuffer([]byte("image data")),
	}
	contentTypes := map[string]string{
		"file1.jpg": "image/jpeg",
	}

	sessionID, _ := uuidHelper.ParseUUID("123e4567e89b12d3a456426614174000")

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

	if repo.created == nil {
		t.Error("expected thread to be saved in repo")
	}

	if repo.created.ID != threadResult.ID {
		t.Error("saved thread does not match returned thread")
	}
}

func TestGetThreadByID_Success(t *testing.T) {
	expectedID, _ := uuidHelper.ParseUUID("123e4567e89b12d3a456426614174001")
	mockThread := &thread.Thread{
		ID:        expectedID,
		Title:     "Test Title",
		Content:   "Test Content",
		ImageURLs: []string{"http://minio:9000/thread/image1.jpg"},
	}

	repo := &mockThreadRepo{
		GetThreadByIDFunc: func(ctx context.Context, id uuidHelper.UUID) (*thread.Thread, error) {
			if id != expectedID {
				t.Errorf("expected ID %s, got %s", expectedID, id)
			}
			return mockThread, nil
		},
	}
	s3 := &mockS3{}
	svc := services.NewThreadService(repo, s3)

	result, err := svc.GetThreadByID(context.TODO(), expectedID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if result == nil {
		t.Fatal("expected thread, got nil")
	}
	if result.ID != expectedID {
		t.Errorf("expected ID %s, got %s", expectedID, result.ID)
	}
	expectedURL := "http://localhost:9000/thread/image1.jpg"
	if result.ImageURLs[0] != expectedURL {
		t.Errorf("expected image URL '%s', got '%s'", expectedURL, result.ImageURLs[0])
	}
}
