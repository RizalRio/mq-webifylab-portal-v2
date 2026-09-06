# 🧪 WebifyLab Portal

> **"We Experiment, You Grow"** — Digital Learning Lab untuk UMKM & Startup Indonesia

[![Build & Deploy](https://github.com/RizalRio/mq-webifylab-portal-v2/actions/workflows/deploy.yml/badge.svg)](https://github.com/RizalRio/mq-webifylab-portal-v2/actions)
[![License](https://img.shields.io/badge/license-Proprietary-blue.svg)](LICENSE)

## 📋 Overview

WebifyLab Portal adalah platform digital yang menghubungkan layanan pembuatan web & mobile app berbasis AI dengan ekosistem produk SaaS WebifyLab. Portal ini menjadi ruang eksperimen dan edukasi digital bagi UMKM dan Startup di Indonesia.

### 🎯 Goals (Phase 1 — MVP)

- ✅ Portal company profile profesional & SEO-friendly
- ✅ Custom CMS untuk manage blog, case study, portfolio
- ✅ Lead management system (mini CRM)
- ✅ Email automation via Resend
- ✅ Arsitektur modular siap untuk SaaS integration

### 🌐 Live URLs

| Service  | URL                           |
| -------- | ----------------------------- |
| Frontend | https://webifylab.my.id       |
| API      | https://api.webifylab.my.id   |
| Admin    | https://webifylab.my.id/admin |

---

## 🏗️ Tech Stack

### Frontend (`/frontend`)

- **Framework:** Next.js 15+ (App Router)
- **Language:** TypeScript 5.x
- **Styling:** Tailwind CSS + shadcn/ui
- **Animation:** Framer Motion
- **Rich Text:** Tiptap (ProseMirror)
- **HTTP Client:** Axios
- **Forms:** React Hook Form + Zod
- **Charts:** Recharts

### Backend (`/backend`)

- **Language:** Go 1.22+
- **HTTP Framework:** Gin
- **ORM:** GORM
- **Database:** PostgreSQL 15
- **Authentication:** JWT (golang-jwt/jwt v5)
- **Email:** Resend
- **Validation:** go-playground/validator
- **Logging:** Zerolog
- **Config:** Viper

### Infrastructure (`/infra`)

- **Containerization:** Docker + Docker Compose
- **Reverse Proxy:** Nginx
- **SSL:** Let's Encrypt + Certbot
- **CI/CD:** GitHub Actions
- **Monitoring:** Uptime Kuma (self-hosted)

---

## 📁 Project Structure

```
webifylab-portal/
├── frontend/              # Next.js Application
│   ├── src/
│   │   ├── app/          # Pages (App Router)
│   │   ├── components/   # React components
│   │   ├── lib/          # Utilities & API client
│   │   ├── hooks/        # Custom hooks
│   │   └── types/        # TypeScript types
│   └── Dockerfile
│
├── backend/               # Golang API
│   ├── cmd/server/       # Entry point
│   ├── internal/
│   │   ├── config/       # Configuration
│   │   ├── models/       # GORM models
│   │   ├── handlers/     # HTTP handlers
│   │   ├── services/     # Business logic
│   │   ├── repositories/ # Data access
│   │   └── middleware/   # Auth, CORS, etc
│   ├── migrations/       # Database migrations
│   └── Dockerfile
│
├── infra/                 # Infrastructure configs
│   ├── nginx/            # Nginx configurations
│   ├── postgres/         # PostgreSQL tuning
│   └── scripts/          # Deployment scripts
│
├── docs/                  # Documentation
│   ├── PRD.md
│   ├── DATABASE_SCHEMA.md
│   ├── API_SPECIFICATION.md
│   └── DEPLOYMENT_GUIDE.md
│
├── .github/workflows/     # CI/CD pipelines
├── docker-compose.yml     # Production orchestration
├── Makefile              # Common commands
└── README.md             # This file
```

---

## 🚀 Quick Start

### Prerequisites

- **Go** ≥ 1.22 ([Download](https://go.dev/dl/))
- **Node.js** ≥ 20 ([Download](https://nodejs.org/))
- **PostgreSQL** 15 ([Download](https://www.postgresql.org/download/))
- **Make** (optional, untuk convenience commands)

### 1. Clone Repository

```bash
git clone https://github.com/RizalRio/mq-webifylab-portal-v2.git
cd webifylab-portal
```

### 2. Setup Environment Variables

```bash
# Copy template
cp .env.example .env

# Edit .env dan isi nilai yang dibutuhkan
nano .env  # atau editor favorit kamu
```

**Variabel wajib:**

- `DB_PASSWORD` — Password PostgreSQL lokal
- `JWT_SECRET` — Generate dengan: `openssl rand -base64 32`

### 3. Setup Database

```bash
# Buat database dan user
psql -U postgres
CREATE USER webifylab_dev WITH PASSWORD 'dev';
CREATE DATABASE webifylab_dev OWNER webifylab_dev;
\q

# Jalankan migrations
make migrate

# Seed initial data
make seed
```

### 4. Install Dependencies

```bash
make deps
```

### 5. Start Development

**Terminal 1 — Backend:**

```bash
make dev-backend
# Server running at http://localhost:8080
```

**Terminal 2 — Frontend:**

```bash
make dev-frontend
# App running at http://localhost:3000
```

### 6. Verify

- **Health Check:** http://localhost:8080/api/v1/health
- **Frontend:** http://localhost:3000
- **Admin:** http://localhost:3000/admin

**Default Admin Credentials:**

- Email: `admin@webifylab.my.id`
- Password: `admin123`

---

## 📦 Available Commands

```bash
make help              # Lihat semua commands
make dev-backend       # Start backend (hot reload)
make dev-frontend      # Start frontend (hot reload)
make migrate           # Run database migrations
make seed              # Seed initial data
make test              # Run all tests
make build             # Build untuk production
make clean             # Clean build artifacts
```

---

## 🔄 Deployment Workflow

Deployment dilakukan otomatis via **GitHub Actions** setiap push ke branch `main`.

### Flow:

```
Laptop (Push code) → GitHub Actions (Build di cloud) → Deploy ke VPS
```

### Monitor Deployment:

https://github.com/RizalRio/mq-webifylab-portal-v2/actions

### Manual Deployment:

```bash
# Commit dan push perubahan
git add .
git commit -m "feat: your message"
git push origin main

# GitHub Actions akan otomatis:
# 1. Build Docker images
# 2. Push ke GitHub Container Registry
# 3. Deploy ke VPS via SSH
```

---

## 📚 Documentation

Dokumentasi lengkap tersedia di folder `/docs`:

| Document                                          | Description                           |
| ------------------------------------------------- | ------------------------------------- |
| [PRD.md](docs/PRD.md)                             | Product Requirements Document         |
| [DATABASE_SCHEMA.md](docs/DATABASE_SCHEMA.md)     | Database design & ERD                 |
| [API_SPECIFICATION.md](docs/API_SPECIFICATION.md) | REST API documentation (52 endpoints) |
| [DEPLOYMENT_GUIDE.md](docs/DEPLOYMENT_GUIDE.md)   | Deployment & maintenance guide        |
| [SPRINT_ROADMAP.md](docs/SPRINT_ROADMAP.md)       | Development timeline & tasks          |

---

## 🧪 API Endpoints

Total: **52 endpoints** across 11 modules

| Module         | Endpoints | Auth            |
| -------------- | --------- | --------------- |
| Authentication | 6         | Public/Required |
| Blog           | 6         | Public/Admin    |
| Case Studies   | 5         | Public/Admin    |
| Portfolio      | 6         | Public/Admin    |
| Services       | 5         | Public/Admin    |
| Leads          | 6         | Public/Admin    |
| Media          | 3         | Admin           |
| Categories     | 4         | Public/Admin    |
| Tags           | 3         | Public/Admin    |
| Dashboard      | 4         | Admin           |
| Settings       | 3         | Public/Admin    |

**Base URL:**

- Production: `https://api.webifylab.my.id/api/v1`
- Development: `http://localhost:8080/api/v1`

---

## 🔒 Security

- ✅ HTTPS everywhere (Let's Encrypt)
- ✅ JWT authentication dengan refresh token
- ✅ Password hashing (bcrypt, cost 10)
- ✅ Rate limiting (Gin middleware + Nginx)
- ✅ CORS strict (only allow webifylab.my.id)
- ✅ Input validation & SQL injection prevention (GORM)
- ✅ File upload validation (type, size)
- ✅ SSH key-based authentication (no password)
- ✅ Firewall (UFW) — only port 22, 80, 443

---

## 📊 Server Specifications

**VPS:**

- CPU: 1 vCPU
- RAM: 1 GB
- Storage: 20 GB SSD
- OS: Ubuntu 22.04 LTS
- Swap: 2 GB (configured)

**Container Memory Limits:**

- Frontend (Next.js): 350 MB
- Backend (Golang): 150 MB
- PostgreSQL: 200 MB
- **Total: ~700 MB / 1 GB**

---

## 🤝 Contributing

Project ini dikelola oleh **Rizal** (System Analyst & Fullstack Developer).

Untuk pertanyaan atau kolaborasi:

- 📧 Email: admin@webifylab.my.id
- 💬 WhatsApp: +62 812-3456-7890
- 🌐 Website: https://webifylab.my.id

---

## 📄 License

Proprietary — All rights reserved © 2026 WebifyLab

---

## 🙏 Acknowledgments

- **Gin Framework** — Fast HTTP framework
- **GORM** — Developer-friendly ORM
- **Next.js** — React framework for production
- **Tailwind CSS** — Utility-first CSS framework
- **Resend** — Email API for developers
- **Let's Encrypt** — Free SSL certificates

---

<div align="center">

**Made with 🧪 by WebifyLab**

_"We Experiment, You Grow"_

</div>
