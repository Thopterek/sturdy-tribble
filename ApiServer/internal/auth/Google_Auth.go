package auth

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
)

type GoogleClaims struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	EmailVerified bool   `json:"email_verified"`
}

func VerifyGoogleToken(ctx context.Context, clientID string, idToken string) (*GoogleClaims, error) {

	provider, err := oidc.NewProvider(
		ctx,
		"https://accounts.google.com",
	)
	if err != nil {
		return nil, err
	}

	verifier := provider.Verifier(
		&oidc.Config{
			ClientID: clientID,
		},
	)

	token, err := verifier.Verify(
		ctx,
		idToken,
	)
	if err != nil {
		return nil, err
	}

	var claims GoogleClaims

	if err := token.Claims(&claims); err != nil {
		return nil, err
	}

	return &claims, nil
}
