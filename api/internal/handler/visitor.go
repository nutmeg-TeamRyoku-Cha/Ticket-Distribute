package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"ticket-app/internal/domain"
	"ticket-app/internal/usecase"
)

type VisitorHandler struct {
	uc usecase.VisitorUsecase
}

func NewVisitorHandler(uc usecase.VisitorUsecase) *VisitorHandler { return &VisitorHandler{uc: uc} }

type createVisitorReq struct {
	Nickname  string `json:"nickname"`
	BirthDate string `json:"birth_date"` // YYYY-MM-DD
	PartySize int    `json:"party_size"`
}

type createVisitorRes struct {
	VisitorID uint64 `json:"visitor_id"`
}

type listVisitorRes struct {
	VisitorID uint64 `json:"visitor_id"`
	Nickname  string `json:"nickname"`
	BirthDate string `json:"birth_date"`
	PartySize int    `json:"party_size"`
}

func (h *VisitorHandler) CreateVisitor(c echo.Context) error {
	var req createVisitorReq
	if err := c.Bind(&req); err != nil || req.Nickname == "" || req.BirthDate == "" || req.PartySize < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	bd, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid birth_date format")
	}
	v := domain.Visitor{Nickname: req.Nickname, BirthDate: bd, PartySize: req.PartySize}
	id, err := h.uc.CreateVisitor(c.Request().Context(), v)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createVisitorRes{VisitorID: id})
}

func (h *VisitorHandler) ListVisitors(c echo.Context) error {
	vs, err := h.uc.ListVisitors(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	out := make([]listVisitorRes, 0, len(vs))
	for _, v := range vs {
		out = append(out, listVisitorRes{
			VisitorID: v.VisitorID,
			Nickname:  v.Nickname,
			BirthDate: v.BirthDate.Format("2006-01-02"),
			PartySize: v.PartySize,
		})
	}
	return c.JSON(http.StatusOK, out)
}
