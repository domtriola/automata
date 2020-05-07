#!/bin/bash

set -euo pipefail

echo "Running unit tests"

go test --race -coverprofile tmp/cp.out $(pwd)/...

echo
echo "Unit tests pass!"
echo "------------------------------------------------"
