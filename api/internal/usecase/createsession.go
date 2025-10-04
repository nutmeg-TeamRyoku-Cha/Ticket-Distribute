package usecase

import (
	"context"
	"ticket-app/internal/domain"
	"time"
)

type CreateSessionUsecase interface {
	CreateSession(ctx context.Context, visitorID uint64) (domain.TokenAndSession, error)
}

type createsessionUsecase struct {
	repo domain.SessionRepository
	now  func() time.Time
}

func NewSessionUsecase(repo domain.SessionRepository) CreateSessionUsecase {
	return &createsessionUsecase{repo: repo, now: time.Now}
}

func (u createsessionUsecase) CreateSession(ctx context.Context, visitorID uint64) (domain.TokenAndSession, error) {
	ts, err := domain.IssueSession(visitorID, u.now())
	if err != nil {
		return domain.TokenAndSession{}, err
	}
	if err := u.repo.Create(ctx, ts.LoginSession); err != nil {
		return domain.TokenAndSession{}, err
	}
	return ts, nil
}
