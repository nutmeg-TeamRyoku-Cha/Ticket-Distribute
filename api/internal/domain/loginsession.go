package domain

import (
	"context"
	"time"

	"ticket-app/internal/auth"
)

type LoginSession struct {
	SessionHash []byte
	VisitorID   uint64
	ExpiresAt   time.Time
}

type TokenAndSession struct {
	Token        string
	LoginSession LoginSession
}

type SessionRepository interface {
	Create(ctx context.Context, s LoginSession) error
}

func IssueSession(visitorID uint64, now time.Time) (TokenAndSession, error) {
	token, hash, err := auth.NewSessionToken()
	if err != nil {
		return TokenAndSession{}, err
	}
	expires := now.UTC().Truncate(time.Microsecond).Add(6 * time.Hour)

	return TokenAndSession{
		Token: token,
		LoginSession: LoginSession{
			SessionHash: hash,
			VisitorID:   visitorID,
			ExpiresAt:   expires,
		},
	}, nil
}
