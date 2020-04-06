package main

import (
	"fmt"
	"log"

	runner "github.com/domtriola/automata-gen/internal/simulation_runner"
)

func main() {
	s, err := runner.NewSimulationRunner("cellular_automata")
	if err != nil {
		log.Fatalln(err)
	}

	filename, err := s.BuildGIF()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Simulation GIF created at:", filename)
}
