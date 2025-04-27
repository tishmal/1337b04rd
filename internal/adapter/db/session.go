package db

import (
	"1337B04RD/internal/domain/entity"
)

func (r *PostgresRepository) Save(session *entity.Session) error {
	_, err := r.db.Exec(`
		INSERT INTO sessions (session_id, expires_at)
		VALUES ($1, $2);
	`, session.ID, session.ExpiresAt)
	return err
}

func (r *PostgresRepository) Get(sessionID string) (*entity.Session, error) {
	row := r.db.QueryRow(`SELECT session_id, expires_at FROM sessions WHERE session_id=$1`, sessionID)
	var s entity.Session
	err := row.Scan(&s.ID, &s.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *PostgresRepository) Delete(sessionID string) error {
	_, err := r.db.Exec(`DELETE FROM sessions WHERE session_id=$1`, sessionID)
	return err
}
