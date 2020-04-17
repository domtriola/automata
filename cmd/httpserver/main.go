package main

import (
	"fmt"
	"net/http"

	"github.com/domtriola/automata/internal/handlers"
	log "github.com/sirupsen/logrus"
)

const port = 8000

func main() {
	http.HandleFunc("/healthcheck", handlers.Healthcheck)
	log.Printf("Starting server at: localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}
