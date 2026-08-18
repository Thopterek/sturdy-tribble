package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Conn       *websocket.Conn
	WriteMutex sync.Mutex
}

func (client *Client) SendJSON(value any) error {
	client.WriteMutex.Lock()
	defer client.WriteMutex.Unlock()

	return client.Conn.WriteJSON(value)
}

func (client *Client) SendError(message string) error {
	return client.SendJSON(ErrorResponse{
		Type:    "error",
		Message: message,
	})
}

var clients = make(map[string]*Client)
var clientsMutex sync.RWMutex
