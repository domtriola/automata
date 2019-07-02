package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/domtriola/automata-gen/simulation"
)

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	query := r.URL.Query()

	fmt.Print("Server says...\n")
	fmt.Printf("Your path is: %q\n", path)
	fmt.Printf("Your your query is: \"%v\"\n", query)

	fmt.Println("Building the simulation...")
	imgName := simulation.Build()
	http.ServeFile(w, r, imgName)
}
