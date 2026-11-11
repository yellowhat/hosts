#!/usr/bin/env bash
# Benchmark blocky (pod published UDP port, see pod.yaml)
set -euo pipefail

# renovate: datasource=docker depName=ghcr.io/tantalor93/dnspyre
DNSPYRE_VER=v3.12.0

args=(
    # getent hosts host.containers.internal
    --server "169.254.1.2:8053"
    --concurrency 1
    --duration 60s
    --no-distribution
    --probability 0.33
    https://raw.githubusercontent.com/Tantalor93/dnspyre/master/data/10000-domains
)
podman run \
    --name dnspyre \
    --interactive \
    --tty \
    --rm \
    "ghcr.io/tantalor93/dnspyre:$DNSPYRE_VER" "${args[@]}"
