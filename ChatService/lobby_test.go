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
