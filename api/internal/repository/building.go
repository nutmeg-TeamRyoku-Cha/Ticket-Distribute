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

func (r *BuildingRepository) List(ctx context.Context) ([]domain.Building, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT building_id, building_name, latitude, longitude
		FROM buildings
		ORDER BY building_id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Building
	for rows.Next() {
		var id uint64
		var buildname string
		var la float64
		var lo float64
		if err := rows.Scan(&id, &buildname, &la, &lo); err != nil {
			return nil, err
		}
		out = append(out, domain.Building{BuildingID: id, BuildingName: buildname, Latitude: &la, Longitude: &lo})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *BuildingRepository) GetByID(ctx context.Context, id uint64) (domain.Building, bool, error) {
	var bid uint64
	var buildname string
	var la float64
	var lo float64
	err := r.DB.QueryRowContext(ctx, `
		SELECT building_id, building_name, latitude, longitude
		FROM buildings
		WHERE building_id = ?
		LIMIT 1
	`, id).Scan(&bid, &buildname, &la, &lo)
	if err == sql.ErrNoRows {
		return domain.Building{}, false, nil
	}
	if err != nil {
		return domain.Building{}, false, err
	}
	return domain.Building{BuildingID: bid, BuildingName: buildname, Latitude: &la, Longitude: &lo}, true, nil
}
