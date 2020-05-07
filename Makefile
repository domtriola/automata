test: unit integration lint

unit:
	scripts/unit.sh

integration:
	test/integration/run.sh

lint:
	scripts/lint.sh

format:
	gofmt -s -w .

clean:
	rm -rf tmp/{*.gif,*.out}

example_automata:
	go run cmd/cli/main.go --width 100 --height 100 --nFrames 100

example_slimemold:
	go run cmd/cli/main.go --sim slime_mold --width 100 --height 100 --nFrames 100

.PHONY: test unit integration lint format example_automata example_slimemold
