package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"ticket-app/external"
	"ticket-app/internal/handler"
	"ticket-app/internal/repository"
	"ticket-app/internal/router"
	"ticket-app/internal/usecase"
)

func main() {
	// Open DataBase
	db, err := external.OpenDB()
	if err != nil {
		log.Fatalf("db open failed: %v", err)
	}
	defer db.Close()

	//Repository
	sessionRepo := repository.NewLoginSessionRepository(db)
	visitorRepo := repository.NewVisitorRepository(db)
	ticketRepo := repository.NewTicketRepository(db)
	//Usecase
	sessionsUC := usecase.NewSessionUsecase(sessionRepo)
	visitorsUC := usecase.NewVisitorUsecase(visitorRepo)
	ticketsUC := usecase.NewTicketUsecase(ticketRepo)
	//Handler
	sessionsH := handler.NewSessionHandler(sessionsUC)
	visitorsH := handler.NewVisitorHandler(visitorsUC)
	ticketsH := handler.NewTicketHandler(ticketsUC)

	// Start Echo Server
	e := router.New(router.Deps{
		SessionHandler: sessionsH,
		VisitorHandler: visitorsH,
		TicketHandler:  ticketsH,
	})
	go func() {
		log.Println("Starting server on :8080")
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	//終了シグナルで終了
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
