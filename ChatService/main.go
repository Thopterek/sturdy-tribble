package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	SenderID    string `json:"sender_id"`
	RecipientID string `json:"recipient_id"`
	Content     string `json:"content"`
}

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const maxMessageLength = 1000

func main() {
	http.HandleFunc("/ChatHealth", healthHandler)
	http.HandleFunc("/ws", websocketHandler)

	fmt.Println("ChatService lauft auf Port 8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Serverfehler:", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ChatService ist erreichbar")
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	if userID == "" {
		http.Error(w, "user_id fehlt", http.StatusBadRequest)
		return
	}

	fmt.Println("Verbindungsanfrage von:", userID)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Websocket-Fehler:", err)
		return
	}

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

	defer func() {
		clientsMutex.Lock()
		delete(clients, userID)
		clientsMutex.Unlock()

		conn.Close()

		fmt.Println("Client entfernt:", userID)
	}()

	fmt.Println("WebSocket verbunden")

	for {
		var message Message

		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println("Verbindung beendet oder ungultiges JSON:", err)
			return
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

		message.SenderID = userID

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
	}
}
