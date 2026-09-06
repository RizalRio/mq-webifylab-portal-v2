# 📄 PRODUCT REQUIREMENTS DOCUMENT (PRD)
## WebifyLab Portal — Phase 1 (MVP)

---

### 1. DOCUMENT INFORMATION

| Field | Detail |
|---|---|
| **Product Name** | WebifyLab Portal |
| **Domain (Frontend)** | webifylab.my.id |
| **Domain (API)** | api.webifylab.my.id |
| **Tagline** | "We Experiment, You Grow" |
| **Document Version** | 1.0.0 (Final) |
| **Date** | 05 September 2026 |
| **Author** | Rizal (System Analyst & Fullstack Developer) |
| **Status** | ✅ Approved — Phase 1 MVP |
| **Architecture** | Golang (Gin + GORM) + Next.js + PostgreSQL + Docker + Nginx |
| **Repository** | Monorepo (frontend + backend + infra + docs) |
| **Email Service** | Resend (✅ Akun sudah tersedia) |

---

### 2. PRODUCT OVERVIEW

#### 2.1 Vision
Menjadi **digital learning lab** yang menghubungkan layanan pembuatan web & mobile app berbasis AI dengan ekosistem produk SaaS WebifyLab, serta menjadi ruang eksperimen dan edukasi digital bagi UMKM dan Startup di Indonesia.

#### 2.2 Brand Philosophy
> **"We Experiment, You Grow"**

WebifyLab bukan sekadar vendor yang jual solusi jadi. Kami adalah **lab** — tempat di mana setiap project adalah eksperimen, setiap tantangan adalah pelajaran, dan setiap klien adalah partner belajar.

**Pilar Filosofi:**
- **Transparansi proses** — dokumentasikan journey (case study, blog) apa adanya, termasuk kegagalan dan pembelajaran.
- **Kolaboratif** — klien diajak memahami teknologi yang dipakai (terutama AI/SPK).
- **Iteratif** — MVP dulu, lalu improve berdasarkan data dan feedback.
- **Edukatif** — setiap konten bertujuan membagi ilmu, bukan sekadar jualan.

#### 2.3 Brand Identity

**Color Palette:**
| Role | Color | Hex | Usage |
|---|---|---|---|
| Primary 1 | Lab Blue | `#2563EB` | Primary CTA, links, brand accent |
| Primary 2 | Experiment Orange | `#F97316` | Secondary CTA, highlights, badges |
| Secondary 1 | Growth Green | `#10B981` | Success states, metrics, results |
| Secondary 2 | Deep Slate | `#0F172A` | Text, headings, contrast |
| Secondary 3 | Light Gray | `#F1F5F9` | Backgrounds, cards, whitespace |

**Typography:**
- **Headings:** Inter / Plus Jakarta Sans
- **Body:** Inter / Open Sans
- **Code:** JetBrains Mono

**Tone of Voice:**
- ✅ Clear, simple, approachable
- ✅ Edukatif tanpa menggurui
- ✅ Transparan tentang proses dan lessons learned
- ✅ Kolaboratif ("mari kita eksperimen bareng")
- ❌ Hindari jargon berlebihan tanpa penjelasan
- ❌ Hindari janji "solusi ajaib"

#### 2.4 Mission
1. Menyediakan layanan pengembangan web dan mobile app yang terintegrasi dengan algoritma AI (SPK, Sistem Pakar, Data Analytics) melalui pendekatan **eksperimen dan belajar** bersama UMKM dan Startup.
2. Membangun ekosistem produk SaaS yang saling terhubung melalui satu portal terpusat, dengan setiap produk sebagai hasil eksperimen nyata.
3. Mendokumentasikan dan mengedukasi pasar melalui konten berkualitas (blog, case study) tentang transformasi digital dan pemanfaatan AI dalam bisnis.

#### 2.5 Product Goals (Phase 1 — 3 Bulan Pertama)

| No | Goal | Success Metric |
|---|---|---|
| G1 | Meluncurkan portal company profile yang profesional, SEO-friendly, dan mencerminkan filosofi "We Experiment, You Grow" | Lighthouse score ≥ 90 (Performance, SEO, Accessibility) |
| G2 | Mendapatkan leads/inquiry dari calon klien UMKM dan Startup | ≥ 10 inquiry/bulan setelah launch |
| G3 | Membangun content engine untuk blog & case study sebagai dokumentasi journey belajar | ≥ 4 artikel/bulan terpublikasi |
| G4 | Menyiapkan arsitektur yang siap untuk integrasi SaaS di masa depan | Arsitektur modular terdokumentasi, SSO-ready, reusable untuk SaaS |

---

### 3. TARGET USERS & PERSONAS

#### 3.1 External Users (Publik)

| Persona | Deskripsi | Kebutuhan Utama | Pain Points |
|---|---|---|---|
| **Budi — Pemilik UMKM** | Pemilik toko/kedai, usia 30-45, melek digital dasar, ingin go-digital | Website affordable, bisa membantu operasional (SPK untuk stok), mudah di-maintain, ada yang mau jelasin teknologi dengan bahasa sederhana | Budget terbatas, bingung mulai dari mana, takut teknologi rumit, pernah dijanjikan "solusi ajaib" |
| **Sari — Co-founder Startup** | Founder startup early-stage, usia 24-32, tech-savvy | MVP cepat, scalable, ada fitur AI/data untuk competitive advantage, partner yang mau diajak brainstorming | Butuh tim dev yang paham AI, deadline ketat, butuh partner bukan sekadar vendor |
| **Andi — Marketing Manager** | Bekerja di perusahaan menengah, mencari vendor untuk project digital | Portfolio meyakinkan, case study terukur, proses transparan dan edukatif | Pernah kecewa dengan vendor sebelumnya, butuh bukti hasil nyata |

#### 3.2 Internal Users

| Persona | Deskripsi | Kebutuhan Utama |
|---|---|---|
| **Rizal — Admin / Owner** | Fullstack dev & system analyst, mengelola seluruh konten dan leads | CMS yang cepat dan intuitif, dashboard leads, analytics sederhana |
| **Future: Content Writer** | Tim/kontributor yang menulis blog dan case study | Akses terbatas ke CMS blog, draft & publish workflow |

---

### 4. PRE-REQUISITES & MIGRATION

#### 4.1 Pre-Requisites (Sebelum Development Dimulai)

| No | Task | Status | Owner |
|---|---|---|---|
| P1 | Backup previous deployment data (Golang + Next.js) | ✅ Done | Rizal |
| P2 | Hapus previous deployment dari VPS (stop containers, remove images, cleanup data) | ⏳ Pending | Rizal |
| P3 | Setup DNS A record untuk `api.webifylab.my.id` → IP VPS | ⏳ Pending | Rizal |
| P4 | Setup akun Resend (verifikasi domain, API key) | ✅ Done | Rizal |
| P5 | Setup SSL certificate untuk `api.webifylab.my.id` via Certbot | ⏳ Pending | Rizal |
| P6 | Setup **monorepo repository** (structure, git, README) | ⏳ Pending | Rizal |
| P7 | Setup environment variables template (.env.example di root) | ⏳ Pending | Rizal |
| P8 | Setup Makefile untuk common commands | ⏳ Pending | Rizal |

#### 4.2 Migration Checklist (Previous → New Deployment)

```bash
# 1. Stop previous containers
docker-compose -f /path/to/old/deployment/docker-compose.yml down

# 2. Backup old database (jika perlu dipelajari/dimigrasi)
docker exec old-postgres pg_dump -U user dbname > old_backup.sql

# 3. Remove old containers & images
docker container prune
docker image prune

# 4. Cleanup old volumes (HATI-HATI: pastikan backup sudah aman)
docker volume rm old_pgdata old_uploads

# 5. Cleanup old files
rm -rf /var/www/old-webifylab/*

# 6. Verify VPS clean state
docker ps  # harus kosong
df -h      # cek storage
free -h    # cek RAM
```

---

### 5. SCOPE & PHASING

#### 5.1 Phase 1 — MVP (Current Scope) ✅
Fokus: **Portal Marketing + Custom CMS + Lead Management**

| Modul | Fitur |
|---|---|
| **Public Pages (Next.js)** | Home, About, Services, Portfolio, Case Studies, Blog, Contact, Products/Ecosystem (placeholder) |
| **Backend API (Golang + Gin + GORM)** | REST API untuk CMS, Lead Management, Media, Authentication |
| **CMS Frontend (Next.js)** | Admin dashboard untuk manage Blog, Case Study, Portfolio, Services, Leads, Media |
| **Email Service** | Resend untuk auto-reply dan notifikasi admin |
| **SEO & Performance** | SSR/SSG, Sitemap, Meta tags, Open Graph, Structured Data |
| **Language** | Bahasa Indonesia saja (multi-language di Phase 2+) |

