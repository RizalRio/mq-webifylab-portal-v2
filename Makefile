# ============================================
# WebifyLab Portal - Makefile
# ============================================
# Usage: make [command]
# ============================================

.PHONY: help dev dev-backend dev-frontend build test clean migrate seed

# Default target
help:
	@echo ""
	@echo "WebifyLab Portal - Available Commands"
	@echo "========================================="
	@echo ""
	@echo "Development (Native - tanpa Docker):"
	@echo "  make dev-backend   - Start backend (Golang, hot reload)"
	@echo "  make dev-frontend  - Start frontend (Next.js, hot reload)"
	@echo "  make dev           - Start both (2 terminal tabs)"
	@echo ""
	@echo "Database:"
	@echo "  make migrate       - Run database migrations"
	@echo "  make seed          - Seed initial data"
	@echo "  make db-shell      - Open PostgreSQL shell (local)"
	@echo "  make db-reset      - Drop & recreate database (DEV ONLY)"
	@echo ""
	@echo "Testing:"
	@echo "  make test-backend  - Run backend tests"
	@echo "  make test-frontend - Run frontend tests"
	@echo "  make test          - Run all tests"
	@echo ""
	@echo "Build (untuk CI/CD):"
	@echo "  make build-backend - Build Golang binary"
	@echo "  make build-frontend- Build Next.js app"
	@echo "  make build         - Build all"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make deps          - Install all dependencies"
	@echo ""
	@echo "Documentation:"
	@echo "  make docs          - List all documentation files"
	@echo ""

# ============================================
# Development (Native)
# ============================================
dev-backend:
	@echo "Starting backend on port 8080..."
	cd backend && go run cmd/api/main.go

dev-frontend:
	@echo "Starting frontend on port 3000..."
	cd frontend && npm run dev

dev:
	@echo "Starting development environment..."
	@echo "  Buka 2 terminal terpisah:"
	@echo "   Terminal 1: make dev-backend"
	@echo "   Terminal 2: make dev-frontend"
	@echo ""
	@echo "Backend:  http://localhost:8080"
	@echo "Frontend: http://localhost:3000"
	@echo "API Docs: http://localhost:8080/api/v1/health"

# ============================================
# Database
# ============================================
migrate:
	@echo "Running migrations..."
	cd backend && go run cmd/server/main.go migrate

seed:
	@echo "Seeding initial data..."
	cd backend && go run cmd/server/main.go seed

db-shell:
	@echo "Opening PostgreSQL shell..."
	psql -U webifylab_dev -d webifylab_dev -h localhost

db-reset:
	@echo "WARNING: This will DROP and RECREATE the database!"
	@read -p "Are you sure? (y/N) " confirm && [ $$confirm = "y" ] || exit 1
	psql -U postgres -c "DROP DATABASE IF EXISTS webifylab_dev;"
	psql -U postgres -c "CREATE DATABASE webifylab_dev OWNER webifylab_dev;"
	@echo "Database reset. Run 'make migrate' to apply schema."

# ============================================
# Testing
# ============================================
test-backend:
	@echo "Running backend tests..."
	cd backend && go test ./... -v

test-frontend:
	@echo "Running frontend tests..."
	cd frontend && npm run test

test: test-backend test-frontend

# ============================================
# Build (untuk CI/CD)
# ============================================
build-backend:
	@echo "Building backend binary..."
	cd backend && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server ./cmd/server/main.go
	@echo "Binary: backend/server"

build-frontend:
	@echo "Building frontend..."
	cd frontend && npm run build
	@echo "Build: frontend/.next/"

build: build-backend build-frontend

# ============================================
# Dependencies
# ============================================
deps:
	@echo "Installing dependencies..."
	cd backend && go mod tidy
	cd frontend && npm install
	@echo "All dependencies installed"

# ============================================
# Cleanup
# ============================================
clean:
	@echo "Cleaning build artifacts..."
	rm -f backend/server
	rm -rf frontend/.next
	rm -rf frontend/out
	@echo "Cleaned"

# ============================================
# Documentation
# ============================================
docs:
	@echo "Documentation files:"
	@echo "  - docs/PRD.md"
	@echo "  - docs/DATABASE_SCHEMA.md"
	@echo "  - docs/API_SPECIFICATION.md"
	@echo "  - docs/DESIGN_SPECIFICATION.md"
	@echo "  - docs/SPRINT_ROADMAP.md"
	@echo "  - docs/DEPLOYMENT_GUIDE.md"

# ============================================
# Deployment (via GitHub Actions)
# ============================================
deploy:
	@echo "Deployment via GitHub Actions..."
	@echo "Pastikan semua perubahan sudah di-commit dan di-push ke branch 'main'"
	@echo "Monitor: https://github.com/RizalRio/mq-webifylab-portal-v2/actions"
	git push origin main