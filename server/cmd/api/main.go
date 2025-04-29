package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/h22k/wordle-turkish-overengineering/server/config"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/bootstrap"
)

func init() {
	// Set the timezone to Turkey
	loc, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		panic("Failed to load timezone: " + err.Error())
	}

	time.Local = loc
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.LoadConfig()
	app := bootstrap.InitApplication(ctx, cfg)

	app.SetRoutes()

	go func() {
		if err := app.Run(); err != nil {
			panic(err)
		}
	}()

	<-sigs

	app.Close()
	cancel()
}
