# 📄 SPRINT ROADMAP & PLANNING DOCUMENT
## WebifyLab Portal — Phase 1 (MVP)

---

### 1. INFORMASI DOKUMEN

| Field | Detail |
|---|---|
| **Produk** | WebifyLab Portal |
| **Tipe Dokumen** | Sprint Roadmap & Planning |
| **Versi** | 1.0.0 |
| **Tanggal** | 05 September 2026 |
| **Durasi Total** | 12 minggu (6 sprint × 2 minggu) |
| **Metodologi** | Agile Scrum (Solo Developer) |
| **Penulis** | Rizal (System Analyst & Fullstack Developer) |
| **Status** | Approved — Ready for Execution |

---

### 2. OVERVIEW SPRINT

#### 2.1 Ringkasan Timeline

```
Minggu:     1    2    3    4    5    6    7    8    9    10   11   12
            │    │    │    │    │    │    │    │    │    │    │    │
Sprint 0:   ████                                                    ← Setup (1 minggu)
Sprint 1:        ██████████                                         ← Backend Foundation
Sprint 2:                   ██████████                              ← Content API
Sprint 3:                              ██████████                   ← Frontend Public
Sprint 4:                                         ██████████        ← Admin Dashboard
Sprint 5:                                                    ██████████ ← Integration
Sprint 6:                                                              ████████ ← Launch
```

#### 2.2 Daftar Sprint

| Sprint | Nama | Durasi | Fokus Utama | Deliverable |
|---|---|---|---|---|
| **Sprint 0** | Pre-Development Setup | 1 minggu | Setup infrastruktur & repository | Monorepo siap, VPS siap |
| **Sprint 1** | Backend Foundation | 2 minggu | Golang project + Auth API | API auth berfungsi |
| **Sprint 2** | Content API | 2 minggu | CRUD API untuk semua content | 30+ endpoints siap |
| **Sprint 3** | Frontend Public | 2 minggu | Next.js + semua halaman publik | 11 halaman publik live |
| **Sprint 4** | Admin Dashboard | 2 minggu | CMS dashboard lengkap | Admin bisa manage content |
| **Sprint 5** | Integration & Polish | 2 minggu | Integrasi + optimasi + email | Semua fitur terintegrasi |
| **Sprint 6** | Testing & Launch | 1 minggu | QA + deployment + launch | Portal live! 🚀 |

**Total Durasi:** 12 minggu

#### 2.3 Milestone Utama

| Milestone | Target Tanggal | Kriteria |
|---|---|---|
| 🏁 **M1: Infrastructure Ready** | Akhir Minggu 1 | Monorepo siap, VPS clean, DNS & SSL aktif |
| 🏁 **M2: API Core Ready** | Akhir Minggu 3 | Auth API + database connected |
| 🏁 **M3: Content API Complete** | Akhir Minggu 5 | Semua CRUD API berfungsi |
| 🏁 **M4: Public Pages Live** | Akhir Minggu 7 | Semua halaman publik bisa diakses |
| 🏁 **M5: Admin Dashboard Complete** | Akhir Minggu 9 | CMS bisa manage semua content |
| 🏁 **M6: Feature Complete** | Akhir Minggu 11 | Semua fitur terintegrasi |
| 🏁 **M7: LAUNCH!** | Akhir Minggu 12 | Portal live di webifylab.my.id |

#### 2.4 Definition of Done (DoD)

Sebuah task dianggap **SELESAI** jika:

- [ ] Kode sudah ditulis dan berjalan tanpa error
- [ ] Kode sudah di-test manual (minimal happy path)
- [ ] Kode sudah di-commit ke git dengan pesan yang jelas
- [ ] Dokumentasi sudah di-update (jika perlu)
- [ ] Tidak ada breaking changes ke fitur lain
- [ ] Acceptance criteria terpenuhi

#### 2.5 Daily Check-in Template (Solo Dev)

Setiap pagi, tanyakan ke diri sendiri:

```
📅 Tanggal: _______________
🎯 Sprint: _______________

✅ Kemarin saya menyelesaikan:
   1. _______________
   2. _______________
   3. _______________

📌 Hari ini saya akan mengerjakan:
   1. _______________
   2. _______________
   3. _______________

🚧 Ada blocker/hambatan:
   1. _______________

⏰ Estimasi waktu kerja hari ini: ___ jam
```

---

### 3. SPRINT 0: PRE-DEVELOPMENT SETUP
**Durasi:** 1 minggu  
**Fokus:** Setup infrastruktur, repository, dan environment development  
**Goal:** Semua fondasi siap sebelum coding dimulai

---

#### 3.1 Task List Sprint 0

##### 🗂️ Kategori: Infrastructure & VPS

- [ ] **TASK-001: Hapus previous deployment dari VPS**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Tidak ada
  - **Detail:**
    - [ ] Stop semua container yang berjalan (`docker-compose down`)
    - [ ] Backup database lama jika perlu (`docker exec ... pg_dump`)
    - [ ] Hapus container dan image lama (`docker container prune`, `docker image prune`)
    - [ ] Hapus volume lama (`docker volume rm`)
    - [ ] Hapus files deployment lama (`rm -rf /var/www/old-webifylab/*`)
    - [ ] Verifikasi VPS dalam keadaan bersih (`docker ps` kosong)
  - **Acceptance Criteria:**
    - VPS dalam keadaan bersih, tidak ada container berjalan
    - Storage dan RAM tersedia untuk deployment baru

- [ ] **TASK-002: Setup DNS A record untuk api.webifylab.my.id**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Tidak ada
  - **Detail:**
    - [ ] Login ke DNS provider (domain registrar)
    - [ ] Tambahkan A record: `api` → IP VPS
    - [ ] Tunggu propagation (1-24 jam)
    - [ ] Verifikasi dengan `nslookup api.webifylab.my.id` atau `dig`
  - **Acceptance Criteria:**
    - `api.webifylab.my.id` resolve ke IP VPS
    - `curl http://api.webifylab.my.id` bisa diakses (meski belum ada content)

- [ ] **TASK-003: Setup SSL certificate untuk api.webifylab.my.id**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-002 (DNS sudah propagate)
  - **Detail:**
    - [ ] Install Certbot jika belum (`sudo apt install certbot`)
    - [ ] Generate SSL untuk api subdomain: `sudo certbot certonly --standalone -d api.webifylab.my.id`
    - [ ] Verifikasi certificate ter-generate di `/etc/letsencrypt/live/api.webifylab.my.id/`
    - [ ] Setup auto-renew: `sudo crontab -e` → `0 3 * * * certbot renew --quiet`
  - **Acceptance Criteria:**
    - SSL certificate valid untuk api.webifylab.my.id
    - Auto-renew sudah dikonfigurasi

- [ ] **TASK-004: Setup swap space di VPS**
  - **Estimasi:** 30 menit
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Tidak ada
  - **Detail:**
    - [ ] Buat swap file 2GB: `sudo fallocate -l 2G /swapfile`
    - [ ] Set permissions: `sudo chmod 600 /swapfile`
    - [ ] Format swap: `sudo mkswap /swapfile`
    - [ ] Enable swap: `sudo swapon /swapfile`
    - [ ] Persist di fstab: `echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab`
    - [ ] Set swappiness: `sudo sysctl vm.swappiness=10`
  - **Acceptance Criteria:**
    - `free -h` menunjukkan swap 2GB aktif
    - Swappiness = 10

- [ ] **TASK-005: Verifikasi Docker dan Nginx di VPS**
  - **Estimasi:** 30 menit
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Tidak ada
  - **Detail:**
    - [ ] Cek Docker: `docker --version` dan `docker compose version`
    - [ ] Cek Nginx: `nginx -v` dan `sudo systemctl status nginx`
    - [ ] Pastikan Nginx berjalan: `sudo systemctl start nginx`
    - [ ] Test Nginx default page: `curl http://localhost`
  - **Acceptance Criteria:**
    - Docker dan Docker Compose terinstall dan berfungsi
    - Nginx berjalan dan bisa serve halaman default

##### 🗂️ Kategori: Repository Setup

- [ ] **TASK-006: Init monorepo repository**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Tidak ada
  - **Detail:**
    - [ ] Buat folder project: `mkdir webifylab-portal && cd webifylab-portal`
    - [ ] Init git: `git init`
    - [ ] Buat struktur folder:
      ```
      webifylab-portal/
      ├── frontend/
      ├── backend/
      ├── infra/
      │   ├── nginx/
      │   ├── postgres/
      │   └── scripts/
      ├── docs/
      └── .github/
      ```
    - [ ] Buat file root:
      - [ ] `README.md` — Project overview
      - [ ] `.gitignore` — Global gitignore
      - [ ] `.env.example` — Environment variables template
      - [ ] `docker-compose.yml` — Docker orchestration
      - [ ] `docker-compose.dev.yml` — Development overrides
      - [ ] `Makefile` — Common commands
    - [ ] Commit awal: `git add . && git commit -m "chore: init monorepo structure"`
  - **Acceptance Criteria:**
    - Struktur folder sesuai dengan Database Schema Document
    - Semua file root sudah ada
    - Git repository ter-initialize

- [ ] **TASK-007: Setup .gitignore**
  - **Estimasi:** 30 menit
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-006
  - **Detail:**
    - [ ] Buat `.gitignore` di root:
      ```
      # Dependencies
      node_modules/
      vendor/
      
      # Build outputs
      .next/
      dist/
      build/
      
      # Environment
      .env
      .env.local
      .env.*.local
      
      # Logs
      *.log
      logs/
      
      # OS files
      .DS_Store
      Thumbs.db
      
      # IDE
      .vscode/
      .idea/
      
      # Docker
      *.tar.gz
      
      # Uploads (local development)
      backend/uploads/*
      !backend/uploads/.gitkeep
      
      # Database
      *.sql.gz
      ```
  - **Acceptance Criteria:**
    - File-file sensitive tidak ter-commit ke git

- [ ] **TASK-008: Setup Makefile dengan common commands**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-006
  - **Detail:**
    - [ ] Buat `Makefile` dengan commands:
      - `make dev` — Start development environment
      - `make build` — Build all Docker images
      - `make up` — Start all containers
      - `make down` — Stop all containers
      - `make logs` — View logs
      - `make migrate` — Run database migrations
      - `make seed` — Seed initial data
      - `make deploy` — Full deployment pipeline
      - `make backup` — Backup database + uploads
    - [ ] Test semua commands berjalan tanpa error
  - **Acceptance Criteria:**
    - Semua commands di Makefile berfungsi
    - Dokumentasi commands ada di README.md

- [ ] **TASK-009: Setup docker-compose.yml**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-006
  - **Detail:**
    - [ ] Buat `docker-compose.yml` dengan 3 services:
      - `frontend` (Next.js)
      - `backend` (Golang)
      - `postgres` (PostgreSQL 15)
    - [ ] Konfigurasi network: `webifylab-network`
    - [ ] Konfigurasi volumes: `pgdata`, `uploads`
    - [ ] Memory limits sesuai PRD:
      - frontend: 350MB
      - backend: 150MB
      - postgres: 200MB
    - [ ] Healthcheck untuk postgres
    - [ ] Test dengan `docker-compose config`
  - **Acceptance Criteria:**
    - `docker-compose config` valid tanpa error
    - Struktur sesuai dengan API Specification Document

- [ ] **TASK-010: Setup .env.example**
  - **Estimasi:** 30 menit
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-006
  - **Detail:**
    - [ ] Buat `.env.example` dengan semua variables yang dibutuhkan:
      ```
      # Application
      APP_ENV=production
      SITE_URL=https://webifylab.my.id
      API_URL=https://api.webifylab.my.id
      
      # Database
      DB_USER=webifylab
      DB_PASSWORD=
      DB_NAME=webifylab
      
      # JWT
      JWT_SECRET=
      
      # Resend
      RESEND_API_KEY=
      ADMIN_EMAIL=admin@webifylab.my.id
      ```
    - [ ] Buat `.env` lokal untuk development (tidak di-commit)
    - [ ] Generate JWT secret: `openssl rand -base64 32`
  - **Acceptance Criteria:**
    - Template .env.example lengkap
    - .env lokal sudah terisi untuk development

##### 🗂️ Kategori: Local Development Environment

- [ ] **TASK-011: Verifikasi local development tools**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Tidak ada
  - **Detail:**
    - [ ] Cek Node.js: `node --version` (minimal v20)
    - [ ] Cek npm: `npm --version`
    - [ ] Cek Go: `go version` (minimal 1.22)
    - [ ] Cek Docker Desktop: `docker --version` (jika pakai)
    - [ ] Cek Git: `git --version`
    - [ ] Cek code editor (VS Code / lainnya)
    - [ ] Install extensions yang dibutuhkan:
      - Go (untuk Golang)
      - ESLint + Prettier (untuk Next.js)
      - Docker (untuk Docker)
  - **Acceptance Criteria:**
    - Semua tools terinstall dengan versi yang sesuai
    - Laptop 16GB RAM cukup untuk build

- [ ] **TASK-012: Setup PostgreSQL lokal untuk development**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-011
  - **Detail:**
    - [ ] Option A: Install PostgreSQL lokal
      - [ ] Install PostgreSQL 15
      - [ ] Buat database: `createdb webifylab_dev`
      - [ ] Buat user: `createuser webifylab_dev`
    - [ ] Option B: Pakai Docker untuk PostgreSQL
      - [ ] `docker run -d --name webifylab-postgres -e POSTGRES_DB=webifylab_dev -e POSTGRES_PASSWORD=dev -p 5432:5432 postgres:15-alpine`
    - [ ] Test koneksi: `psql -U webifylab_dev -d webifylab_dev -h localhost`
  - **Acceptance Criteria:**
    - PostgreSQL bisa diakses dari lokal
    - Database `webifylab_dev` siap digunakan

##### 🗂️ Kategori: Documentation

- [ ] **TASK-013: Update README.md dengan project overview**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-006
  - **Detail:**
    - [ ] Tulis overview project
    - [ ] Tulis tech stack
    - [ ] Tulis cara setup local development
    - [ ] Tulis cara deployment
    - [ ] Tulis struktur folder
    - [ ] Link ke dokumen lain (PRD, Database Schema, API Spec, Design Spec)
  - **Acceptance Criteria:**
    - README.md informatif dan lengkap
    - New developer bisa setup project hanya dengan membaca README

