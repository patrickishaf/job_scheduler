package common

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func InitLogger() *slog.Logger {
	if logger != nil {
		return logger
	}
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return logger
}

func GetLogger() *slog.Logger {
	if logger != nil {
		return InitLogger()
	}
	return logger
}
