#!/bin/bash

set -euo pipefail

echo "Running unit tests"

echo "testing internal/..."

go test $(pwd)/internal/...

echo "testing pkg/..."

go test $(pwd)/pkg/...

echo
echo "Unit tests pass!"
echo "------------------------------------------------"
