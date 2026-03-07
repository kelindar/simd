#!/bin/bash
set -euo pipefail

cd "$(dirname "$0")"
go run ./templates/main.go

# requires gocc: go install github.com/kelindar/gocc/cmd/gocc@latest

#gocc simd_avx2_amd64.c --arch avx2 -O3 --package simd --output ../
#gocc simd_neon_arm64.c --arch neon -O3 --package simd --output ../
#gocc simd_apple_arm64.c --arch apple -O3 --package simd --output ../

