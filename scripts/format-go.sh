#!/usr/bin/env bash

set -e

for path in "$@"; do
  if [ -d "$path" ]; then
    # gofumpt can format entire directories, but it skips auto-generated files.
    find "$path" -type f -name '*.go' -print0 | xargs -0 gofumpt -l -w
  else
    gofumpt -l -w "$path"
  fi
done

gci write "$@" -s standard -s 'prefix(mkuznets.com/go/sfs)' -s default -s blank -s dot
