package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"ticket-app/external"
)

func main() {
	// Open DataBase
	db, err := external.OpenDB()
	if err != nil {
		log.Fatalf("db open failed: %v", err)
	}
	defer db.Close()

	// Start Echo Server
	e := echo.New()
	log.Println("Starting server on :8080")
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
