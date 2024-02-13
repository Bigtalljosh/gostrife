package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Bigtalljosh/gostrife/internal/services"
	"github.com/Bigtalljosh/gostrife/internal/web/customers"
	"github.com/Bigtalljosh/gostrife/internal/web/landing"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
		
	SetupLogging()
	LoadConfig()

	r := mux.NewRouter()

	customerService := &services.InMemoryCustomerService{}

	customerController := customers.NewCustomerController(customerService)
	landingController := landing.NewLandingController()
	
	customerController.RegisterRoutes(r)
	landingController.RegisterRoutes(r)

	server := &http.Server{
		Addr: viper.GetString("server.port"),
		Handler: r,
	}

	go func() {
		log.Info().
		Str("Server running on port", viper.GetString("server.port"))
        
        if err := server.ListenAndServe(); err != nil {
            log.Fatal().Err(err)
        }
    }()

    // Gracefully shutdown the server when receiving an interrupt signal
    // This is currently GPT magic for me I need to understand this better
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c
    log.Info().Msg("Shutting down server...")	

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Error().Err(err)
    }

    log.Info().Msg("Server gracefully stopped")
}

func SetupLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Add file and line number to log
	log.Logger = log.With().Caller().Logger()

    log.Print("logger set up")
}

func LoadConfig(){
	viper.SetConfigFile("./config.json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatal().Err(err).Msg("Config file not found")
		} else {
			// Config file was found but another error was produced
			log.Fatal().Err(err).Msg("Config found, error parsing")
		}
	}
	// Config file found and successfully parsed
}
