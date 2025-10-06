package domain

import (
	"context"
	"time"
)

type Ticket struct {
	TicketID       uint64
	VisitorID      uint64
	ProjectID      uint64
	Status         string
	EntryStartTime *time.Time
	EntryEndTime   *time.Time
}

type BuildingBrief struct {
	BuildingID   uint64
	BuildingName string
	Latitude     *float64
	Longitude    *float64
}

type ProjectBrief struct {
	ProjectID      uint64
	ProjectName    string
	RequiresTicket bool
	StartTime      time.Time
	EndTime        *time.Time
	Building       BuildingBrief
}

type TicketWithProject struct {
	TicketID       uint64
	VisitorID      uint64
	Status         string
	EntryStartTime *time.Time
	EntryEndTime   *time.Time
	Project        ProjectBrief
}

type TicketRepository interface {
	Create(ctx context.Context, ticket Ticket) (uint64, error)
	ListByVisitor(ctx context.Context, visitorID uint64) ([]Ticket, error)
	ListByVisitorWithProject(ctx context.Context, visitorID uint64) ([]TicketWithProject, error)
	GetByID(ctx context.Context, id uint64) (Ticket, bool, error)
	ListAll(ctx context.Context) ([]Ticket, error)
	UpdateStatus(ctx context.Context, id uint64, status string) error
}