- [ ] **TASK-014: Simpan semua dokumen di folder /docs**
  - **Estimasi:** 30 menit
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-006
  - **Detail:**
    - [ ] Copy PRD ke `/docs/PRD.md`
    - [ ] Copy Database Schema ke `/docs/DATABASE_SCHEMA.md`
    - [ ] Copy API Specification ke `/docs/API_SPECIFICATION.md`
    - [ ] Copy Design Specification ke `/docs/DESIGN_SPECIFICATION.md`
    - [ ] Copy Sprint Roadmap ini ke `/docs/SPRINT_ROADMAP.md`
    - [ ] Commit semua dokumen
  - **Acceptance Criteria:**
    - Semua dokumen terpusat di `/docs`
    - Mudah diakses saat development

---

#### 3.2 Sprint 0 Summary

| Kategori | Jumlah Task | Estimasi Total |
|---|---|---|
| Infrastructure & VPS | 5 tasks | 6 jam |
| Repository Setup | 5 tasks | 7 jam |
| Local Development | 2 tasks | 2 jam |
| Documentation | 2 tasks | 1.5 jam |
| **TOTAL** | **14 tasks** | **~16.5 jam** |

**Sprint 0 Progress Tracker:**

```
┌─────────────────────────────────────────────────────────────┐
│  SPRINT 0 PROGRESS                                          │
├─────────────────────────────────────────────────────────────┤
│  [ ] TASK-001: Hapus previous deployment                    │
│  [ ] TASK-002: Setup DNS A record                           │
│  [ ] TASK-003: Setup SSL certificate                        │
│  [ ] TASK-004: Setup swap space                             │
│  [ ] TASK-005: Verifikasi Docker & Nginx                    │
│  [ ] TASK-006: Init monorepo repository                     │
│  [ ] TASK-007: Setup .gitignore                             │
│  [ ] TASK-008: Setup Makefile                               │
│  [ ] TASK-009: Setup docker-compose.yml                     │
│  [ ] TASK-010: Setup .env.example                           │
│  [ ] TASK-011: Verifikasi local tools                       │
│  [ ] TASK-012: Setup PostgreSQL lokal                       │
│  [ ] TASK-013: Update README.md                             │
│  [ ] TASK-014: Simpan dokumen di /docs                      │
├─────────────────────────────────────────────────────────────┤
│  PROGRESS: 0/14 tasks (0%)                                  │
│  ████████████████████████████████████████████████████████   │
└─────────────────────────────────────────────────────────────┘
```

---

### 4. SPRINT 1: BACKEND FOUNDATION
**Durasi:** 2 minggu  
**Fokus:** Setup Golang project, database connection, dan Auth API  
**Goal:** Backend API dasar berfungsi dengan authentication

---

#### 4.1 Task List Sprint 1

##### 🗂️ Kategori: Golang Project Setup

- [ ] **TASK-101: Init Golang project structure**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 0 selesai
  - **Detail:**
    - [ ] Buat folder `backend/`
    - [ ] Init Go module: `go mod init webifylab`
    - [ ] Buat struktur folder sesuai Database Schema Document:
      ```
      backend/
      ├── cmd/
      │   └── server/
      │       └── main.go
      ├── internal/
      │   ├── config/
      │   ├── models/
      │   ├── handlers/
      │   ├── services/
      │   ├── repositories/
      │   ├── middleware/
      │   ├── dto/
      │   │   ├── request/
      │   │   └── response/
      │   └── routes/
      ├── migrations/
      ├── pkg/
      │   ├── response/
      │   ├── validator/
      │   └── hasher/
      ├── uploads/
      ├── Dockerfile
      ├── go.mod
      └── go.sum
      ```
    - [ ] Install dependencies:
      ```bash
      go get github.com/gin-gonic/gin
      go get gorm.io/gorm
      go get gorm.io/driver/postgres
      go get github.com/golang-jwt/jwt/v5
      go get github.com/google/uuid
      go get golang.org/x/crypto/bcrypt
      go get github.com/spf13/viper
      go get github.com/rs/zerolog
      go get github.com/golang-migrate/migrate/v4
      go get github.com/resend/resend-go/v2
      ```
    - [ ] Buat `main.go` dengan basic Gin server
    - [ ] Test server berjalan: `go run cmd/server/main.go`
  - **Acceptance Criteria:**
    - Server Gin berjalan di port 8080
    - Health check endpoint `/api/v1/health` return 200
    - Struktur folder sesuai dengan dokumentasi

- [ ] **TASK-102: Setup configuration dengan Viper**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `internal/config/config.go`
    - [ ] Load environment variables dari `.env`
    - [ ] Definisi struct Config:
      ```go
      type Config struct {
          AppEnv      string
          AppPort     string
          DBHost      string
          DBPort      string
          DBUser      string
          DBPassword  string
          DBName      string
          JWTSecret   string
          JWTExpiry   time.Duration
          ResendKey   string
          AdminEmail  string
          FrontendURL string
      }
      ```
    - [ ] Test config loading
  - **Acceptance Criteria:**
    - Config ter-load dari environment variables
    - Default values berfungsi jika env tidak ada

- [ ] **TASK-103: Setup database connection dengan GORM**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-102
  - **Detail:**
    - [ ] Buat `internal/config/database.go`
    - [ ] Setup GORM connection ke PostgreSQL
    - [ ] Setup connection pool settings
    - [ ] Test koneksi berhasil
    - [ ] Setup auto-migrate untuk development (disable untuk production)
  - **Acceptance Criteria:**
    - Koneksi ke PostgreSQL berhasil
    - Connection pool configured
    - Bisa execute query sederhana

- [ ] **TASK-104: Setup logging dengan Zerolog**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `pkg/logger/logger.go`
    - [ ] Setup Zerolog dengan level berdasarkan environment
    - [ ] Integrasikan dengan Gin middleware untuk request logging
    - [ ] Format log JSON untuk production, pretty untuk development
  - **Acceptance Criteria:**
    - Semua request ter-log dengan method, path, status, duration
    - Error ter-log dengan stack trace

##### 🗂️ Kategori: Database Models

- [ ] **TASK-105: Buat semua GORM models**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-103
  - **Detail:**
    - [ ] Buat model `User` di `internal/models/user.go`
    - [ ] Buat model `RefreshToken` di `internal/models/refresh_token.go`
    - [ ] Buat model `PasswordReset` di `internal/models/password_reset.go`
    - [ ] Buat model `Category` di `internal/models/category.go`
    - [ ] Buat model `Tag` di `internal/models/tag.go`
    - [ ] Buat model `BlogPost` di `internal/models/blog_post.go`
    - [ ] Buat model `PostTag` di `internal/models/post_tag.go`
    - [ ] Buat model `CaseStudy` di `internal/models/case_study.go`
    - [ ] Buat model `CaseStudyTag` di `internal/models/case_study_tag.go`
    - [ ] Buat model `Portfolio` di `internal/models/portfolio.go`
    - [ ] Buat model `Service` di `internal/models/service.go`
    - [ ] Buat model `Lead` di `internal/models/lead.go`
    - [ ] Buat model `Media` di `internal/models/media.go`
    - [ ] Buat model `SiteSetting` di `internal/models/site_setting.go`
    - [ ] Pastikan semua models sesuai dengan Database Schema Document
  - **Acceptance Criteria:**
    - Semua 14 models terdefinisi
    - Relationships antar models benar
    - GORM tags sesuai dengan schema

- [ ] **TASK-106: Buat database migrations**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-105
  - **Detail:**
    - [ ] Install golang-migrate CLI
    - [ ] Buat migration files:
      - `000001_init_schema.up.sql`
      - `000001_init_schema.down.sql`
      - `000002_seed_data.up.sql`
      - `000002_seed_data.down.sql`
    - [ ] Copy SQL dari Database Schema Document
    - [ ] Test migrate up: `migrate -path migrations -database "postgresql://..." up`
    - [ ] Test migrate down: `migrate -path migrations -database "postgresql://..." down`
    - [ ] Buat CLI command untuk migrate: `./server migrate`
  - **Acceptance Criteria:**
    - Migration berhasil dijalankan tanpa error
    - Semua tabel ter-create sesuai schema
    - Rollback berfungsi

- [ ] **TASK-107: Seed initial data**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-106
  - **Detail:**
    - [ ] Buat seeder untuk admin user default
    - [ ] Buat seeder untuk categories default
    - [ ] Buat seeder untuk tags default
    - [ ] Buat seeder untuk services default
    - [ ] Buat seeder untuk site settings default
    - [ ] Buat CLI command untuk seed: `./server seed`
    - [ ] Test seed berhasil
  - **Acceptance Criteria:**
    - Admin user default bisa login
    - Categories, tags, services, settings ter-seed
    - Seed idempotent (tidak duplicate jika dijalankan 2x)

##### 🗂️ Kategori: Middleware

- [ ] **TASK-108: Buat CORS middleware**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `internal/middleware/cors.go`
    - [ ] Allow origin: `https://webifylab.my.id` dan `http://localhost:3000` (dev)
    - [ ] Allow methods: GET, POST, PUT, DELETE, OPTIONS
    - [ ] Allow headers: Content-Type, Authorization
    - [ ] Allow credentials: true
    - [ ] Test CORS berfungsi
  - **Acceptance Criteria:**
    - Request dari frontend diterima
    - Request dari origin lain ditolak

- [ ] **TASK-109: Buat JWT middleware**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `internal/middleware/auth.go`
    - [ ] Parse JWT dari Authorization header
    - [ ] Validate token signature dan expiry
    - [ ] Extract user ID dan role dari token
    - [ ] Set user context di Gin
    - [ ] Return 401 jika token invalid/expired
    - [ ] Buat middleware role check: `RequireRole("admin")`
  - **Acceptance Criteria:**
    - Request tanpa token → 401
    - Request dengan token expired → 401
    - Request dengan token valid → lanjut ke handler
    - Role check berfungsi

- [ ] **TASK-110: Buat rate limiting middleware**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `internal/middleware/ratelimit.go`
    - [ ] Implementasi rate limiting per IP
    - [ ] Konfigurasi limit per endpoint:
      - `/auth/login`: 5 req/menit
      - `/auth/forgot-password`: 3 req/jam
      - `/leads`: 5 req/jam
      - Default: 100 req/menit
    - [ ] Return 429 jika limit exceeded
    - [ ] Set headers: X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset
  - **Acceptance Criteria:**
    - Rate limiting berfungsi per endpoint
    - Response 429 dengan retry_after
    - Headers ter-set dengan benar

##### 🗂️ Kategori: Auth API

- [ ] **TASK-111: Implementasi login endpoint**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-109, TASK-105
  - **Detail:**
    - [ ] Buat `internal/handlers/auth.go`
    - [ ] Buat `internal/services/auth_service.go`
    - [ ] Buat `internal/repositories/user_repo.go`
    - [ ] Implementasi `POST /api/v1/auth/login`:
      - [ ] Validate email dan password
      - [ ] Cari user by email
      - [ ] Verify password dengan bcrypt
      - [ ] Generate access token (JWT, 24 jam)
      - [ ] Generate refresh token (random string, 7 hari)
      - [ ] Simpan refresh token ke database
      - [ ] Return tokens + user data
    - [ ] Test dengan Postman/curl
  - **Acceptance Criteria:**
    - Login dengan valid credentials → 200 + tokens
    - Login dengan invalid credentials → 401
    - Login dengan email tidak terdaftar → 401
    - Response format sesuai API Specification

- [ ] **TASK-112: Implementasi refresh token endpoint**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-111
  - **Detail:**
    - [ ] Implementasi `POST /api/v1/auth/refresh`:
      - [ ] Validate refresh token
      - [ ] Cari refresh token di database
      - [ ] Cek expiry
      - [ ] Generate access token baru
      - [ ] Return new access token
    - [ ] Test dengan Postman/curl
  - **Acceptance Criteria:**
    - Refresh dengan valid token → 200 + new access token
    - Refresh dengan invalid token → 401
    - Refresh dengan expired token → 401

- [ ] **TASK-113: Implementasi logout endpoint**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-111
  - **Detail:**
    - [ ] Implementasi `POST /api/v1/auth/logout`:
      - [ ] Delete refresh token dari database (optional)
      - [ ] Return success message
    - [ ] Test dengan Postman/curl
  - **Acceptance Criteria:**
    - Logout berhasil → 200
    - Refresh token tidak bisa dipakai lagi

- [ ] **TASK-114: Implementasi get current user endpoint**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-109
  - **Detail:**
    - [ ] Implementasi `GET /api/v1/auth/me`:
      - [ ] Get user ID dari JWT context
      - [ ] Fetch user dari database
      - [ ] Return user profile (tanpa password)
    - [ ] Test dengan Postman/curl
  - **Acceptance Criteria:**
    - Get me dengan valid token → 200 + user data
    - Get me tanpa token → 401
    - Password tidak ter-expose di response

- [ ] **TASK-115: Implementasi forgot password endpoint**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-105, TASK-111
  - **Detail:**
    - [ ] Implementasi `POST /api/v1/auth/forgot-password`:
      - [ ] Validate email
      - [ ] Cari user by email
      - [ ] Jika user ada, generate reset token
      - [ ] Simpan reset token ke database (expiry 1 jam)
      - [ ] Kirim email via Resend (async)
      - [ ] Return success message (selalu success untuk prevent enumeration)
    - [ ] Test dengan Postman/curl
  - **Acceptance Criteria:**
    - Forgot password dengan email terdaftar → 200 + email terkirim
    - Forgot password dengan email tidak terdaftar → 200 (tidak expose)
    - Reset token expired dalam 1 jam

- [ ] **TASK-116: Implementasi reset password endpoint**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-115
  - **Detail:**
    - [ ] Implementasi `POST /api/v1/auth/reset-password`:
      - [ ] Validate token dan password baru
      - [ ] Cari reset token di database
      - [ ] Cek expiry dan unused
      - [ ] Update password user (bcrypt hash)
      - [ ] Mark token sebagai used
      - [ ] Invalidate semua refresh tokens user
    - [ ] Test dengan Postman/curl
  - **Acceptance Criteria:**
    - Reset dengan valid token → 200 + password ter-update
    - Reset dengan expired token → 422
    - Reset dengan used token → 422
    - Login dengan password baru berhasil

##### 🗂️ Kategori: Utilities

- [ ] **TASK-117: Buat standard response helper**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `pkg/response/response.go`
    - [ ] Implementasi helper functions:
      - `Success(c, message, data)`
      - `Created(c, message, data)`
      - `Error(c, status, message, errors)`
      - `Paginated(c, message, data, pagination)`
    - [ ] Definisi struct Response sesuai API Specification
    - [ ] Test semua helpers
  - **Acceptance Criteria:**
    - Response format konsisten di semua endpoints
    - Pagination structure sesuai API Specification

