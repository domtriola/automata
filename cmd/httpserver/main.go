package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/domtriola/automata-gen/internal/handlers"
)

const port = 8000

func main() {
	http.HandleFunc("/healthcheck", handlers.Healthcheck)
	log.Printf("Starting server at: localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}
