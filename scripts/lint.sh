#!/bin/bash

set -euo pipefail

echo "Running lint tests"

golangci-lint run

echo
echo "Lint tests pass!"
echo "------------------------------------------------"
