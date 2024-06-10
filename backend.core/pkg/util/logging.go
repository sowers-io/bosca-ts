package util

import (
	"bosca.io/pkg/configuration"
	"log/slog"
	"os"
)

func InitializeLogging(cfg configuration.Configuration) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.SetDefault(logger)
}
