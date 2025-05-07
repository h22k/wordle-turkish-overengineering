package game

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/event"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/validator"
	response "github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/echo"
	commonGame "github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/game"
	"github.com/labstack/echo/v4"
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

func (h *Handler) GetGame() echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionId, _ := c.Cookie("session_id")
		if sessionId == nil {
			sessionId = &http.Cookie{Value: ""}
		}

		game, err := h.gameService.GetGameInfo(c.Request().Context(), sessionId.Value)
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

func (h *Handler) MakeGuess() echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionId, _ := c.Cookie("session_id")
		if sessionId == nil {
			sessionId = &http.Cookie{Value: ""}
		}

		var req commonGame.MakeGuessRequest
		if err := c.Bind(&req); err != nil {
			return response.BadRequest(c, err)
		}

		if err := h.iv.Validate(req); err != nil {
			return response.BadRequest(c, err)
		}

		guess, err := h.gameService.MakeGuess(c.Request().Context(), sessionId.Value, req.Guess)
		if err != nil {
			return response.BadRequest(c, err)
		}

		return response.Created(c, commonGame.GuessedWordResponse{
			Word:    guess.Guess.String(),
			Letters: commonGame.LettersToView(guess.Letters),
		})
	}
}

func (h *Handler) Sse(broker *event.Broker) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().Header().Set("X-Accel-Buffering", "no")
		c.Response().Header().Set("Transfer-Encoding", "chunked")
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		name := c.Param("name")
		ch := broker.Subscribe()

		fmt.Println("SSE connection opened:" + name)
		defer fmt.Println("SSE connection closed:" + name)

		c.Response().Writer.WriteHeader(http.StatusOK)

		keepAliveTicker := time.NewTicker(1 * time.Second)
		defer keepAliveTicker.Stop()

		for {
			select {
			case <-c.Request().Context().Done():
				broker.Unsubscribe(ch)
				return nil

			case ev, ok := <-ch:
				if !ok {
					log.Printf("Channel closed\nname: %s\n", name)
					return nil
				}

				if _, err := fmt.Fprintf(c.Response().Writer, "event: game_created\ndata: %s\n\n", ev.Name()); err != nil {
					log.Printf("Error while writing Data: %v\n", err)
					return err
				}
				c.Response().Flush()

			case <-keepAliveTicker.C:
				if _, err := fmt.Fprintf(c.Response().Writer, ":keep-alive\n\n"); err != nil {
					log.Printf("Error while writing keep-alive: %v\n", err)
					return err
				}
				c.Response().Flush()
			}
		}
	}
}
