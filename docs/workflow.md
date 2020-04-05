# Workflow

## Dev Setup

* Install

```bash
go get github.com/domtriola/automata-gen
```

* Start Server

```bash
go run cmd/httpserver/main.go
```

## Example Endpoints

* [5 species, all directions](http://localhost:8000/gen?width=400&height=400&nFrames=800&delay=2&nSpecies=5&threshold=3)
