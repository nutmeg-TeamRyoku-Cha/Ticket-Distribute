package handler

import (
	"fmt"
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
	VisitorID      uint64  `json:"visitor_id"`
	ProjectID      uint64  `json:"project_id"`
	EntryStartTime *string `json:"entry_start_time,omitempty"`
	EntryEndTime   *string `json:"entry_end_time,omitempty"`
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

type buildingRes struct {
	BuildingID   uint64 `json:"building_id"`
	BuildingName string `json:"building_name"`
	Latitude     string `json:"latitude,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
}

type projectRes struct {
	ProjectID      uint64      `json:"project_id"`
	ProjectName    string      `json:"project_name"`
	Building       buildingRes `json:"building"`
	RequiresTicket bool        `json:"requires_ticket"` // 必要なら "resuires_ticket"
	StartTime      string      `json:"start_time"`
	EndTime        string      `json:"end_time,omitempty"`
}

type ticketWithProjectRes struct {
	TicketID       uint64     `json:"ticket_id"`
	VisitorID      uint64     `json:"visitor_id"`
	Project        projectRes `json:"project"`
	Status         string     `json:"status"`
	EntryStartTime *string    `json:"entry_start_time"`
	EntryEndTime   *string    `json:"entry_end_time"`
}

type updateTicketStatusReq struct {
	Status string `json:"status"`
}

func toRFC3339Ptr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.UTC().Format(time.RFC3339)
	return &s
}

func floatToStrPtr(v *float64) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%.3f", *v)
}

func (h *TicketHandler) CreateTicket(c echo.Context) error {
	var req createTicketReq
	if err := c.Bind(&req); err != nil || req.VisitorID == 0 || req.ProjectID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	// 時間文字列(RFC3339形式)を time.Time 型に変換する
	var est, eet *time.Time
	if req.EntryStartTime != nil {
		t, err := time.Parse(time.RFC3339, *req.EntryStartTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid entry_start_time format")
		}
		est = &t
	}
	if req.EntryEndTime != nil {
		t, err := time.Parse(time.RFC3339, *req.EntryEndTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid entry_end_time format")
		}
		eet = &t
	}

	t := domain.Ticket{
		VisitorID:      req.VisitorID,
		ProjectID:      req.ProjectID,
		Status:         "issued",
		EntryStartTime: est,
		EntryEndTime:   eet,
	}
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

func ticketJoinedToRes(in []domain.TicketWithProject) []ticketWithProjectRes {
	out := make([]ticketWithProjectRes, 0, len(in))
	for _, r := range in {
		br := buildingRes{
			BuildingID:   r.Project.Building.BuildingID,
			BuildingName: r.Project.Building.BuildingName,
		}
		if s := floatToStrPtr(r.Project.Building.Latitude); s != "" {
			br.Latitude = s
		}
		if s := floatToStrPtr(r.Project.Building.Longitude); s != "" {
			br.Longitude = s
		}

		var endStr string
		if r.Project.EndTime != nil {
			endStr = r.Project.EndTime.UTC().Format(time.RFC3339)
		}

		pr := projectRes{
			ProjectID:      r.Project.ProjectID,
			ProjectName:    r.Project.ProjectName,
			Building:       br,
			RequiresTicket: r.Project.RequiresTicket,
			StartTime:      r.Project.StartTime.UTC().Format(time.RFC3339),
			EndTime:        endStr,
		}

		out = append(out, ticketWithProjectRes{
			TicketID:       r.TicketID,
			VisitorID:      r.VisitorID,
			Project:        pr,
			Status:         r.Status,
			EntryStartTime: toRFC3339Ptr(r.EntryStartTime),
			EntryEndTime:   toRFC3339Ptr(r.EntryEndTime),
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
	recs, err := h.uc.ListTicketsByVisitorWithProject(c.Request().Context(), vid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ticketJoinedToRes(recs))
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

func (h *TicketHandler) UpdateTicketStatus(c echo.Context) error {
	// URLからチケットIDを取得
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ticket id")
	}

	// リクエストボディから新しいステータスを取得
	var req updateTicketStatusReq
	if err := c.Bind(&req); err != nil || req.Status == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request: status is required")
	}

	// Usecaseを呼び出してステータスを更新
	if err := h.uc.UpdateTicketStatus(c.Request().Context(), id, req.Status); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
