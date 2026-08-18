package main

import (
	"fmt"
	"net/http"
)

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
