.PHONY: setup dev build test deploy clean help

# ========================================
# Webifylab Monorepo - Root Makefile
# ========================================

help: ## Show this help
	@echo "Webifylab Monorepo Commands:"
	@echo ""
	@echo "  make setup       - Install all dependencies"
	@echo "  make dev-web     - Start Astro dev server"
	@echo "  make dev-api     - Start Go API server (V1.5)"
	@echo "  make build-web   - Build Astro for production"
	@echo "  make build-api   - Build Go binary"
	@echo "  make test-web    - Run web tests"
	@echo "  make test-api    - Run API tests"
	@echo "  make deploy      - Deploy to VPS"
	@echo "  make clean       - Clean build artifacts"
	@echo ""

setup: ## Install all dependencies
	cd apps/web && npm install
	cd apps/api && go mod download

dev-web: ## Start Astro dev server
	cd apps/web && npm run dev

dev-api: ## Start Go API server
	cd apps/api && go run cmd/server/main.go

build-web: ## Build Astro for production
	cd apps/web && npm run build

build-api: ## Build Go binary
	cd apps/api && go build -o webifylab-api cmd/server/main.go

test-web: ## Run web tests
	cd apps/web && npm run test

test-api: ## Run API tests
	cd apps/api && go test ./...

deploy: ## Deploy to VPS
	@echo "Run deploy script..."
	powershell -File infra/scripts/deploy.ps1

clean: ## Clean build artifacts
	rm -rf apps/web/dist
	rm -rf apps/web/.astro
	rm -f apps/api/webifylab-api
	rm -f apps/api/webifylab-api.exe
	@echo "Cleaned!"
