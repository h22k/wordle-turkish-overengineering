package bootstrap

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/h22k/wordle-turkish-overengineering/server/internal/application/game/checker"
	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/game"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/presentation/http/fiber/adapter"
)

func initChecker(uc *usecase) *domain.WordCheckerChain {
	return domain.NewWordCheckerChain(
		checker.NewWordLenChecker(),
		checker.NewDatabaseWordWriterChecker(
			checker.NewDatabaseWordChecker(uc.VocableQuery()),
			checker.NewTdkWordChecker(adapter.NewTdkClient(&http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // TODO: remove this in production.
				},
				Timeout: 5 * time.Second,
			})),
			uc.AddWordCommand(),
		),
	)
}
