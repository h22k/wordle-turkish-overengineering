package game

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/validator"
	response "github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/fiber"
)

type Handler struct {
	gameService *Service

	iv validator.InputValidator
}

func NewHandler(gameService *Service, v validator.InputValidator) *Handler {
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

		return response.Success(c, ActiveGameResponse{
			MaxGuesses:     game.MaxWordGuesses,
			IsGameFinished: game.GuessedCorrectly() || game.GuessExceeded(),
			Guesses:        guessesToResponse(game.WordGuesses),
		})
	}
}

func (h *Handler) MakeGuess() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionId := c.Cookies("session_id", "")

		var req MakeGuessRequest
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

		return response.Created(c, GuessedWordResponse{
			Word:    guess.Guess.String(),
			Letters: lettersToView(guess.Letters),
		})
	}
}
