#!/bin/bash

set -euo pipefail

OUTPUT_PATH="test/tmp/TODDOM.gif"

rm test/tmp/*.gif

go run cmd/cli/main.go

if [ ! -f $OUTPUT_PATH ]; then
  echo "$OUTPUT_PATH not found"
  exit 1
fi
