package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	RecipientID string `json:"recipient_id"`
	Content string `json:"content"`
}

type Client struct {
	ID string
	Conn *websocket.Conn
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

	defer conn.Close()

	fmt.Println("WebSocket verbunden")

	for {
		var message Message

		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Println("Verbindung beendet oder ungultiges JSON:", err)
			return
		}

		fmt.Println("Empfanger:", message.RecipientID)
		fmt.Println("Nachricht:", message.Content)

		err = conn.WriteJSON(message)
		if err != nil {
			fmt.Println("Fehler beim Senden:", err)
			return
		}
	}
}