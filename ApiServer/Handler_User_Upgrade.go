package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/tsironi93/WebServer/internal/auth"
)

type Payload struct {
	Event string `json:"event"`
	Data  Data   `json:"data"`
}

type Data struct {
	UserID uuid.UUID `json:"user_id"`
}

// HandlerUserUpgradeToRed godoc
// @Summary Handle Polka webhook for user upgrade
// @Description Receives Polka webhook events and upgrades a user to 'Chirpy Red'. Expects Authorization header with ApiKey.
// @Tags webhooks, users
// @Accept json
// @Produce json
// @Param Authorization header string true "ApiKey <key>"
// @Param payload body Payload true "Webhook payload"
// @Success 204
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/polka/webhooks [post]
func (cfg *apiConf) HandlerUserUpgradeToRed(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	param := Payload{}
	if err := decoder.Decode(&param); err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldnt decode request", err)
		return
	}

	k, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if k != cfg.PolkaKey {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if param.Event != "user.upgraded" {
		respondWithError(w, http.StatusNoContent, "", nil)
		return
	}

	if err := cfg.db.UserUpgradeToChirpRed(r.Context(), param.Data.UserID); err != nil {
		respondWithError(w, http.StatusNotFound, "failed to upgraded", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
