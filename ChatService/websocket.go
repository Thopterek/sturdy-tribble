package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	pongWait           = 60 * time.Second
	pingPeriod         = 30 * time.Second
	maxClientEventSize = 2048
)

func websocketHandler(
	w http.ResponseWriter,
	r *http.Request,
	jwtSecret string,
) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Websocket-Fehler:", err)
		return
	}

	conn.SetReadLimit(maxClientEventSize)

	err = conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		fmt.Println("Pong-Timeout konnte nicht gesetzt werden:", err)
		conn.Close()
		return
	}

	conn.SetPongHandler(func(string) error {
		return conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	var authRequest AuthRequest

	err = conn.ReadJSON(&authRequest)
	if err != nil {
		fmt.Println("Auth-Nachricht konnte nicht gelesen werden:", err)
		conn.Close()
		return
	}

	if authRequest.Type != "auth" || authRequest.Token == "" {
		conn.WriteJSON(ErrorResponse{
			Type:    "error",
			Message: "Authentifzierung fehlt",
		})
		conn.Close()
		return
	}

	userID, allowedLobbyID, err := validateJWT(authRequest.Token, jwtSecret)
	if err != nil {
		fmt.Println("JWT abgelehnt:", err)

		conn.WriteJSON(ErrorResponse{
			Type:    "error",
			Message: "Token ist ungultig oder abgelaufen",
		})
		conn.Close()
		return
	}

	err = conn.SetReadDeadline(time.Time{})
	if err != nil {
		fmt.Println("Auth-Timeout konnte nicht entfernt werden:", err)
		conn.Close()
		return
	}

	fmt.Println("Benutzer authentifiziert:", userID)

	client := &Client{
		ID:   userID,
		Conn: conn,
	}

	clientsMutex.Lock()

	_, alreadyConnected := clients[userID]
	if alreadyConnected {
		clientsMutex.Unlock()

		fmt.Println("Client bereits verbunden:", userID)
		conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(
				websocket.ClosePolicyViolation,
				"Benutzer ist bereits verbunden",
			),
		)
		conn.Close()
		fmt.Println("Doppelte Verbindung geschlossen:", userID)
		return
	}
	clients[userID] = client
	clientsMutex.Unlock()

	fmt.Println("Client gespeichert:", userID)

	var currentLobby *Lobby

	defer func() {
		if currentLobby != nil {
			currentLobby.RemoveClient(userID)
			fmt.Println("Client aus Lobby entfernt:", currentLobby.ID)
		}

		clientsMutex.Lock()
		delete(clients, userID)
		clientsMutex.Unlock()

		conn.Close()

		fmt.Println("Client entfernt:", userID)
	}()

	err = client.SendJSON(AuthResponse{
		Type:   "auth_success",
		UserID: userID,
	})
	if err != nil {
		fmt.Println("Auth-Bestatigung konnte nicht gesendet werden:", err)
		return
	}

	fmt.Println("WebSocket authentifiziert und bereit")

	pingDone := make(chan struct{})
	defer close(pingDone)

	go func() {
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := client.SendPing()
				if err != nil {
					fmt.Println("Websocket-Ping fehlgeschlagen:", userID, err)
					conn.Close()
					return
				}
			case <-pingDone:
				return
			}
		}
	}()

	for {
		var event ClientEvent

		err := conn.ReadJSON(&event)
		if err != nil {
			fmt.Println("Verbindung beendet oder ungultiges JSON:", err)
			return
		}

		if event.Type == "join_lobby" {
			if event.LobbyID != allowedLobbyID {
				err = client.SendError(
					"Client ist nicht fur diese Lobby berechtig",
				)
				if err != nil {
					return
				}
				continue
			}
			if currentLobby != nil {
				err = client.SendError("Client ist bereits einer Lobby beigetreten")
				if err != nil {
					return
				}
				continue
			}

			lobby, joinerr := joinLobby(client, event.LobbyID)
			if joinerr != nil {
				err = client.SendError(joinerr.Error())
				if err != nil {
					return
				}
				continue
			}

			currentLobby = lobby

			err = client.SendJSON(LobbyJoinedResponse{
				Type:    "lobby_joined",
				LobbyID: lobby.ID,
			})
			if err != nil {
				fmt.Println("Lobby-Bestatigung konnte nicht gesendet werden:", err)
				return
			}
			fmt.Println("Client ist Lobby beigetreten:", lobby.ID)
			continue
		}

		if event.Type != "message" {
			err = client.SendError("Unbekannter Event-Type")
			if err != nil {
				return
			}
			continue
		}

		if currentLobby == nil {
			err = client.SendError("Client ist keiner Lobby beigetreten")
			if err != nil {
				return
			}
			continue
		}

		parsedLobbyID, parseErr := uuid.Parse(event.LobbyID)
		if parseErr != nil {
			err = client.SendError("Lobby-ID ist keine gultige UUID")
			if err != nil {
				return
			}
			continue
		}

		if parsedLobbyID.String() != currentLobby.ID {
			err = client.SendError("Client ist dieser Lobby nicht beigetreten")
			if err != nil {
				return
			}
			continue
		}

		if strings.TrimSpace(event.Content) == "" {
			fmt.Println("Nachricht abgelehnt: Inhalt ist leer")

			err = client.SendError("Nachrichteninhalt darf nicht leer sein")
			if err != nil {
				fmt.Println("Fehler beim senden der Fehlermeldung:", err)
				return
			}
			continue
		}

		if len(event.Content) > maxMessageLength {
			fmt.Println("Nachricht abgelehnt: Inhalt ist zu lang")

			err = client.SendError("Nachricht darf maximal 1000 bytes lang sein")
			if err != nil {
				fmt.Println("Fehler beim Sender der Fehlermeldung:", err)
				return
			}
			continue
		}

		message := Message{
			ID:        uuid.NewString(),
			Type:      "message",
			SenderID:  userID,
			LobbyID:   currentLobby.ID,
			Content:   event.Content,
			CreatedAt: time.Now().UTC(),
		}

		fmt.Println("Sender:", message.SenderID)
		fmt.Println("Lobby:", message.LobbyID)
		fmt.Println("Nachricht:", message.Content)

		for _, recipient := range currentLobby.ClientsSnapshot() {
			outgoingMessage := message

			if recipient.ID == userID {
				outgoingMessage.Type = "message_sent"
			}

			sendErr := recipient.SendJSON(outgoingMessage)
			if sendErr != nil {
				fmt.Println(
					"Nachricht konnte nicht an Client gesendet werden:",
					recipient.ID,
					sendErr,
				)
				if recipient.ID == userID {
					return
				}
			}
		}
	}
}
