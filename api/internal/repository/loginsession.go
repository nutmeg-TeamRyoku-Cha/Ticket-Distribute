package repository

import (
	"context"
	"database/sql"

	"ticket-app/internal/auth"
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

func (r *LoginSessionRepository) ResolveVisitorByToken(ctx context.Context, token string) (uint64, bool, error) {
	h, err := auth.SessionHashFromToken(token)
	if err != nil {
		return 0, false, err
	}
	var visitorID uint64
	err = r.DB.QueryRowContext(ctx, `
        SELECT visitor_id
        FROM login_sessions
        WHERE session_hash = ? AND expires_at > UTC_TIMESTAMP(6)
        LIMIT 1
    `, h).Scan(&visitorID)
	if err == sql.ErrNoRows {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, err
	}
	return visitorID, true, nil
}
