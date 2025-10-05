package usecase

import (
	"context"
	"errors"
	"fmt"
	"ticket-app/internal/domain"
	"time"
)

var ErrInvalidOrExpired = errors.New("invalid or expired session")

type SessionUsecase interface {
	CreateSession(ctx context.Context, visitorID uint64) (domain.TokenAndSession, error)
	ResolveVisitorProfile(ctx context.Context, token string) (domain.Visitor, error)
}

type sessionUsecase struct {
	repo domain.SessionRepository
	now  func() time.Time
}

func NewSessionUsecase(repo domain.SessionRepository) SessionUsecase {
	return &sessionUsecase{repo: repo, now: time.Now}
}

func (u sessionUsecase) CreateSession(ctx context.Context, visitorID uint64) (domain.TokenAndSession, error) {
	ts, err := domain.IssueSession(visitorID, u.now())
	if err != nil {
		return domain.TokenAndSession{}, err
	}
	if err := u.repo.Create(ctx, ts.LoginSession); err != nil {
		return domain.TokenAndSession{}, err
	}
	return ts, nil
}

func (u *sessionUsecase) ResolveVisitorProfile(ctx context.Context, token string) (domain.Visitor, error) {
	v, err := u.repo.VisitorProfByToken(ctx, token)
	if err != nil {
		if errors.Is(err, domain.ErrSessionNotFound) {
			return domain.Visitor{}, ErrInvalidOrExpired
		}
		return domain.Visitor{}, fmt.Errorf("resolve profile failed: %w", err)
	}
	return v, nil
}
