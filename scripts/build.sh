#!/usr/bin/env bash
#
# This script builds the application from source.

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

cd $DIR

# Delete the old dir
echo "==> Removing old directory..."
rm -f bin/*
mkdir -p bin/

# Instruct gox to build statically linked binaries
export CGO_ENABLED=0

# Set module download mode to readonly to not implicitly update go.mod
export GOFLAGS="-mod=readonly"

# Ensure all remote modules are downloaded and cached before build so that
# the concurrent builds launched by gox won't race to redundantly download them.
echo "==> Download remote modules..."
go mod download

# Build!
echo "==> Building..."
for cmd in $(ls cmd); do
  go build -o bin/$cmd cmd/$cmd/main.go
done

# Done!
echo "==> Results:"
ls bin | awk '{printf "    - %s\n", $1}'
