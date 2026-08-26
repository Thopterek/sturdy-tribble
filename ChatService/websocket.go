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

	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("Auth-Timeout konnte nicht gesetzt werden:", err)
		conn.Close()
		return
	}

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

	userID, err := validateJWT(authRequest.Token, jwtSecret)
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
			fmt.Println("Client aus Lobby entfernen:", currentLobby.ID)
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

	for {
		var event ClientEvent

		err := conn.ReadJSON(&event)
		if err != nil {
			fmt.Println("Verbindung beendet oder ungultiges JSON:", err)
			return
		}

		if event.Type == "join_lobby" {
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

		message := Message{
			RecipientID: event.RecipientID,
			Content:     event.Content,
		}

		if message.RecipientID == "" {
			fmt.Println("Nachricht abgelehnt: Empfanger fehlt")

			err = client.SendError("Empfanger fehlt")
			if err != nil {
				fmt.Println("Fehler beim senden der Fehlermeldung:", err)
				return
			}
			continue
		}

		if strings.TrimSpace(message.Content) == "" {
			fmt.Println("Nachricht abgelehnt: Inhalt ist leer")

			err = client.SendError("Nachrichteninhalt darf nicht leer sein")
			if err != nil {
				fmt.Println("Fehler beim Senden der Fehlermeldung:", err)
				return
			}
			continue
		}

		if len(message.Content) > maxMessageLength {
			fmt.Println("Nachricht abgelehnt: Inhalt ist zu lang")

			err = client.SendError("Nachricht darf maximal 1000 Bytes lang sein")
			if err != nil {
				fmt.Println("Fehler beim Senden der Fehlermeldung:", err)
				return
			}
			continue
		}

		message.ID = uuid.NewString()
		message.Type = "message"
		message.SenderID = userID
		message.CreatedAt = time.Now().UTC()

		fmt.Println("Sender:", message.SenderID)
		fmt.Println("Empfanger:", message.RecipientID)
		fmt.Println("Nachricht:", message.Content)

		clientsMutex.RLock()
		recipient, found := clients[message.RecipientID]
		clientsMutex.RUnlock()

		if !found {
			fmt.Println("Empfanger nicht verbunden:", message.RecipientID)

			err = client.SendError("Empfanger ist nicht verbunden")
			if err != nil {
				fmt.Println("Fehler beim Sender der Fehlermeldung:", err)
				return
			}
			continue
		}

		err = recipient.SendJSON(message)
		if err != nil {
			fmt.Println("Fehler beim Weiterleiten", err)
			continue
		}

		sentConfirmation := message
		sentConfirmation.Type = "message_sent"

		err = client.SendJSON(sentConfirmation)
		if err != nil {
			fmt.Println("Fehler beim Senden der Bestatigung:", err)
			return
		}
	}
}
