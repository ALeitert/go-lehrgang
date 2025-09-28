PROJ_DIR=$(shell pwd)


run:
	go run cmd/main.go --config=config/config.yaml


### Linting ###

check-golangci-lint:
	$(eval GOLANGCI_LINT_VERSION=$(shell curl -s https://api.github.com/repos/golangci/golangci-lint/releases/latest | jq -r '.name'))
	@./bin/golangci-lint --version | grep -qF "$(GOLANGCI_LINT_VERSION:1)" || { \
		echo "golangci-lint not found or version mismatch. Installing..."; \
		$(MAKE) install-golangci-lint; \
	}

install-golangci-lint: ## install golang lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

lint-check: check-golangci-lint
	$(PROJ_DIR)/bin/golangci-lint run --config $(PROJ_DIR)/.golangci.yaml


### Unit Tests ###

ut:
	go test -parallel=1 -race ./...
