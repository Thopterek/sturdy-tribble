package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConf) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

// HandlerMetrics godoc
// @Summary Admin metrics page
// @Description Returns a simple HTML page with metrics. Admin-only endpoint in this app.
// @Tags admin
// @Accept html
// @Produce html
// @Success 200 {string} string
// @Router /admin/metrics [get]
func (cfg *apiConf) HandlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	adminMetrics := fmt.Sprintf(`
	<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>
	`, cfg.fileserverHits.Load())

	w.Write([]byte(adminMetrics))
}
