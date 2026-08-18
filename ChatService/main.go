package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	SenderID    string `json:"sender_id"`
	RecipientID string `json:"recipient_id"`
	Content     string `json:"content"`
}

type Client struct {
	ID         string
	Conn       *websocket.Conn
	WriteMutex sync.Mutex
}

var clients = make(map[string]*Client)
var clientsMutex sync.RWMutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

		message.SenderID = userID

		fmt.Println("Sender:", message.SenderID)
		fmt.Println("Empfanger:", message.RecipientID)
		fmt.Println("Nachricht:", message.Content)

		clientsMutex.RLock()
		recipient, found := clients[message.RecipientID]
		clientsMutex.RUnlock()

		if !found {
			fmt.Println("Empfanger nicht verbunden:", message.RecipientID)
			continue
		}

		recipient.WriteMutex.Lock()
		err = recipient.Conn.WriteJSON(message)
		recipient.WriteMutex.Unlock()
		if err != nil {
			fmt.Println("Fehler beim Weiterleiten", err)
			continue
		}
	}
}
