#!/usr/bin/env bash
# Benchmark providers
set -euo pipefail

providers=(
    https://dns.quad9.net/dns-query
    quic://9.9.9.9:853
    https://blank.dnsforge.de/dns-query
    quic://blank.dnsforge.de
    https://unfiltered.joindns4.eu/dns-query
)
for provider in "${providers[@]}"; do
    echo "[INFO] $provider"
    args=(
        --server "$provider"
        # Quad9 rejects HTTP/1.1 with 505, force HTTP/2
        --doh-protocol 2
        --concurrency 1
        --number 3
        --no-distribution
        --probability 0.33
        https://raw.githubusercontent.com/Tantalor93/dnspyre/master/data/1000-domains
    )

    podman run \
        --name "blocky" \
        --interactive \
        --tty \
        --rm \
        ghcr.io/tantalor93/dnspyre:v3.12.0 "${args[@]}"
done
