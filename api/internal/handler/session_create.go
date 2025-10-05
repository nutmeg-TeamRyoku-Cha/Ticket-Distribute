package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"ticket-app/internal/usecase"
)

const SessionCookieName = "session_token"

type CreateSessionHandler struct {
	uc usecase.CreateSessionUsecase
}

func NewSessionHandler(uc usecase.CreateSessionUsecase) *CreateSessionHandler {
	return &CreateSessionHandler{uc: uc}
}

type createSessionReq struct {
	VisitorID uint64 `json:"visitor_id"`
}

type createSessionRes struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

func (h *CreateSessionHandler) CreateSession(c echo.Context) error {
	var req createSessionReq
	if err := c.Bind(&req); err != nil || req.VisitorID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	ts, err := h.uc.CreateSession(c.Request().Context(), req.VisitorID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:     SessionCookieName,
		Value:    ts.Token,
		Path:     "/",
		Expires:  ts.LoginSession.ExpiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	})

	return c.JSON(http.StatusCreated, createSessionRes{
		Token:     ts.Token,
		ExpiresAt: ts.LoginSession.ExpiresAt.Format(time.RFC3339),
	})
}
