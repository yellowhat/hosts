#!/usr/bin/env bash
# Test if hostnames are blocked or not
set -euo pipefail

podman run \
    --interactive \
    --tty \
    --rm \
    --network host \
    --volume "$PWD:/data" \
    --security-opt label=disable \
    --workdir /data \
    alpine sh -c "apk add bash bind-tools && /data/scripts/test_inside.sh"
