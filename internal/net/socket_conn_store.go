package net

import "github.com/gorilla/websocket"

type SocketConnectionStore struct {
	store map[*websocket.Conn]bool
}

func CreateSocketConnStore() *SocketConnectionStore {
	return &SocketConnectionStore{
		store: make(map[*websocket.Conn]bool),
	}
}

func (this *SocketConnectionStore) Add(conn *websocket.Conn) {
	this.store[conn] = true
}

func (this *SocketConnectionStore) Clear() {
	for c := range this.store {
		delete(this.store, c)
	}
}

func (this *SocketConnectionStore) Remove(conn *websocket.Conn) {
	_, exists := this.store[conn]
	if exists {
		delete(this.store, conn)
	}
}
