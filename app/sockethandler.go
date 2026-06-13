package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/net"
)

type socketHandler struct {
	logger     *slog.Logger
	srv        *service
	wsUpgrader *websocket.Upgrader
}

func CreateSocketHandler(cfg *config.SocketConfig, srv *service, logger *slog.Logger) *socketHandler {
	childLogger := logger.With("handler", "socketHandler")
	return &socketHandler{
		logger: childLogger,
		srv:    srv,
		wsUpgrader: &websocket.Upgrader{
			ReadBufferSize:  cfg.ReadBufferSize,
			WriteBufferSize: cfg.WriteBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (this *socketHandler) handleConnection(c *gin.Context) {
	logger := this.logger.With("caller", "socketHandler.handleConnection")

	logger.Info("a socket connection has been established")
	conn, err := this.wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "failed to connect to websocket")
		return
	}

	this.srv.StoreSocketConn(conn)

	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("error is: %s\n", err.Error())
			if err.Error() == "websocket: close 1001 (going away)" {
				logger.Info("repeated read on socket connection", "err", err.Error())
				break
			}

			closeError, isCloseError := err.(*websocket.CloseError)
			if isCloseError && closeError.Code == 1005 {
				logger.Info("socket closed. terminating listener")
				break
			}

			logger.Error("socket event error", "err", err.Error())
			break
		}

		net.SendSocketMessage(conn, &net.SocketMessage{
			Event:   net.SOCKET_EVENT_ACK,
			Payload: string(msgBytes),
		})
	}
}
