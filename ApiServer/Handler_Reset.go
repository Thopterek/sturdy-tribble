package main

import (
	"fmt"
	"net/http"
)

// HandlerResetHits godoc
// @Summary Reset application hits and delete all users (dev only)
// @Description Resets in-memory hits counter and deletes all users from the database. Only available on dev platform.
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin/reset [post]
func (cfg *apiConf) HandlerResetHits(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Forbidden", fmt.Errorf("NOPE"))
		return
	}

	cfg.fileserverHits.Store(0)
	if err := cfg.db.DeleteAllUsers(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to reset the database: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	hits := fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())
	w.Write([]byte(hits))
}
