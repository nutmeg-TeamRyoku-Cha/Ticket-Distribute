package usecase

import (
	"context"
	"ticket-app/internal/domain"
	"time"
)

type VisitorUsecase interface {
	CreateVisitor(ctx context.Context, v domain.Visitor) (uint64, error)
	ListVisitors(ctx context.Context) ([]domain.Visitor, error)
}

type visitorUsecase struct {
	repo domain.VisitorRepository
	now  func() time.Time
}

func NewVisitorUsecase(repo domain.VisitorRepository) VisitorUsecase {
	return &visitorUsecase{repo: repo, now: time.Now}
}

func (u *visitorUsecase) CreateVisitor(ctx context.Context, v domain.Visitor) (uint64, error) {
	// ensure birth_date is normalized
	v.BirthDate = v.BirthDate.UTC()
	return u.repo.Create(ctx, v)
}

func (u *visitorUsecase) ListVisitors(ctx context.Context) ([]domain.Visitor, error) {
	return u.repo.List(ctx)
}
