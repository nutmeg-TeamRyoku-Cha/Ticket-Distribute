package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"ticket-app/internal/usecase"
)

type SessionHandler struct {
	uc usecase.SessionUsecase
}

func NewSessionHandler(uc usecase.SessionUsecase) *SessionHandler {
	return &SessionHandler{uc: uc}
}

type createSessionReq struct {
	VisitorID uint64 `json:"visitor_id"`
}

type createSessionRes struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

func (h *SessionHandler) CreateSession(c echo.Context) error {
	var req createSessionReq
	if err := c.Bind(&req); err != nil || req.VisitorID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	ts, err := h.uc.CreateSession(c.Request().Context(), req.VisitorID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createSessionRes{
		Token:     ts.Token,
		ExpiresAt: ts.LoginSession.ExpiresAt.Format(time.RFC3339),
	})
}
