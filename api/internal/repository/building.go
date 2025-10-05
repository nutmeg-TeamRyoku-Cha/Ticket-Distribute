package repository

import (
	"context"
	"database/sql"

	"ticket-app/internal/domain"
)

type BuildingRepository struct{ DB *sql.DB }

func NewBuildingRepository(db *sql.DB) *BuildingRepository {
	return &BuildingRepository{DB: db}
}

func (r *BuildingRepository) Create(ctx context.Context, b domain.Building) (uint64, error) {
	var lat any
	var lon any
	if b.Latitude != nil {
		lat = *b.Latitude // 実値
	} else {
		lat = nil // NULL
	}
	if b.Longitude != nil {
		lon = *b.Longitude
	} else {
		lon = nil
	}

	res, err := r.DB.ExecContext(ctx, `
		INSERT INTO buildings (building_name, latitude, longitude)
		VALUES (?, ?, ?)
	`, b.BuildingName, lat, lon)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
