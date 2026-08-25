package main

import "testing"

func TestGetOrCreateLobby(t *testing.T) {
	lobbiesMutex.Lock()
	lobbies = make(map[string]*Lobby)
	lobbiesMutex.Unlock()

	lobbyID := "test-lobby"

	firstLobby := getOrCreateLobby(lobbyID)
	secondLobby := getOrCreateLobby(lobbyID)

	if firstLobby == nil {
		t.Fatal("Lobby wurde nicht erstellt")
	}

	if firstLobby.ID != lobbyID {
		t.Fatalf(
			"Falsche Lobby-ID: erwartet %s, erhalten %s", lobbyID, firstLobby.ID)
	}

	if firstLobby.ConnectedClients == nil {
		t.Fatal("ConnectedCLients-Map wurde nicht initialisiert")
	}

	if firstLobby != secondLobby {
		t.Fatal("Bestehende Lobby wurde nicht wiederverwendet")
	}
}

func TestLobbyClientManagement(t *testing.T) {
	lobby := &Lobby{
		ID:               "test-lobby",
		ConnectedClients: make(map[string]*Client),
	}

	client := &Client{
		ID: "test-user",
	}

	lobby.AddClient(client)

	lobby.Mutex.RLock()
	storedClient, exists := lobby.ConnectedClients[client.ID]
	lobby.Mutex.RUnlock()

	if !exists {
		t.Fatal("CLient wurde nicht zur Lobby hinzugefugt")
	}

	if storedClient != client {
		t.Fatal("In der Lobby wurde nicht derselbe Client gespeichert")
	}

	lobby.RemoveClient(client.ID)

	lobby.Mutex.RLock()
	_, exists = lobby.ConnectedClients[client.ID]
	lobby.Mutex.RUnlock()

	if exists {
		t.Fatal("CLient wurde nicht aus der Lobby entfernt")
	}
}
