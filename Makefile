test: lint unit integration

unit:
	scripts/unit.sh

integration:
	test/integration/run.sh

lint:
	scripts/lint.sh

format:
	gofmt -s -w .

example:
	go run cmd/cli/main.go
