package main

import "net/http"

// HandlerReadiness godoc
// @Summary Health/readiness check
// @Description Returns 200 OK if the service is ready
// @Tags health
// @Accept plain
// @Produce plain
// @Success 200 {string} string
// @Router /api/healthz [get]
func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
