package main

import (
	"encoding/json"
	"net/http"

	"github.com/tsironi93/WebServer/internal/auth"
	"github.com/tsironi93/WebServer/internal/database"
)

type RequestUserUpdate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseUserUpdate struct {
	Email string `json:"email"`
}

// HandlerUserUpdate godoc
// @Summary Update user password/email
// @Description Update the authenticated user's password and/or email. Requires Bearer JWT token.
// @Tags users, auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <JWT token>"
// @Param user body RequestUserUpdate true "User update payload"
// @Success 200 {object} ResponseUserUpdate
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users [put]
func (cfg *apiConf) HandlerUserUpdate(w http.ResponseWriter, r *http.Request) {
	bearer, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "no bearer header", err)
		return
	}

	user, err := auth.ValidateJWT(bearer, cfg.JWTSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid JWT", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	update := RequestUserUpdate{}
	if err := decoder.Decode(&update); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}

	pass, err := auth.HashPassword(update.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "couldnt hash the password", err)
		return
	}

	err = cfg.db.UserUpdatePassword(r.Context(), database.UserUpdatePasswordParams{
		ID:             user,
		HashedPassword: pass,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "user doesnt exists", err)
		return
	}

	respondWithJSON(w, http.StatusOK, ResponseUserUpdate{
		Email: update.Email,
	})
}
