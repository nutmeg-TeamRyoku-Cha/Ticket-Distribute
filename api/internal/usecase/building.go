package usecase

import (
	"context"

	"ticket-app/internal/domain"
)

type BuildingUsecase interface {
	CreateBuilding(ctx context.Context, b domain.Building) (uint64, error)
}

type buildingUsecase struct {
	repo domain.BuildingRepository
}

func NewBuildingUsecase(repo domain.BuildingRepository) BuildingUsecase {
	return &buildingUsecase{repo: repo}
}

func (u *buildingUsecase) CreateBuilding(ctx context.Context, b domain.Building) (uint64, error) {
	return u.repo.Create(ctx, b)
}
