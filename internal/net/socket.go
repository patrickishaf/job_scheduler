package net

import (
	"github.com/gorilla/websocket"
)

type SocketEvent string

const (
	SOCKET_EVENT_ACK        SocketEvent = "ack"
	SOCKET_EVENT_CONNECT    SocketEvent = "connect"
	SOCKET_EVENT_DISCONNECT SocketEvent = "disconnect"
	SOCKET_EVENT_ERROR      SocketEvent = "error"
)

type SocketMessage struct {
	Event   SocketEvent `json:"event"`
	Payload any         `json:"data"`
}

func SendSocketMessage(conn *websocket.Conn, msg *SocketMessage) error {
	err := conn.WriteJSON(msg)
	return err
}
