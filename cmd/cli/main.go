package main

import (
	"fmt"
	"log"

	"github.com/domtriola/automata-gen/internal/models"
)

func main() {
	s := models.NewSimulationRunner("cellular_automata")
	filename, err := s.BuildGIF()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Simulation GIF created at:", filename)
}
