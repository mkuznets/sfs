#!/usr/bin/env bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd)"

# Make sure go install'ed tools are available in PATH
. "$SCRIPT_DIR/set-path.sh"

cd "$SCRIPT_DIR/.."  # back to root

go install github.com/swaggo/swag/cmd/swag@v1.16.2
go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.3

swag init -g internal/api/api.go --output internal/api/swagger

# Add newline if it's missing
[ -n "$(tail -c1 internal/api/swagger/swagger.json)" ] && echo >> internal/api/swagger/swagger.json

swagger generate client --template=stratoscale --spec internal/api/swagger/swagger.json --name SimpleFeedService --strict-responders --target ./api

go generate ./...
go mod tidy

swag fmt --exclude internal/api/resources.go
