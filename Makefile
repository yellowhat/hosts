.DEFAULT_GOAL := help
SHELL         := /bin/bash
MAKEFLAGS     += --no-print-directory
MKFILE_DIR    := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

.PHONY: adlist
adlist: ## Build container
	@echo "[INFO] Generating new adlist"
	@python adlist.py
	@echo ""

.PHONY: run-blocky
run-blocky: ## Run blocky container detached
	@echo "[INFO] Run blocky container..."
	./scripts/run_blocky.sh

.PHONY: test
test: ## Run app and redis container
	$(MAKE) run-blocky
	@echo "[INFO] Run test"
	./scripts/test.sh

.PHONY: clean
clean: ## Clean cached files and running containers
	docker rm --force blocky
	shopt -s globstar
	rm --recursive --force --verbose \
		**/.venv \
		**/__pycache__ \
		**/.pytest_cache \
		**/.pytype \
		**/.mypy_cache

.PHONY: help
help: ## Makefile Help Page
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[\/\%a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-21s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST) 2>/dev/null

