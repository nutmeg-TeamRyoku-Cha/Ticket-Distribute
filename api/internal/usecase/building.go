package usecase

import (
	"context"

	"ticket-app/internal/domain"
)

type BuildingUsecase interface {
	CreateBuilding(ctx context.Context, b domain.Building) (uint64, error)
	ListBuildings(ctx context.Context) ([]domain.Building, error)
	GetBuildingByID(ctx context.Context, id uint64) (domain.Building, bool, error)
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

func (u *buildingUsecase) ListBuildings(ctx context.Context) ([]domain.Building, error) {
	return u.repo.List(ctx)
}

func (u *buildingUsecase) GetBuildingByID(ctx context.Context, id uint64) (domain.Building, bool, error) {
	return u.repo.GetByID(ctx, id)
}