#### 5.2 Phase 2 — SaaS Integration (Future) 🔮
| Modul | Fitur |
|---|---|
| **SSO Gateway (Golang)** | Unified login dari portal ke semua SaaS (OAuth2/OIDC) |
| **Client Portal** | Klien login untuk lihat status project, invoice, akses SaaS |
| **Product Marketplace** | Katalog SaaS dengan pricing, trial, dan onboarding |
| **Multi-language** | Support Bahasa Indonesia & English (i18n) |

#### 5.3 Phase 3 — Scale (Future) 🔮
| Modul | Fitur |
|---|---|
| **Multi-tenant Dashboard** | Dashboard terpusat untuk semua SaaS metrics |
| **AI Chatbot** | Chatbot untuk lead qualification di portal |
| **API Gateway** | Centralized API untuk semua SaaS |

---

### 6. MONOREPO STRUCTURE

#### 6.1 Repository Structure

```
webifylab-portal/                    # Root monorepo
│
├── 📄 README.md                     # Project overview, setup guide
├── 📄 .gitignore                    # Global gitignore
├── 📄 .env.example                  # Environment variables template
├── 📄 docker-compose.yml            # Orchestrate all services
├── 📄 docker-compose.dev.yml        # Development overrides
├── 📄 Makefile                      # Common commands (build, deploy, etc)
├── 📄 LICENSE                       # Optional
│
├── 📁 frontend/                     # Next.js Application
│   ├── 📄 README.md
│   ├── 📄 Dockerfile
│   ├── 📄 package.json
│   ├── 📄 tsconfig.json
│   ├── 📄 tailwind.config.ts
│   ├── 📄 next.config.ts
│   ├── 📄 .env.example
│   ├── 📁 src/
│   │   ├── 📁 app/                  # Next.js App Router
│   │   │   ├── 📁 (public)/         # Public pages (home, about, blog, etc)
│   │   │   ├── 📁 admin/            # Admin dashboard
│   │   │   ├── 📁 api/              # Next.js API routes (optional)
│   │   │   ├── 📄 layout.tsx
│   │   │   └── 📄 page.tsx
│   │   ├── 📁 components/           # React components
│   │   │   ├── 📁 ui/               # shadcn/ui components
│   │   │   ├── 📁 public/           # Public page components
│   │   │   └── 📁 admin/            # Admin components
│   │   ├── 📁 lib/                  # Utilities, helpers
│   │   │   ├── 📄 api.ts            # API client (axios)
│   │   │   ├── 📄 utils.ts
│   │   │   └── 📄 constants.ts
│   │   ├── 📁 hooks/                # Custom React hooks
│   │   ├── 📁 types/                # TypeScript types
│   │   └── 📁 styles/               # Global styles
│   └── 📁 public/                   # Static assets
│
├── 📁 backend/                      # Golang Application (Gin + GORM)
│   ├── 📄 README.md
│   ├── 📄 Dockerfile
│   ├── 📄 go.mod
│   ├── 📄 go.sum
│   ├── 📄 .env.example
│   ├── 📁 cmd/
│   │   └── 📁 server/
│   │       └── 📄 main.go           # Entry point
│   ├── 📁 internal/
│   │   ├── 📁 config/               # Viper config loader
│   │   ├── 📁 models/               # GORM models
│   │   ├── 📁 handlers/             # Gin handlers (controllers)
│   │   ├── 📁 services/             # Business logic
│   │   ├── 📁 repositories/         # Data access (GORM)
│   │   ├── 📁 middleware/           # Gin middleware (auth, cors, etc)
│   │   ├── 📁 dto/                  # Request/Response DTOs
│   │   │   ├── 📁 request/
│   │   │   └── 📁 response/
│   │   └── 📁 routes/               # Route registration
│   ├── 📁 migrations/               # golang-migrate files
│   ├── 📁 pkg/                      # Shared utilities
│   │   ├── 📁 response/             # Standard response format
│   │   ├── 📁 validator/            # Validation helpers
│   │   └── 📁 hasher/               # Password hashing
│   └── 📁 uploads/                  # Uploaded files (gitignored)
│
├── 📁 infra/                        # Infrastructure configs
│   ├── 📁 nginx/
│   │   ├── 📄 webifylab.conf        # Nginx config (frontend)
│   │   └── 📄 api.conf              # Nginx config (backend API)
│   ├── 📁 postgres/
│   │   └── 📄 postgresql.conf       # PostgreSQL tuning
│   └── 📁 scripts/
│       ├── 📄 deploy.sh             # Deployment script
│       ├── 📄 backup.sh             # Backup script
│       └── 📄 restore.sh            # Restore script
│
├── 📁 docs/                         # Documentation
│   ├── 📄 PRD.md                    # This document
│   ├── 📄 ERD.md                    # Database schema
│   ├── 📄 API.md                    # API specification
│   ├── 📄 DEPLOYMENT.md             # Deployment guide
│   └── 📁 diagrams/                 # Architecture diagrams
│
└── 📁 .github/                      # GitHub workflows (optional)
    └── 📁 workflows/
        └── 📄 ci.yml
```

#### 6.2 Monorepo Benefits

| Benefit | Description |
|---|---|
| **Single Source of Truth** | Semua configs (Docker, Nginx, env) di satu tempat |
| **Atomic Commits** | Frontend + backend changes dalam satu commit |
| **Shared Documentation** | PRD, ERD, API spec terpusat di `/docs` |
| **Simplified Deployment** | Satu `docker-compose.yml` orchestrate semua services |
| **Easy Onboarding** | New developer cukup clone satu repo, baca README |
| **Version Consistency** | Semua components versioned together |
| **Reusable Scripts** | Makefile & scripts bisa dipakai untuk SaaS berikutnya |

#### 6.3 Makefile (Root)

```makefile
# WebifyLab Portal - Makefile
# Common commands for development and deployment

.PHONY: help dev build deploy up down logs clean

# Default target
help:
	@echo "WebifyLab Portal - Available Commands"
	@echo ""
	@echo "Development:"
	@echo "  make dev          - Start development environment"
	@echo "  make dev-frontend - Start frontend only (hot reload)"
	@echo "  make dev-backend  - Start backend only (hot reload)"
	@echo ""
	@echo "Build:"
	@echo "  make build        - Build all Docker images"
	@echo "  make build-frontend - Build frontend image only"
	@echo "  make build-backend  - Build backend image only"
	@echo ""
	@echo "Deployment:"
	@echo "  make deploy       - Deploy to VPS (build + upload + start)"
	@echo "  make deploy-images - Build and save images for upload"
	@echo "  make upload       - Upload images to VPS"
	@echo "  make start        - Start containers on VPS"
	@echo ""
	@echo "Management:"
	@echo "  make up           - Start all containers"
	@echo "  make down         - Stop all containers"
	@echo "  make restart      - Restart all containers"
	@echo "  make logs         - View logs (all services)"
	@echo "  make logs-backend - View backend logs"
	@echo "  make logs-frontend - View frontend logs"
	@echo ""
	@echo "Database:"
	@echo "  make migrate      - Run database migrations"
	@echo "  make seed         - Seed initial data"
	@echo "  make db-shell     - Open PostgreSQL shell"
	@echo "  make backup       - Backup database + uploads"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean        - Remove containers, images, volumes"
	@echo "  make prune        - Prune unused Docker resources"

# Development
dev:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build

dev-frontend:
	cd frontend && npm run dev

dev-backend:
	cd backend && go run cmd/server/main.go

# Build
build:
	docker-compose build

build-frontend:
	docker build -t webifylab-frontend -f frontend/Dockerfile frontend/

build-backend:
	docker build -t webifylab-backend -f backend/Dockerfile backend/

# Deployment
deploy: build-frontend build-backend deploy-images upload start

deploy-images:
	docker save webifylab-frontend | gzip > frontend.tar.gz
	docker save webifylab-backend | gzip > backend.tar.gz

upload:
	scp frontend.tar.gz backend.tar.gz docker-compose.yml user@vps:/var/www/webifylab/
	scp -r infra/ user@vps:/var/www/webifylab/

start:
	ssh user@vps 'cd /var/www/webifylab && docker load -i frontend.tar.gz && docker load -i backend.tar.gz && docker-compose up -d'

# Management
up:
	docker-compose up -d

down:
	docker-compose down

restart:
	docker-compose restart

logs:
	docker-compose logs -f

logs-backend:
	docker-compose logs -f backend

logs-frontend:
	docker-compose logs -f frontend

# Database
migrate:
	docker-compose exec backend ./server migrate

seed:
	docker-compose exec backend ./server seed

db-shell:
	docker-compose exec postgres psql -U webifylab

backup:
	./infra/scripts/backup.sh

# Cleanup
clean:
	docker-compose down -v --rmi all

prune:
	docker system prune -af
```

#### 6.4 Root docker-compose.yml

