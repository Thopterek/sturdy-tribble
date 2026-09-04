package main

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthResponse struct {
	Type   string `json:"type"`
	UserID string `json:"user_id"`
}

type AuthRequest struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

type ChatClaims struct {
	LobbyID string `json:"lobby_id"`
	jwt.RegisteredClaims
}

func validateJWT(tokenString string, secret string) (string, string, error) {
	claims := &ChatClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Alg(),
		}),
		jwt.WithIssuer("chirpy"),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", errors.New("token ist ungultig")
	}

	if claims.Subject == "" {
		return "", "", errors.New("token nthalt keine user-id")
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return "", "", errors.New(
			"user-id im token ist keine gulgitge UUID",
		)
	}

	if claims.LobbyID == "" {
		return "", "", errors.New("token enthalt keine lobby-id")
	}

	lobbyID, err := uuid.Parse(claims.LobbyID)
	if err != nil {
		return "", "", errors.New(
			"lobby-id im token ist keine guktige UUID",
		)
	}
	return userID.String(), lobbyID.String(), nil
}
