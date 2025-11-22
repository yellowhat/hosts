#!/usr/bin/env bash
# Test if hostnames are blocked or not
set -euo pipefail

docker run \
    --name blocky-test \
    --interactive \
    --tty \
    --rm \
    --network host \
    --volume .:/data:z \
    --security-opt label=disable \
    --workdir /data \
    --entrypoint ash \
    docker.io/alpine:latest \
    -c "apk add bash bind-tools && /data/scripts/test_inside.sh"
