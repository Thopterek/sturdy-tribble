package main

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type Lobby struct {
	ID               string
	ConnectedClients map[string]*Client
	Mutex            sync.RWMutex
}

type JoinLobbyRequest struct {
	Type    string `json:"type"`
	LobbyID string `json:"lobby_id"`
}

type LobbyJoinedResponse struct {
	Type    string `json:"type"`
	LobbyID string `json:"lobby_id"`
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

func (lobby *Lobby) AddClient(client *Client) {
	lobby.Mutex.Lock()
	defer lobby.Mutex.Unlock()

	lobby.ConnectedClients[client.ID] = client
}

func (lobby *Lobby) RemoveClient(userID string) {
	lobby.Mutex.Lock()
	defer lobby.Mutex.Unlock()

	delete(lobby.ConnectedClients, userID)
}

func removeClientFromLobby(lobby *Lobby, userID string) {
	lobbiesMutex.Lock()
	defer lobbiesMutex.Unlock()

	lobby.Mutex.Lock()
	defer lobby.Mutex.Unlock()

	delete(lobby.ConnectedClients, userID)

	if len(lobby.ConnectedClients) != 0 {
		return
	}

	storedLobby, exists := lobbies[lobby.ID]
	if exists && storedLobby == lobby {
		delete(lobbies, lobby.ID)
	}
}

func (lobby *Lobby) ClientsSnapshot() []*Client {
	lobby.Mutex.RLock()
	defer lobby.Mutex.RUnlock()

	clients := make([]*Client, 0, len(lobby.ConnectedClients))

	for _, client := range lobby.ConnectedClients {
		clients = append(clients, client)
	}
	return clients
}

func joinLobby(client *Client, lobbyID string) (*Lobby, error) {
	parsedLobbyID, err := uuid.Parse(lobbyID)
	if err != nil {
		return nil, errors.New("lobby-id ist keine gultige UUID")
	}

	normalizedLobbyID := parsedLobbyID.String()

	lobbiesMutex.Lock()
	defer lobbiesMutex.Unlock()

	lobby, exists := lobbies[normalizedLobbyID]
	if !exists {
		lobby = &Lobby{
			ID:               normalizedLobbyID,
			ConnectedClients: make(map[string]*Client),
		}
		lobbies[normalizedLobbyID] = lobby
	}

	lobby.Mutex.Lock()
	lobby.ConnectedClients[client.ID] = client
	lobby.Mutex.Unlock()

	return lobby, nil
}