- [ ] **TASK-118: Buat validation helper**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `pkg/validator/validator.go`
    - [ ] Setup go-playground/validator
    - [ ] Custom validation rules:
      - `email`: valid email format
      - `phone`: valid phone format
      - `slug`: valid slug format
      - `enum`: valid enum value
    - [ ] Format validation errors sesuai API Specification
  - **Acceptance Criteria:**
    - Validation errors ter-format dengan jelas
    - Field-level errors bisa diidentifikasi

- [ ] **TASK-119: Buat password hasher**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `pkg/hasher/hasher.go`
    - [ ] Implementasi bcrypt hash dan verify
    - [ ] Cost factor: 10
    - [ ] Test hash dan verify
  - **Acceptance Criteria:**
    - Password ter-hash dengan bcrypt
    - Verify password berfungsi
    - Plain text password tidak pernah disimpan

- [ ] **TASK-120: Setup Dockerfile untuk backend**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-101
  - **Detail:**
    - [ ] Buat `backend/Dockerfile` dengan multi-stage build:
      - Build stage: `golang:1.22-alpine`
      - Runtime stage: `alpine:3.19`
    - [ ] Copy binary dan migrations ke runtime
    - [ ] Set environment variables
    - [ ] Expose port 8080
    - [ ] Test build: `docker build -t webifylab-backend backend/`
    - [ ] Test run: `docker run -p 8080:8080 webifylab-backend`
  - **Acceptance Criteria:**
    - Docker image berhasil di-build
    - Container berjalan dan serve API
    - Image size < 100MB

---

#### 4.2 Sprint 1 Summary

| Kategori | Jumlah Task | Estimasi Total |
|---|---|---|
| Golang Project Setup | 4 tasks | 9 jam |
| Database Models | 3 tasks | 9 jam |
| Middleware | 3 tasks | 6 jam |
| Auth API | 6 tasks | 12 jam |
| Utilities | 4 tasks | 7 jam |
| **TOTAL** | **20 tasks** | **~43 jam** |

**Sprint 1 Progress Tracker:**

```
┌─────────────────────────────────────────────────────────────┐
│  SPRINT 1 PROGRESS                                          │
├─────────────────────────────────────────────────────────────┤
│  Golang Project Setup:                                      │
│  [ ] TASK-101: Init Golang project structure                │
│  [ ] TASK-102: Setup configuration dengan Viper             │
│  [ ] TASK-103: Setup database connection dengan GORM        │
│  [ ] TASK-104: Setup logging dengan Zerolog                 │
│                                                             │
│  Database Models:                                           │
│  [ ] TASK-105: Buat semua GORM models                       │
│  [ ] TASK-106: Buat database migrations                     │
│  [ ] TASK-107: Seed initial data                            │
│                                                             │
│  Middleware:                                                │
│  [ ] TASK-108: Buat CORS middleware                         │
│  [ ] TASK-109: Buat JWT middleware                          │
│  [ ] TASK-110: Buat rate limiting middleware                │
│                                                             │
│  Auth API:                                                  │
│  [ ] TASK-111: Implementasi login endpoint                  │
│  [ ] TASK-112: Implementasi refresh token endpoint          │
│  [ ] TASK-113: Implementasi logout endpoint                 │
│  [ ] TASK-114: Implementasi get current user endpoint       │
│  [ ] TASK-115: Implementasi forgot password endpoint        │
│  [ ] TASK-116: Implementasi reset password endpoint         │
│                                                             │
│  Utilities:                                                 │
│  [ ] TASK-117: Buat standard response helper                │
│  [ ] TASK-118: Buat validation helper                       │
│  [ ] TASK-119: Buat password hasher                         │
│  [ ] TASK-120: Setup Dockerfile untuk backend               │
├─────────────────────────────────────────────────────────────┤
│  PROGRESS: 0/20 tasks (0%)                                  │
│  ████████████████████████████████████████████████████████   │
└─────────────────────────────────────────────────────────────┘
```

---

### 5. SPRINT 2: CONTENT API
**Durasi:** 2 minggu  
**Fokus:** CRUD API untuk semua content (Blog, Case Study, Portfolio, Services, Leads, Media)  
**Goal:** Semua content API berfungsi dan siap di-consume frontend

---

#### 5.1 Task List Sprint 2

##### 🗂️ Kategori: Blog API

- [ ] **TASK-201: Implementasi Blog repository**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/blog_repo.go`
    - [ ] Implementasi methods:
      - `List(ctx, params) ([]BlogPost, int64, error)` — dengan pagination, filter, search
      - `GetBySlug(ctx, slug) (*BlogPost, error)`
      - `GetByID(ctx, id) (*BlogPost, error)`
      - `Create(ctx, post) (*BlogPost, error)`
      - `Update(ctx, id, post) (*BlogPost, error)`
      - `Delete(ctx, id) error` — soft delete
      - `ListAdmin(ctx, params) ([]BlogPost, int64, error)` — termasuk drafts
    - [ ] Preload relationships: Author, Category, Tags
    - [ ] Test semua methods
  - **Acceptance Criteria:**
    - Semua CRUD operations berfungsi
    - Pagination, filter, search berfungsi
    - Soft delete berfungsi

- [ ] **TASK-202: Implementasi Blog service**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-201
  - **Detail:**
    - [ ] Buat `internal/services/blog_service.go`
    - [ ] Business logic:
      - Auto-generate slug dari title jika tidak disediakan
      - Cek slug uniqueness
      - Set published_at saat status berubah ke published
      - Calculate read time dari content length
      - Handle tags association
    - [ ] Test semua business logic
  - **Acceptance Criteria:**
    - Slug auto-generation berfungsi
    - Published_at ter-set saat publish
    - Read time ter-calculate

- [ ] **TASK-203: Implementasi Blog handlers**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-202
  - **Detail:**
    - [ ] Buat `internal/handlers/blog.go`
    - [ ] Implementasi handlers:
      - `List` — GET /blog (public, only published)
      - `GetBySlug` — GET /blog/:slug (public)
      - `ListAdmin` — GET /blog/admin (admin, all statuses)
      - `Create` — POST /blog (admin)
      - `Update` — PUT /blog/:id (admin)
      - `Delete` — DELETE /blog/:id (admin)
    - [ ] Bind request body ke DTO
    - [ ] Validate request
    - [ ] Call service
    - [ ] Return response sesuai API Specification
    - [ ] Test semua endpoints dengan Postman/curl
  - **Acceptance Criteria:**
    - Semua 6 endpoints berfungsi
    - Response format sesuai API Specification
    - Validation errors ter-handle

##### 🗂️ Kategori: Case Study API

- [ ] **TASK-204: Implementasi Case Study repository**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/case_study_repo.go`
    - [ ] Implementasi methods:
      - `List(ctx, params) ([]CaseStudy, int64, error)` — filter by industry, ai_type
      - `GetBySlug(ctx, slug) (*CaseStudy, error)`
      - `Create(ctx, caseStudy) (*CaseStudy, error)`
      - `Update(ctx, id, caseStudy) (*CaseStudy, error)`
      - `Delete(ctx, id) error` — soft delete
    - [ ] Preload relationships: Tags
    - [ ] Test semua methods
  - **Acceptance Criteria:**
    - Semua CRUD operations berfungsi
    - Filter by industry dan ai_type berfungsi
    - JSONB fields (metrics, images, ai_type) ter-handle

- [ ] **TASK-205: Implementasi Case Study service**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-204
  - **Detail:**
    - [ ] Buat `internal/services/case_study_service.go`
    - [ ] Business logic:
      - Auto-generate slug
      - Validate ai_type values
      - Handle tags association
    - [ ] Test semua business logic
  - **Acceptance Criteria:**
    - Slug auto-generation berfungsi
    - ai_type validation berfungsi

- [ ] **TASK-206: Implementasi Case Study handlers**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-205
  - **Detail:**
    - [ ] Buat `internal/handlers/case_study.go`
    - [ ] Implementasi handlers:
      - `List` — GET /case-studies (public)
      - `GetBySlug` — GET /case-studies/:slug (public)
      - `Create` — POST /case-studies (admin)
      - `Update` — PUT /case-studies/:id (admin)
      - `Delete` — DELETE /case-studies/:id (admin)
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 5 endpoints berfungsi
    - Response format sesuai API Specification

##### 🗂️ Kategori: Portfolio API

- [ ] **TASK-207: Implementasi Portfolio repository**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/portfolio_repo.go`
    - [ ] Implementasi methods:
      - `List(ctx, params) ([]Portfolio, int64, error)` — filter by category
      - `GetBySlug(ctx, slug) (*Portfolio, error)`
      - `Create(ctx, portfolio) (*Portfolio, error)`
      - `Update(ctx, id, portfolio) (*Portfolio, error)`
      - `Delete(ctx, id) error` — soft delete
      - `Reorder(ctx, items) error`
    - [ ] Order by `order_index`
    - [ ] Test semua methods
  - **Acceptance Criteria:**
    - Semua CRUD operations berfungsi
    - Filter by category berfungsi
    - Reorder berfungsi

- [ ] **TASK-208: Implementasi Portfolio service**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-207
  - **Detail:**
    - [ ] Buat `internal/services/portfolio_service.go`
    - [ ] Business logic:
      - Auto-generate slug
      - Validate category values
    - [ ] Test semua business logic
  - **Acceptance Criteria:**
    - Slug auto-generation berfungsi
    - Category validation berfungsi

- [ ] **TASK-209: Implementasi Portfolio handlers**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-208
  - **Detail:**
    - [ ] Buat `internal/handlers/portfolio.go`
    - [ ] Implementasi handlers:
      - `List` — GET /portfolio (public)
      - `GetBySlug` — GET /portfolio/:slug (public)
      - `Create` — POST /portfolio (admin)
      - `Update` — PUT /portfolio/:id (admin)
      - `Reorder` — PUT /portfolio/reorder (admin)
      - `Delete` — DELETE /portfolio/:id (admin)
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 6 endpoints berfungsi
    - Response format sesuai API Specification

##### 🗂️ Kategori: Services API

- [ ] **TASK-210: Implementasi Services repository, service, dan handlers**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/service_repo.go`
    - [ ] Buat `internal/services/service_service.go`
    - [ ] Buat `internal/handlers/service.go`
    - [ ] Implementasi endpoints:
      - `List` — GET /services (public, only active)
      - `GetBySlug` — GET /services/:slug (public)
      - `Create` — POST /services (admin)
      - `Update` — PUT /services/:id (admin)
      - `Delete` — DELETE /services/:id (admin) — hard delete
    - [ ] Order by `order_index`
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 5 endpoints berfungsi
    - Only active services ter-return di public endpoint
    - Order by order_index

##### 🗂️ Kategori: Leads API

- [ ] **TASK-211: Implementasi Leads repository**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/lead_repo.go`
    - [ ] Implementasi methods:
      - `List(ctx, params) ([]Lead, int64, error)` — filter by status, service_type
      - `GetByID(ctx, id) (*Lead, error)`
      - `Create(ctx, lead) (*Lead, error)`
      - `Update(ctx, id, lead) (*Lead, error)`
      - `Delete(ctx, id) error` — hard delete
      - `CountByStatus(ctx) (map[string]int64, error)`
      - `CountByMonth(ctx, months) ([]MonthlyCount, error)`
      - `Export(ctx, params) ([]Lead, error)`
    - [ ] Test semua methods
  - **Acceptance Criteria:**
    - Semua CRUD operations berfungsi
    - Statistics methods berfungsi

- [ ] **TASK-212: Implementasi Leads service dengan email integration**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-211
  - **Detail:**
    - [ ] Buat `internal/services/lead_service.go`
    - [ ] Business logic:
      - Validate lead data
      - Simpan lead ke database
      - Kirim auto-reply email ke user via Resend
      - Kirim notification email ke admin via Resend
      - Email sending async (goroutine)
    - [ ] Buat email templates:
      - Auto-reply ke user
      - Notification ke admin
    - [ ] Test semua business logic
  - **Acceptance Criteria:**
    - Lead ter-simpan ke database
    - Auto-reply email terkirim ke user
    - Notification email terkirim ke admin
    - Email sending tidak blocking

- [ ] **TASK-213: Implementasi Leads handlers**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-212
  - **Detail:**
    - [ ] Buat `internal/handlers/lead.go`
    - [ ] Implementasi handlers:
      - `Submit` — POST /leads (public, rate-limited)
      - `List` — GET /leads (admin)
      - `Export` — GET /leads/export (admin, CSV download)
      - `GetByID` — GET /leads/:id (admin)
      - `Update` — PUT /leads/:id (admin)
      - `Delete` — DELETE /leads/:id (admin)
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 6 endpoints berfungsi
    - Rate limiting berfungsi di POST /leads
    - CSV export berfungsi

##### 🗂️ Kategori: Media API

- [ ] **TASK-214: Implementasi Media repository**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/media_repo.go`
    - [ ] Implementasi methods:
      - `List(ctx, params) ([]Media, int64, error)`
      - `GetByID(ctx, id) (*Media, error)`
      - `Create(ctx, media) (*Media, error)`
      - `Delete(ctx, id) error`
    - [ ] Test semua methods
  - **Acceptance Criteria:**
    - Semua CRUD operations berfungsi

- [ ] **TASK-215: Implementasi Media service dengan image processing**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-214
  - **Detail:**
    - [ ] Buat `internal/services/media_service.go`
    - [ ] Business logic:
      - Validate file type dan size
      - Generate unique filename dengan UUID
      - Compress image (quality 80%)
      - Generate WebP version (optional)
      - Extract dimensions (width, height)
      - Save file ke `/uploads/media/`
      - Simpan metadata ke database
      - Delete file dari storage saat delete record
    - [ ] Test semua business logic
  - **Acceptance Criteria:**
    - File ter-upload dan ter-compress
    - Metadata ter-simpan ke database
    - File ter-delete saat record di-delete
    - Only JPG, PNG, WebP diterima
    - Max 5MB enforced

- [ ] **TASK-216: Implementasi Media handlers**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-215
  - **Detail:**
    - [ ] Buat `internal/handlers/media.go`
    - [ ] Implementasi handlers:
      - `Upload` — POST /media/upload (admin, multipart/form-data)
      - `List` — GET /media (admin)
      - `Delete` — DELETE /media/:id (admin)
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 3 endpoints berfungsi
    - File upload berfungsi
    - Error handling untuk invalid files

##### 🗂️ Kategori: Categories & Tags API

