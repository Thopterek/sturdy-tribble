package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	jwtSecret := os.Getenv("SECRET")
	if jwtSecret == "" {
		log.Fatal("SECRET muss gesetzt sein")
	}
	
	http.HandleFunc("/ChatHealth", healthHandler)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocketHandler(w, r, jwtSecret)
	})

	fmt.Println("ChatService lauft auf Port 8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Serverfehler:", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ChatService ist erreichbar")
}
