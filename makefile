SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage


run: ## Start application
	$(GOCMD) run ./cmd -b 0.0.0.0
