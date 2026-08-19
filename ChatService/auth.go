package main

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthRequest struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

func validateJWT(tokenString string, secret string) (string, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		},
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Alg(),
		}),
	)
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("token ist ungultig")
	}

	if claims.Subject == "" {
		return "", errors.New("token enthalt keine user-id")
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return "", errors.New("user-id im token ist keine gultige UUID")
	}

	return userID.String(), nil
}
