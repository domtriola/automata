package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	query := r.URL.Query()

	fmt.Fprint(w, "Server says...\n")
	fmt.Fprintf(w, "Your path is: %q\n", path)
	fmt.Fprintf(w, "Your your query is: \"%v\"\n", query)
}