- [ ] **TASK-217: Implementasi Categories repository, service, dan handlers**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/category_repo.go`
    - [ ] Buat `internal/services/category_service.go`
    - [ ] Buat `internal/handlers/category.go`
    - [ ] Implementasi endpoints:
      - `List` — GET /categories (public, filter by type)
      - `Create` — POST /categories (admin)
      - `Update` — PUT /categories/:id (admin)
      - `Delete` — DELETE /categories/:id (admin)
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 4 endpoints berfungsi
    - Filter by type berfungsi

- [ ] **TASK-218: Implementasi Tags repository, service, dan handlers**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/tag_repo.go`
    - [ ] Buat `internal/services/tag_service.go`
    - [ ] Buat `internal/handlers/tag.go`
    - [ ] Implementasi endpoints:
      - `List` — GET /tags (public)
      - `Create` — POST /tags (admin)
      - `Delete` — DELETE /tags/:id (admin)
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 3 endpoints berfungsi
    - Tag cascade delete berfungsi

##### 🗂️ Kategori: Dashboard & Settings API

- [ ] **TASK-219: Implementasi Dashboard handlers**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-201, TASK-204, TASK-207, TASK-211
  - **Detail:**
    - [ ] Buat `internal/handlers/dashboard.go`
    - [ ] Implementasi endpoints:
      - `Stats` — GET /dashboard/stats (admin)
      - `LeadsChart` — GET /dashboard/leads-chart (admin)
      - `RecentLeads` — GET /dashboard/recent-leads (admin)
      - `RecentPosts` — GET /dashboard/recent-posts (admin)
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 4 endpoints berfungsi
    - Statistics akurat

- [ ] **TASK-220: Implementasi Settings repository, service, dan handlers**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** Sprint 1 selesai
  - **Detail:**
    - [ ] Buat `internal/repositories/setting_repo.go`
    - [ ] Buat `internal/services/setting_service.go`
    - [ ] Buat `internal/handlers/setting.go`
    - [ ] Implementasi endpoints:
      - `GetPublic` — GET /settings (public)
      - `GetAdmin` — GET /settings/admin (admin)
      - `Update` — PUT /settings (admin, bulk update)
    - [ ] Test semua endpoints
  - **Acceptance Criteria:**
    - Semua 3 endpoints berfungsi
    - Bulk update berfungsi

##### 🗂️ Kategori: Route Registration & Integration

- [ ] **TASK-221: Register semua routes di Gin**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-203, TASK-206, TASK-209, TASK-210, TASK-213, TASK-216, TASK-217, TASK-218, TASK-219, TASK-220
  - **Detail:**
    - [ ] Buat `internal/routes/routes.go`
    - [ ] Register semua endpoints sesuai API Specification
    - [ ] Group routes: public, protected, admin
    - [ ] Apply middleware yang sesuai
    - [ ] Test semua routes ter-register
  - **Acceptance Criteria:**
    - Semua 52 endpoints ter-register
    - Middleware ter-apply dengan benar
    - Route grouping benar

- [ ] **TASK-222: Setup health check endpoint**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-103
  - **Detail:**
    - [ ] Buat `internal/handlers/health.go`
    - [ ] Implementasi `GET /api/v1/health`:
      - [ ] Cek database connection
      - [ ] Return status, timestamp, version, uptime
    - [ ] Test endpoint
  - **Acceptance Criteria:**
    - Health check return 200 jika semua sehat
    - Health check return 503 jika ada masalah

- [ ] **TASK-223: Integration testing semua endpoints**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-221
  - **Detail:**
    - [ ] Test semua 52 endpoints dengan Postman/curl
    - [ ] Verifikasi response format sesuai API Specification
    - [ ] Test error cases
    - [ ] Test rate limiting
    - [ ] Test authentication flow
    - [ ] Dokumentasikan hasil testing
  - **Acceptance Criteria:**
    - Semua endpoints berfungsi sesuai API Specification
    - Error handling benar
    - Tidak ada breaking changes

---

#### 5.2 Sprint 2 Summary

| Kategori | Jumlah Task | Estimasi Total |
|---|---|---|
| Blog API | 3 tasks | 8 jam |
| Case Study API | 3 tasks | 7 jam |
| Portfolio API | 3 tasks | 5 jam |
| Services API | 1 task | 3 jam |
| Leads API | 3 tasks | 7 jam |
| Media API | 3 tasks | 8 jam |
| Categories & Tags API | 2 tasks | 4 jam |
| Dashboard & Settings API | 2 tasks | 5 jam |
| Route Registration | 3 tasks | 7 jam |
| **TOTAL** | **23 tasks** | **~54 jam** |

**Sprint 2 Progress Tracker:**

```
┌─────────────────────────────────────────────────────────────┐
│  SPRINT 2 PROGRESS                                          │
├─────────────────────────────────────────────────────────────┤
│  Blog API:                                                  │
│  [ ] TASK-201: Implementasi Blog repository                 │
│  [ ] TASK-202: Implementasi Blog service                    │
│  [ ] TASK-203: Implementasi Blog handlers                   │
│                                                             │
│  Case Study API:                                            │
│  [ ] TASK-204: Implementasi Case Study repository           │
│  [ ] TASK-205: Implementasi Case Study service              │
│  [ ] TASK-206: Implementasi Case Study handlers             │
│                                                             │
│  Portfolio API:                                             │
│  [ ] TASK-207: Implementasi Portfolio repository            │
│  [ ] TASK-208: Implementasi Portfolio service               │
│  [ ] TASK-209: Implementasi Portfolio handlers              │
│                                                             │
│  Services API:                                              │
│  [ ] TASK-210: Implementasi Services API                    │
│                                                             │
│  Leads API:                                                 │
│  [ ] TASK-211: Implementasi Leads repository                │
│  [ ] TASK-212: Implementasi Leads service + email           │
│  [ ] TASK-213: Implementasi Leads handlers                  │
│                                                             │
│  Media API:                                                 │
│  [ ] TASK-214: Implementasi Media repository                │
│  [ ] TASK-215: Implementasi Media service + image proc      │
│  [ ] TASK-216: Implementasi Media handlers                  │
│                                                             │
│  Categories & Tags API:                                     │
│  [ ] TASK-217: Implementasi Categories API                  │
│  [ ] TASK-218: Implementasi Tags API                        │
│                                                             │
│  Dashboard & Settings API:                                  │
│  [ ] TASK-219: Implementasi Dashboard handlers              │
│  [ ] TASK-220: Implementasi Settings API                    │
│                                                             │
│  Route Registration & Integration:                          │
│  [ ] TASK-221: Register semua routes                        │
│  [ ] TASK-222: Setup health check endpoint                  │
│  [ ] TASK-223: Integration testing semua endpoints          │
├─────────────────────────────────────────────────────────────┤
│  PROGRESS: 0/23 tasks (0%)                                  │
│  ████████████████████████████████████████████████████████   │
└─────────────────────────────────────────────────────────────┘
```

---

### 6. SPRINT 3: FRONTEND PUBLIC
**Durasi:** 2 minggu  
**Fokus:** Setup Next.js dan semua halaman publik  
**Goal:** Semua halaman publik live dan bisa diakses visitor

---

#### 6.1 Task List Sprint 3

##### 🗂️ Kategori: Next.js Project Setup

- [ ] **TASK-301: Init Next.js project**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 2 selesai
  - **Detail:**
    - [ ] Buat folder `frontend/`
    - [ ] Init Next.js: `npx create-next-app@latest frontend --typescript --tailwind --eslint --app --src-dir`
    - [ ] Install dependencies:
      ```bash
      npm install axios framer-motion lucide-react
      npm install react-hook-form zod @hookform/resolvers
      npm install @tiptap/react @tiptap/starter-kit
      npm install recharts
      npm install clsx tailwind-merge
      ```
    - [ ] Setup folder structure:
      ```
      frontend/
      ├── src/
      │   ├── app/
      │   │   ├── (public)/
      │   │   ├── admin/
      │   │   └── api/
      │   ├── components/
      │   │   ├── ui/
      │   │   ├── public/
      │   │   └── admin/
      │   ├── lib/
      │   ├── hooks/
      │   ├── types/
      │   └── styles/
      └── public/
      ```
    - [ ] Test server berjalan: `npm run dev`
  - **Acceptance Criteria:**
    - Next.js server berjalan di port 3000
    - Default page bisa diakses
    - TypeScript dan Tailwind berfungsi

- [ ] **TASK-302: Setup Tailwind config dengan brand colors**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-301
  - **Detail:**
    - [ ] Update `tailwind.config.ts` dengan color palette dari Design Spec
    - [ ] Setup font family: Inter, Plus Jakarta Sans, JetBrains Mono
    - [ ] Test colors berfungsi
  - **Acceptance Criteria:**
    - Brand colors bisa dipakai di semua components
    - Fonts ter-load dengan benar

- [ ] **TASK-303: Setup API client dengan Axios**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-301
  - **Detail:**
    - [ ] Buat `src/lib/api.ts`
    - [ ] Setup Axios instance dengan baseURL dari env
    - [ ] Setup request interceptor (auth token)
    - [ ] Setup response interceptor (error handling, token refresh)
    - [ ] Buat API services untuk setiap module:
      - `src/lib/services/blog.ts`
      - `src/lib/services/case-study.ts`
      - `src/lib/services/portfolio.ts`
      - `src/lib/services/service.ts`
      - `src/lib/services/lead.ts`
      - `src/lib/services/auth.ts`
      - `src/lib/services/settings.ts`
    - [ ] Test API client berfungsi
  - **Acceptance Criteria:**
    - API client bisa call backend API
    - Error handling berfungsi
    - Token refresh otomatis saat expired

- [ ] **TASK-304: Setup TypeScript types**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-301
  - **Detail:**
    - [ ] Buat `src/types/` folder
    - [ ] Definisi types untuk semua entities:
      - `User`, `BlogPost`, `CaseStudy`, `Portfolio`, `Service`, `Lead`, `Media`, `Category`, `Tag`, `SiteSetting`
    - [ ] Definisi types untuk API responses:
      - `ApiResponse<T>`, `PaginatedResponse<T>`, `ErrorResponse`
    - [ ] Definisi types untuk forms:
      - `CreateBlogInput`, `UpdateBlogInput`, `SubmitLeadInput`, dll
  - **Acceptance Criteria:**
    - Semua types terdefinisi
    - Types sesuai dengan Database Schema dan API Specification

##### 🗂️ Kategori: Layout & Navigation

- [ ] **TASK-305: Buat Header component**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-302
  - **Detail:**
    - [ ] Buat `src/components/public/header.tsx`
    - [ ] Implementasi desktop navigation
    - [ ] Implementasi mobile hamburger menu
    - [ ] Sticky header dengan backdrop blur saat scroll
    - [ ] CTA buttons: "Mulai Eksperimen", "Login"
    - [ ] Active state untuk current page
    - [ ] Responsive behavior
  - **Acceptance Criteria:**
    - Header berfungsi di desktop dan mobile
    - Navigation links berfungsi
    - Mobile menu buka/tutup smooth
    - Active state ter-highlight

- [ ] **TASK-306: Buat Footer component**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-302
  - **Detail:**
    - [ ] Buat `src/components/public/footer.tsx`
    - [ ] Implementasi footer layout sesuai Design Spec
    - [ ] Sections: Services, Company, Resources, Contact
    - [ ] Social media links
    - [ ] Copyright notice
    - [ ] Data dari `GET /settings`
  - **Acceptance Criteria:**
    - Footer berfungsi dan responsive
    - Links berfungsi
    - Data ter-load dari API

- [ ] **TASK-307: Buat public layout**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-305, TASK-306
  - **Detail:**
    - [ ] Buat `src/app/(public)/layout.tsx`
    - [ ] Integrasikan Header dan Footer
    - [ ] Setup metadata (title, description, OG tags)
    - [ ] Setup Google Analytics / Plausible (optional)
  - **Acceptance Criteria:**
    - Semua public pages pakai layout ini
    - Header dan Footer konsisten di semua halaman
    - Metadata ter-set

##### 🗂️ Kategori: UI Components

- [ ] **TASK-308: Buat Button component**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-302
  - **Detail:**
    - [ ] Buat `src/components/ui/button.tsx`
    - [ ] Variants: primary, secondary, ghost, outline
    - [ ] Sizes: sm, md, lg
    - [ ] States: default, hover, active, disabled, loading
    - [ ] Loading state dengan spinner
  - **Acceptance Criteria:**
    - Semua variants berfungsi
    - Loading state berfungsi
    - Accessibility: focus visible

- [ ] **TASK-309: Buat Card components**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-302
  - **Detail:**
    - [ ] Buat `src/components/public/blog-card.tsx`
    - [ ] Buat `src/components/public/case-study-card.tsx`
    - [ ] Buat `src/components/public/portfolio-card.tsx`
    - [ ] Buat `src/components/public/service-card.tsx`
    - [ ] Implementasi hover effects
    - [ ] Implementasi image lazy loading
  - **Acceptance Criteria:**
    - Semua card components berfungsi
    - Hover effects smooth
    - Images lazy load

- [ ] **TASK-310: Buat Badge component**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-302
  - **Detail:**
    - [ ] Buat `src/components/ui/badge.tsx`
    - [ ] Variants: category, status, ai-type
    - [ ] Colors sesuai Design Spec
  - **Acceptance Criteria:**
    - Semua variants berfungsi
    - Colors sesuai Design Spec

- [ ] **TASK-311: Buat Form components**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-302
  - **Detail:**
    - [ ] Buat `src/components/ui/input.tsx`
    - [ ] Buat `src/components/ui/textarea.tsx`
    - [ ] Buat `src/components/ui/select.tsx`
    - [ ] Buat `src/components/ui/label.tsx`
    - [ ] Error states dengan messages
    - [ ] Integration dengan React Hook Form
  - **Acceptance Criteria:**
    - Semua form components berfungsi
    - Error states ter-display
    - Integration dengan RHF berfungsi

- [ ] **TASK-312: Buat Skeleton loading component**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-302
  - **Detail:**
    - [ ] Buat `src/components/ui/skeleton.tsx`
    - [ ] Variants: text, image, card
    - [ ] Pulse animation
  - **Acceptance Criteria:**
    - Skeleton loading berfungsi
    - Animation smooth

- [ ] **TASK-313: Buat Pagination component**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-302
  - **Detail:**
    - [ ] Buat `src/components/ui/pagination.tsx`
    - [ ] Previous/Next buttons
    - [ ] Page numbers
    - [ ] Disabled states
    - [ ] Responsive (hide page numbers di mobile)
  - **Acceptance Criteria:**
    - Pagination berfungsi
    - Navigation antar pages berfungsi

