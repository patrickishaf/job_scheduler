package net

import (
	"github.com/gorilla/websocket"
)

type SocketEvent string

const (
	SOCKET_EVENT_ACK           SocketEvent = "ack"
	SOCKET_EVENT_CONNECT       SocketEvent = "connect"
	SOCKET_EVENT_DISCONNECT    SocketEvent = "disconnect"
	SOCKET_EVENT_DLQ_UPDATED   SocketEvent = "dlq_updated"
	SOCKET_EVENT_ERROR         SocketEvent = "error"
	SOCKET_EVENT_JOB_CREATED   SocketEvent = "job_created"
	SOCKET_EVENT_JOB_QUEUED    SocketEvent = "job_queued"
	SOCKET_EVENT_JOB_RUNNING   SocketEvent = "job_running"
	SOCKET_EVENT_JOB_COMPLETED SocketEvent = "job_completed"
	SOCKET_EVENT_JOB_FAILED    SocketEvent = "job_failed"
	SOCKET_EVENT_JOB_SCHEDULED SocketEvent = "job_scheduled"
)

type SocketMessage struct {
	Event   SocketEvent `json:"event"`
	Payload any         `json:"data"`
}

func BroadcastSocketMessage(conns []*websocket.Conn, msg *SocketMessage) []string {
	var errorMessages []string
	for _, conn := range conns {
		err := conn.WriteJSON(msg)
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
	}
	if len(errorMessages) == 0 {
		return nil
	}
	return errorMessages
}

func SendSocketMessage(conn *websocket.Conn, msg *SocketMessage) error {
	err := conn.WriteJSON(msg)
	return err
}
