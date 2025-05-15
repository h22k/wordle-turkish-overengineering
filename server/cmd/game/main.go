package main

import (
	"context"
	"fmt"
	"log"

	"github.com/h22k/wordle-turkish-overengineering/server/config"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/bootstrap"
)

func main() {
	fmt.Println("Worker started")
	cfg := config.LoadConfig()
	ctx := context.Background()
	app := bootstrap.InitApplication(ctx, cfg)
	defer app.Close()

	gameService := app.AppService().GameService()

	lastGame, err := gameService.LastGame(ctx)
	if err != nil {
		log.Printf("Error getting last game: %v", err)
	} else {
		err = gameService.MakeGameInactive(ctx, lastGame)
		if err != nil {
			log.Printf("Error making game inactive: %v", err)
			return
		}
	}

	// Get a random word for the new game
	res, err := gameService.CreateGame(ctx)
	if err != nil {
		log.Printf("Error getting random word: %v", err)
		return
	}

	log.Printf("Created new game with ID: %s, lenght: %d", res.ID, res.WordLength)
}
