# Workflow

## Dev Setup

Prerequisites: [xcode](https://developer.apple.com/library/archive/technotes/tn2339/_index.html) (if mac), [go](https://golang.org/doc/install), [golangci-lint](https://github.com/golangci/golangci-lint)

### Install

```bash
go get github.com/domtriola/automata
```

### Test

* All tests: `make test`
* Unit tests: `make unit`
* Integration tests: `make integration`
* Lint tests: `make lint`

## Generating Simulations

### Cli

Basic example: `make example`
