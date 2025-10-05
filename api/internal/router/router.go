package router

import (
	"net/http"

	"ticket-app/internal/handler"

	"github.com/labstack/echo/v4"
)

type Deps struct {
	SessionHandler *handler.SessionHandler
	VisitorHandler *handler.VisitorHandler
}

func New(d Deps) *echo.Echo {
	e := echo.New()

	e.GET("/healthz", func(c echo.Context) error { return c.NoContent(http.StatusOK) })

	// Session関連のAPI
	e.POST("/sessions", d.SessionHandler.CreateSession)
	e.GET("/sessions/me", d.SessionHandler.Me)

	// Visitor関連のAPI
	e.GET("/visitors", d.VisitorHandler.ListVisitors)
	e.GET("/visitors/:id", d.VisitorHandler.GetVisitor)
	e.POST("/visitors", d.VisitorHandler.CreateVisitor)

	return e
}
