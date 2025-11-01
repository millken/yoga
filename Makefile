# Yoga Library Makefile
# Provides convenient targets for building and managing the Yoga library

.PHONY: help build clean check-deps update-submodule build-current build-all install-headers test

# Default target
.DEFAULT_GOAL := help

# Project directories
SCRIPTS_DIR := scripts
YOGA_DIR := yoga
INCLUDE_DIR := include
LIBS_DIR := _libs

# Check if Python script is available
PYTHON_SCRIPT := $(SCRIPTS_DIR)/build_yoga.py
SHELL_SCRIPT := $(SCRIPTS_DIR)/build_yoga.sh

help: ## Show this help message
	@echo "Yoga Library Build System"
	@echo "========================="
	@echo ""
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "Usage examples:"
	@echo "  make build              # Build for current platform"
	@echo "  make check-deps         # Check dependencies"
	@echo "  make clean              # Clean build artifacts"
	@echo "  make build-all          # Build for all platforms (requires Python)"

check-deps: ## Check build dependencies
	@if command -v python3 >/dev/null 2>&1 && [ -f "$(PYTHON_SCRIPT)" ]; then \
		python3 $(PYTHON_SCRIPT) --skip-deps-check; \
	else \
		$(SHELL_SCRIPT) --check-deps; \
	fi

update-submodule: ## Update yoga submodule
	@echo "Updating yoga submodule..."
	git submodule update --init --recursive

clean: ## Clean build directories and artifacts
	@echo "Cleaning build directories..."
	rm -rf $(YOGA_DIR)/build
	@echo "Clean completed"

build-current: ## Build for current platform only
	@echo "Building Yoga library for current platform..."
	@if command -v python3 >/dev/null 2>&1 && [ -f "$(PYTHON_SCRIPT)" ]; then \
		python3 $(PYTHON_SCRIPT) --platform-only; \
	else \
		$(SHELL_SCRIPT) --current-platform; \
	fi

build-all: ## Build for all supported platforms (requires Python)
	@echo "Building Yoga library for all platforms..."
	@if command -v python3 >/dev/null 2>&1 && [ -f "$(PYTHON_SCRIPT)" ]; then \
		python3 $(PYTHON_SCRIPT) --all-platforms; \
	else \
		echo "Python3 not available. Using shell script for limited builds..."; \
		$(SHELL_SCRIPT) --all-platforms; \
	fi

build: build-current ## Alias for build-current

install-headers: ## Install header files only
	@echo "Installing header files..."
	@if [ -d "$(YOGA_DIR)/yoga" ]; then \
		rm -rf $(INCLUDE_DIR)/yoga; \
		cp -r $(YOGA_DIR)/yoga $(INCLUDE_DIR)/; \
		echo "Headers installed to $(INCLUDE_DIR)/yoga/"; \
	else \
		echo "Error: Yoga source directory not found. Run 'make update-submodule' first."; \
		exit 1; \
	fi

test: ## Run Go tests
	@echo "Running Go tests..."
	go test -v ./...

benchmark: ## Run Go benchmarks
	@echo "Running Go benchmarks..."
	go test -bench=. -benchmem ./...

lint: ## Run Go linter
	@echo "Running Go linter..."
	if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Running basic go vet..."; \
		go vet ./...; \
	fi

format: ## Format Go code
	@echo "Formatting Go code..."
	if command -v goimports >/dev/null 2>&1; then \
		goimports -w .; \
	else \
		echo "goimports not found. Running go fmt..."; \
		go fmt ./...; \
	fi

ci: check-deps build-current test lint ## Run CI pipeline (deps, build, test, lint)

dev: format build-current test ## Development workflow (format, build, test)

# Platform-specific builds
build-linux-amd64: ## Build for Linux amd64 (requires cross-compilation setup)
	@echo "Building for Linux amd64..."
	@if command -v python3 >/dev/null 2>&1 && [ -f "$(PYTHON_SCRIPT)" ]; then \
		python3 $(PYTHON_SCRIPT) --platform linux amd64; \
	else \
		echo "Platform-specific builds require Python script"; \
		exit 1; \
	fi

build-darwin-arm64: ## Build for macOS ARM64 (requires cross-compilation setup)
	@echo "Building for macOS ARM64..."
	@if command -v python3 >/dev/null 2>&1 && [ -f "$(PYTHON_SCRIPT)" ]; then \
		python3 $(PYTHON_SCRIPT) --platform darwin arm64; \
	else \
		echo "Platform-specific builds require Python script"; \
		exit 1; \
	fi

