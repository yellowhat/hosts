#!/usr/bin/env bash
# Test if hostnames are blocked or not
set -euo pipefail

docker run \
    --interactive \
    --rm \
    --network host \
    --volume "$PWD:/data" \
    --security-opt label=disable \
    --workdir /data \
    docker.io/alpine:latest \
    sh -c "apk add bash bind-tools && /data/scripts/test_inside.sh"