```yaml
# docker-compose.yml (Root monorepo)
version: '3.8'

services:
  # Next.js Frontend
  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    container_name: webifylab-frontend
    restart: unless-stopped
    environment:
      - NODE_ENV=production
      - NEXT_PUBLIC_API_URL=${API_URL:-https://api.webifylab.my.id}
      - NEXT_PUBLIC_SITE_URL=${SITE_URL:-https://webifylab.my.id}
    networks:
      - webifylab-network
    deploy:
      resources:
        limits:
          memory: 350M

  # Golang Backend (Gin + GORM)
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    container_name: webifylab-backend
    restart: unless-stopped
    env_file:
      - .env
    environment:
      - APP_ENV=${APP_ENV:-production}
      - APP_PORT=8080
      - APP_HOST=0.0.0.0
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=${DB_USER:-webifylab}
      - DB_PASSWORD=${DB_PASSWORD}
      - DB_NAME=${DB_NAME:-webifylab}
      - DB_SSLMODE=disable
      - JWT_SECRET=${JWT_SECRET}
      - JWT_EXPIRY=24h
      - REFRESH_TOKEN_EXPIRY=168h
      - RESEND_API_KEY=${RESEND_API_KEY}
      - ADMIN_EMAIL=${ADMIN_EMAIL}
      - FRONTEND_URL=${SITE_URL:-https://webifylab.my.id}
      - CORS_ORIGIN=${SITE_URL:-https://webifylab.my.id}
      - UPLOAD_DIR=/app/uploads
      - MAX_UPLOAD_SIZE=5242880
    depends_on:
      postgres:
        condition: service_healthy
    volumes:
      - uploads:/app/uploads
    networks:
      - webifylab-network
    deploy:
      resources:
        limits:
          memory: 150M

  # PostgreSQL Database
  postgres:
    image: postgres:15-alpine
    container_name: webifylab-postgres
    restart: unless-stopped
    environment:
      - POSTGRES_DB=${DB_NAME:-webifylab}
      - POSTGRES_USER=${DB_USER:-webifylab}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
    volumes:
      - pgdata:/var/lib/postgresql/data
      - ./infra/postgres/postgresql.conf:/etc/postgresql/postgresql.conf
    command: postgres -c config_file=/etc/postgresql/postgresql.conf
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${DB_USER:-webifylab}"]
      interval: 10s
      timeout: 5s
      retries: 5
    networks:
      - webifylab-network
    deploy:
      resources:
        limits:
          memory: 200M

networks:
  webifylab-network:
    driver: bridge

volumes:
  pgdata:
    driver: local
  uploads:
    driver: local
```

#### 6.5 Root .env.example

```bash
# ===================================
# WebifyLab Portal - Environment Variables
# ===================================
# Copy this file to .env and fill in the values
# cp .env.example .env

# -----------------------------------
# Application
# -----------------------------------
APP_ENV=production
SITE_URL=https://webifylab.my.id
API_URL=https://api.webifylab.my.id

# -----------------------------------
# Database (PostgreSQL)
# -----------------------------------
DB_USER=webifylab
DB_PASSWORD=your_secure_password_here_min_16_chars
DB_NAME=webifylab

# -----------------------------------
# JWT Authentication
# -----------------------------------
# Generate with: openssl rand -base64 32
JWT_SECRET=your_jwt_secret_here_min_32_chars

# -----------------------------------
# Resend Email Service
# -----------------------------------
# Get from: https://resend.com/api-keys
RESEND_API_KEY=re_your_resend_api_key_here
ADMIN_EMAIL=admin@webifylab.my.id

# -----------------------------------
# Optional: Upload Limits
# -----------------------------------
MAX_UPLOAD_SIZE=5242880  # 5MB in bytes
```

#### 6.6 Development Workflow (Local)

**Option A: Full Docker (Recommended untuk testing production-like)**
```bash
# Di laptop (16GB RAM)
cd webifylab-portal
make dev  # Start all services with hot reload
```

**Option B: Native Development (Lebih cepat untuk development)**
```bash
# Terminal 1: PostgreSQL (via Docker)
docker-compose up postgres

# Terminal 2: Backend (Golang native)
cd backend
go run cmd/server/main.go

# Terminal 3: Frontend (Next.js native)
cd frontend
npm run dev
```

#### 6.7 Deployment Workflow (Production)

```bash
# Di laptop (16GB RAM)

# 1. Build Docker images
make build

# 2. Save images to tar files
make deploy-images

# 3. Upload ke VPS
make upload

# 4. Di VPS, start containers
make start

# 5. Run migrations & seed
ssh user@vps 'cd /var/www/webifylab && docker-compose exec backend ./server migrate && docker-compose exec backend ./server seed'

# 6. Verify
curl https://api.webifylab.my.id/api/v1/health
```

#### 6.8 Shared Resources

**Shared Types (Optional - Future Enhancement):**
Untuk menjaga konsistensi antara frontend dan backend, bisa generate TypeScript types dari Golang structs menggunakan tools seperti:
- OpenAPI/Swagger → TypeScript codegen
- Atau manual sync via `/shared/types/` folder

**Shared Constants:**
- Status enums (lead status, post status)
- Validation rules
- Error codes

Dokumentasikan di `/docs/CONTRACTS.md` untuk referensi bersama.

---

### 7. FUNCTIONAL REQUIREMENTS (Detail)

#### 7.1 Modul Publik — Public Pages (Next.js Frontend)

##### FR-1.1: Home Page
| ID | Requirement | Priority |
|---|---|---|
| FR-1.1.1 | Hero section dengan headline "We Experiment, You Grow", sub-headline, CTA button ("Mulai Eksperimen", "Lihat Portfolio") | P0 |
| FR-1.1.2 | Section "Layanan Unggulan" menampilkan 3-4 layanan utama dengan ikon dan deskripsi singkat | P0 |
| FR-1.1.3 | Section "Kenapa WebifyLab?" — USP: integrasi AI (SPK, Sistem Pakar, Data Analytics) + pendekatan kolaboratif & edukatif | P0 |
| FR-1.1.4 | Section "Eksperimen Terbaru" — showcase 3-6 project/case study terbaru | P0 |
| FR-1.1.5 | Section "Catatan Belajar" — 3 blog post terbaru (tone: sharing, bukan jualan) | P1 |
| FR-1.1.6 | Section "Client Logos / Testimonials" | P1 |
| FR-1.1.7 | Section CTA akhir — "Siap Eksperimen Bareng?" dengan form singkat atau link ke contact | P0 |
| FR-1.1.8 | Responsive design (mobile-first) | P0 |

##### FR-1.2: About Page
| ID | Requirement | Priority |
|---|---|---|
| FR-1.2.1 | Cerita/visi WebifyLab dengan narasi "We Experiment, You Grow" | P0 |
| FR-1.2.2 | Filosofi kerja: transparan, kolaboratif, iteratif, edukatif | P0 |
| FR-1.2.3 | Tim section (foto, nama, role, keahlian, "apa yang sedang saya pelajari sekarang") | P1 |
| FR-1.2.4 | Tech stack yang digunakan (visual badges) | P1 |
| FR-1.2.5 | Timeline/milestone perusahaan — fokus pada "apa yang kami pelajari" di setiap fase | P2 |

##### FR-1.3: Services Page
| ID | Requirement | Priority |
|---|---|---|
| FR-1.3.1 | Daftar layanan: Web Development, Mobile App Development, UI/UX Design, AI Integration (SPK, Sistem Pakar, Data Analytics) | P0 |
| FR-1.3.2 | Setiap layanan memiliki: deskripsi, use case untuk UMKM/Startup, estimasi timeline, tech stack, **"apa yang akan kita pelajari bersama"** section | P0 |
| FR-1.3.3 | CTA per layanan — "Diskusikan Eksperimen Ini" → contact form dengan pre-filled service | P0 |
| FR-1.3.4 | Section "Bagaimana Kami Bekerja" — workflow (Discovery → Design → Dev → Testing → Deploy → Learn & Iterate) | P1 |

##### FR-1.4: Portfolio Page
| ID | Requirement | Priority |
|---|---|---|
| FR-1.4.1 | Grid/gallery project dengan filter berdasarkan kategori (Web, Mobile, AI, E-commerce, dll) | P0 |
| FR-1.4.2 | Setiap item menampilkan: thumbnail, judul, kategori, tech stack badges | P0 |
| FR-1.4.3 | Detail page per project: overview, challenge, solution, tech stack, screenshots, link live, **"lessons learned"** section | P0 |
| FR-1.4.4 | Related projects di bagian bawah detail page | P2 |

##### FR-1.5: Case Studies Page
| ID | Requirement | Priority |
|---|---|---|
| FR-1.5.1 | List case study dengan format card: judul, klien (industri), hasil utama (angka/metrics) | P0 |
| FR-1.5.2 | Detail case study dengan struktur: **Client Profile → Challenge → Approach → Solution (termasuk algoritma AI) → Results → Lessons Learned → Testimonial** | P0 |
| FR-1.5.3 | Tag/filter berdasarkan industri dan jenis AI (SPK, Sistem Pakar, Data Analytics) | P1 |
| FR-1.5.4 | CTA "Ingin eksperimen serupa?" di setiap case study | P0 |

