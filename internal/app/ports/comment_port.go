package ports

import (
	"context"

	"1337b04rd/internal/domain/comment"

	uuidHelper "1337b04rd/internal/app/common/utils"
)

type CommentPort interface {
	CreateComment(ctx context.Context, c *comment.Comment) error
	GetCommentsByThreadID(ctx context.Context, threadID uuidHelper.UUID) ([]*comment.Comment, error)
}
