#!/bin/bash

set -euo pipefail

rm -f test/tmp/*.gif

echo "it creates a basic simulation"

OUTPUT_PATH="test/tmp/test_sim.gif"

go run cmd/cli/main.go --out $OUTPUT_PATH --width 10 --height 10 --nFrames 100

if [ ! -f $OUTPUT_PATH ]; then
  echo "$OUTPUT_PATH not found"
  exit 1
fi

echo "ok"
echo "--------"

echo "it creates a cellular automata simulation with custom params"

OUTPUT_PATH="test/tmp/ca_custom_params.gif"

go run cmd/cli/main.go --out $OUTPUT_PATH --width 4 --height 10 --nFrames 2 --nSpecies 2 --pThreshold 2 --pDirs "n,e,w,s"

if [ ! -f $OUTPUT_PATH ]; then
  echo "$OUTPUT_PATH not found"
  exit 1
fi

echo "ok"
echo "--------"

echo "it creates a slime mold simulation with custom params"

OUTPUT_PATH="test/tmp/sm_custom_params.gif"

go run cmd/cli/main.go --out $OUTPUT_PATH --sim slime_mold --width 4 --height 10 --nFrames 2 --sDecay 0.6 --sSpread 0.1 --sReach 5 --sDeg 45

if [ ! -f $OUTPUT_PATH ]; then
  echo "$OUTPUT_PATH not found"
  exit 1
fi

echo "ok"
echo "--------"

echo "cli tests pass"
echo "----------------------"
