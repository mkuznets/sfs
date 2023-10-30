#!/usr/bin/env bash

PATH="$PATH"

GOPATH="$(go env GOPATH)"
GOBIN="$(go env GOBIN)"

if [ "$GOBIN" ]; then
  PATH="$GOBIN:$PATH"
elif [ "$GOPATH" ]; then
  PATH="$GOPATH/bin:$PATH"
else
  PATH="$HOME/go/bin:$PATH"
fi

export PATH=$PATH
