package usecase

import (
	"context"
	"ticket-app/internal/domain"
	"time"
)

type TicketUsecase interface {
	CreateTicket(ctx context.Context, t domain.Ticket) (uint64, error)
	ListTicketsByVisitor(ctx context.Context, visitorID uint64) ([]domain.Ticket, error)
	ListTicketsByVisitorWithProject(ctx context.Context, visitorID uint64) ([]domain.TicketWithProject, error)
	GetTicketByID(ctx context.Context, id uint64) (domain.Ticket, bool, error)
	ListAllTickets(ctx context.Context) ([]domain.Ticket, error)
	UpdateTicketStatus(ctx context.Context, id uint64, status string) error
}

type ticketUsecase struct {
	repo domain.TicketRepository
	now  func() time.Time
}

func NewTicketUsecase(repo domain.TicketRepository) TicketUsecase {
	return &ticketUsecase{repo: repo, now: time.Now}
}

func (u *ticketUsecase) CreateTicket(ctx context.Context, t domain.Ticket) (uint64, error) {
	return u.repo.Create(ctx, t)
}

func (u *ticketUsecase) GetTicketByID(ctx context.Context, id uint64) (domain.Ticket, bool, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *ticketUsecase) ListAllTickets(ctx context.Context) ([]domain.Ticket, error) {
	return u.repo.ListAll(ctx)
}

func (u *ticketUsecase) ListTicketsByVisitor(ctx context.Context, visitorID uint64) ([]domain.Ticket, error) {
	return u.repo.ListByVisitor(ctx, visitorID)
}

func (u *ticketUsecase) ListTicketsByVisitorWithProject(ctx context.Context, visitorID uint64) ([]domain.TicketWithProject, error) {
	return u.repo.ListByVisitorWithProject(ctx, visitorID)
}

func (u *ticketUsecase) UpdateTicketStatus(ctx context.Context, id uint64, status string) error {
	return u.repo.UpdateStatus(ctx, id, status)
}
