package main

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func TestValidateJWT(t *testing.T) {
	secret := "test-secret"
	expectedUserID := uuid.New()
	expectedLobbyID := uuid.New()

	claims := ChatClaims{
		LobbyID: expectedLobbyID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "chirpy",
			Subject:   expectedUserID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("Token konnte nicht erstellt werden: %v", err)
	}

	userID, lobbyID, err := validateJWT(signedToken, secret)
	if err != nil {
		t.Fatalf("Token wurde abfgelehnt: %v", err)
	}

	if userID != expectedUserID.String() {
		t.Fatalf(
			"Falsche. User-ID: erwartet %s, erhalten %s",
			expectedUserID.String(),
			userID,
		)
	}

	if lobbyID != expectedLobbyID.String() {
		t.Fatalf(
			"Falsche Lobby-ID: erwartet %s, erhalten %s",
			expectedLobbyID.String(),
			lobbyID,
		)
	}
}
