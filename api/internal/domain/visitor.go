package domain

import (
	"context"
	"time"
)

type Visitor struct {
	VisitorID uint64
	Nickname  string
	BirthDate time.Time
	PartySize int
}

type VisitorRepository interface {
	Create(ctx context.Context, v Visitor) (uint64, error)
	List(ctx context.Context) ([]Visitor, error)
}