##### FR-1.6: Blog / Articles Page
| ID | Requirement | Priority |
|---|---|---|
| FR-1.6.1 | Blog listing dengan pagination (10 post/page) | P0 |
| FR-1.6.2 | Filter by category (Tutorial, Business, AI & Data, Eksperimen, Lessons Learned) | P0 |
| FR-1.6.3 | Search bar untuk artikel | P1 |
| FR-1.6.4 | Detail artikel: title, author, date, read time, cover image, rich text content, tags | P0 |
| FR-1.6.5 | Table of Contents (auto-generated dari headings) | P1 |
| FR-1.6.6 | Code syntax highlighting (karena artikel teknis) | P1 |
| FR-1.6.7 | Related articles di akhir post | P2 |
| FR-1.6.8 | Share buttons (LinkedIn, Twitter/X, WhatsApp) | P1 |
| FR-1.6.9 | Reading progress bar | P2 |

##### FR-1.7: Contact Page
| ID | Requirement | Priority |
|---|---|---|
| FR-1.7.1 | Form inquiry: Nama, Email, No. WhatsApp, Jenis Layanan (dropdown), Budget Range (dropdown), Deskripsi Project (textarea) | P0 |
| FR-1.7.2 | Validasi form (client-side + server-side via Golang API) | P0 |
| FR-1.7.3 | Success message setelah submit + auto-reply email via Resend ke user | P0 |
| FR-1.7.4 | Notifikasi email via Resend ke admin saat ada inquiry baru | P0 |
| FR-1.7.5 | Informasi kontak: email, WhatsApp (link langsung ke chat), social media links | P0 |
| FR-1.7.6 | Embedded map (Google Maps / OpenStreetMap) — opsional | P2 |

##### FR-1.8: Products / Ecosystem Page (Placeholder)
| ID | Requirement | Priority |
|---|---|---|
| FR-1.8.1 | Halaman yang menampilkan daftar produk SaaS WebifyLab | P1 |
| FR-1.8.2 | Setiap produk: nama, deskripsi singkat, status (Live/Coming Soon/Experimenting), link ke subdomain | P1 |
| FR-1.8.3 | Section "Eksperimen Selanjutnya" untuk membangun anticipation | P2 |

---

#### 7.2 Modul Backend — Golang REST API (Gin + GORM)

##### FR-2.1: Authentication API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.1.1 | `POST /api/v1/auth/login` — Login dengan email & password, return JWT | P0 |
| FR-2.1.2 | `POST /api/v1/auth/logout` — Invalidate token (optional, client-side sufficient) | P1 |
| FR-2.1.3 | `POST /api/v1/auth/refresh` — Refresh JWT token | P1 |
| FR-2.1.4 | `POST /api/v1/auth/forgot-password` — Kirim email reset password via Resend | P1 |
| FR-2.1.5 | `POST /api/v1/auth/reset-password` — Reset password dengan token | P1 |
| FR-2.1.6 | `GET /api/v1/auth/me` — Get current user profile | P0 |
| FR-2.1.7 | JWT middleware untuk protected routes (Gin middleware) | P0 |
| FR-2.1.8 | Role-based access control: `super_admin`, `admin`, `editor` | P1 |

##### FR-2.2: Blog API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.2.1 | `GET /api/v1/blog` — List blog posts (pagination, filter by category, search) | P0 |
| FR-2.2.2 | `GET /api/v1/blog/:slug` — Get blog post by slug | P0 |
| FR-2.2.3 | `POST /api/v1/blog` — Create blog post (admin) | P0 |
| FR-2.2.4 | `PUT /api/v1/blog/:id` — Update blog post (admin) | P0 |
| FR-2.2.5 | `DELETE /api/v1/blog/:id` — Delete blog post (admin) | P0 |
| FR-2.2.6 | Support status: draft, published, archived | P0 |
| FR-2.2.7 | Auto-generate slug from title | P0 |

##### FR-2.3: Case Study API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.3.1 | `GET /api/v1/case-studies` — List case studies (pagination, filter by industry/ai_type) | P0 |
| FR-2.3.2 | `GET /api/v1/case-studies/:slug` — Get case study by slug | P0 |
| FR-2.3.3 | `POST /api/v1/case-studies` — Create case study (admin) | P0 |
| FR-2.3.4 | `PUT /api/v1/case-studies/:id` — Update case study (admin) | P0 |
| FR-2.3.5 | `DELETE /api/v1/case-studies/:id` — Delete case study (admin) | P0 |
| FR-2.3.6 | Field terstruktur: client_name, industry, challenge, approach, solution, ai_type, results, lessons, testimonial, metrics | P0 |

##### FR-2.4: Portfolio API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.4.1 | `GET /api/v1/portfolio` — List portfolio items (pagination, filter by category) | P0 |
| FR-2.4.2 | `GET /api/v1/portfolio/:slug` — Get portfolio item by slug | P0 |
| FR-2.4.3 | `POST /api/v1/portfolio` — Create portfolio item (admin) | P0 |
| FR-2.4.4 | `PUT /api/v1/portfolio/:id` — Update portfolio item (admin) | P0 |
| FR-2.4.5 | `DELETE /api/v1/portfolio/:id` — Delete portfolio item (admin) | P0 |
| FR-2.4.6 | Support tech stack multi-select, ordering | P1 |

##### FR-2.5: Services API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.5.1 | `GET /api/v1/services` — List services | P0 |
| FR-2.5.2 | `GET /api/v1/services/:slug` — Get service by slug | P0 |
| FR-2.5.3 | `POST /api/v1/services` — Create service (admin) | P0 |
| FR-2.5.4 | `PUT /api/v1/services/:id` — Update service (admin) | P0 |
| FR-2.5.5 | `DELETE /api/v1/services/:id` — Delete service (admin) | P0 |

##### FR-2.6: Leads API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.6.1 | `POST /api/v1/leads` — Submit inquiry (public, rate-limited) | P0 |
| FR-2.6.2 | `GET /api/v1/leads` — List leads (admin, pagination, filter by status) | P0 |
| FR-2.6.3 | `GET /api/v1/leads/:id` — Get lead detail (admin) | P0 |
| FR-2.6.4 | `PUT /api/v1/leads/:id` — Update lead status/notes (admin) | P0 |
| FR-2.6.5 | `DELETE /api/v1/leads/:id` — Delete lead (admin) | P1 |
| FR-2.6.6 | Trigger auto-reply email via Resend saat lead baru dibuat | P0 |
| FR-2.6.7 | Trigger notification email ke admin via Resend | P0 |
| FR-2.6.8 | Rate limiting: max 5 submissions per IP per hour (Gin middleware) | P0 |

##### FR-2.7: Media API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.7.1 | `POST /api/v1/media/upload` — Upload image (admin, max 5MB) | P0 |
| FR-2.7.2 | `GET /api/v1/media` — List media (admin, pagination) | P1 |
| FR-2.7.3 | `DELETE /api/v1/media/:id` — Delete media (admin) | P0 |
| FR-2.7.4 | Auto-compress & generate WebP version | P1 |
| FR-2.7.5 | Validate file type (JPG, PNG, WebP only) | P0 |
| FR-2.7.6 | Serve static files via Nginx (bukan lewat Golang) | P0 |

##### FR-2.8: Categories & Tags API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.8.1 | `GET /api/v1/categories` — List categories (filter by type: blog/portfolio) | P0 |
| FR-2.8.2 | `POST /api/v1/categories` — Create category (admin) | P0 |
| FR-2.8.3 | `PUT /api/v1/categories/:id` — Update category (admin) | P0 |
| FR-2.8.4 | `DELETE /api/v1/categories/:id` — Delete category (admin) | P0 |
| FR-2.8.5 | `GET /api/v1/tags` — List tags | P0 |
| FR-2.8.6 | `POST /api/v1/tags` — Create tag (admin) | P0 |

##### FR-2.9: Dashboard API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.9.1 | `GET /api/v1/dashboard/stats` — Get overview stats (total blog, case study, portfolio, leads this month) | P0 |
| FR-2.9.2 | `GET /api/v1/dashboard/leads-chart` — Get leads per month (last 12 months) | P1 |
| FR-2.9.3 | `GET /api/v1/dashboard/recent-leads` — Get 5 recent leads | P0 |
| FR-2.9.4 | `GET /api/v1/dashboard/recent-posts` — Get 5 recent blog posts | P1 |

##### FR-2.10: Settings API
| ID | Requirement | Priority |
|---|---|---|
| FR-2.10.1 | `GET /api/v1/settings` — Get site settings (SEO, contact info) | P0 |
| FR-2.10.2 | `PUT /api/v1/settings` — Update site settings (admin) | P0 |

