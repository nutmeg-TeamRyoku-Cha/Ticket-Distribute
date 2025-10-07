package usecase

import (
	"context"
	"ticket-app/internal/domain"
	"time"
)

// ProjectUsecase defines the interface for project business logic.
type ProjectUsecase interface {
	CreateProject(ctx context.Context, p domain.Project) (uint64, error)
	ListAllProjects(ctx context.Context) ([]domain.Project, error)
	GetProjectByID(ctx context.Context, id uint64) (domain.Project, bool, error)
	UpdateRemainingTickets(ctx context.Context, id uint64, remainingTickets uint) error
	ListProjectsResolved(ctx context.Context) ([]domain.ProjectBrief, error)
}

type projectUsecase struct {
	repo domain.ProjectRepository
	now  func() time.Time
}

// NewProjectUsecase creates a new ProjectUsecase.
func NewProjectUsecase(repo domain.ProjectRepository) ProjectUsecase {
	return &projectUsecase{repo: repo, now: time.Now}
}

func (u *projectUsecase) GetProjectByID(ctx context.Context, id uint64) (domain.Project, bool, error) {
	return u.repo.GetByID(ctx, id)
}

// CreateProject handles the logic for creating a new project.
func (u *projectUsecase) CreateProject(ctx context.Context, p domain.Project) (uint64, error) {
	// Business logic, such as validation, can be added here.
	return u.repo.Create(ctx, p)
}

// ListAllProjects handles the logic for listing all projects.
func (u *projectUsecase) ListAllProjects(ctx context.Context) ([]domain.Project, error) {
	return u.repo.ListAll(ctx)
}

// UpdateRemainingTickets handles the logic for updating remaining tickets.
func (u *projectUsecase) UpdateRemainingTickets(ctx context.Context, id uint64, remainingTickets uint) error {
	return u.repo.UpdateRemainingTickets(ctx, id, remainingTickets)
}

func (u *projectUsecase) ListProjectsResolved(ctx context.Context) ([]domain.ProjectBrief, error) {
	return u.repo.ListResolved(ctx)
}
