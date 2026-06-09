package app

import "log/slog"

type service struct {
	logger *slog.Logger
}

func CreateService(logger *slog.Logger) *service {
	return &service{
		logger: logger,
	}
}
