.DEFAULT_GOAL:=help

.PHONY: build
build: clean ## Build binary
	@go build -o zombiego *.go

.PHONY: clean
clean: ## Clean up binary
	@rm -f zombiego

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Play game (via go run)
	@go run *.go
