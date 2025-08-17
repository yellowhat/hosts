set shell := ["bash", "-o", "errexit", "-o", "nounset", "-o", "pipefail", "-c"]

@_:
    just --list --unsorted

# Build container
adlist:
    echo "[INFO] Generating new adlist"
    python adlist.py
    echo ""

# Run blocky container detached
run:
    @echo "[INFO] Run blocky container..."
    bash ./scripts/run_blocky.sh

# Run blocky and test
test: run
    @echo "[INFO] Run test"
    bash scripts/test.sh

# Clean cached files and running containers
clean:
    docker rm --force blocky
    shopt -s globstar
    rm --recursive --force --verbose \
    	**/.venv \
    	**/__pycache__ \
    	**/.pytest_cache \
    	**/.pytype \
    	**/.mypy_cache
