#!/bin/bash

set -euo pipefail

echo "Running unit tests"

go test $(pwd)/...

echo
echo "Unit tests pass!"
echo "------------------------------------------------"
