package game

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/event"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/validator"
	response "github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/fiber"
	commonGame "github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/game"
)

type Handler struct {
	gameService *commonGame.Service
	iv          validator.InputValidator
}

func NewHandler(gameService *commonGame.Service, v validator.InputValidator) *Handler {
	return &Handler{
		gameService: gameService,
		iv:          v,
	}
}

func (h *Handler) GetGame() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionId := c.Cookies("session_id", "")

		game, err := h.gameService.GetGameInfo(c.Context(), sessionId)
		if err != nil {
			return response.BadRequest(c, err)
		}

		return response.Success(c, commonGame.ActiveGameResponse{
			MaxGuesses:     game.MaxWordGuesses,
			IsGameFinished: game.IsFinished(),
			Guesses:        commonGame.GuessesToResponse(game.WordGuesses),
		})
	}
}

func (h *Handler) MakeGuess() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionId := c.Cookies("session_id", "")

		var req commonGame.MakeGuessRequest
		if err := c.BodyParser(&req); err != nil {
			return response.BadRequest(c, err)
		}

		if err := h.iv.Validate(req); err != nil {
			return response.BadRequest(c, err)
		}

		guess, err := h.gameService.MakeGuess(c.Context(), sessionId, req.Guess)
		if err != nil {
			return response.BadRequest(c, err)
		}

		return response.Created(c, commonGame.GuessedWordResponse{
			Word:    guess.Guess.String(),
			Letters: commonGame.LettersToView(guess.Letters),
		})
	}
}

func (h *Handler) Sse(broker *event.Broker) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("X-Accel-Buffering", "no")
		c.Set("Transfer-Encoding", "chunked")
		c.Set("Access-Control-Allow-Origin", "*")

		name := c.Params("name")

		ch := broker.Subscribe()

		fmt.Println("SSE connection opened:" + name)
		defer fmt.Println("SSE connection closed:" + name)

		c.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
			keepAliveTicker := time.NewTicker(1 * time.Second)

			for loop := true; loop; {
				select {

				case ev, ok := <-ch:
					if !ok {
						log.Printf("Channel closed\nname: %s\n", name)
						keepAliveTicker.Stop()
						loop = false
						break
					}

					// send sse formatted message
					_, err := fmt.Fprintf(w, "event: game_created\ndata: %s\n\n", ev.Name())

					if err != nil {
						log.Printf("Error while writing Data: %v\n", err)
						continue
					}

					err = w.Flush()
					if err != nil {
						log.Printf("Error while flushing Data: %v\nname:%s", err, name)
						keepAliveTicker.Stop()
						loop = false
						broker.Unsubscribe(ch)
						break
					}
				case <-keepAliveTicker.C:
					_, err := fmt.Fprintf(w, ":keep-alive\n\n")
					if err != nil {
						log.Printf("Error while writing keep-alive: %v\n", err)
						keepAliveTicker.Stop()
						loop = false
						broker.Unsubscribe(ch)
						break
					}

					err = w.Flush()
					if err != nil {
						log.Printf("Error while flushing keep-alive: %v, %s\n", err, name)
						log.Printf("Closing connection\nname: %s\n", name)
						keepAliveTicker.Stop()
						loop = false
						broker.Unsubscribe(ch)
						break
					}
				}
			}
		})

		return nil
	}
}
