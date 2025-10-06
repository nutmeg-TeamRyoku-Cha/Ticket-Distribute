package repository

import (
	"context"
	"database/sql"
	"ticket-app/internal/domain"
	"time"
)

// ProjectRepository implements the domain.ProjectRepository interface.
type ProjectRepository struct{ DB *sql.DB }

// NewProjectRepository creates a new ProjectRepository.
func NewProjectRepository(db *sql.DB) *ProjectRepository { return &ProjectRepository{DB: db} }

// GetByID retrieves a single project from the database by its ID.
func (r *ProjectRepository) GetByID(ctx context.Context, id uint64) (domain.Project, bool, error) {
	var p domain.Project
	var st, et time.Time
	err := r.DB.QueryRowContext(ctx, `
        SELECT project_id, project_name, building_id, requires_ticket, remaining_tickets, start_time, end_time
        FROM projects
        WHERE project_id = ?
    `, id).Scan(&p.ProjectID, &p.ProjectName, &p.BuildingID, &p.RequiresTicket, &p.RemainingTickets, &st, &et)

	if err == sql.ErrNoRows {
		return domain.Project{}, false, nil
	}
	if err != nil {
		return domain.Project{}, false, err
	}
	p.StartTime = st
	p.EndTime = et
	return p, true, nil
}

// Create inserts a new project into the database.
func (r *ProjectRepository) Create(ctx context.Context, p domain.Project) (uint64, error) {
	res, err := r.DB.ExecContext(ctx, `
        INSERT INTO projects (project_name, building_id, requires_ticket, remaining_tickets, start_time, end_time)
        VALUES (?, ?, ?, ?, ?, ?)
    `, p.ProjectName, p.BuildingID, p.RequiresTicket, p.RemainingTickets, p.StartTime, p.EndTime)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// ListAll retrieves all projects from the database.
func (r *ProjectRepository) ListAll(ctx context.Context) ([]domain.Project, error) {
	rows, err := r.DB.QueryContext(ctx, `
        SELECT project_id, project_name, building_id, requires_ticket, remaining_tickets, start_time, end_time
        FROM projects
        ORDER BY project_id ASC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Project
	for rows.Next() {
		var p domain.Project
		var st, et time.Time
		if err := rows.Scan(&p.ProjectID, &p.ProjectName, &p.BuildingID, &p.RequiresTicket, &p.RemainingTickets, &st, &et); err != nil {
			return nil, err
		}
		p.StartTime = st
		p.EndTime = et
		out = append(out, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateRemainingTickets updates the number of remaining tickets for a specific project.
func (r *ProjectRepository) UpdateRemainingTickets(ctx context.Context, id uint64, remainingTickets uint) error {
	_, err := r.DB.ExecContext(ctx, `
        UPDATE projects SET remaining_tickets = ? WHERE project_id = ?
    `, remainingTickets, id)
	return err
}
