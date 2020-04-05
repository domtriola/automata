package handlers

import "net/http"

// Healthcheck returns a 200 to indicate that the http server is alive and active.
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
