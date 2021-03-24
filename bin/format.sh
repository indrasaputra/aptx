#!/usr/bin/env bash

set -euo pipefail

goimports -w -local github.com/indrasaputra/url-shortener $(go list -f {{.Dir}} ./...)
gofmt -s -w .

for file in `find . -name '*.proto'`; do
    clang-format -i ${file}
done
