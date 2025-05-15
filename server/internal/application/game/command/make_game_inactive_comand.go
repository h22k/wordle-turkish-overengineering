package command

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
)

type MakeGameInactiveCommand struct {
	gameRepo domain.GameRepository
}

func NewMakeGameInactiveCommand(gameRepo domain.GameRepository) *MakeGameInactiveCommand {
	return &MakeGameInactiveCommand{
		gameRepo: gameRepo,
	}
}

func (m *MakeGameInactiveCommand) Execute(ctx context.Context, game domain.Game) error {
	return m.gameRepo.MakeGameInactive(ctx, game.ID)
}
