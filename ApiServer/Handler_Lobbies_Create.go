package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/tsironi93/WebServer/internal/auth"
	"github.com/tsironi93/WebServer/internal/database"
)

func (cfg *apiConf) HandlerLobbiesCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		LobbyName        string `json:"lobby_name"`
		ShortDescription string `json:"short_description"`
		GameMap          string `json:"game_map"`
	}

	params := parameters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request body", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "missing authorization header", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.JWTSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid token", err)
		return
	}

	lobby, err := cfg.db.CreateLobby(r.Context(), database.CreateLobbyParams{
		LobbyName:        params.LobbyName,
		GameMaster:       userID,
		ShortDescription: sql.NullString{},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create lobby", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, lobby)
}
