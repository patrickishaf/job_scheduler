package app

import "log/slog"

type socketHandler struct {
	logger *slog.Logger
	srv    *service
}

func CreateSocketHandler(srv *service, logger *slog.Logger) *socketHandler {
	return &socketHandler{
		logger: logger,
		srv:    srv,
	}
}
