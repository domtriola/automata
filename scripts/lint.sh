#!/bin/bash

set -euo pipefail

echo "Running lint tests"

golangci-lint run \
  -E bodyclose \
  -E depguard \
  -E dogsled \
  -E dupl \
  -E gochecknoglobals \
  -E gochecknoinits \
  -E gocognit \
  -E goconst \
  -E gocritic \
  -E gocyclo \
  -E gofmt \
  -E goimports \
  -E golint \
  -E goprintffuncname \
  -E gosec \
  -E interfacer \
  -E lll \
  -E maligned \
  -E misspell \
  -E nakedret \
  -E prealloc \
  -E rowserrcheck \
  -E scopelint \
  -E stylecheck \
  -E unconvert \
  -E whitespace \
  -E wsl \
  # -E funlen \
  # -E unparam \
  # -E godox \

echo
echo "Lint tests pass!"
echo "------------------------------------------------"
