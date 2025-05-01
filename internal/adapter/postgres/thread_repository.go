package postgres

import (
	"1337b04rd/internal/app/common/logger"
	"1337b04rd/internal/domain/errors"
	"1337b04rd/internal/domain/thread"
	"context"
	"database/sql"
	"time"

	uuidHelper "1337b04rd/internal/app/common/utils"

	"github.com/lib/pq"
)

type ThreadRepository struct {
	db *sql.DB
}

func NewThreadRepository(db *sql.DB) *ThreadRepository {
	return &ThreadRepository{db: db}
}

func (r *ThreadRepository) CreateThread(ctx context.Context, t *thread.Thread) error {
	if err := ctx.Err(); err != nil {
		logger.Error("context error while creating thread", "error", err, "thread_id", t.ID)
		return err
	}

	_, err := r.db.ExecContext(ctx, CreateThread,
		t.ID.String(),
		t.Title,
		t.Content,
		pq.Array(t.ImageURLs),
		t.SessionID.String(),
		t.CreatedAt,
		t.LastCommented,
		t.IsDeleted,
	)
	if err != nil {
		logger.Error("failed to execute create thread query", "error", err, "thread_id", t.ID)
	}
	return err
}

func (r *ThreadRepository) GetThreadByID(ctx context.Context, id uuidHelper.UUID) (*thread.Thread, error) {
	if err := ctx.Err(); err != nil {
		logger.Error("context error while getting thread by ID", "error", err, "thread_id", id)
		return nil, err
	}

	row := r.db.QueryRowContext(ctx, GetThreadByID, id.String())
	t, err := scanThread(row)
	if err == sql.ErrNoRows {
		return nil, errors.ErrThreadNotFound
	}
	if err != nil {
		logger.Error("failed to scan thread row", "error", err, "thread_id", id)
	}
	return t, err
}

func (r *ThreadRepository) UpdateThread(ctx context.Context, t *thread.Thread) error {
	if err := ctx.Err(); err != nil {
		logger.Error("context error while updating thread", "error", err, "thread_id", t.ID)
		return err
	}

	_, err := r.db.ExecContext(ctx, UpdateThread,
		t.ID.String(),
		t.Title,
		t.Content,
		pq.Array(t.ImageURLs),
		t.SessionID.String(),
		t.CreatedAt,
		t.LastCommented,
		t.IsDeleted,
	)
	if err != nil {
		logger.Error("failed to execute update thread query", "error", err, "thread_id", t.ID)
	}
	return err
}

func (r *ThreadRepository) ListActiveThreads(ctx context.Context) ([]*thread.Thread, error) {
	if err := ctx.Err(); err != nil {
		logger.Error("context error while listing active threads", "error", err)
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, ListActiveThreads)
	if err != nil {
		logger.Error("failed to execute list active threads query", "error", err)
		return nil, err
	}
	defer rows.Close()

	var threads []*thread.Thread
	for rows.Next() {
		t, err := scanThread(rows)
		if err != nil {
			logger.Error("failed to scan thread", "error", err)
			return nil, err
		}
		threads = append(threads, t)
	}

	if err := rows.Err(); err != nil {
		logger.Error("error occurred during rows iteration for active threads", "error", err)
		return nil, err
	}

	return threads, nil
}

func (r *ThreadRepository) ListAllThreads(ctx context.Context) ([]*thread.Thread, error) {
	if err := ctx.Err(); err != nil {
		logger.Error("context error while listing all threads", "error", err)
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, ListAllThreads)
	if err != nil {
		logger.Error("failed to execute list all threads query", "error", err)
		return nil, err
	}
	defer rows.Close()

	var threads []*thread.Thread
	for rows.Next() {
		t, err := scanThread(rows)
		if err != nil {
			logger.Error("failed to scan thread", "error", err)
			return nil, err
		}
		threads = append(threads, t)
	}

	if err := rows.Err(); err != nil {
		logger.Error("error occurred during rows iteration for all threads", "error", err)
		return nil, err
	}

	return threads, nil
}

func scanThread(scanner interface {
	Scan(dest ...interface{}) error
}) (*thread.Thread, error) {
	t := &thread.Thread{}
	var (
		imageURLs     pq.StringArray
		lastCommented sql.NullTime
		idStr         string
		sessionIDStr  string
	)

	err := scanner.Scan(
		&idStr,
		&t.Title,
		&t.Content,
		&imageURLs,
		&sessionIDStr,
		&t.CreatedAt,
		&lastCommented,
		&t.IsDeleted,
	)
	if err != nil {
		logger.Error("failed to scan thread row", "error", err)
		return nil, err
	}

	t.ID, err = uuidHelper.ParseUUID(idStr)
	if err != nil {
		logger.Error("invalid UUID format for ID", "value", idStr, "error", err)
		return nil, err
	}

	t.SessionID, err = uuidHelper.ParseUUID(sessionIDStr)
	if err != nil {
		logger.Error("invalid UUID format for session_id", "value", sessionIDStr, "error", err)
		return nil, err
	}

	if lastCommented.Valid {
		t.LastCommented = &lastCommented.Time
	}

	t.ImageURLs = []string(imageURLs)
	return t, nil
}

func (r *ThreadRepository) LikeAdd(ctx context.Context, threadID, sessionID uuidHelper.UUID) error {
	var exists bool

	err := r.db.QueryRowContext(ctx, LikeExist, threadID.String(), sessionID.String()).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	now := time.Now()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // проброс паники дальше
		} else if err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, ThreadLikesInsert, threadID.String(), sessionID.String(), now)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, LikeAddThread, threadID.String())
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *ThreadRepository) GetCountLikes(ctx context.Context, threadID uuidHelper.UUID) (int, error) {
	if err := ctx.Err(); err != nil {
		logger.Error("context error while getting likes", "error", err)
		return 0, err
	}

	var likes int
	err := r.db.QueryRowContext(ctx, LikesAll, threadID.String()).Scan(&likes)
	if err != nil {
		logger.Error("failed to get likes for thread", "error", err, "thread_id", threadID)
		return 0, err
	}

	return likes, nil
}
