package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/duohedron/orders/internal/api"
	"github.com/duohedron/orders/internal/config"
	"github.com/duohedron/orders/internal/orders"
)

func main() {
	conf := config.Load()

	writer := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: zerolog.TimeFieldFormat}
	logger := zerolog.New(writer).With().Timestamp().Logger()
	logLevel, err := zerolog.ParseLevel(conf.LogLevel)
	if err != nil {
		logger.Panic().Err(err).Str("level", conf.LogLevel).Msg("Invalid log level")
	}
	zerolog.SetGlobalLevel(logLevel)
	log.Logger = logger

	e := echo.New()

	store, err := orders.NewStore(conf.DatabaseURL)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	events := make(chan orders.Event, 10)
	svc := orders.NewService(store, events)
	go orders.StartWorker(events)

	api.RegisterRoutes(e, svc)

	go func() {
		if err := e.Start(conf.Address); err != nil {
			log.Info().Err(err).Msg("shutdown")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	e.Shutdown(ctx)
}
