package repository

import (
	"context"
	"database/sql"

	"ticket-app/internal/domain"
)

type LoginSessionRepository struct{ DB *sql.DB }

func NewLoginSessionRepository(db *sql.DB) *LoginSessionRepository {
	return &LoginSessionRepository{DB: db}
}

func (r *LoginSessionRepository) Create(ctx context.Context, s domain.LoginSession) error {
	_, err := r.DB.ExecContext(ctx, `
		INSERT INTO login_sessions (session_hash, visitor_id, expires_at)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
			visitor_id = VALUES(visitor_id),
			expires_at = VALUES(expires_at)
	`, s.SessionHash, s.VisitorID, s.ExpiresAt)
	return err
}
