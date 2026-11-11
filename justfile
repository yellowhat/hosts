# renovate: datasource=docker depName=docker.io/golang
GOLANG_VER := "1.27.4"

set shell := ["bash", "-o", "errexit", "-o", "nounset", "-o", "pipefail", "-c"]

@_:
    just --list --unsorted

# Start interactive go container
env:
    podman run \
        --interactive \
        --tty \
        --rm \
        --volume .:/data:z \
        --workdir /data \
        "docker.io/golang:{{ GOLANG_VER }}"

# Generate new adlist
adlist:
    (cd go && go build -o /tmp/out)
    rm -rf raw
    /tmp/out

# Run blocky container detached
run:
    bash scripts/run_blocky.sh

# Run blocky and test
test: clean run && clean
    bash scripts/test.sh

# Clean cached files and running containers
clean:
    docker rm --force blocky
    rm --recursive --force --verbose \
        **/.venv \
        **/__pycache__ \
        **/.pytest_cache \
        **/.pytype \
        **/.mypy_cache

# Start environment to benchmark
[group("benchmark")]
[working-directory: "./bench"]
bench-env:
    podman kube play pod.yaml --replace

# Start benchmark
[group("benchmark")]
[working-directory: "./bench"]
bench-run:
    bash 00_dnspyre_blocky.sh

# Stop pod used for benchmark
[group("benchmark")]
[working-directory: "./bench"]
bench-stop:
    podman kube down pod.yaml
