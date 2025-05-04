package unit

import (
	"1337b04rd/internal/adapter/postgres"
	"1337b04rd/internal/domain/errors"
	"context"
	"fmt"
	"testing"

	"1337b04rd/internal/app/common/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestThreadRepository_LikeAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := &postgres.ThreadRepository{}

	threadID, err := utils.NewUUID()
	if err != nil {
		require.ErrorIs(t, err, errors.ErrInvalidThreadID)
	}
	sessionID, err := utils.NewUUID()
	if err != nil {
		require.ErrorIs(t, err, errors.ErrInvalidSessionID)
	}

	t.Run("successfully add like", func(t *testing.T) {
		mock.ExpectQuery("SELECT EXISTS").
			WithArgs(threadID.String(), sessionID.String()).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

		mock.ExpectBegin()

		mock.ExpectExec("INSERT INTO thread_likes").
			WithArgs(threadID, sessionID, sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectExec("UPDATE threads SET likes = likes").
			WithArgs(threadID.String()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		err := repo.LikeAdd(context.Background(), threadID, sessionID)
		require.NoError(t, err)
	})

	t.Run("already liked", func(t *testing.T) {
		mock.ExpectQuery("SELECT EXISTS").
			WithArgs(threadID.String(), sessionID.String()).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

		err := repo.LikeAdd(context.Background(), threadID, sessionID)
		require.ErrorIs(t, err, errors.ErrAlreadyLiked)
	})

	t.Run("foreign key violation", func(t *testing.T) {
		mock.ExpectQuery("SELECT EXISTS").
			WithArgs(threadID.String(), sessionID.String()).
			WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

		mock.ExpectBegin()

		mock.ExpectExec("INSERT INTO thread_likes").
			WithArgs(threadID, sessionID, sqlmock.AnyArg()).
			WillReturnError(fmt.Errorf("pq: insert or update on table violates foreign key"))

		mock.ExpectRollback()

		err := repo.LikeAdd(context.Background(), threadID, sessionID)
		require.ErrorContains(t, err, "foreign key")
	})
}
