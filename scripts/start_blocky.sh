#!/usr/bin/env bash
set -euo pipefail

podman run \
    --interactive \
    --tty \
    --rm \
    --publish 5353:5353/tcp \
    --publish 5353:5353/udp \
    --publish 8053:8053/tcp \
    --publish 8053:8053/udp \
    --publish 8080:8080/tcp \
    --volume ./config.yml:/app/config.yml:ro \
    --volume ./adlist.txt:/app/adlist.txt:ro \
    --security-opt label=disable \
    ghcr.io/0xerr0r/blocky:v0.23
