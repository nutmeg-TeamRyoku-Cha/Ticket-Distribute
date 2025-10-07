package domain

import (
	"context"
	"time"
)

// Project defines the structure for a project.
type Project struct {
	ProjectID        uint64
	ProjectName      string
	BuildingID       uint64
	RequiresTicket   bool
	RemainingTickets uint
	StartTime        time.Time
	EndTime          time.Time
}

// ProjectRepository defines the interface for project database operations.
type ProjectRepository interface {
	Create(ctx context.Context, project Project) (uint64, error)
	ListAll(ctx context.Context) ([]Project, error)
	GetByID(ctx context.Context, id uint64) (Project, bool, error)
	UpdateRemainingTickets(ctx context.Context, id uint64, remainingTickets uint) error
	ListResolved(ctx context.Context) ([]ProjectBrief, error)
}
