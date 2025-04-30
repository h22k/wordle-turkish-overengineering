package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/checker"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/fiber/adapter"
)

func initChecker(uc *usecase) *domain.WordCheckerChain {
	return domain.NewWordCheckerChain(
		checker.NewWordLenChecker(),
		checker.NewDatabaseWordWriterChecker(
			checker.NewDatabaseWordChecker(uc.VocableQuery()),
			checker.NewTdkWordChecker(adapter.NewTdkClient(fiber.Get)),
			uc.AddWordCommand(),
		),
	)
}
