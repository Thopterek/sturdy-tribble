package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func createTestToken(t *testing.T, secret, userID string) string {
	t.Helper()

	claims := jwt.RegisteredClaims{
		Issuer:		"chirpy",
		Subject:	userID,
		ExpiresAt:	jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Test-Token konnte nicht erstellt werden: %v", err)
	}
	return signedToken
}

func connectTestClient(
	t *testing.T,
	websocketURL string,
	token string,
) *websocket.Conn {
	t.Helper()

	conn, _, err := websocket.DefaultDialer.Dial(websocketURL, nil)
	if err != nil {
		t.Fatalf("WebSocket-Verbindung fehlgeschlagen: %v", err)
	}

	t.Cleanup(func() {
		conn.Close()
	})

	err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		t.Fatalf("Read-Deadline konnte nicht gesetzt werden: %v", err)
	}

	err = conn.WriteJSON(AuthRequest{
		Type:	"auth",
		Token:	token,
	})
	if err != nil {
		t.Fatalf("Auth-Nachricht konnte nicht gesendet werden: %v", err)
	}

	var response AuthResponse

	err = conn.ReadJSON(&response)
	if err != nil {
		t.Fatalf("Auth-Antwort konnte nicht gelesen werden: %v", err)
	}

	if response.Type != "auth_success" {
		t.Fatalf(
			"Unerwartete Auth-Antwort: erwartet auth_success, erhalten %s",
			response.Type,
		)
	}
	return conn
}

func joinTestLobby(
	t *testing.T,
	conn *websocket.Conn,
	lobbyID string,
) {
	t.Helper()

	err := conn.WriteJSON(ClientEvent{
		Type:	"join_lobby",
		LobbyID:	"lobbyID",
	})
	if err != nil {
		t.Fatalf("Lobby-Beitritt konnte nicht gesendet werden: %v", err)
	}

	var response LobbyJoinedResponse

	err = conn.ReadJSON(&response)
	if err != nil {
		t.Fatalf("Lobby-Antwort konnte nicht gesendet werden: %v", err)
	}

	if response.Type != "lobby_joined" {
		t.Fatalf(
			"Unerwarteter Antworttyp: erwartet lobby_joined, erhalten %s",
			response.Type,
		)
	}

	if response.LobbyID != lobbyID {
		t.Fatalf(
			"Falsche Lobby-ID: erwartet %s, erhalten %s",
			lobbyID,
			response.LobbyID,
		)
	}
}

