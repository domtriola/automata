package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/domtriola/automata-gen/internal/runner"
)

type args struct {
	sim string
}

func main() {
	a := collectArgs()

	s, err := runner.New(a.sim)
	if err != nil {
		log.Fatalln(err)
	}

	filename, err := s.BuildGIF()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Simulation GIF created at:", filename)
}

func collectArgs() *args {
	sim := flag.String("sim", "cellular_automata", "The type of simulation to generate")
	flag.Parse()

	return &args{
		sim: *sim,
	}
}
