package app

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	logger *slog.Logger
	srv    *service
}

func CreateHTTPHandler(srv *service, logger *slog.Logger) *httpHandler {
	return &httpHandler{
		logger: logger,
		srv:    srv,
	}
}

func (this *httpHandler) GetInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "interesting")
}
