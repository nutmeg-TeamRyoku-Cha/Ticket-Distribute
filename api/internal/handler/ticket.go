package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"ticket-app/internal/domain"
	"ticket-app/internal/usecase"
)

type TicketHandler struct {
	uc usecase.TicketUsecase
}

func NewTicketHandler(uc usecase.TicketUsecase) *TicketHandler { return &TicketHandler{uc: uc} }

type createTicketReq struct {
	VisitorID uint64 `json:"visitor_id"`
	ProjectID uint64 `json:"project_id"`
}

type createTicketRes struct {
	TicketID uint64 `json:"ticket_id"`
}

type ticketRes struct {
	TicketID       uint64  `json:"ticket_id"`
	VisitorID      uint64  `json:"visitor_id"`
	ProjectID      uint64  `json:"project_id"`
	Status         string  `json:"status"`
	EntryStartTime *string `json:"entry_start_time"`
	EntryEndTime   *string `json:"entry_end_time"`
}

func (h *TicketHandler) CreateTicket(c echo.Context) error {
	var req createTicketReq
	if err := c.Bind(&req); err != nil || req.VisitorID == 0 || req.ProjectID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	t := domain.Ticket{VisitorID: req.VisitorID, ProjectID: req.ProjectID, Status: "issued"}
	id, err := h.uc.CreateTicket(c.Request().Context(), t)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createTicketRes{TicketID: id})
}

func ticketsToRes(ts []domain.Ticket) []ticketRes {
	out := make([]ticketRes, 0, len(ts))
	for _, t := range ts {
		var est, eet *string
		if t.EntryStartTime != nil {
			s := t.EntryStartTime.Format(time.RFC3339)
			est = &s
		}
		if t.EntryEndTime != nil {
			s := t.EntryEndTime.Format(time.RFC3339)
			eet = &s
		}
		out = append(out, ticketRes{
			TicketID:       t.TicketID,
			VisitorID:      t.VisitorID,
			ProjectID:      t.ProjectID,
			Status:         t.Status,
			EntryStartTime: est,
			EntryEndTime:   eet,
		})
	}
	return out
}

// ListTickets handles GET /tickets?visitor_id=1
func (h *TicketHandler) ListTickets(c echo.Context) error {
	ts, err := h.uc.ListAllTickets(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ticketsToRes(ts))
}

// ListTicketsByVisitorPath handles GET /tickets/visitor/:id
func (h *TicketHandler) ListTicketsByVisitorPath(c echo.Context) error {
	idStr := c.Param("id")
	vid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || vid == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid visitor_id")
	}
	ts, err := h.uc.ListTicketsByVisitor(c.Request().Context(), vid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ticketsToRes(ts))
}

func (h *TicketHandler) GetTicket(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	t, ok, err := h.uc.GetTicketByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "not found")
	}
	res := ticketsToRes([]domain.Ticket{t})
	if len(res) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, res[0])
}
