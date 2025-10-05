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

type TicketRepository interface {
	Create(ctx context.Context, ticket Ticket) (uint64, error)
	ListByVisitor(ctx context.Context, visitorID uint64) ([]Ticket, error)
	GetByID(ctx context.Context, id uint64) (Ticket, bool, error)
	ListAll(ctx context.Context) ([]Ticket, error)
}
