package usecase

import (
	"context"
	"errors"
	"fmt"

	"ticket-app/internal/domain"
)

var ErrInvalidOrExpired = errors.New("invalid or expired session")

type ResolveSessionUsecase interface {
	ResolveVisitor(ctx context.Context, token string) (uint64, error)
}

type resolveSessionUsecase struct {
	repo domain.SessionRepository
}

func NewResolveSessionUsecase(repo domain.SessionRepository) ResolveSessionUsecase {
	return &resolveSessionUsecase{repo: repo}
}

func (u resolveSessionUsecase) ResolveVisitor(ctx context.Context, token string) (uint64, error) {
	visitorID, ok, err := u.repo.ResolveVisitorByToken(ctx, token)
	if err != nil {
		return 0, fmt.Errorf("resolve failed: %w", err)
	}
	if !ok {
		return 0, ErrInvalidOrExpired
	}
	return visitorID, nil
}
