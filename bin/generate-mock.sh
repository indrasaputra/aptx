#!/usr/bin/env bash

set -euo pipefail

for file in `find . -name '*.go' | grep -v proto `; do
    if `grep -q 'interface {' ${file}`; then
        mockgen -source=${file} -destination=test/mock/${file}
    fi
done