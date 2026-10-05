#!/bin/sh
set -eu

BIN="bin/herdr-gitlab-issues"

mkdir -p bin
go build -o "$BIN" ./cmd/herdr-gitlab-issues
echo "built $BIN with $(go version)"
