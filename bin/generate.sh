#!/usr/bin/env bash

set -euo pipefail

IMAGE_DIR=bin

buf lint
buf breaking --against '.git#branch=main'
buf generate

mkdir -p ${IMAGE_DIR}
buf build -o ${IMAGE_DIR}/image.bin