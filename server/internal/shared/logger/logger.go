package logger

import (
	"log/slog"
	"os"

	"github.com/h22k/wordle-turkish-overengineering/server/config"
)

var Logger *slog.Logger

func InitLogger(cfg config.Config) {
	var handler slog.Handler

	if cfg.IsProd() {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}
