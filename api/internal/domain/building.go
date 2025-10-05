package domain

import "context"

type Building struct {
	BuildingID   uint64
	BuildingName string
	Latitude     *float64
	Longitude    *float64
}

type BuildingRepository interface {
	Create(ctx context.Context, b Building) (uint64, error)
}
