#!/usr/bin/env bash

set -e

CODE_DIRS="api cmd internal sql"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd)"

# Make sure go install'ed tools are available in PATH
. "$SCRIPT_DIR/set-path.sh"

cd "$SCRIPT_DIR/.." # back to root

go install github.com/daixiang0/gci@v0.10.1
go install mvdan.cc/gofumpt@v0.5.0

# shellcheck disable=SC2086 # we do want word splitting here to pass CODE_DIRS as multiple args
"$SCRIPT_DIR"/format-go.sh $CODE_DIRS

# Sometimes gofumpt needs to be called twice to ensure stable formatting
# shellcheck disable=SC2086
"$SCRIPT_DIR"/format-go.sh $CODE_DIRS
