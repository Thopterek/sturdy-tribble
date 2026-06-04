package main

import (
	"net/http"
	"time"

	"github.com/tsironi93/WebServer/internal/auth"
)

type RefreshResp struct {
	Token string `json:"token"`
}

// HandlerTokenRefresh godoc
// @Summary Refresh JWT token
// @Description Exchanges a refresh token (provided in Authorization header) for a new JWT token.
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <refresh token>"
// @Success 200 {object} RefreshResp
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/refresh [post]
func (cfg *apiConf) HandlerTokenRefresh(w http.ResponseWriter, r *http.Request) {

	bearer, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "couldnt get token", err)
		return
	}

	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), bearer)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "token not in the database", err)
		return
	}

	expires := time.Hour
	token, err := auth.MakeJWT(user, cfg.JWTSecret, expires)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to create token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, RefreshResp{
		Token: token,
	})
}
