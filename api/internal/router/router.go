package router

import (
	"net/http"

	"ticket-app/internal/handler"

	"github.com/labstack/echo/v4"
)

type Deps struct {
	SessionHandler *handler.SessionHandler
}

func New(d Deps) *echo.Echo {
	e := echo.New()

	e.GET("/healthz", func(c echo.Context) error { return c.NoContent(http.StatusOK) })
	e.POST("/sessions", d.SessionHandler.CreateSession)

	return e
}
