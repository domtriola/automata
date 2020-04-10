test: lint unit

unit:
	scripts/unit.sh

lint:
	scripts/lint.sh

example:
	go run cmd/cli/main.go
