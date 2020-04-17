#!/bin/bash

set -euo pipefail

echo "Running integration tests"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source $DIR/cli_test.sh

echo
echo "Integration tests pass!"
echo "------------------------------------------------"
