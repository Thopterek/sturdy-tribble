package main

import (
	"net/http"

	"github.com/google/uuid"
)

// HandlerChirpsGetSingle godoc
// @Summary Get a single chirp
// @Description Get a chirp by its UUID
// @Tags chirps
// @Accept json
// @Produce json
// @Param chirpID path string true "Chirp UUID"
// @Success 200 {object} Chirp
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/chirps/{chirpID} [get]
func (cfg *apiConf) HandlerChirpsGetSingle(w http.ResponseWriter, r *http.Request) {
	uuidString := r.PathValue("chirpID")
	chirpUUID, err := uuid.Parse(uuidString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed parsing uuid", err)
		return
	}

	chirp, err := cfg.db.GetSingleChirp(r.Context(), chirpUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "chirp not in the database", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
