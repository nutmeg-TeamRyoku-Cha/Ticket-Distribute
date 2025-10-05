package handler

import (
	"net/http"

	"ticket-app/internal/usecase"

	"github.com/labstack/echo/v4"
)

type ResolveSessionHandler struct {
	uc usecase.ResolveSessionUsecase
}

func NewResolveSessionHandler(uc usecase.ResolveSessionUsecase) *ResolveSessionHandler {
	return &ResolveSessionHandler{uc: uc}
}

type resolveSessionRes struct {
	VisitorID uint64 `json:"visitor_id"`
}

func (h *ResolveSessionHandler) ResolveVisitor(c echo.Context) error {
	cookie, err := c.Cookie(SessionCookieName)
	if err != nil || cookie.Value == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing session")
	}
	visitorID, err := h.uc.ResolveVisitor(c.Request().Context(), cookie.Value)
	if err != nil {
		if err == usecase.ErrInvalidOrExpired {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired session")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resolveSessionRes{VisitorID: visitorID})
}
