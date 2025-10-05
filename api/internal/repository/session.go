package repository

import (
	"context"
	"database/sql"
	"time"

	"ticket-app/internal/auth"
	"ticket-app/internal/domain"
)

type LoginSessionRepository struct{ DB *sql.DB }

func NewSessionRepository(db *sql.DB) *LoginSessionRepository {
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

func (r *LoginSessionRepository) VisitorProfByToken(ctx context.Context, token string) (domain.Visitor, error) {
	h, err := auth.SessionHashFromToken(token)
	if err != nil {
		return domain.Visitor{}, err
	}

	var v domain.Visitor
	var birth sql.NullTime
	var party sql.NullInt64

	err = r.DB.QueryRowContext(ctx, `
		SELECT v.visitor_id, v.nickname, v.birth_date, v.party_size
		FROM login_sessions s
		JOIN visitors v ON v.visitor_id = s.visitor_id
		WHERE s.session_hash = ? AND s.expires_at > UTC_TIMESTAMP(6)
		LIMIT 1
	`, h).Scan(&v.VisitorID, &v.Nickname, &birth, &party)
	if err == sql.ErrNoRows {
		return domain.Visitor{}, domain.ErrSessionNotFound
	}
	if err != nil {
		return domain.Visitor{}, err
	}

	if birth.Valid {
		v.BirthDate = birth.Time
	} else {
		v.BirthDate = time.Time{} // ゼロ時刻（必要に応じて意味を決める）
	}
	if party.Valid {
		v.PartySize = int(party.Int64)
	} else {
		v.PartySize = 0
	}

	return v, nil
}
