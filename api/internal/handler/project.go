package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"ticket-app/internal/domain"
	"ticket-app/internal/usecase"
	"time"

	"github.com/labstack/echo/v4"
)

// ProjectHandler handles HTTP requests for projects.
type ProjectHandler struct {
	uc usecase.ProjectUsecase
}

// NewProjectHandler creates a new ProjectHandler.
func NewProjectHandler(uc usecase.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{uc: uc}
}

// --- Request and Response Structs ---

type createProjectReq struct {
	ProjectName      string    `json:"project_name"`
	BuildingID       uint64    `json:"building_id"`
	RequiresTicket   bool      `json:"requires_ticket"`
	RemainingTickets uint      `json:"remaining_tickets"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
}

type createProjectRes struct {
	ProjectID uint64 `json:"project_id"`
}

type projectRes struct {
	ProjectID        uint64    `json:"project_id"`
	ProjectName      string    `json:"project_name"`
	BuildingID       uint64    `json:"building_id"`
	RequiresTicket   bool      `json:"requires_ticket"`
	RemainingTickets uint      `json:"remaining_tickets"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
}

type updateRemainingTicketsReq struct {
	RemainingTickets uint `json:"remaining_tickets"`
}

type buildingResolvedRes struct {
	BuildingID   uint64 `json:"building_id"`
	BuildingName string `json:"building_name"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
}

type projectResolvedRes struct {
	ProjectID      uint64              `json:"project_id"`
	ProjectName    string              `json:"project_name"`
	Building       buildingResolvedRes `json:"building"`
	RequiresTicket bool                `json:"requires_ticket"`
	StartTime      string              `json:"start_time"`
	EndTime        string              `json:"end_time,omitempty"`
}

func floatToStr(v *float64) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%.6f", *v)
}

// CreateProject handles POST /projects
func (h *ProjectHandler) CreateProject(c echo.Context) error {
	var req createProjectReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request: "+err.Error())
	}

	// Basic validation
	if req.ProjectName == "" || req.BuildingID == 0 || req.StartTime.IsZero() || req.EndTime.IsZero() || req.EndTime.Before(req.StartTime) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request: missing required fields or invalid times")
	}

	p := domain.Project{
		ProjectName:      req.ProjectName,
		BuildingID:       req.BuildingID,
		RequiresTicket:   req.RequiresTicket,
		RemainingTickets: req.RemainingTickets,
		StartTime:        req.StartTime,
		EndTime:          req.EndTime,
	}

	id, err := h.uc.CreateProject(c.Request().Context(), p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createProjectRes{ProjectID: id})
}

// GetProject handles GET /projects/:id
func (h *ProjectHandler) GetProject(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid project id")
	}

	p, ok, err := h.uc.GetProjectByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "project not found")
	}

	res := projectToRes(p)
	return c.JSON(http.StatusOK, res)
}

// ListProjects handles GET /projects
func (h *ProjectHandler) ListProjects(c echo.Context) error {
	ps, err := h.uc.ListAllProjects(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, projectsToRes(ps))
}

// UpdateRemainingTickets handles PATCH /projects/:id/remaining_tickets
func (h *ProjectHandler) UpdateRemainingTickets(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid project id")
	}

	var req updateRemainingTicketsReq
	if err := c.Bind(&req); err != nil {
		// Note: The default value for uint is 0, so a check for negative values isn't necessary.
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request: "+err.Error())
	}

	if err := h.uc.UpdateRemainingTickets(c.Request().Context(), id, req.RemainingTickets); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// projectsToRes converts a slice of domain.Project to a slice of projectRes.
func projectsToRes(ps []domain.Project) []projectRes {
	out := make([]projectRes, 0, len(ps))
	for _, p := range ps {
		out = append(out, projectRes{
			ProjectID:        p.ProjectID,
			ProjectName:      p.ProjectName,
			BuildingID:       p.BuildingID,
			RequiresTicket:   p.RequiresTicket,
			RemainingTickets: p.RemainingTickets,
			StartTime:        p.StartTime,
			EndTime:          p.EndTime,
		})
	}
	return out
}

func projectToRes(p domain.Project) projectRes {
	return projectRes{
		ProjectID:        p.ProjectID,
		ProjectName:      p.ProjectName,
		BuildingID:       p.BuildingID,
		RequiresTicket:   p.RequiresTicket,
		RemainingTickets: p.RemainingTickets,
		StartTime:        p.StartTime,
		EndTime:          p.EndTime,
	}
}

func toProjectResolvedRes(in []domain.ProjectBrief) []projectResolvedRes {
	out := make([]projectResolvedRes, 0, len(in))
	for _, p := range in {
		br := buildingResolvedRes{
			BuildingID:   p.Building.BuildingID,
			BuildingName: p.Building.BuildingName,
		}
		if s := floatToStr(p.Building.Latitude); s != "" {
			br.Latitude = s
		}
		if s := floatToStr(p.Building.Longitude); s != "" {
			br.Longitude = s
		}
		var endStr string
		if p.EndTime != nil {
			endStr = p.EndTime.UTC().Format(time.RFC3339)
		}
		out = append(out, projectResolvedRes{
			ProjectID:      p.ProjectID,
			ProjectName:    p.ProjectName,
			Building:       br,
			RequiresTicket: p.RequiresTicket,
			StartTime:      p.StartTime.UTC().Format(time.RFC3339),
			EndTime:        endStr,
		})
	}
	return out
}

func (h *ProjectHandler) ListProjectsResolved(c echo.Context) error {
	recs, err := h.uc.ListProjectsResolved(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, toProjectResolvedRes(recs))
}
