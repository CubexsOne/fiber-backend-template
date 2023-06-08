.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Starts backend service
	@./scripts/make-run.sh

test: ## Starts backend service with unit-tests
	@./scripts/make-unit-test.sh

test-ci: ## Test run for ci pipeline
	@./scripts/make-ci-unit-test.sh

build: ## Build backend service
	@./scripts/make-build.sh

publish: ## Push backend to registry
	@./scripts/make-publish.sh