##### 🗂️ Kategori: Public Pages

- [ ] **TASK-314: Buat Home page**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-307, TASK-309
  - **Detail:**
    - [ ] Buat `src/app/(public)/page.tsx`
    - [ ] Implementasi sections sesuai Design Spec:
      - [ ] Hero section dengan headline "We Experiment, You Grow"
      - [ ] Layanan Unggulan (GET /services)
      - [ ] Kenapa WebifyLab?
      - [ ] Eksperimen Terbaru (GET /case-studies?limit=3)
      - [ ] Catatan Belajar (GET /blog?limit=3)
      - [ ] Client Logos / Testimonials
      - [ ] CTA Section
    - [ ] Framer Motion animations (fade-in on scroll)
    - [ ] Responsive design
    - [ ] Loading states (skeleton)
    - [ ] Error states
  - **Acceptance Criteria:**
    - Semua sections ter-render
    - Data ter-load dari API
    - Animations smooth
    - Responsive di semua breakpoints
    - Lighthouse score ≥ 90

- [ ] **TASK-315: Buat About page**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-307
  - **Detail:**
    - [ ] Buat `src/app/(public)/about/page.tsx`
    - [ ] Implementasi sections:
      - [ ] Hero section
      - [ ] Filosofi Kami (4 cards)
      - [ ] Tim Kami
      - [ ] Tech Stack Kami
      - [ ] CTA Section
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Semua sections ter-render
    - Responsive di semua breakpoints

- [ ] **TASK-316: Buat Services page**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-307, TASK-309
  - **Detail:**
    - [ ] Buat `src/app/(public)/services/page.tsx`
    - [ ] Fetch services dari `GET /services`
    - [ ] Implementasi service cards
    - [ ] Implementasi "Bagaimana Kami Bekerja" section
    - [ ] CTA per service → Contact page dengan pre-filled service
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Services ter-load dari API
    - CTA links berfungsi
    - Responsive di semua breakpoints

- [ ] **TASK-317: Buat Portfolio page**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-307, TASK-309, TASK-313
  - **Detail:**
    - [ ] Buat `src/app/(public)/portfolio/page.tsx`
    - [ ] Fetch portfolios dari `GET /portfolio`
    - [ ] Implementasi filter tabs (Semua, Web, Mobile, AI, E-commerce)
    - [ ] Implementasi portfolio cards grid
    - [ ] Implementasi pagination
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Portfolios ter-load dari API
    - Filter berfungsi
    - Pagination berfungsi
    - Responsive di semua breakpoints

- [ ] **TASK-318: Buat Portfolio detail page**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-317
  - **Detail:**
    - [ ] Buat `src/app/(public)/portfolio/[slug]/page.tsx`
    - [ ] Fetch portfolio dari `GET /portfolio/:slug`
    - [ ] Implementasi sections:
      - [ ] Title + meta + badges
      - [ ] Hero image
      - [ ] Overview
      - [ ] Gallery
      - [ ] Lessons Learned
      - [ ] Related portfolios
    - [ ] Back link ke Portfolio page
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Portfolio detail ter-load dari API
    - Semua sections ter-render
    - Related portfolios ter-load
    - 404 page untuk slug tidak ada

- [ ] **TASK-319: Buat Case Studies page**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-307, TASK-309, TASK-313
  - **Detail:**
    - [ ] Buat `src/app/(public)/case-studies/page.tsx`
    - [ ] Fetch case studies dari `GET /case-studies`
    - [ ] Implementasi filter by industry dan ai_type
    - [ ] Implementasi case study cards
    - [ ] Implementasi pagination
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Case studies ter-load dari API
    - Filter berfungsi
    - Pagination berfungsi
    - Responsive di semua breakpoints

- [ ] **TASK-320: Buat Case Study detail page**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-319
  - **Detail:**
    - [ ] Buat `src/app/(public)/case-studies/[slug]/page.tsx`
    - [ ] Fetch case study dari `GET /case-studies/:slug`
    - [ ] Implementasi sections:
      - [ ] Title + meta + badges
      - [ ] Challenge
      - [ ] Approach
      - [ ] Solution + images
      - [ ] Results + metrics cards
      - [ ] Lessons Learned
      - [ ] Testimonial
      - [ ] CTA "Ingin Eksperimen Serupa?"
    - [ ] Back link ke Case Studies page
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Case study detail ter-load dari API
    - Semua sections ter-render
    - Metrics cards ter-render
    - 404 page untuk slug tidak ada

- [ ] **TASK-321: Buat Blog page**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-307, TASK-309, TASK-313
  - **Detail:**
    - [ ] Buat `src/app/(public)/blog/page.tsx`
    - [ ] Fetch blog posts dari `GET /blog`
    - [ ] Implementasi search bar
    - [ ] Implementasi filter by category
    - [ ] Implementasi blog cards grid
    - [ ] Implementasi pagination
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Blog posts ter-load dari API
    - Search berfungsi
    - Filter berfungsi
    - Pagination berfungsi
    - Responsive di semua breakpoints

- [ ] **TASK-322: Buat Blog detail page**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-321
  - **Detail:**
    - [ ] Buat `src/app/(public)/blog/[slug]/page.tsx`
    - [ ] Fetch blog post dari `GET /blog/:slug`
    - [ ] Implementasi sections:
      - [ ] Title + meta + badges
      - [ ] Cover image
      - [ ] Table of Contents (auto-generated, sticky di desktop)
      - [ ] Article content (rich text rendering)
      - [ ] Code syntax highlighting
      - [ ] Related articles
      - [ ] Share buttons
    - [ ] Reading progress bar
    - [ ] Back link ke Blog page
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Blog post ter-load dari API
    - Semua sections ter-render
    - Code blocks ter-highlight
    - TOC auto-generated
    - Share buttons berfungsi
    - 404 page untuk slug tidak ada

- [ ] **TASK-323: Buat Contact page**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-307, TASK-311
  - **Detail:**
    - [ ] Buat `src/app/(public)/contact/page.tsx`
    - [ ] Implementasi contact form dengan React Hook Form + Zod
    - [ ] Fields: Nama, Email, WhatsApp, Jenis Layanan, Budget Range, Deskripsi
    - [ ] Client-side validation
    - [ ] Submit ke `POST /leads`
    - [ ] Success message setelah submit
    - [ ] Error handling
    - [ ] Contact info section (dari GET /settings)
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Form berfungsi dengan validation
    - Submit berhasil ke API
    - Success message ter-display
    - Error handling berfungsi
    - Responsive di semua breakpoints

- [ ] **TASK-324: Buat Products page (placeholder)**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-307
  - **Detail:**
    - [ ] Buat `src/app/(public)/products/page.tsx`
    - [ ] Implementasi placeholder untuk SaaS products
    - [ ] Section "Eksperimen Selanjutnya"
    - [ ] Responsive design
  - **Acceptance Criteria:**
    - Page ter-render
    - Responsive di semua breakpoints

- [ ] **TASK-325: Buat 404 page**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-307
  - **Detail:**
    - [ ] Buat `src/app/not-found.tsx`
    - [ ] Design sesuai brand
    - [ ] Link kembali ke Home
  - **Acceptance Criteria:**
    - 404 page ter-render untuk routes tidak ada
    - Link kembali berfungsi

##### 🗂️ Kategori: SEO & Performance

- [ ] **TASK-326: Setup metadata untuk semua pages**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-314 s/d TASK-325
  - **Detail:**
    - [ ] Setup metadata di setiap page:
      - title
      - description
      - openGraph
      - twitter
    - [ ] Dynamic metadata untuk detail pages (blog, case study, portfolio)
    - [ ] Canonical URLs
  - **Acceptance Criteria:**
    - Semua pages punya metadata unik
    - OG tags ter-set
    - Canonical URLs benar

- [ ] **TASK-327: Setup sitemap dan robots.txt**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-326
  - **Detail:**
    - [ ] Buat `src/app/sitemap.ts` — dynamic sitemap
    - [ ] Buat `src/app/robots.ts` — robots.txt
    - [ ] Include semua public pages
    - [ ] Exclude admin pages
  - **Acceptance Criteria:**
    - sitemap.xml ter-generate
    - robots.txt ter-generate
    - Admin pages ter-exclude

- [ ] **TASK-328: Setup structured data (JSON-LD)**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-326
  - **Detail:**
    - [ ] Organization schema di Home page
    - [ ] Article schema di Blog detail
    - [ ] Service schema di Services page
    - [ ] FAQ schema (jika ada FAQ)
  - **Acceptance Criteria:**
    - Structured data valid (test dengan Google Rich Results Test)
    - Semua schemas ter-include

- [ ] **TASK-329: Optimize images dan fonts**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-314 s/d TASK-325
  - **Detail:**
    - [ ] Pakai Next.js Image component untuk semua images
    - [ ] Setup lazy loading
    - [ ] Setup responsive sizes
    - [ ] Optimize fonts dengan next/font
    - [ ] Preload critical fonts
  - **Acceptance Criteria:**
    - Images lazy load
    - Fonts ter-load tanpa FOUT
    - Lighthouse Performance ≥ 90

- [ ] **TASK-330: Setup Dockerfile untuk frontend**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-301
  - **Detail:**
    - [ ] Buat `frontend/Dockerfile` dengan multi-stage build:
      - Build stage: `node:20-alpine`
      - Runtime stage: `node:20-alpine` dengan standalone output
    - [ ] Setup next.config.ts dengan `output: 'standalone'`
    - [ ] Test build: `docker build -t webifylab-frontend frontend/`
    - [ ] Test run: `docker run -p 3000:3000 webifylab-frontend`
  - **Acceptance Criteria:**
    - Docker image berhasil di-build
    - Container berjalan dan serve frontend
    - Image size < 200MB

---

#### 6.2 Sprint 3 Summary

| Kategori | Jumlah Task | Estimasi Total |
|---|---|---|
| Next.js Project Setup | 4 tasks | 7 jam |
| Layout & Navigation | 3 tasks | 7 jam |
| UI Components | 6 tasks | 12 jam |
| Public Pages | 12 tasks | 31 jam |
| SEO & Performance | 5 tasks | 9 jam |
| **TOTAL** | **30 tasks** | **~66 jam** |

**Sprint 3 Progress Tracker:**

```
┌─────────────────────────────────────────────────────────────┐
│  SPRINT 3 PROGRESS                                          │
├─────────────────────────────────────────────────────────────┤
│  Next.js Project Setup:                                     │
│  [ ] TASK-301: Init Next.js project                         │
│  [ ] TASK-302: Setup Tailwind config                        │
│  [ ] TASK-303: Setup API client                             │
│  [ ] TASK-304: Setup TypeScript types                       │
│                                                             │
│  Layout & Navigation:                                       │
│  [ ] TASK-305: Buat Header component                        │
│  [ ] TASK-306: Buat Footer component                        │
│  [ ] TASK-307: Buat public layout                           │
│                                                             │
│  UI Components:                                             │
│  [ ] TASK-308: Buat Button component                        │
│  [ ] TASK-309: Buat Card components                         │
│  [ ] TASK-310: Buat Badge component                         │
│  [ ] TASK-311: Buat Form components                         │
│  [ ] TASK-312: Buat Skeleton loading                        │
│  [ ] TASK-313: Buat Pagination component                    │
│                                                             │
│  Public Pages:                                              │
│  [ ] TASK-314: Buat Home page                               │
│  [ ] TASK-315: Buat About page                              │
│  [ ] TASK-316: Buat Services page                           │
│  [ ] TASK-317: Buat Portfolio page                          │
│  [ ] TASK-318: Buat Portfolio detail page                   │
│  [ ] TASK-319: Buat Case Studies page                       │
│  [ ] TASK-320: Buat Case Study detail page                  │
│  [ ] TASK-321: Buat Blog page                               │
│  [ ] TASK-322: Buat Blog detail page                        │
│  [ ] TASK-323: Buat Contact page                            │
│  [ ] TASK-324: Buat Products page                           │
│  [ ] TASK-325: Buat 404 page                                │
│                                                             │
│  SEO & Performance:                                         │
│  [ ] TASK-326: Setup metadata                               │
│  [ ] TASK-327: Setup sitemap & robots.txt                   │
│  [ ] TASK-328: Setup structured data                        │
│  [ ] TASK-329: Optimize images & fonts                      │
│  [ ] TASK-330: Setup Dockerfile frontend                    │
├─────────────────────────────────────────────────────────────┤
│  PROGRESS: 0/30 tasks (0%)                                  │
│  ████████████████████████████████████████████████████████   │
└─────────────────────────────────────────────────────────────┘
```

---

### 7. SPRINT 4: ADMIN DASHBOARD
**Durasi:** 2 minggu  
**Fokus:** CMS dashboard lengkap untuk manage semua content  
**Goal:** Admin bisa manage blog, case study, portfolio, leads, dan media

---

#### 7.1 Task List Sprint 4

##### 🗂️ Kategori: Admin Layout & Auth

- [ ] **TASK-401: Buat Admin layout**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 3 selesai
  - **Detail:**
    - [ ] Buat `src/app/admin/layout.tsx`
    - [ ] Implementasi sidebar dengan navigation:
      - Dashboard
      - Blog
      - Case Studies
      - Portfolio
      - Leads
      - Media
      - Settings
    - [ ] Implementasi header dengan search + user menu
    - [ ] Sidebar collapsible di mobile
    - [ ] Dark sidebar (Deep Slate background)
    - [ ] Active state untuk current page
  - **Acceptance Criteria:**
    - Admin layout berfungsi
    - Sidebar navigation berfungsi
    - Responsive di mobile (hamburger)
    - Active state ter-highlight

- [ ] **TASK-402: Buat Login page**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-303, TASK-311
  - **Detail:**
    - [ ] Buat `src/app/admin/login/page.tsx`
    - [ ] Implementasi login form:
      - Email
      - Password
    - [ ] Client-side validation
    - [ ] Submit ke `POST /auth/login`
    - [ ] Simpan tokens di httpOnly cookies
    - [ ] Redirect ke /admin/dashboard setelah login
    - [ ] Error handling untuk invalid credentials
    - [ ] Loading state
  - **Acceptance Criteria:**
    - Login dengan valid credentials → redirect ke dashboard
    - Login dengan invalid credentials → error message
    - Loading state ter-display
    - Tokens ter-simpan di cookies

