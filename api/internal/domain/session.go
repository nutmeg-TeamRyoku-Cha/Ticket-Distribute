package domain

import (
	"context"
	"errors"
	"time"

	"ticket-app/internal/auth"
)

var ErrSessionNotFound = errors.New("session not found or expired")

type Session struct {
	SessionHash []byte
	VisitorID   uint64
	ExpiresAt   time.Time
}

type TokenAndSession struct {
	Token   string
	Session Session
}

type SessionRepository interface {
	Create(ctx context.Context, s Session) error
	VisitorProfByToken(ctx context.Context, token string) (Visitor, error)
}

// ユーザーに見せない情報なので、Domainで実装
func IssueSession(visitorID uint64, now time.Time) (TokenAndSession, error) {
	token, hash, err := auth.NewSessionToken()
	if err != nil {
		return TokenAndSession{}, err
	}
	expires := now.UTC().Truncate(time.Microsecond).Add(6 * time.Hour)

	return TokenAndSession{
		Token: token,
		Session: Session{
			SessionHash: hash,
			VisitorID:   visitorID,
			ExpiresAt:   expires,
		},
	}, nil
}