##### FR-2.11: API Standards
| ID | Requirement | Priority |
|---|---|---|
| FR-2.11.1 | Consistent JSON response format: `{ success, data, message, errors }` | P0 |
| FR-2.11.2 | Proper HTTP status codes (200, 201, 400, 401, 403, 404, 500) | P0 |
| FR-2.11.3 | CORS configuration (allow `https://webifylab.my.id` only) | P0 |
| FR-2.11.4 | Request validation with clear error messages (go-playground/validator) | P0 |
| FR-2.11.5 | API versioning (`/api/v1/`) | P0 |
| FR-2.11.6 | Request logging (method, path, status, duration) via Gin middleware | P1 |
| FR-2.11.7 | Health check endpoint: `GET /api/v1/health` | P0 |

---

#### 7.3 Modul Admin — CMS Dashboard (Next.js Frontend)

##### FR-3.1: Authentication
| ID | Requirement | Priority |
|---|---|---|
| FR-3.1.1 | Login page dengan email & password | P0 |
| FR-3.1.2 | JWT stored in httpOnly cookie (secure) | P0 |
| FR-3.1.3 | Protected routes (middleware check) | P0 |
| FR-3.1.4 | Forgot password flow | P1 |
| FR-3.1.5 | Logout functionality | P0 |

##### FR-3.2: Dashboard Overview
| ID | Requirement | Priority |
|---|---|---|
| FR-3.2.1 | Statistik ringkas: total blog, total case study, total portfolio, total leads (bulan ini) | P0 |
| FR-3.2.2 | Grafik leads per bulan (bar chart) | P1 |
| FR-3.2.3 | List 5 leads terbaru dengan status | P0 |
| FR-3.2.4 | List 5 artikel terbaru (draft/published) | P1 |

##### FR-3.3: Blog Management
| ID | Requirement | Priority |
|---|---|---|
| FR-3.3.1 | CRUD blog post via Golang API | P0 |
| FR-3.3.2 | Rich text editor (Tiptap) dengan support: headings, bold/italic, lists, links, images, code blocks, tables | P0 |
| FR-3.3.3 | Upload & manage cover image | P0 |
| FR-3.3.4 | Status: Draft, Published, Archived | P0 |
| FR-3.3.5 | Slug auto-generation dari judul (editable) | P0 |
| FR-3.3.6 | Meta SEO: custom title, description, OG image | P1 |
| FR-3.3.7 | Category & Tags management | P0 |
| FR-3.3.8 | Preview before publish | P1 |

##### FR-3.4: Case Study Management
| ID | Requirement | Priority |
|---|---|---|
| FR-3.4.1 | CRUD case study dengan field terstruktur via Golang API | P0 |
| FR-3.4.2 | Upload multiple images/screenshots | P0 |
| FR-3.4.3 | Field untuk metrics (angka hasil) | P0 |
| FR-3.4.4 | Tag AI type: SPK, Sistem Pakar, Data Analytics, Machine Learning | P1 |
| FR-3.4.5 | Status: Draft, Published, Archived | P0 |

##### FR-3.5: Portfolio Management
| ID | Requirement | Priority |
|---|---|---|
| FR-3.5.1 | CRUD portfolio item via Golang API | P0 |
| FR-3.5.2 | Field: title, description, category, tech stack (multi-select), live URL, images, lessons learned | P0 |
| FR-3.5.3 | Upload thumbnail + gallery | P0 |
| FR-3.5.4 | Ordering/sorting (drag & drop atau manual) | P1 |

##### FR-3.6: Lead Management (Mini CRM)
| ID | Requirement | Priority |
|---|---|---|
| FR-3.6.1 | List semua leads dari Golang API | P0 |
| FR-3.6.2 | Detail lead: semua data form, timestamp, IP address | P0 |
| FR-3.6.3 | Status pipeline: New → Contacted (via WA) → Negotiation → Proposal Sent → Won → Lost | P0 |
| FR-3.6.4 | Catatan internal (notes) per lead | P1 |
| FR-3.6.5 | Filter & search leads | P1 |
| FR-3.6.6 | Quick action: "Copy WhatsApp number" | P1 |
| FR-3.6.7 | Export leads to CSV | P2 |

##### FR-3.7: Media Manager
| ID | Requirement | Priority |
|---|---|---|
| FR-3.7.1 | Upload images via Golang API (JPG, PNG, WebP) dengan auto-compression | P0 |
| FR-3.7.2 | Gallery view dengan search & filter | P1 |
| FR-3.7.3 | Copy URL to clipboard (untuk paste di rich text editor) | P1 |
| FR-3.7.4 | Delete media | P0 |

##### FR-3.8: Settings
| ID | Requirement | Priority |
|---|---|---|
| FR-3.8.1 | Site settings: site name, tagline, contact info, social media links | P0 |
| FR-3.8.2 | SEO settings: default meta title, description, OG image | P1 |

---

#### 7.4 Modul Sistem — SEO & Performance

##### FR-4.1: SEO
| ID | Requirement | Priority |
|---|---|---|
| FR-4.1.1 | Server-Side Rendering (SSR) atau Static Site Generation (SSG) untuk semua public pages | P0 |
| FR-4.1.2 | Dynamic meta tags (title, description, OG, Twitter Card) per halaman | P0 |
| FR-4.1.3 | Auto-generate XML Sitemap | P0 |
| FR-4.1.4 | robots.txt | P0 |
| FR-4.1.5 | Structured Data (JSON-LD): Organization, Article, Service, FAQ | P1 |
| FR-4.1.6 | Canonical URLs | P0 |
| FR-4.1.7 | Clean URL structure (`/blog/slug`, `/case-studies/slug`, `/portfolio/slug`) | P0 |
| FR-4.1.8 | hreflang tags (placeholder untuk multi-language di Phase 2) | P2 |

##### FR-4.2: Performance
| ID | Requirement | Priority |
|---|---|---|
| FR-4.2.1 | Image optimization (lazy loading, responsive sizes, WebP) | P0 |
| FR-4.2.2 | Code splitting & dynamic imports | P0 |
| FR-4.2.3 | Font optimization (preload, font-display: swap) | P1 |
| FR-4.2.4 | Caching strategy (CDN, browser cache headers) | P1 |
| FR-4.2.5 | API response caching untuk public endpoints (optional) | P2 |

---

### 8. NON-FUNCTIONAL REQUIREMENTS

