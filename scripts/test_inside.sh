#!/usr/bin/env bash
# Test if hostnames are blocked or not
set -euo pipefail

SERVER=127.0.0.1
PORT=5353

readarray -t domains_block <blocklist.txt
readarray -t domains_allow <allowlist.txt

echo "[FIRST REQUEST]"
dig +retry=10 "@$SERVER" -p "$PORT" google.com
echo ""

echo "[BLOCKED]"
for domain in "${domains_block[@]}"; do
    if [ -z "$domain" ] || [[ "$domain" == \#* ]]; then
        continue
    fi
    ip=$(dig +short "@$SERVER" -p "$PORT" "$domain")
    if [[ "$ip" == "0.0.0.0" ]]; then
        echo "blocked: $domain"
    else
        echo "not blocked: $domain"
        exit 1
    fi
done
echo ""

echo "[NOT BLOCKED]"
for domain in "${domains_allow[@]}"; do
    if [ -z "$domain" ] || [[ "$domain" == \#* ]]; then
        continue
    fi
    ip=$(dig +short "@$SERVER" -p "$PORT" "$domain")
    if [[ "$ip" == "0.0.0.0" ]]; then
        echo "blocked: $domain"
        exit 1
    else
        echo "not blocked: $domain"
    fi
done
echo ""

echo "[INFO] Success!!"
