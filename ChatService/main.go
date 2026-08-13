package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Websocket-Fehler:", err)
		return
	}

	defer conn.Close()

	fmt.Println("WebSocket verbunden")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Verbindung beendet:", err)
			return
		}

		fmt.Println("Nachrichtentyp:", messageType)
		fmt.Println("Nachricht:", string(message))

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Fehler beim Senden:", err)
			return
		}
	}
}