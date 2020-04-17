#!/bin/bash

set -euo pipefail

OUTPUT_PATH="test/tmp/test_sim.gif"

rm -f test/tmp/*.gif

go run cmd/cli/main.go --out $OUTPUT_PATH

if [ ! -f $OUTPUT_PATH ]; then
  echo "$OUTPUT_PATH not found"
  exit 1
fi
