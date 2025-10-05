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

func NewSessionHandler(uc usecase.SessionUsecase) *SessionHandler { return &SessionHandler{uc: uc} }

type createSessionReq struct {
	VisitorID uint64 `json:"visitor_id"`
}

type createSessionRes struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

type resolveVisitorRes struct {
	VisitorID uint64    `json:"visitor_id"`
	Nickname  string    `json:"nickname"`
	BirthDate time.Time `json:"birth_date"`
	PartySize int       `json:"party_size"`
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

	c.SetCookie(&http.Cookie{
		Name:     "session_token",
		Value:    ts.Token,
		Path:     "/",
		Expires:  ts.Session.ExpiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})

	return c.JSON(http.StatusCreated, createSessionRes{
		Token:     ts.Token,
		ExpiresAt: ts.Session.ExpiresAt.Format(time.RFC3339),
	})
}

func (h *SessionHandler) Me(c echo.Context) error {
	cookie, err := c.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing session")
	}
	v, err := h.uc.ResolveVisitorProfile(c.Request().Context(), cookie.Value)
	if err != nil {
		if err == usecase.ErrInvalidOrExpired {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired session")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resolveVisitorRes{
		VisitorID: v.VisitorID,
		Nickname:  v.Nickname,
		BirthDate: v.BirthDate,
		PartySize: v.PartySize,
	})
}
