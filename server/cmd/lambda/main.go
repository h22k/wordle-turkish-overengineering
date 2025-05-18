package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/h22k/wordle-turkish-overengineering/server/config"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/bootstrap"
)

func createGame(ctx context.Context, event json.RawMessage) error {
	log.Println("Worker started")
	cfg := config.LoadConfig()
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
			return err
		}
	}

	// Get a random word for the new game
	res, err := gameService.CreateGame(ctx)
	if err != nil {
		log.Printf("Error getting random word: %v", err)
		return err
	}

	log.Printf("Created new game with ID: %s, lenght: %d", res.ID, res.WordLength)
	return nil
}

func main() {
	lambda.Start(createGame)
}
