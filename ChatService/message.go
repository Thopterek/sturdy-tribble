package main

import "time"

const maxMessageLength = 1000

type Message struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	SenderID    string    `json:"sender_id"`
	RecipientID string    `json:"recipient_id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"created_at"`
}

type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
