package main

import "sync"

type Lobby struct {
	ID               string
	ConnectedClients map[string]*Client
	Mutex            sync.RWMutex
}

var lobbies = make(map[string]*Lobby)
var lobbiesMutex sync.RWMutex

func getOrCreateLobby(lobbyID string) *Lobby {
	lobbiesMutex.Lock()
	defer lobbiesMutex.Unlock()

	lobby, exists := lobbies[lobbyID]
	if exists {
		return lobby
	}

	lobby = &Lobby{
		ID:               lobbyID,
		ConnectedClients: make(map[string]*Client),
	}

	lobbies[lobbyID] = lobby

	return lobby
}
