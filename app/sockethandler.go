package app

import (
	"fmt"
	"log"
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
	conn, err := this.wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, "failed to connect to websocket")
	}

	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			closeError, isCloseError := err.(*websocket.CloseError)
			if isCloseError && closeError.Code == 1005 {
				log.Printf("socket closed. terminating listener")
				break
			}
			fmt.Printf("error. %s\n", err.Error())
			continue
		}

		net.SendSocketMessage(conn, &net.SocketMessage{
			Event:   net.SOCKET_EVENT_ACK,
			Payload: string(msgBytes),
		})
	}
}
