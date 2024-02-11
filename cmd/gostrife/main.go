package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Bigtalljosh/gostrife/internal/services"
	// TODO: Ask Antony the correct way to do modules, because this can't be it
	customers "github.com/Bigtalljosh/gostrife/internal/web/customers"
	landing "github.com/Bigtalljosh/gostrife/internal/web/landing"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	customerService := &services.InMemoryCustomerService{}

	customerController := customers.NewCustomerController(customerService)
	landingController := landing.NewLandingController()
	
	customerController.RegisterRoutes(r)
	landingController.RegisterRoutes(r)

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	go func() {
        fmt.Println("Server running on port 8080")

        if err := server.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

	// Gracefully shutdown the server when receiving an interrupt signal
    // This is currently GPT magic for me I need to understand this better
	c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    fmt.Println("\nShutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Server gracefully stopped")
}