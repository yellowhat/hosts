#!/usr/bin/env bash
# Test if hostnames are blocked or not
set -euo pipefail

SERVER=127.0.0.1
PORT=5353

readarray -t domains_block <blocklist.txt
readarray -t domains_allow <allowlist.txt

echo "[INFO] First request"
dig +retry=10 "@$SERVER" -p "$PORT" google.com
echo ""

echo "[INFO] Blocked"
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

echo "[INFO] Not blocked"
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

dnssec_status() {
    dig +dnssec +time=10 +noall +comments "@$SERVER" -p "$PORT" "$1" |
        grep -o 'status: [A-Z]*' | cut -d' ' -f2
}

echo "[INFO] DNSSEC"
# Signed domain must resolve
status=$(dnssec_status cloudflare.com)
if [[ "$status" != "NOERROR" ]]; then
    echo "dnssec valid domain failed: cloudflare.com ($status)"
    exit 1
fi
echo "resolved: cloudflare.com"
# Domain with broken signatures must be rejected
status=$(dnssec_status dnssec-failed.org)
if [[ "$status" != "SERVFAIL" ]]; then
    echo "dnssec invalid domain not rejected: dnssec-failed.org ($status)"
    exit 1
fi
echo "rejected: dnssec-failed.org"
echo ""

echo "[INFO] Success!!"