- [ ] **TASK-403: Buat auth middleware untuk admin routes**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-402
  - **Detail:**
    - [ ] Buat `src/middleware.ts` (Next.js middleware)
    - [ ] Check auth cookie untuk semua /admin routes (kecuali /admin/login)
    - [ ] Redirect ke /admin/login jika tidak authenticated
    - [ ] Auto-refresh token jika expired
  - **Acceptance Criteria:**
    - Access /admin tanpa login → redirect ke /admin/login
    - Access /admin dengan valid token → lanjut
    - Token expired → auto-refresh atau redirect ke login

- [ ] **TASK-404: Buat ProtectedRoute component**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-403
  - **Detail:**
    - [ ] Buat `src/components/admin/protected-route.tsx`
    - [ ] Wrap admin pages dengan component ini
    - [ ] Show loading saat check auth
    - [ ] Redirect ke login jika tidak authenticated
  - **Acceptance Criteria:**
    - Semua admin pages ter-protect
    - Loading state ter-display

##### 🗂️ Kategori: Dashboard Overview

- [ ] **TASK-405: Buat Admin Dashboard page**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-401, TASK-303
  - **Detail:**
    - [ ] Buat `src/app/admin/dashboard/page.tsx`
    - [ ] Implementasi sections:
      - [ ] Stats cards (total blog, case study, portfolio, leads)
      - [ ] Leads chart (bar chart, last 12 months) — pakai Recharts
      - [ ] Recent leads list
      - [ ] Recent posts list
    - [ ] Fetch data dari:
      - `GET /dashboard/stats`
      - `GET /dashboard/leads-chart`
      - `GET /dashboard/recent-leads`
      - `GET /dashboard/recent-posts`
    - [ ] Loading states (skeleton)
    - [ ] Error states
  - **Acceptance Criteria:**
    - Semua stats ter-display
    - Chart ter-render
    - Recent lists ter-display
    - Loading dan error states berfungsi

##### 🗂️ Kategori: Blog Management

- [ ] **TASK-406: Buat Blog list page (admin)**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-401, TASK-303
  - **Detail:**
    - [ ] Buat `src/app/admin/blog/page.tsx`
    - [ ] Implementasi:
      - [ ] Table dengan columns: Title, Status, Author, Date, Actions
      - [ ] Filter tabs: Semua, Published, Draft, Archived
      - [ ] Search bar
      - [ ] Pagination
      - [ ] Actions: Edit, Delete
      - [ ] Button "Tulis Artikel Baru"
    - [ ] Fetch data dari `GET /blog/admin`
    - [ ] Delete dengan confirmation modal
    - [ ] Loading states
  - **Acceptance Criteria:**
    - Blog list ter-display
    - Filter berfungsi
    - Search berfungsi
    - Pagination berfungsi
    - Delete berfungsi dengan confirmation

- [ ] **TASK-407: Buat Blog create/edit page**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-406
  - **Detail:**
    - [ ] Buat `src/app/admin/blog/create/page.tsx`
    - [ ] Buat `src/app/admin/blog/[id]/edit/page.tsx`
    - [ ] Implementasi form dengan fields:
      - [ ] Title (auto-generate slug)
      - [ ] Slug (editable)
      - [ ] Content (rich text editor — Tiptap)
      - [ ] Excerpt
      - [ ] Cover image (upload via Media Manager)
      - [ ] Category (select)
      - [ ] Tags (multi-select)
      - [ ] Status (draft/published/archived)
      - [ ] Meta title
      - [ ] Meta description
    - [ ] Rich text editor dengan Tiptap:
      - Headings, bold, italic
      - Lists
      - Links
      - Images
      - Code blocks
      - Tables
    - [ ] Preview before publish
    - [ ] Save as draft / Publish buttons
    - [ ] Auto-save draft (optional)
  - **Acceptance Criteria:**
    - Form berfungsi dengan semua fields
    - Rich text editor berfungsi
    - Image upload berfungsi
    - Preview berfungsi
    - Save dan publish berfungsi

- [ ] **TASK-408: Integrasikan Blog CRUD dengan API**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-407
  - **Detail:**
    - [ ] Integrasikan create dengan `POST /blog`
    - [ ] Integrasikan update dengan `PUT /blog/:id`
    - [ ] Integrasikan delete dengan `DELETE /blog/:id`
    - [ ] Success/error toasts
    - [ ] Redirect setelah save
  - **Acceptance Criteria:**
    - Create berhasil → redirect ke list + toast success
    - Update berhasil → toast success
    - Delete berhasil → toast success + refresh list

##### 🗂️ Kategori: Case Study Management

- [ ] **TASK-409: Buat Case Study list page (admin)**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-401, TASK-303
  - **Detail:**
    - [ ] Buat `src/app/admin/case-studies/page.tsx`
    - [ ] Implementasi table dengan columns: Title, Client, Industry, Status, Date, Actions
    - [ ] Filter by status
    - [ ] Pagination
    - [ ] Actions: Edit, Delete
    - [ ] Button "Tambah Case Study"
  - **Acceptance Criteria:**
    - Case study list ter-display
    - Filter berfungsi
    - Pagination berfungsi
    - Delete berfungsi

- [ ] **TASK-410: Buat Case Study create/edit page**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-409
  - **Detail:**
    - [ ] Buat `src/app/admin/case-studies/create/page.tsx`
    - [ ] Buat `src/app/admin/case-studies/[id]/edit/page.tsx`
    - [ ] Implementasi form dengan fields terstruktur:
      - [ ] Title
      - [ ] Slug
      - [ ] Client Name
      - [ ] Industry
      - [ ] Challenge (rich text)
      - [ ] Approach (rich text)
      - [ ] Solution (rich text)
      - [ ] AI Type (multi-select: SPK, Sistem Pakar, Data Analytics, ML)
      - [ ] Results (rich text)
      - [ ] Lessons (rich text)
      - [ ] Testimonial (textarea)
      - [ ] Metrics (dynamic key-value fields)
      - [ ] Cover image
      - [ ] Images (multiple upload)
      - [ ] Tags
      - [ ] Status
      - [ ] Meta title
      - [ ] Meta description
    - [ ] Metrics dynamic form (add/remove fields)
    - [ ] Preview before publish
  - **Acceptance Criteria:**
    - Form berfungsi dengan semua fields
    - Metrics dynamic form berfungsi
    - Multiple image upload berfungsi
    - Save dan publish berfungsi

- [ ] **TASK-411: Integrasikan Case Study CRUD dengan API**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-410
  - **Detail:**
    - [ ] Integrasikan create dengan `POST /case-studies`
    - [ ] Integrasikan update dengan `PUT /case-studies/:id`
    - [ ] Integrasikan delete dengan `DELETE /case-studies/:id`
    - [ ] Success/error toasts
  - **Acceptance Criteria:**
    - Semua CRUD operations berfungsi
    - Toasts ter-display

##### 🗂️ Kategori: Portfolio Management

- [ ] **TASK-412: Buat Portfolio list page (admin)**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-401, TASK-303
  - **Detail:**
    - [ ] Buat `src/app/admin/portfolio/page.tsx`
    - [ ] Implementasi table/grid dengan columns: Title, Category, Status, Order, Actions
    - [ ] Filter by status dan category
    - [ ] Pagination
    - [ ] Actions: Edit, Delete
    - [ ] Button "Tambah Portfolio"
    - [ ] Drag & drop untuk reorder (optional)
  - **Acceptance Criteria:**
    - Portfolio list ter-display
    - Filter berfungsi
    - Pagination berfungsi
    - Delete berfungsi

- [ ] **TASK-413: Buat Portfolio create/edit page**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-412
  - **Detail:**
    - [ ] Buat `src/app/admin/portfolio/create/page.tsx`
    - [ ] Buat `src/app/admin/portfolio/[id]/edit/page.tsx`
    - [ ] Implementasi form dengan fields:
      - [ ] Title
      - [ ] Slug
      - [ ] Description (rich text)
      - [ ] Category (select)
      - [ ] Tech Stack (dynamic tags input)
      - [ ] Live URL
      - [ ] Thumbnail (upload)
      - [ ] Images (multiple upload)
      - [ ] Lessons (rich text)
      - [ ] Order Index
      - [ ] Status
      - [ ] Meta title
      - [ ] Meta description
    - [ ] Tech stack dynamic input (add/remove)
  - **Acceptance Criteria:**
    - Form berfungsi dengan semua fields
    - Tech stack dynamic input berfungsi
    - Multiple image upload berfungsi

- [ ] **TASK-414: Integrasikan Portfolio CRUD dengan API**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-413
  - **Detail:**
    - [ ] Integrasikan create dengan `POST /portfolio`
    - [ ] Integrasikan update dengan `PUT /portfolio/:id`
    - [ ] Integrasikan delete dengan `DELETE /portfolio/:id`
    - [ ] Integrasikan reorder dengan `PUT /portfolio/reorder`
    - [ ] Success/error toasts
  - **Acceptance Criteria:**
    - Semua CRUD operations berfungsi
    - Reorder berfungsi

##### 🗂️ Kategori: Lead Management

- [ ] **TASK-415: Buat Leads list page (admin)**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-401, TASK-303
  - **Detail:**
    - [ ] Buat `src/app/admin/leads/page.tsx`
    - [ ] Implementasi table dengan columns: Name, Email, Service, Budget, Status, Date, Actions
    - [ ] Filter tabs: Semua, New, Contacted, Negotiation, Proposal Sent, Won, Lost
    - [ ] Search bar
    - [ ] Pagination
    - [ ] Actions: View, Copy WhatsApp
    - [ ] Button "Export CSV"
    - [ ] Status badges dengan colors
  - **Acceptance Criteria:**
    - Leads list ter-display
    - Filter berfungsi
    - Search berfungsi
    - Pagination berfungsi
    - Copy WhatsApp berfungsi
    - Export CSV berfungsi

- [ ] **TASK-416: Buat Lead detail modal**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-415
  - **Detail:**
    - [ ] Buat `src/components/admin/lead-detail-modal.tsx`
    - [ ] Implementasi modal dengan fields:
      - [ ] Name, Email, WhatsApp (dengan copy button)
      - [ ] Service Type, Budget Range
      - [ ] Description
      - [ ] Status (select untuk update)
      - [ ] Notes (textarea untuk update)
      - [ ] IP Address, User Agent
      - [ ] Created At
    - [ ] Update status dan notes via `PUT /leads/:id`
    - [ ] Save button
  - **Acceptance Criteria:**
    - Modal berfungsi
    - Semua data ter-display
    - Status dan notes bisa di-update
    - Copy WhatsApp berfungsi

##### 🗂️ Kategori: Media Management

- [ ] **TASK-417: Buat Media Manager page**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-401, TASK-303
  - **Detail:**
    - [ ] Buat `src/app/admin/media/page.tsx`
    - [ ] Implementasi:
      - [ ] Upload button (drag & drop + click)
      - [ ] Media grid dengan thumbnails
      - [ ] Filter by mime type
      - [ ] Search bar
      - [ ] Pagination
      - [ ] Actions: Copy URL, Delete
    - [ ] Upload ke `POST /media/upload`
    - [ ] Delete dengan confirmation modal
    - [ ] Upload progress indicator
    - [ ] File size dan type validation
  - **Acceptance Criteria:**
    - Media list ter-display
    - Upload berfungsi dengan progress
    - Copy URL berfungsi
    - Delete berfungsi
    - Validation berfungsi

##### 🗂️ Kategori: Settings

- [ ] **TASK-418: Buat Settings page**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-401, TASK-303
  - **Detail:**
    - [ ] Buat `src/app/admin/settings/page.tsx`
    - [ ] Implementasi form dengan fields:
      - [ ] Site Name
      - [ ] Tagline
      - [ ] Contact Email
      - [ ] Contact WhatsApp
      - [ ] Social Media Links (Instagram, LinkedIn, GitHub)
      - [ ] SEO Default Title
      - [ ] SEO Default Description
      - [ ] SEO OG Image
    - [ ] Fetch data dari `GET /settings/admin`
    - [ ] Save ke `PUT /settings`
    - [ ] Success/error toasts
  - **Acceptance Criteria:**
    - Settings ter-load dari API
    - Save berfungsi
    - Toasts ter-display

##### 🗂️ Kategori: Categories & Tags Management

- [ ] **TASK-419: Buat Categories & Tags management**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-401, TASK-303
  - **Detail:**
    - [ ] Buat `src/app/admin/categories/page.tsx`
    - [ ] Implementasi:
      - [ ] List categories dengan filter by type
      - [ ] Add category form
      - [ ] Edit category
      - [ ] Delete category
    - [ ] Buat `src/app/admin/tags/page.tsx`
    - [ ] Implementasi:
      - [ ] List tags
      - [ ] Add tag form
      - [ ] Delete tag
    - [ ] Integrasikan dengan API
  - **Acceptance Criteria:**
    - Categories dan tags ter-display
    - CRUD berfungsi
    - Validation berfungsi

##### 🗂️ Kategori: Polish & UX

- [ ] **TASK-420: Buat Toast notification system**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-401
  - **Detail:**
    - [ ] Buat `src/components/ui/toast.tsx`
    - [ ] Implementasi toast types: success, error, warning, info
    - [ ] Auto-dismiss setelah 5 detik
    - [ ] Manual dismiss
    - [ ] Position: top-right
    - [ ] Integrasikan dengan semua CRUD operations
  - **Acceptance Criteria:**
    - Toasts ter-display untuk semua operations
    - Auto-dismiss berfungsi
    - Manual dismiss berfungsi

- [ ] **TASK-421: Buat Confirmation modal component**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-401
  - **Detail:**
    - [ ] Buat `src/components/ui/confirm-modal.tsx`
    - [ ] Implementasi untuk delete confirmations
    - [ ] Customizable title, message, actions
    - [ ] Integrasikan dengan semua delete operations
  - **Acceptance Criteria:**
    - Modal berfungsi
    - Digunakan di semua delete operations

- [ ] **TASK-422: Responsive polish untuk admin**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-401 s/d TASK-419
  - **Detail:**
    - [ ] Test semua admin pages di mobile
    - [ ] Fix layout issues
    - [ ] Sidebar collapsible di mobile
    - [ ] Tables responsive (horizontal scroll)
    - [ ] Forms responsive
  - **Acceptance Criteria:**
    - Semua admin pages responsive
    - Tidak ada layout issues di mobile
    - Tables bisa di-scroll horizontal

---

#### 7.2 Sprint 4 Summary

