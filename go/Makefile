.DEFAULT_GOAL := help
SHELL          = /usr/bin/env bash -o pipefail
.SHELLFLAGS    = -ec
# renovate: datasource=docker depName=docker.io/golang
GOLANG_VER     = 1.24.2

.PHONY: it
it: ## Start interactive go container
	podman run \
		--interactive \
		--tty \
		--rm \
		--volume ".:/data:z" \
		--workdir /data \
    	"docker.io/golang:$(GOLANG_VER)"

.PHONY: help
help: ## Makefile Help Page
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[\/\%a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-21s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST) 2>/dev/null
