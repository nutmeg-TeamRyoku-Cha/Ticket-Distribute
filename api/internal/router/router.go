package router

import (
	"net/http"

	"ticket-app/internal/handler"

	"github.com/labstack/echo/v4"
)

type Deps struct {
	SessionHandler  *handler.SessionHandler
	VisitorHandler  *handler.VisitorHandler
	TicketHandler   *handler.TicketHandler
	BuildingHandler *handler.BuildingHandler
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
	// Ticket関連のAPI
	e.POST("/tickets", d.TicketHandler.CreateTicket)                        // チケット作成
	e.GET("/tickets/visitor/:id", d.TicketHandler.ListTicketsByVisitorPath) // 特定の訪問者のチケット一覧 idの訪問者のチケット一覧を取得
	e.PATCH("/tickets/:id/status", d.TicketHandler.UpdateTicketStatus)      // チケットstatusのステータス更新 idのチケットのstatusを更新
	e.GET("/tickets/:id", d.TicketHandler.GetTicket)                        // 特定のチケット取得 idのチケットを取得
	e.GET("/tickets", d.TicketHandler.ListTickets)                          // チケット一覧取得
	// Building関連のAPI
	e.POST("/buildings", d.BuildingHandler.CreateBuilding)
	e.GET("/buildings", d.BuildingHandler.ListBuildings)
	e.GET("/buildings/:id", d.BuildingHandler.GetBuilding)
	return e
}