| Kategori | Jumlah Task | Estimasi Total |
|---|---|---|
| Admin Layout & Auth | 4 tasks | 8 jam |
| Dashboard Overview | 1 task | 3 jam |
| Blog Management | 3 tasks | 9 jam |
| Case Study Management | 3 tasks | 8 jam |
| Portfolio Management | 3 tasks | 7 jam |
| Lead Management | 2 tasks | 5 jam |
| Media Management | 1 task | 3 jam |
| Settings | 1 task | 2 jam |
| Categories & Tags | 1 task | 2 jam |
| Polish & UX | 3 tasks | 5 jam |
| **TOTAL** | **22 tasks** | **~52 jam** |

**Sprint 4 Progress Tracker:**

```
┌─────────────────────────────────────────────────────────────┐
│  SPRINT 4 PROGRESS                                          │
├─────────────────────────────────────────────────────────────┤
│  Admin Layout & Auth:                                       │
│  [ ] TASK-401: Buat Admin layout                            │
│  [ ] TASK-402: Buat Login page                              │
│  [ ] TASK-403: Buat auth middleware                         │
│  [ ] TASK-404: Buat ProtectedRoute component                │
│                                                             │
│  Dashboard Overview:                                        │
│  [ ] TASK-405: Buat Admin Dashboard page                    │
│                                                             │
│  Blog Management:                                           │
│  [ ] TASK-406: Buat Blog list page                          │
│  [ ] TASK-407: Buat Blog create/edit page                   │
│  [ ] TASK-408: Integrasikan Blog CRUD                       │
│                                                             │
│  Case Study Management:                                     │
│  [ ] TASK-409: Buat Case Study list page                    │
│  [ ] TASK-410: Buat Case Study create/edit page             │
│  [ ] TASK-411: Integrasikan Case Study CRUD                 │
│                                                             │
│  Portfolio Management:                                      │
│  [ ] TASK-412: Buat Portfolio list page                     │
│  [ ] TASK-413: Buat Portfolio create/edit page              │
│  [ ] TASK-414: Integrasikan Portfolio CRUD                  │
│                                                             │
│  Lead Management:                                           │
│  [ ] TASK-415: Buat Leads list page                         │
│  [ ] TASK-416: Buat Lead detail modal                       │
│                                                             │
│  Media Management:                                          │
│  [ ] TASK-417: Buat Media Manager page                      │
│                                                             │
│  Settings:                                                  │
│  [ ] TASK-418: Buat Settings page                           │
│                                                             │
│  Categories & Tags:                                         │
│  [ ] TASK-419: Buat Categories & Tags management            │
│                                                             │
│  Polish & UX:                                               │
│  [ ] TASK-420: Buat Toast notification system               │
│  [ ] TASK-421: Buat Confirmation modal                      │
│  [ ] TASK-422: Responsive polish untuk admin                │
├─────────────────────────────────────────────────────────────┤
│  PROGRESS: 0/22 tasks (0%)                                  │
│  ████████████████████████████████████████████████████████   │
└─────────────────────────────────────────────────────────────┘
```

---

### 8. SPRINT 5: INTEGRATION & POLISH
**Durasi:** 2 minggu  
**Fokus:** Integrasi semua fitur, email, SEO, performance, dan bug fixing  
**Goal:** Semua fitur terintegrasi dan siap untuk testing

---

#### 8.1 Task List Sprint 5

##### 🗂️ Kategori: Email Integration

- [ ] **TASK-501: Setup Resend email service**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 4 selesai
  - **Detail:**
    - [ ] Verifikasi domain di Resend dashboard
    - [ ] Generate API key
    - [ ] Setup environment variable `RESEND_API_KEY`
    - [ ] Test koneksi ke Resend API
  - **Acceptance Criteria:**
    - Resend API key valid
    - Bisa kirim test email

- [ ] **TASK-502: Buat email templates**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-501
  - **Detail:**
    - [ ] Buat email template untuk:
      - [ ] Auto-reply ke user (setelah submit inquiry)
      - [ ] Notification ke admin (lead baru)
      - [ ] Forgot password
      - [ ] Reset password success
    - [ ] Design email sesuai brand (HTML)
    - [ ] Test render di email clients
  - **Acceptance Criteria:**
    - Semua templates ter-design
    - Render benar di email clients

- [ ] **TASK-503: Implementasi email sending di Golang**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-502
  - **Detail:**
    - [ ] Buat `internal/services/email_service.go`
    - [ ] Implementasi methods:
      - `SendAutoReply(ctx, lead) error`
      - `SendAdminNotification(ctx, lead) error`
      - `SendForgotPassword(ctx, user, token) error`
      - `SendResetSuccess(ctx, user) error`
    - [ ] Async sending dengan goroutine
    - [ ] Error logging jika gagal
    - [ ] Retry logic (optional)
    - [ ] Test semua email types
  - **Acceptance Criteria:**
    - Semua email types terkirim
    - Async sending tidak blocking
    - Errors ter-log

##### 🗂️ Kategori: Frontend-Backend Integration

- [ ] **TASK-504: End-to-end integration testing**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-503
  - **Detail:**
    - [ ] Test semua user flows:
      - [ ] Visitor → baca blog → submit inquiry → admin terima email
      - [ ] Admin login → create blog → publish → visitor bisa baca
      - [ ] Admin login → create case study → publish → visitor bisa baca
      - [ ] Admin login → upload media → pakai di blog
      - [ ] Admin login → manage leads → update status
      - [ ] Admin login → update settings → frontend ter-update
    - [ ] Test error cases
    - [ ] Test loading states
    - [ ] Dokumentasikan bugs yang ditemukan
  - **Acceptance Criteria:**
    - Semua user flows berfungsi
    - Tidak ada breaking bugs
    - Error handling benar

- [ ] **TASK-505: Fix integration bugs**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-504
  - **Detail:**
    - [ ] Fix semua bugs yang ditemukan di TASK-504
    - [ ] Re-test setelah fix
    - [ ] Update dokumentasi jika perlu
  - **Acceptance Criteria:**
    - Semua bugs ter-fix
    - Re-test passed

##### 🗂️ Kategori: SEO Optimization

- [ ] **TASK-506: SEO audit dan optimization**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-504
  - **Detail:**
    - [ ] Run Lighthouse audit untuk semua pages
    - [ ] Fix issues:
      - [ ] Missing meta descriptions
      - [ ] Missing alt text
      - [ ] Low contrast
      - [ ] Missing canonical
      - [ ] Slow pages
    - [ ] Verify sitemap dan robots.txt
    - [ ] Test structured data dengan Google Rich Results Test
    - [ ] Submit sitemap ke Google Search Console
  - **Acceptance Criteria:**
    - Lighthouse SEO score ≥ 90 untuk semua pages
    - Structured data valid
    - Sitemap ter-submit

- [ ] **TASK-507: Setup Google Search Console**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-506
  - **Detail:**
    - [ ] Daftarkan webifylab.my.id di Google Search Console
    - [ ] Verifikasi ownership
    - [ ] Submit sitemap
    - [ ] Monitor indexing
  - **Acceptance Criteria:**
    - Site ter-verifikasi di GSC
    - Sitemap ter-submit
    - Indexing dimulai

##### 🗂️ Kategori: Performance Optimization

- [ ] **TASK-508: Performance audit dan optimization**
  - **Estimasi:** 3 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-504
  - **Detail:**
    - [ ] Run Lighthouse performance audit
    - [ ] Fix issues:
      - [ ] Large images → compress, WebP
      - [ ] Unused JavaScript → code splitting
      - [ ] Slow API calls → caching
      - [ ] Render-blocking resources → defer
    - [ ] Setup caching headers untuk static assets
    - [ ] Test API response times
    - [ ] Optimize database queries (add indexes jika perlu)
  - **Acceptance Criteria:**
    - Lighthouse Performance ≥ 90 untuk semua pages
    - API response time < 500ms (P95)
    - TTFB < 500ms

- [ ] **TASK-509: Setup monitoring dan logging**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-508
  - **Detail:**
    - [ ] Setup Uptime Kuma (self-hosted) untuk monitoring
    - [ ] Monitor:
      - webifylab.my.id (frontend)
      - api.webifylab.my.id/health (backend)
    - [ ] Setup alert jika downtime (email/WhatsApp)
    - [ ] Setup log rotation di VPS
  - **Acceptance Criteria:**
    - Uptime Kuma berjalan
    - Alerts ter-configure
    - Logs ter-rotate

##### 🗂️ Kategori: Security Audit

- [ ] **TASK-510: Security checklist**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-504
  - **Detail:**
    - [ ] Verifikasi HTTPS aktif untuk semua domains
    - [ ] Verifikasi security headers:
      - [ ] X-Frame-Options
      - [ ] X-Content-Type-Options
      - [ ] X-XSS-Protection
      - [ ] Referrer-Policy
    - [ ] Verifikasi CORS hanya allow webifylab.my.id
    - [ ] Verifikasi rate limiting berfungsi
    - [ ] Verifikasi file upload validation
    - [ ] Verifikasi SQL injection prevention (GORM)
    - [ ] Verifikasi password hashing (bcrypt)
    - [ ] Verifikasi JWT expiry
    - [ ] Verifikasi tidak ada sensitive data ter-expose di logs
    - [ ] Verifikasi .env tidak ter-commit ke git
  - **Acceptance Criteria:**
    - Semua security checks passed
    - Tidak ada vulnerabilities critical

##### 🗂️ Kategori: Content & Data

- [ ] **TASK-511: Buat content initial**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-504
  - **Detail:**
    - [ ] Tulis 3 blog posts awal:
      - [ ] "Perkenalan WebifyLab: We Experiment, You Grow"
      - [ ] "Kenapa UMKM Butuh Website di 2026"
      - [ ] "Mengenal SPK: Sistem Pendukung Keputusan untuk Bisnis"
    - [ ] Buat 1 case study (bisa dari project sebelumnya atau fiktif untuk demo)
    - [ ] Upload portfolio items (jika ada project sebelumnya)
    - [ ] Update services descriptions
    - [ ] Update site settings
  - **Acceptance Criteria:**
    - 3 blog posts ter-publish
    - 1 case study ter-publish
    - Portfolio items ter-upload
    - Services dan settings ter-update

- [ ] **TASK-512: Test content management workflow**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-511
  - **Detail:**
    - [ ] Test create, edit, publish, archive blog post
    - [ ] Test create, edit, publish case study
    - [ ] Test upload media dan pakai di content
    - [ ] Test update settings
    - [ ] Verifikasi semua content ter-render benar di frontend
  - **Acceptance Criteria:**
    - Semua workflow berfungsi
    - Content ter-render benar

##### 🗂️ Kategori: Final Polish

- [ ] **TASK-513: Cross-browser testing**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-504
  - **Detail:**
    - [ ] Test di Chrome
    - [ ] Test di Firefox
    - [ ] Test di Safari (jika ada Mac)
    - [ ] Test di Edge
    - [ ] Fix issues yang ditemukan
  - **Acceptance Criteria:**
    - Semua browsers berfungsi tanpa issues
    - Layout konsisten

- [ ] **TASK-514: Mobile testing**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-504
  - **Detail:**
    - [ ] Test semua pages di mobile (iPhone/Android)
    - [ ] Test form submission di mobile
    - [ ] Test navigation di mobile
    - [ ] Test admin dashboard di mobile
    - [ ] Fix issues yang ditemukan
  - **Acceptance Criteria:**
    - Semua pages berfungsi di mobile
    - Forms berfungsi
    - Navigation berfungsi

- [ ] **TASK-515: Accessibility testing**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-504
  - **Detail:**
    - [ ] Run WAVE accessibility checker
    - [ ] Test keyboard navigation
    - [ ] Test screen reader (VoiceOver/NVDA)
    - [ ] Verify color contrast
    - [ ] Verify alt text untuk semua images
    - [ ] Verify focus states
    - [ ] Fix issues yang ditemukan
  - **Acceptance Criteria:**
    - Tidak ada accessibility errors critical
    - Keyboard navigation berfungsi
    - Contrast ratio ≥ 4.5:1

- [ ] **TASK-516: Final bug fixing**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-513, TASK-514, TASK-515
  - **Detail:**
    - [ ] Fix semua bugs yang ditemukan di testing
    - [ ] Re-test setelah fix
    - [ ] Update dokumentasi jika perlu
  - **Acceptance Criteria:**
    - Semua bugs ter-fix
    - Re-test passed

---

#### 8.2 Sprint 5 Summary

| Kategori | Jumlah Task | Estimasi Total |
|---|---|---|
| Email Integration | 3 tasks | 8 jam |
| Frontend-Backend Integration | 2 tasks | 8 jam |
| SEO Optimization | 2 tasks | 4 jam |
| Performance Optimization | 2 tasks | 5 jam |
| Security Audit | 1 task | 2 jam |
| Content & Data | 2 tasks | 6 jam |
| Final Polish | 4 tasks | 10 jam |
| **TOTAL** | **16 tasks** | **~43 jam** |

**Sprint 5 Progress Tracker:**

```
┌─────────────────────────────────────────────────────────────┐
│  SPRINT 5 PROGRESS                                          │
├─────────────────────────────────────────────────────────────┤
│  Email Integration:                                         │
│  [ ] TASK-501: Setup Resend email service                   │
│  [ ] TASK-502: Buat email templates                         │
│  [ ] TASK-503: Implementasi email sending                   │
│                                                             │
│  Frontend-Backend Integration:                              │
│  [ ] TASK-504: End-to-end integration testing               │
│  [ ] TASK-505: Fix integration bugs                         │
│                                                             │
│  SEO Optimization:                                          │
│  [ ] TASK-506: SEO audit dan optimization                   │
│  [ ] TASK-507: Setup Google Search Console                  │
│                                                             │
│  Performance Optimization:                                  │
│  [ ] TASK-508: Performance audit dan optimization           │
│  [ ] TASK-509: Setup monitoring dan logging                 │
│                                                             │
│  Security Audit:                                            │
│  [ ] TASK-510: Security checklist                           │
│                                                             │
│  Content & Data:                                            │
│  [ ] TASK-511: Buat content initial                         │
│  [ ] TASK-512: Test content management workflow             │
│                                                             │
│  Final Polish:                                              │
│  [ ] TASK-513: Cross-browser testing                        │
│  [ ] TASK-514: Mobile testing                               │
│  [ ] TASK-515: Accessibility testing                        │
│  [ ] TASK-516: Final bug fixing                             │
├─────────────────────────────────────────────────────────────┤
│  PROGRESS: 0/16 tasks (0%)                                  │
│  ████████████████████████████████████████████████████████   │
└─────────────────────────────────────────────────────────────┘
```

---

### 9. SPRINT 6: TESTING & LAUNCH
**Durasi:** 1 minggu  
**Fokus:** Final testing, deployment, dan launch  
**Goal:** Portal live di webifylab.my.id! 🚀