| ID | Category | Requirement | Target |
|---|---|---|---|
| NFR-1 | **Performance** | First Contentful Paint (FCP) | < 1.5s |
| NFR-2 | **Performance** | Largest Contentful Paint (LCP) | < 2.5s |
| NFR-3 | **Performance** | Cumulative Layout Shift (CLS) | < 0.1 |
| NFR-4 | **Performance** | Lighthouse Performance Score | ≥ 90 |
| NFR-5 | **Performance** | API Response Time (P95) | < 500ms |
| NFR-6 | **Security** | HTTPS everywhere (via Nginx + Let's Encrypt) | Mandatory |
| NFR-7 | **Security** | Input sanitization & SQL injection prevention (GORM) | Mandatory |
| NFR-8 | **Security** | Rate limiting pada form inquiry & login (Gin middleware) | Mandatory |
| NFR-9 | **Security** | CSRF protection | Mandatory |
| NFR-10 | **Security** | File upload validation (type, size) | Mandatory |
| NFR-11 | **Security** | Password hashing (bcrypt) | Mandatory |
| NFR-12 | **Security** | JWT with short expiry + refresh token | Mandatory |
| NFR-13 | **Security** | CORS strict (only allow `https://webifylab.my.id`) | Mandatory |
| NFR-14 | **Scalability** | Arsitektur modular, siap untuk SSO integration | Required |
| NFR-15 | **Scalability** | Database terpisah untuk setiap SaaS | Required |
| NFR-16 | **Scalability** | Siap untuk multi-language (i18n) di Phase 2 | Required |
| NFR-17 | **Scalability** | Golang backend reusable untuk SaaS-saas berikutnya | Required |
| NFR-18 | **Scalability** | Monorepo structure reusable untuk SaaS-saas berikutnya | Required |
| NFR-19 | **Availability** | Uptime | ≥ 99.5% |
| NFR-20 | **Availability** | Auto-restart container saat crash | Required |
| NFR-21 | **Accessibility** | WCAG 2.1 Level AA compliance | Target |
| NFR-22 | **Browser Support** | Chrome, Firefox, Safari, Edge (2 latest versions) | Required |
| NFR-23 | **Responsive** | Mobile (320px+), Tablet (768px+), Desktop (1024px+) | Required |

---

### 9. HIGH-LEVEL USER FLOWS

#### 9.1 Flow: Calon Klien → Inquiry
```
Pengunjung → webifylab.my.id → Baca Services/Case Study → 
Klik CTA "Mulai Eksperimen" → Isi Form Inquiry → 
Submit → POST https://api.webifylab.my.id/api/v1/leads (Gin) → 
GORM simpan ke DB → Trigger email via Resend (auto-reply + notifikasi admin) → 
Next.js tampilkan success message → 
Admin terima email → Admin copy No. WA → 
Admin follow-up manual via WhatsApp → Admin update status lead di dashboard
```

#### 9.2 Flow: Pengunjung → Blog Reader
```
Pengunjung → webifylab.my.id / Google Search → Blog Listing (GET /api/v1/blog) → 
Filter/Search Artikel → Baca Artikel (GET /api/v1/blog/:slug) → 
Share / Baca Related Articles → (Konversi ke Contact)
```

#### 9.3 Flow: Admin → Publish Case Study
```
Admin Login → POST /api/v1/auth/login (Gin) → JWT token → 
Dashboard → Menu Case Study → 
Create New → Isi Field Terstruktur (Client, Challenge, Solution, AI Used, Results, Lessons Learned) → 
Upload Screenshots (POST /api/v1/media/upload) → 
Preview → Publish (POST /api/v1/case-studies) → 
Live di webifylab.my.id/case-studies/slug
```

#### 9.4 Flow: Admin → Manage Leads
```
Admin Login → Dashboard → Menu Leads (GET /api/v1/leads) → 
Lihat List Leads → Filter by Status → 
Klik Lead Detail (GET /api/v1/leads/:id) → Baca Deskripsi Project → 
Klik "Copy WA Number" → Follow up via WhatsApp (manual) → 
Update Status (PUT /api/v1/leads/:id) → Tambah Notes
```

#### 9.5 Flow: Development (Local)
```
Developer (Laptop 16GB) → 
Clone monorepo → cp .env.example .env → 
Option A: make dev (full Docker) → semua services jalan
Option B: Native dev (PostgreSQL Docker + Golang native + Next.js native)
→ Develop & test → Commit ke git
```

#### 9.6 Flow: Deployment (Production)
```
Developer (Laptop 16GB) → 
make build → make deploy-images → make upload → make start → 
Run migrations & seed → Verify health check → 
Site live di webifylab.my.id + api.webifylab.my.id
```

---

### 10. INFORMATION ARCHITECTURE

```
webifylab.my.id/  (Next.js Frontend)
├── /                        → Home
├── /about                   → About Us (filosofi "We Experiment, You Grow")
├── /services                → Services (list)
│   └── /services/[slug]     → Service Detail (opsional)
├── /portfolio               → Portfolio (grid + filter)
│   └── /portfolio/[slug]    → Portfolio Detail (+ lessons learned)
├── /case-studies            → Case Studies (list)
│   └── /case-studies/[slug] → Case Study Detail (+ lessons learned)
├── /blog                    → Blog (list + search + filter)
│   └── /blog/[slug]         → Blog Post Detail
├── /products                → Products / Ecosystem (SaaS list)
├── /contact                 → Contact + Inquiry Form
│
├── /admin                   → Admin Dashboard
│   ├── /admin/login         → Login
│   ├── /admin/dashboard     → Overview
│   ├── /admin/blog          → Blog Management
│   ├── /admin/case-studies  → Case Study Management
│   ├── /admin/portfolio     → Portfolio Management
│   ├── /admin/leads         → Lead Management
│   ├── /admin/media         → Media Manager
│   └── /admin/settings      → Settings (SEO, profile)
│
api.webifylab.my.id/  (Golang Backend - Gin + GORM)
├── /api/v1/auth/*           → Authentication endpoints
├── /api/v1/blog/*           → Blog endpoints
├── /api/v1/case-studies/*   → Case study endpoints
├── /api/v1/portfolio/*      → Portfolio endpoints
├── /api/v1/services/*       → Services endpoints
├── /api/v1/leads/*          → Leads endpoints
├── /api/v1/media/*          → Media endpoints
├── /api/v1/categories/*     → Categories endpoints
├── /api/v1/tags/*           → Tags endpoints
├── /api/v1/dashboard/*      → Dashboard stats endpoints
├── /api/v1/settings/*       → Settings endpoints
└── /api/v1/health           → Health check
```

---

### 11. HIGH-LEVEL DATA MODEL (ERD Overview)

```
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│    users     │     │   blog_posts │     │  categories  │
├──────────────┤     ├──────────────┤     ├──────────────┤
│ id (UUID)    │     │ id (UUID)    │     │ id (UUID)    │
│ name         │     │ title        │     │ name         │
│ email        │     │ slug         │     │ slug         │
│ password     │     │ content      │     │ type (blog/  │
│ role         │     │ cover_image  │     │   portfolio) │
│ created_at   │     │ excerpt      │     └──────┬───────┘
│ updated_at   │     │ author_id(FK)│            │
└──────────────┘     │ category_id  │◄───────────┘
                     │ status       │
┌──────────────┐     │ meta_title   │     ┌──────────────┐
│    leads     │     │ meta_desc    │     │    tags      │
├──────────────┤     │ published_at │     ├──────────────┤
│ id (UUID)    │     │ created_at   │     │ id (UUID)    │
│ name         │     │ updated_at   │     │ name         │
│ email        │     └──────┬───────┘     │ slug         │
│ phone (WA)   │            │             └──────┬───────┘
│ service_type │     ┌──────┴───────┐            │
│ budget_range │     │ post_tags    │            │
│ description  │     ├──────────────┤     ┌──────┴───────┐
│ status       │     │ post_id (FK) │     │ case_studies │
│ notes        │     │ tag_id (FK)  │     ├──────────────┤
│ ip_address   │     └──────────────┘     │ id (UUID)    │
│ created_at   │                          │ client_name  │
│ updated_at   │                          │ industry     │
└──────────────┘                          │ challenge    │
                                          │ approach     │
┌──────────────┐     ┌──────────────┐     │ solution     │
│  portfolio   │     │   services   │     │ ai_type      │
├──────────────┤     ├──────────────┤     │ results      │
│ id (UUID)    │     │ id (UUID)    │     │ lessons      │
│ title        │     │ name         │     │ testimonial  │
│ slug         │     │ slug         │     │ metrics      │
│ description  │     │ description  │     │ images       │
│ category     │     │ icon         │     │ status       │
│ tech_stack   │     │ order        │     │ slug         │
│ live_url     │     │ is_active    │     │ created_at   │
│ images       │     └──────────────┘     │ updated_at   │
│ thumbnail    │                          └──────────────┘
│ lessons      │
│ status       │     ┌──────────────┐
│ order        │     │    media     │
│ created_at   │     ├──────────────┤
│ updated_at   │     │ id (UUID)    │
└──────────────┘     │ filename     │
                     │ url          │
                     │ mime_type    │
                     │ size         │
                     │ uploaded_by  │
                     │ created_at   │
                     └──────────────┘

┌──────────────────┐
│  site_settings   │
├──────────────────┤
│ id (UUID)        │
│ key              │
│ value            │
│ updated_at       │
└──────────────────┘
```

**Catatan:**
- Semua primary key menggunakan UUID (aman untuk SSO di masa depan)
- Timestamp menggunakan `TIMESTAMPTZ` (timezone-aware)
- GORM akan handle auto-migration di development, manual migration di production

---

### 12. TECH STACK & INFRASTRUCTURE

#### 12.1 Tech Stack

##### Frontend (Next.js) — `/frontend`
| Layer | Technology | Version | Alasan |
|---|---|---|---|
| **Framework** | Next.js (App Router) | 15+ | SSR/SSG untuk SEO, React ecosystem |
| **Language** | TypeScript | 5.x | Type safety, maintainability |
| **Styling** | Tailwind CSS + shadcn/ui | Latest | Rapid development, consistent design system |
| **Animation** | Framer Motion | Latest | Premium feel untuk agency website |
| **Rich Text Editor** | Tiptap (ProseMirror) | Latest | Extensible, headless, support code blocks |
| **HTTP Client** | Axios | Latest | Call Golang API |
| **Forms** | React Hook Form + Zod | Latest | Validation, type-safe |
| **Charts** | Recharts | Latest | Dashboard grafik leads |
| **Icons** | Lucide React | Latest | Consistent iconography |

##### Backend (Golang) — `/backend`
| Layer | Technology | Version | Alasan |
|---|---|---|---|
| **Language** | Go | 1.22+ | Lightweight, fast, low memory footprint |
| **HTTP Framework** | **Gin** | Latest | ✅ Populer, ecosystem luas, middleware kaya |
| **ORM** | **GORM** | Latest | ✅ Developer-friendly, auto-migration, hooks |
| **DB Driver** | gorm.io/driver/postgres | Latest | PostgreSQL driver untuk GORM |
| **Validation** | go-playground/validator | Latest | Standard validation library |
| **JWT** | golang-jwt/jwt | v5 | Standard JWT implementation |
| **Password Hash** | golang.org/x/crypto/bcrypt | Latest | Secure password hashing |
| **Email** | **resend/resend-go** | Latest | ✅ Official Resend Go client |
| **Image Processing** | disintegration/imaging | Latest | Auto-compress, resize |
| **Logging** | rs/zerolog | Latest | Structured logging, low overhead |
| **Config** | spf13/viper | Latest | Environment & config management |
| **Migration** | golang-migrate/migrate | Latest | Database schema migration |
| **CORS** | gin-contrib/cors | Latest | Gin CORS middleware |
| **Rate Limit** | gin-contrib/limiter atau ulule/limiter | Latest | Rate limiting middleware |
| **UUID** | google/uuid | Latest | UUID generation |

##### Infrastructure — `/infra`
| Layer | Technology | Alasan |
|---|---|---|
| **Database** | PostgreSQL 15 | Relational, robust, scalable |
| **Reverse Proxy** | Nginx | ✅ Sudah terinstall, familiar, handle SSL |
| **SSL** | Let's Encrypt + Certbot | Free, auto-renew |
| **Containerization** | Docker + Docker Compose | Consistent environment, easy deploy |
| **Process Management** | Docker restart policy | Auto-restart container saat crash |
| **Monitoring** | Uptime Kuma (self-hosted) | Free, lightweight uptime monitoring |
| **Backup** | pg_dump + cron | Daily database backup |
| **Email Service** | **Resend** | ✅ Akun sudah tersedia |

##### Monorepo Tooling — Root
| Layer | Technology | Alasan |
|---|---|---|
| **Orchestration** | Docker Compose | Manage multi-container setup |
| **Automation** | Makefile | Common commands, simplify workflow |
| **Documentation** | Markdown + `/docs` folder | Centralized documentation |
| **Version Control** | Git | Atomic commits, version history |
| **Environment** | .env di root | Shared environment variables |

#### 12.2 Server Infrastructure

**Current Spec:**
- CPU: 1 vCPU
- RAM: 1 GB
- Storage: 20 GB SSD
- Bandwidth: Unmetered
- Zone: Bogor T3
- OS: Ubuntu 22.04 LTS / Debian 12
- Existing: Docker installed, Nginx installed
- Previous deployment: Golang + Next.js (backup ✅, belum dihapus ⏳)

**Deployment Architecture:**

```
┌──────────────────────────────────────────────────────────┐
│                  VPS (1GB RAM, Bogor T3)                 │
│                                                          │
│  ┌────────────────────────────────────────────────────┐ │
│  │              Nginx (Host Network)                  │ │
│  │  - Reverse proxy                                   │ │
│  │  - SSL termination (Let's Encrypt)                 │ │
│  │  - 2 server blocks:                                │ │
│  │    • webifylab.my.id → proxy to :3000 (Next.js)    │ │
│  │    • api.webifylab.my.id → proxy to :8080 (Gin)    │ │
│  │  - Serve static files (uploads, .next/static)      │ │
│  │  - ~20MB RAM                                       │ │
│  └──────────────┬──────────────────┬──────────────────┘ │
│                 │                  │                     │
│                 ▼                  ▼                     │
│  ┌─────────────────────┐  ┌─────────────────────┐       │
│  │  Docker: Next.js    │  │  Docker: Golang     │       │
│  │  Container          │  │  Container (Gin)    │       │
│  │  - Port 3000        │  │  - Port 8080        │       │
│  │  - SSR/SSG          │  │  - REST API         │       │
│  │  - Admin CMS        │  │  - GORM + Postgres  │       │
│  │  - ~300MB RAM       │  │  - Resend email     │       │
│  │  (limit: 350MB)     │  │  - ~100MB RAM       │       │
│  └─────────────────────┘  │  (limit: 150MB)     │       │
│                            └─────────────────────┘       │
│                                                          │
│  ┌────────────────────────────────────────────────────┐ │
│  │         Docker: PostgreSQL Container               │ │
│  │  - Port 5432 (internal only)                       │ │
│  │  - Tuned for low memory                            │ │
│  │  - ~128MB RAM (limit: 200MB)                       │ │
│  └────────────────────────────────────────────────────┘ │
│                                                          │
│  Total Container RAM: ~530MB / 1GB                      │
│  Buffer for OS + spike: ~470MB                          │
└──────────────────────────────────────────────────────────┘
```

**Docker Compose Configuration:** (lihat Section 6.4)

**Golang Backend Structure:** (lihat Section 6.1 - `/backend` folder)

**Dockerfile — Golang Backend (Multi-stage Build):**

```dockerfile
# backend/Dockerfile
# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Install dependencies for image processing
RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/server

# Runtime stage
FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations

RUN mkdir -p /app/uploads

EXPOSE 8080
CMD ["./server"]
```

**Dockerfile — Next.js Frontend:**

```dockerfile
# frontend/Dockerfile
# Build stage
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

# Runtime stage
FROM node:20-alpine
WORKDIR /app
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public
EXPOSE 3000
ENV NODE_ENV=production
ENV HOSTNAME="0.0.0.0"
CMD ["node", "server.js"]
```

**PostgreSQL Tuning for 1GB RAM:**

```ini
# infra/postgres/postgresql.conf

# Memory settings (total ~128MB)
shared_buffers = 128MB
effective_cache_size = 256MB
work_mem = 4MB
maintenance_work_mem = 64MB

# WAL settings
wal_buffers = 4MB

# Connection settings
max_connections = 20

# Logging
log_min_duration_statement = 1000

# Checkpoints
checkpoint_completion_target = 0.9
```

**Nginx Configuration:**

```nginx
# infra/nginx/webifylab.conf

# Redirect HTTP to HTTPS
server {
    listen 80;
    server_name webifylab.my.id api.webifylab.my.id;
    return 301 https://$server_name$request_uri;
}

# Frontend (Next.js) - webifylab.my.id
server {
    listen 443 ssl http2;
    server_name webifylab.my.id;

    ssl_certificate /etc/letsencrypt/live/webifylab.my.id/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/webifylab.my.id/privkey.pem;

    # Security headers
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;

    # Gzip
    gzip on;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml;

    # Uploaded media files (served directly)
    location /uploads/ {
        alias /var/www/webifylab/uploads/;
        expires 30d;
        access_log off;
    }

    # All other requests to Next.js
    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;
    }
}
```

```nginx
# infra/nginx/api.conf

# Backend API (Golang + Gin) - api.webifylab.my.id
server {
    listen 443 ssl http2;
    server_name api.webifylab.my.id;

    ssl_certificate /etc/letsencrypt/live/api.webifylab.my.id/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.webifylab.my.id/privkey.pem;

    # Security headers
    add_header X-Frame-Options "DENY" always;
    add_header X-Content-Type-Options "nosniff" always;

    # Rate limiting
    limit_req_zone $binary_remote_addr zone=api:10m rate=10r/s;
    limit_req zone=api burst=20 nodelay;

    location / {
        proxy_pass http://localhost:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # Timeouts
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # Uploaded files (served directly by Nginx)
    location /uploads/ {
        alias /var/www/webifylab/uploads/;
        expires 30d;
        access_log off;
    }
}
```

**System Optimization:**

```bash
# Swap space (2GB) - critical for 1GB RAM VPS
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab

# Swappiness (use swap only when necessary)
sudo sysctl vm.swappiness=10
echo 'vm.swappiness=10' | sudo tee -a /etc/sysctl.conf
```

**Deployment Script:**

```bash
#!/bin/bash
# infra/scripts/deploy.sh

set -e

echo "🚀 Starting deployment..."

# Build images
echo "📦 Building Docker images..."
make build

# Save images
echo "💾 Saving images..."
make deploy-images

# Upload to VPS
echo "📤 Uploading to VPS..."
make upload

# Start containers
echo "▶️  Starting containers..."
make start

# Run migrations
echo "🗄️  Running migrations..."
ssh user@vps 'cd /var/www/webifylab && docker-compose exec backend ./server migrate'

# Cleanup
echo "🧹 Cleaning up..."
rm -f frontend.tar.gz backend.tar.gz

echo "✅ Deployment completed!"
```

**Backup Script:**

```bash
#!/bin/bash
# infra/scripts/backup.sh

BACKUP_DIR="/var/backups/webifylab"
DATE=$(date +%Y%m%d_%H%M%S)
mkdir -p $BACKUP_DIR

echo "[$(date)] Starting backup..."

# Backup database
docker-compose -f /var/www/webifylab/docker-compose.yml exec -T postgres \
  pg_dump -U webifylab webifylab | gzip > $BACKUP_DIR/db_$DATE.sql.gz

# Backup uploads
tar -czf $BACKUP_DIR/uploads_$DATE.tar.gz /var/www/webifylab/uploads/

# Keep last 7 days
find $BACKUP_DIR -type f -mtime +7 -delete

echo "[$(date)] Backup completed: db_$DATE.sql.gz, uploads_$DATE.tar.gz"
```

**Environment Variables Template:** (lihat Section 6.5)

**Monitoring Commands:**

```bash
# Check container status
docker-compose ps

# Check resource usage
docker stats

# Check logs
docker-compose logs -f backend
docker-compose logs -f frontend

# Health check
curl https://api.webifylab.my.id/api/v1/health

# Database connection
docker-compose exec postgres psql -U webifylab -c "SELECT count(*) FROM pg_stat_activity;"
```

**Risk Mitigation:**

| Risk | Mitigation |
|---|---|
| OOM (Out of Memory) saat traffic spike | Swap 2GB, container memory limits, monitoring alert |
| Container crash | Docker restart policy: `unless-stopped` |
| Database corruption | Daily backup + test restore quarterly |
| SSL certificate expiry | Certbot auto-renew + cron monitoring |
| Deployment downtime | Blue-green deployment (future), maintenance mode |
| Build process gagal di VPS | Build di laptop, upload artifact |
| Golang binary incompatible | Multi-stage Docker build, target linux/amd64 |
| Resend API limit (100 emails/hari) | Monitor usage, upgrade plan jika perlu |
| GORM migration error | Test migration di local dulu, backup DB sebelum migrate |
| Monorepo complexity | Clear folder structure, Makefile, documentation |

---

### 13. CONSTRAINTS & ASSUMPTIONS

#### Constraints
| No | Constraint |
|---|---|
| C1 | Budget terbatas (solo developer / small team) — pilih tech stack yang free-tier friendly |
| C2 | Timeline MVP: 8-12 minggu |
| C3 | Portal dan SaaS harus menggunakan database terpisah |
| C4 | Konten blog & case study akan ditulis dalam **Bahasa Indonesia** (multi-language di Phase 2) |
| C5 | WhatsApp notification untuk lead dilakukan secara **manual** (tidak menggunakan API berbayar) |
| C6 | Domain frontend: **webifylab.my.id** |
| C7 | Domain API: **api.webifylab.my.id** (subdomain terpisah) |
| C8 | Server: VPS 1 vCPU, 1 GB RAM, 20 GB SSD, Bogor T3 |
| C9 | Build process dilakukan di laptop lokal (16GB RAM), bukan di server |
| C10 | Deployment menggunakan **Docker + Docker Compose** |
| C11 | Reverse proxy menggunakan **Nginx** (sudah terinstall di VPS) |
| C12 | Backend menggunakan **Golang** dengan framework **Gin** dan ORM **GORM** |
| C13 | Frontend menggunakan **Next.js** (SSR/SSG untuk SEO) |
| C14 | Email service menggunakan **Resend** (akun sudah tersedia) |
| C15 | Previous deployment (Golang + Next.js) sudah di-backup, akan dihapus sebelum deployment baru |
| C16 | Repository menggunakan **monorepo** structure (frontend + backend + infra + docs) |

#### Assumptions
| No | Assumption |
|---|---|
| A1 | Volume traffic awal rendah (< 10K visits/bulan), sehingga single server sudah cukup |
| A2 | Jumlah admin/editor awal 1-3 orang |
| A3 | SaaS pertama akan dibangun setelah portal MVP selesai |
| A4 | SSO akan diimplementasikan saat sudah ada ≥ 2 SaaS yang live |
| A5 | Klien UMKM lebih responsif via WhatsApp daripada email, sehingga follow-up manual via WA lebih efektif |
| A6 | Brand guideline akan dikembangkan secara iteratif seiring dengan launch portal |
| A7 | Multi-language (ID/EN) akan diimplementasikan di Phase 2 setelah portal stabil |
| A8 | Laptop lokal memiliki RAM 16GB, cukup untuk build Docker images |
| A9 | VPS memiliki Docker dan Nginx terinstall, siap digunakan |
| A10 | User familiar dengan basic Linux command, Nginx, dan Docker dasar |
| A11 | Arsitektur Golang (Gin + GORM) + Next.js + Docker ini akan menjadi **template** untuk SaaS-saas berikutnya |
| A12 | Resend free tier (3000 emails/bulan, 100 emails/hari) cukup untuk Phase 1 |
| A13 | DNS propagation untuk `api.webifylab.my.id` akan selesai dalam 24-48 jam |
| A14 | Previous deployment data tidak perlu dimigrasi ke database baru (start fresh) |
| A15 | Monorepo structure akan direuse untuk SaaS-saas berikutnya (template) |

---

### 14. SUCCESS METRICS & KPI

| Metric | Target (3 Bulan Post-Launch) | Cara Ukur |
|---|---|---|
| Organic Traffic | 500+ visits/bulan | Google Analytics / Plausible |
| Blog Articles Published | 12+ artikel | CMS Dashboard |
| Case Studies Published | 3+ case study | CMS Dashboard |
| Inquiry/Leads | 10+ leads/bulan | Lead Dashboard |
| Lead Conversion Rate | ≥ 10% (inquiry → project) | Manual tracking |
| Lighthouse Score | ≥ 90 semua kategori | Lighthouse CI |
| Uptime | ≥ 99.5% | Uptime Kuma |
| API Response Time (P95) | < 500ms | Golang logging (zerolog) |
| Container Memory Usage | < 80% of limit | `docker stats` |
| Email Delivery Rate | ≥ 98% | Resend dashboard |
| API Error Rate | < 1% | Golang logging |
| Deployment Time | < 10 menit | Makefile automation |

---

### 15. OPEN QUESTIONS & DECISIONS

| No | Question | Status |
|---|---|---|
| 1 | ~~Apakah domain sudah dibeli?~~ | ✅ webifylab.my.id |
| 2 | ~~Apakah WhatsApp notification manual?~~ | ✅ Manual |
| 3 | ~~Apakah multi-language di Phase 1?~~ | ✅ Phase 2 |
| 4 | ~~Apakah upgrade VPS ke 2GB RAM?~~ | ✅ Tetap 1GB dengan optimasi |
| 5 | ~~Apakah tagline sudah final?~~ | ✅ "We Experiment, You Grow" |
| 6 | ~~Apakah brand guideline sudah cocok?~~ | ✅ Sudah cocok |
| 7 | ~~Apakah pakai Docker container?~~ | ✅ Docker + Docker Compose |
| 8 | ~~Email service?~~ | ✅ Resend (akun sudah ada) |
| 9 | ~~Backend technology?~~ | ✅ Golang + Gin + GORM |
| 10 | ~~Backup data previous deployment?~~ | ✅ Sudah ada |
| 11 | ~~API domain structure?~~ | ✅ Subdomain terpisah (api.webifylab.my.id) |
| 12 | ~~Golang framework?~~ | ✅ Gin |
| 13 | ~~Golang ORM?~~ | ✅ GORM |
| 14 | ~~Repository structure?~~ | ✅ **Monorepo** (frontend + backend + infra + docs) |
| 15 | Hapus previous deployment dari VPS? | ⏳ **Action item sebelum dev dimulai** |
| 16 | Setup DNS A record untuk api.webifylab.my.id? | ⏳ **Action item sebelum dev dimulai** |
| 17 | Setup SSL certificate untuk api.webifylab.my.id? | ⏳ **Action item sebelum dev dimulai** |
| 18 | Seed data untuk initial admin user? | ❓ Pending (recommend: ya, via CLI command) |
| 19 | Git hosting (GitHub/GitLab/self-hosted)? | ❓ Pending |

---

### 16. NEXT STEPS

Setelah PRD ini di-approve, langkah selanjutnya adalah:

#### 🚨 Immediate Actions (Pre-Development)
1. **Hapus previous deployment** dari VPS (ikuti checklist di Section 4.2)
2. **Setup DNS A record** untuk `api.webifylab.my.id` → IP VPS
3. **Setup SSL certificate** untuk `api.webifylab.my.id` via Certbot
4. **Verifikasi domain di Resend** dan generate API key
5. **Init monorepo repository** (structure, git, README, Makefile)

#### 📋 Development Phase
1. **🗄️ Detailed Database Schema** — Full ERD dengan data types, constraints, indexes, GORM models
2. **🔌 API Specification (OpenAPI/Swagger)** — Detail semua endpoints, request/response schema
3. **📐 Wireframe / UI Design** — Low-fi wireframe untuk semua halaman utama
4. **🏗️ Project Setup** — Init monorepo, setup Next.js, Golang (Gin + GORM), Docker, Nginx configs
5. **🔐 Authentication Implementation** — JWT flow, middleware, protected routes
6. **📝 Sprint Planning** — Breakdown ke dalam sprint 2 mingguan

#### 🚀 Deployment Phase
1. **🚀 Deployment Pipeline** — Setup build & deployment workflow via Makefile
2. **📊 Monitoring Setup** — Uptime Kuma, logging, backup cron
3. **🧪 Testing & QA** — End-to-end testing, performance testing
4. **🎯 Launch** — Soft launch → monitoring → full launch