build-darwin-amd64: ## Build for macOS amd64 (requires cross-compilation setup)
	@echo "Building for macOS amd64..."
	@if command -v python3 >/dev/null 2>&1 && [ -f "$(PYTHON_SCRIPT)" ]; then \
		python3 $(PYTHON_SCRIPT) --platform darwin amd64; \
	else \
		echo "Platform-specific builds require Python script"; \
		exit 1; \
	fi

build-windows-x64: ## Build for Windows x64 (requires cross-compilation setup)
	@echo "Building for Windows x64..."
	@if command -v python3 >/dev/null 2>&1 && [ -f "$(PYTHON_SCRIPT)" ]; then \
		python3 $(PYTHON_SCRIPT) --platform windows x64; \
	else \
		echo "Platform-specific builds require Python script"; \
		exit 1; \
	fi

build-windows-amd64: ## Build for Windows amd64 (alias for x64)
	@echo "Building for Windows amd64..."
	@$(MAKE) build-windows-x64

# Docker-based cross-compilation targets
docker-build-image: ## Build Docker image for cross-compilation
	@echo "Building Docker image for cross-compilation..."
	@cd scripts && ./docker_build.sh --build-image

docker-build-linux-amd64: ## Build for Linux amd64 using Docker
	@echo "Building for Linux amd64 using Docker..."
	@cd scripts && ./docker_build.sh --platform linux-amd64

docker-build-linux-arm64: ## Build for Linux ARM64 using Docker
	@echo "Building for Linux ARM64 using Docker..."
	@cd scripts && ./docker_build.sh --platform linux-arm64

docker-build-windows-amd64: ## Build for Windows amd64 using Docker
	@echo "Building for Windows amd64 using Docker..."
	@cd scripts && ./docker_build.sh --platform windows-amd64

docker-build-all: ## Build for all platforms using Docker
	@echo "Building for all platforms using Docker..."
	@cd scripts && ./docker_build.sh --all-platforms

docker-clean: ## Remove Docker build image
	@echo "Cleaning Docker build image..."
	@cd scripts && ./docker_build.sh --clean

# Advanced targets
rebuild: clean build-current ## Clean and rebuild current platform

rebuild-all: clean build-all ## Clean and rebuild all platforms

status: ## Show build status
	@echo "Yoga Library Build Status"
	@echo "========================"
	@echo ""
	@echo "Directories:"
	@echo "  Include dir: $(INCLUDE_DIR)"
	@echo "  Libs dir:    $(LIBS_DIR)"
	@echo "  Yoga dir:    $(YOGA_DIR)"
	@echo ""
	@echo "Available libraries:"
	@if [ -d "$(LIBS_DIR)" ]; then \
		find $(LIBS_DIR) -name "*.a" -o -name "*.lib" | sort; \
	else \
		echo "  No libraries found"; \
	fi
	@echo ""
	@echo "Headers:"
	@if [ -d "$(INCLUDE_DIR)/yoga" ]; then \
		echo "  Yoga headers: $(shell find $(INCLUDE_DIR)/yoga -name "*.h" | wc -l) files"; \
	else \
		echo "  No yoga headers found"; \
	fi

# Development utilities
setup-dev: ## Set up development environment
	@echo "Setting up development environment..."
	@echo "Checking dependencies..."
	@$(MAKE) check-deps
	@echo ""
	@echo "Updating submodule..."
	@$(MAKE) update-submodule
	@echo ""
	@echo "Building library..."
	@$(MAKE) build
	@echo ""
	@echo "Running tests..."
	@$(MAKE) test
	@echo ""
	@echo "Development environment setup completed!"

version: ## Show version information
	@echo "Yoga Library Build System"
	@echo "========================"
	@echo ""
	@echo "Go version:"
	@go version
	@echo ""
	@echo "CMake version:"
	@cmake --version | head -1
	@echo ""
	@echo "Git version:"
	@git --version
	@echo ""
	@echo "Yoga submodule:"
	@if [ -d "$(YOGA_DIR)/.git" ]; then \
		cd $(YOGA_DIR) && git log -1 --oneline; \
	else \
		echo "  Submodule not initialized"; \
	fi