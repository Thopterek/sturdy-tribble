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
		Issuer:    "chirpy",
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
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
		Type:  "auth",
		Token: token,
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
		Type:    "join_lobby",
		LobbyID: lobbyID,
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

func TestWebSocketLobbyBroadcast(t *testing.T) {
	secret := "test-secret"

	clientsMutex.Lock()
	clients = make(map[string]*Client)
	clientsMutex.Unlock()

	lobbiesMutex.Lock()
	lobbies = make(map[string]*Lobby)
	lobbiesMutex.Unlock()

	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			websocketHandler(w, r, secret)
		},
	))
	defer server.Close()

	websocketURL := "ws" + strings.TrimPrefix(server.URL, "http")

	user1ID := uuid.NewString()
	user2ID := uuid.NewString()

	user1Token := createTestToken(t, secret, user1ID)
	user2Token := createTestToken(t, secret, user2ID)

	user1 := connectTestClient(t, websocketURL, user1Token)
	user2 := connectTestClient(t, websocketURL, user2Token)

	lobbyID := uuid.NewString()

	joinTestLobby(t, user1, lobbyID)
	joinTestLobby(t, user2, lobbyID)

	content := "Hallo and die gesamte Lobby"

	err := user1.WriteJSON(ClientEvent{
		Type:    "message",
		LobbyID: lobbyID,
		Content: content,
	})
	if err != nil {
		t.Fatalf("Lobby-Nachricht konnte nicht gesendet werden: %v", err)
	}

	var senderMessage Message

	err = user1.ReadJSON(&senderMessage)
	if err != nil {
		t.Fatalf("Sender-Bestatigung konnte nicht gelesen werden: %v", err)
	}

	var recipientMessage Message

	err = user2.ReadJSON(&recipientMessage)
	if err != nil {
		t.Fatalf("Empfanger-Nachricht konnte nicht gelesen werden: %v", err)
	}

	if senderMessage.Type != "message_sent" {
		t.Fatalf(
			"Sender erhielt falschen Typ: erwartet message_sent, erhalten %s",
			senderMessage.Type,
		)
	}

	if recipientMessage.Type != "message" {
		t.Fatalf(
			"Empfanger erhielt falschen Typ: erwartet message, erhalten %s",
			recipientMessage.Type,
		)
	}

	if senderMessage.ID == "" {
		t.Fatal("Nachrichten-ID fehlt")
	}

	if senderMessage.ID != recipientMessage.ID {
		t.Fatal("Sender und Empfanger erhielten unterschiedliche Nachrichten-IDs")
	}

	if senderMessage.SenderID != user1ID {
		t.Fatalf(
			"Falscher Sender-ID: erwartet %s, erhalten %s",
			user1ID,
			senderMessage.SenderID,
		)
	}

	if senderMessage.LobbyID != lobbyID || recipientMessage.LobbyID != lobbyID {
		t.Fatal("Nachricht enthalt eine falsche Lobby-ID")
	}

	if senderMessage.Content != content || recipientMessage.Content != content {
		t.Fatal("Nachrichteninhalt stimmt nicht uberein")
	}

	if senderMessage.CreatedAt.IsZero() || recipientMessage.CreatedAt.IsZero() {
		t.Fatal("Erstellungszeitpunkt fehlt")
	}

	if !senderMessage.CreatedAt.Equal(recipientMessage.CreatedAt) {
		t.Fatal("Sender und Empfanger erhalten untershcieliche Zeitstempel")
	}

}