---

#### 9.1 Task List Sprint 6

##### 🗂️ Kategori: Final Testing

- [ ] **TASK-601: Full regression testing**
  - **Estimasi:** 4 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** Sprint 5 selesai
  - **Detail:**
    - [ ] Test semua 52 API endpoints
    - [ ] Test semua 11 public pages
    - [ ] Test semua admin pages
    - [ ] Test semua user flows end-to-end
    - [ ] Test error cases
    - [ ] Test edge cases
    - [ ] Dokumentasikan hasil testing
  - **Acceptance Criteria:**
    - Semua tests passed
    - Tidak ada breaking bugs
    - Performance acceptable

- [ ] **TASK-602: Load testing (optional)**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟢 Rendah
  - **Dependencies:** TASK-601
  - **Detail:**
    - [ ] Test dengan 100 concurrent users (pakai k6 atau ab)
    - [ ] Monitor response times
    - [ ] Monitor memory usage
    - [ ] Identifikasi bottlenecks
  - **Acceptance Criteria:**
    - Server tidak crash dengan 100 concurrent users
    - Response times acceptable

##### 🗂️ Kategori: Deployment

- [ ] **TASK-603: Final deployment preparation**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-601
  - **Detail:**
    - [ ] Backup database terakhir (jika ada data)
    - [ ] Verifikasi .env production sudah benar
    - [ ] Verifikasi SSL certificates valid
    - [ ] Verifikasi DNS records benar
    - [ ] Verifikasi Nginx configs benar
    - [ ] Verifikasi docker-compose.yml benar
    - [ ] Test deployment di staging (optional)
  - **Acceptance Criteria:**
    - Semua preparation selesai
    - Tidak ada configuration errors

- [ ] **TASK-604: Deploy ke production**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-603
  - **Detail:**
    - [ ] Build Docker images di laptop:
      ```bash
      make build
      ```
    - [ ] Save images:
      ```bash
      make deploy-images
      ```
    - [ ] Upload ke VPS:
      ```bash
      make upload
      ```
    - [ ] Start containers:
      ```bash
      make start
      ```
    - [ ] Run migrations:
      ```bash
      ssh user@vps 'cd /var/www/webifylab && docker-compose exec backend ./server migrate'
      ```
    - [ ] Seed data (jika fresh install):
      ```bash
      ssh user@vps 'cd /var/www/webifylab && docker-compose exec backend ./server seed'
      ```
    - [ ] Verify health check:
      ```bash
      curl https://api.webifylab.my.id/api/v1/health
      ```
    - [ ] Verify frontend:
      ```bash
      curl https://webifylab.my.id
      ```
  - **Acceptance Criteria:**
    - Containers berjalan
    - Health check return 200
    - Frontend bisa diakses
    - API bisa diakses

- [ ] **TASK-605: Setup Nginx production configs**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-604
  - **Detail:**
    - [ ] Copy Nginx configs dari `infra/nginx/` ke VPS
    - [ ] Enable sites:
      ```bash
      sudo ln -s /etc/nginx/sites-available/webifylab /etc/nginx/sites-enabled/
      sudo ln -s /etc/nginx/sites-available/api /etc/nginx/sites-enabled/
      ```
    - [ ] Test configs:
      ```bash
      sudo nginx -t
      ```
    - [ ] Reload Nginx:
      ```bash
      sudo systemctl reload nginx
      ```
    - [ ] Verify HTTPS berfungsi
    - [ ] Verify redirects berfungsi (HTTP → HTTPS)
  - **Acceptance Criteria:**
    - Nginx configs valid
    - HTTPS berfungsi untuk kedua domains
    - Redirects berfungsi

- [ ] **TASK-606: Setup backup automation**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-604
  - **Detail:**
    - [ ] Copy backup script dari `infra/scripts/backup.sh` ke VPS
    - [ ] Make executable: `chmod +x backup.sh`
    - [ ] Setup cron job untuk daily backup:
      ```bash
      sudo crontab -e
      # Tambahkan:
      0 2 * * * /var/www/webifylab/infra/scripts/backup.sh
      ```
    - [ ] Test backup manual
    - [ ] Verify backup files ter-create
  - **Acceptance Criteria:**
    - Backup script berfungsi
    - Cron job ter-configure
    - Backup files ter-create

- [ ] **TASK-607: Setup monitoring alerts**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-604
  - **Detail:**
    - [ ] Setup Uptime Kuma (jika belum)
    - [ ] Add monitors:
      - webifylab.my.id (HTTP, interval 1 menit)
      - api.webifylab.my.id/api/v1/health (HTTP, interval 1 menit)
    - [ ] Setup notifications (email/WhatsApp)
    - [ ] Test alert berfungsi
  - **Acceptance Criteria:**
    - Monitoring aktif
    - Alerts ter-configure
    - Test alert terkirim

##### 🗂️ Kategori: Launch

- [ ] **TASK-608: Final pre-launch checklist**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-604 s/d TASK-607
  - **Detail:**
    - [ ] Verifikasi semua pages berfungsi
    - [ ] Verifikasi contact form berfungsi
    - [ ] Verifikasi admin dashboard berfungsi
    - [ ] Verifikasi email notifications berfungsi
    - [ ] Verifikasi backup berfungsi
    - [ ] Verifikasi monitoring berfungsi
    - [ ] Verifikasi SSL valid
    - [ ] Verifikasi tidak ada error di logs
    - [ ] Verifikasi Lighthouse scores ≥ 90
  - **Acceptance Criteria:**
    - Semua checklist items passed
    - Tidak ada issues critical

- [ ] **TASK-609: LAUNCH! 🚀**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-608
  - **Detail:**
    - [ ] Umumkan launch di social media (optional)
    - [ ] Share ke network (optional)
    - [ ] Monitor traffic dan errors
    - [ ] Celebrate! 🎉
  - **Acceptance Criteria:**
    - Portal live di webifylab.my.id
    - API live di api.webifylab.my.id
    - Admin dashboard bisa diakses

- [ ] **TASK-610: Post-launch monitoring**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🔴 Tinggi
  - **Dependencies:** TASK-609
  - **Detail:**
    - [ ] Monitor logs untuk errors
    - [ ] Monitor memory usage
    - [ ] Monitor response times
    - [ ] Monitor email delivery
    - [ ] Fix issues yang muncul (hotfix)
    - [ ] Dokumentasikan lessons learned
  - **Acceptance Criteria:**
    - Tidak ada critical errors
    - Performance stabil
    - Lessons learned ter-dokumentasi

##### 🗂️ Kategori: Documentation

- [ ] **TASK-611: Update semua dokumentasi**
  - **Estimasi:** 2 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-609
  - **Detail:**
    - [ ] Update README.md dengan final setup
    - [ ] Update DEPLOYMENT.md dengan production steps
    - [ ] Dokumentasikan credentials (secure storage)
    - [ ] Dokumentasikan architecture decisions
    - [ ] Dokumentasikan known issues (jika ada)
  - **Acceptance Criteria:**
    - Semua dokumentasi up-to-date
    - Credentials ter-store securely
    - Known issues ter-dokumentasi

- [ ] **TASK-612: Retrospective Sprint 6**
  - **Estimasi:** 1 jam
  - **Prioritas:** 🟡 Sedang
  - **Dependencies:** TASK-610
  - **Detail:**
    - [ ] Review apa yang berjalan baik
    - [ ] Review apa yang bisa di-improve
    - [ ] Dokumentasikan lessons learned
    - [ ] Plan untuk Phase 2 (SaaS integration)
  - **Acceptance Criteria:**
    - Retrospective ter-dokumentasi
    - Lessons learned ter-capture
    - Phase 2 planning dimulai

---

#### 9.2 Sprint 6 Summary

| Kategori | Jumlah Task | Estimasi Total |
|---|---|---|
| Final Testing | 2 tasks | 6 jam |
| Deployment | 5 tasks | 9 jam |
| Launch | 3 tasks | 4 jam |
| Documentation | 2 tasks | 3 jam |
| **TOTAL** | **12 tasks** | **~22 jam** |

**Sprint 6 Progress Tracker:**

```
┌─────────────────────────────────────────────────────────────┐
│  SPRINT 6 PROGRESS                                          │
├─────────────────────────────────────────────────────────────┤
│  Final Testing:                                             │
│  [ ] TASK-601: Full regression testing                      │
│  [ ] TASK-602: Load testing (optional)                      │
│                                                             │
│  Deployment:                                                │
│  [ ] TASK-603: Final deployment preparation                 │
│  [ ] TASK-604: Deploy ke production                         │
│  [ ] TASK-605: Setup Nginx production configs               │
│  [ ] TASK-606: Setup backup automation                      │
│  [ ] TASK-607: Setup monitoring alerts                      │
│                                                             │
│  Launch:                                                    │
│  [ ] TASK-608: Final pre-launch checklist                   │
│  [ ] TASK-609: LAUNCH! 🚀                                   │
│  [ ] TASK-610: Post-launch monitoring                       │
│                                                             │
│  Documentation:                                             │
│  [ ] TASK-611: Update semua dokumentasi                     │
│  [ ] TASK-612: Retrospective Sprint 6                       │
├─────────────────────────────────────────────────────────────┤
│  PROGRESS: 0/12 tasks (0%)                                  │
│  ████████████████████████████████████████████████████████   │
└─────────────────────────────────────────────────────────────┘
```

---

### 10. OVERALL PROGRESS TRACKER

#### 10.1 Ringkasan Semua Sprint

| Sprint | Nama | Tasks | Estimasi | Status | Progress |
|---|---|---|---|---|---|
| Sprint 0 | Pre-Development Setup | 14 | 16.5 jam | ⬜ Belum Mulai | 0% |
| Sprint 1 | Backend Foundation | 20 | 43 jam | ⬜ Belum Mulai | 0% |
| Sprint 2 | Content API | 23 | 54 jam | ⬜ Belum Mulai | 0% |
| Sprint 3 | Frontend Public | 30 | 66 jam | ⬜ Belum Mulai | 0% |
| Sprint 4 | Admin Dashboard | 22 | 52 jam | ⬜ Belum Mulai | 0% |
| Sprint 5 | Integration & Polish | 16 | 43 jam | ⬜ Belum Mulai | 0% |
| Sprint 6 | Testing & Launch | 12 | 22 jam | ⬜ Belum Mulai | 0% |
| **TOTAL** | | **137 tasks** | **~296.5 jam** | | **0%** |

#### 10.2 Master Progress Tracker

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    WEBIFYLAB PORTAL - MASTER TRACKER                    │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  SPRINT 0: Pre-Development Setup                                        │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 0/14 (0%)                    │
│                                                                         │
│  SPRINT 1: Backend Foundation                                           │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 0/20 (0%)                    │
│                                                                         │
│  SPRINT 2: Content API                                                  │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 0/23 (0%)                    │
│                                                                         │
│  SPRINT 3: Frontend Public                                              │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 0/30 (0%)                    │
│                                                                         │
│  SPRINT 4: Admin Dashboard                                              │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 0/22 (0%)                    │
│                                                                         │
│  SPRINT 5: Integration & Polish                                         │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 0/16 (0%)                    │
│                                                                         │
│  SPRINT 6: Testing & Launch                                             │
│  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ 0/12 (0%)                    │
│                                                                         │
├─────────────────────────────────────────────────────────────────────────┤
│  OVERALL PROGRESS: 0/137 tasks (0%)                                     │
│  ████████████████████████████████████████████████████████████████████   │
│                                                                         │
│  ESTIMASI TOTAL: ~296.5 jam                                             │
│  JAM TERPAKAI: 0 jam                                                    │
│  SISA ESTIMASI: ~296.5 jam                                              │
│                                                                         │
│  STATUS: ⬜ BELUM DIMULAI                                                │
│  TARGET LAUNCH: Minggu 12                                               │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

#### 10.3 Weekly Review Template

Setiap akhir minggu, isi review ini:

```
📅 Minggu: _______________
🎯 Sprint: _______________

✅ TASKS SELESAI MINGGU INI:
   1. _______________
   2. _______________
   3. _______________

⏳ TASKS BELUM SELESAI:
   1. _______________ (alasan: _______________)
   2. _______________ (alasan: _______________)

📊 PROGRESS MINGGU INI: ___/___ tasks (___%)

🚧 HAMBATAN:
   1. _______________

💡 PELAJARAN:
   1. _______________

🎯 FOKUS MINGGU DEPAN:
   1. _______________
   2. _______________
   3. _______________

⚠️ RISIKO:
   1. _______________ (mitigasi: _______________)
```

---

### 11. RISK REGISTER

| No | Risiko | Probabilitas | Dampak | Mitigasi |
|---|---|---|---|---|
| R1 | Server 1GB RAM tidak cukup saat build | Tinggi | Tinggi | Build di laptop, deploy artifact saja |
| R2 | DNS propagation lambat | Sedang | Rendah | Setup DNS di awal Sprint 0, tunggu 24-48 jam |
| R3 | Resend API limit (100 emails/hari) | Rendah | Sedang | Monitor usage, upgrade plan jika perlu |
| R4 | Scope creep (fitur tambahan) | Tinggi | Sedang | Disiplin dengan PRD, fitur baru masuk Phase 2 |
| R5 | Burnout (solo dev) | Sedang | Tinggi | Istirahat cukup, jangan kerja > 8 jam/hari |
| R6 | Bug critical saat launch | Sedang | Tinggi | Testing thorough di Sprint 5 & 6 |
| R7 | Database corruption | Rendah | Tinggi | Daily backup, test restore |
| R8 | SSL certificate expired | Rendah | Tinggi | Auto-renew dengan Certbot + monitoring |
| R9 | Golang/Next.js version incompatibility | Rendah | Sedang | Lock versions di go.mod dan package.json |
| R10 | Kehilangan data saat deployment | Rendah | Tinggi | Backup sebelum deployment |

---

### 12. DEFINITION OF DONE (DoD) — RECAP

Sebuah task dianggap **SELESAI** jika:

- [ ] Kode sudah ditulis dan berjalan tanpa error
- [ ] Kode sudah di-test manual (minimal happy path)
- [ ] Kode sudah di-commit ke git dengan pesan yang jelas
- [ ] Dokumentasi sudah di-update (jika perlu)
- [ ] Tidak ada breaking changes ke fitur lain
- [ ] Acceptance criteria terpenuhi

---

### 13. CHANGELOG

| Versi | Tanggal | Perubahan |
|---|---|---|
| 1.0.0 | 2026-09-05 | Initial release — 137 tasks dalam 7 sprint |
