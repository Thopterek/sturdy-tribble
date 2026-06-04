package main

import (
	"net/http"

	"github.com/tsironi93/WebServer/internal/auth"
)

// HandlerTokenRevoke godoc
// @Summary Revoke a refresh token
// @Description Revokes a refresh token provided in the Authorization header.
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <refresh token>"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/revoke [post]
func (cfg *apiConf) HandlerTokenRevoke(w http.ResponseWriter, r *http.Request) {
	bearer, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "couldnt get token", err)
		return
	}

	if err := cfg.db.Revoke(r.Context(), bearer); err != nil {
		respondWithError(w, http.StatusUnauthorized, "token not int he database", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
