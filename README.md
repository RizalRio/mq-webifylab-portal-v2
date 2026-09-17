# 🚀 Webifylab Portal V2

> **Dari Ide Menjadi Ekosistem Digital**

Landing page profesional untuk **Webifylab** — software house yang spesialis dalam pengembangan aplikasi, desain grafis, dan ekosistem SaaS berbasis AI & Data.

[![Astro](https://img.shields.io/badge/Astro-7.x-FF5D01?style=flat-square&logo=astro&logoColor=white)](https://astro.build)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-4.x-06B6D4?style=flat-square&logo=tailwindcss&logoColor=white)](https://tailwindcss.com)
[![TypeScript](https://img.shields.io/badge/TypeScript-Ready-3178C6?style=flat-square&logo=typescript&logoColor=white)](https://www.typescriptlang.org)
[![License](https://img.shields.io/badge/License-Private-gray?style=flat-square)](./LICENSE)

---

## 📋 Daftar Isi

- [Overview](#-overview)
- [Tech Stack](#-tech-stack)
- [Struktur Monorepo](#-struktur-monorepo)
- [Quick Start](#-quick-start)
- [Sections Landing Page](#-sections-landing-page)
- [Konfigurasi](#-konfigurasi)
- [Deployment](#-deployment)
- [Roadmap](#-roadmap)
- [Dokumentasi](#-dokumentasi)

---

## 🎯 Overview

**Webifylab Landing Page V1** adalah aset digital pertama yang berfungsi sebagai wajah dan pintu masuk utama bagi calon klien. Dibangun dengan prinsip **Static-First** menggunakan Astro SSG untuk performa maksimal di VPS 1GB RAM.

### Target KPIs

| Metric | Target |
|--------|--------|
| Page Load Time | < 2 detik |
| Lighthouse Performance | > 90 |
| Bounce Rate | < 60% |
| Leads per bulan | ≥ 3 |

### Tiga Pilar Layanan

1. **Pengembangan Aplikasi** — Web, Mobile, API, Legacy Modernization
2. **Desain Grafis & Web Design** — UI/UX, Brand Identity, Design System
3. **SaaS & Ekosistem Digital** — Produk SaaS, AI & Data, Workflow Automation

---

## 🛠️ Tech Stack

### Frontend (V1 — Active)

| Layer | Teknologi | Versi |
|-------|-----------|-------|
| Framework | [Astro](https://astro.build) | 7.x |
| Styling | [Tailwind CSS](https://tailwindcss.com) | 4.x (CSS-first config) |
| Language | TypeScript | 5.x |
| Fonts | Inter + Plus Jakarta Sans | via Google Fonts |
| Form Handler | [Formspree](https://formspree.io) | Free tier |
| Analytics | Plausible / Google Analytics 4 | - |

### Backend (V1.5 — Planned)

| Layer | Teknologi | Versi |
|-------|-----------|-------|
| Language | [Go](https://go.dev) | 1.22+ |
| Router | Chi | Latest |
| Database | PostgreSQL | 15+ |
| ORM | sqlx / pgx | Latest |

### Infrastructure

| Layer | Teknologi |
|-------|-----------|
| VPS | 1 vCPU, 1GB RAM, 20GB SSD (Ubuntu 22.04) |
| Web Server | Nginx 1.24+ |
| SSL | Let's Encrypt + Certbot |
| Process Manager | Systemd (V1.5) |
| Monitoring | UptimeRobot + htop |

---

## 📁 Struktur Monorepo

```
mq-webifylab-portal-v2/
├── .github/
│   └── workflows/
│       └── deploy.yml              # CI/CD pipeline (GitHub Actions)
│
├── apps/
│   ├── web/                        # 🌐 Astro Frontend (V1 — Active)
│   │   ├── src/
│   │   │   ├── components/
│   │   │   │   ├── layout/         # Header.astro, Footer.astro
│   │   │   │   └── sections/       # Hero, Problem, Services, ...
│   │   │   ├── layouts/
│   │   │   │   └── Layout.astro    # Base layout + SEO meta tags
│   │   │   ├── lib/
│   │   │   │   └── constants.ts    # Site config & content data
│   │   │   ├── pages/
│   │   │   │   ├── index.astro     # Landing page utama
│   │   │   │   └── 404.astro       # Custom 404 page
│   │   │   └── styles/
│   │   │       └── global.css      # Design tokens + Tailwind v4
│   │   ├── public/                 # Static assets (favicon, robots.txt)
│   │   ├── astro.config.mjs
│   │   └── package.json
│   │
│   └── api/                        # ⚙️ Golang API (V1.5 — Planned)
│       ├── cmd/server/main.go
│       ├── internal/
│       │   ├── handlers/           # contact.go, health.go
│       │   ├── models/
│       │   ├── services/
│       │   └── middleware/
│       ├── migrations/
│       ├── go.mod
│       └── Makefile
│
├── infra/
│   ├── nginx/
│   │   └── webifylab.conf          # Nginx reverse proxy config
│   ├── systemd/
│   │   └── webifylab-api.service   # Systemd unit (V1.5)
│   └── scripts/
│       ├── setup-vps.sh            # VPS initial setup script
│       ├── deploy.sh               # Deployment script
│       └── backup.sh               # Backup script
│
├── docs/                           # 📚 Project Documentation
│   ├── prd.md                      # Product Requirements Document
│   ├── tdd.md                      # Technical Design Document
│   ├── design.md                   # Design System
│   ├── wireframe-ui.md             # Wireframe & UI Spec
│   ├── content-copywriting.md      # Content & Copywriting
│   ├── database-schema.md          # DB Schema (V1.5 & V2 ready)
│   ├── seo-strategy.md             # SEO Strategy
│   ├── deployment-runbook.md       # Deployment & Operations
│   └── progress-tracker.md         # Project Progress
│
├── .env.example
├── .gitignore
├── Makefile                        # Root Makefile (shortcut commands)
└── README.md
```

---

## ⚡ Quick Start

### Prerequisites

- **Node.js** >= 22.12.0
- **npm** atau **pnpm**
- **Git**

### Instalasi & Menjalankan

```bash
# 1. Clone repository
git clone https://github.com/RizalRio/mq-webifylab-portal-v2.git
cd mq-webifylab-portal-v2

# 2. Masuk ke frontend app
cd apps/web

# 3. Install dependencies
npm install

# 4. Jalankan dev server
npm run dev
# → http://localhost:4321
```

### Root Makefile Shortcuts

```bash
# Dari root directory
make dev-web       # Jalankan Astro dev server
make build-web     # Build production frontend
make dev-api       # Jalankan Golang API (V1.5)
```

### Scripts Frontend

```bash
cd apps/web

npm run dev        # Dev server (http://localhost:4321)
npm run build      # Build untuk production (output: dist/)
npm run preview    # Preview production build secara lokal
```

---

## 🗂️ Sections Landing Page

Landing page terdiri dari **9 section aktif** (S8 Testimonial disembunyikan di V1):

| ID | Section | Status | Keterangan |
|----|---------|--------|------------|
| S1 | **Navbar** | ✅ Done | Sticky, hamburger mobile, transparent→solid scroll |
| S2 | **Hero** | ✅ Done | Headline gradient, SVG illustration, dual CTA |
| S3 | **Problem** | ✅ Done | 3 pain point cards |
| S4 | **Services** | ✅ Done | 3 layanan + feature checklist |
| S5 | **Why Webifylab** | ✅ Done | 2×2 grid keunggulan |
| S6 | **Approach** | ✅ Done | 4 steps metodologi + tech stack chips |
| S7 | **Portfolio** | ✅ Done | 3 "Coming Soon" placeholder cards |
| S8 | **Testimonial** | ⏸️ Hidden | Ditunda sampai ada testimonial nyata |
| S9 | **Contact / CTA** | ✅ Done | Form Formspree + validasi + WA fallback |
| S10 | **Footer** | ✅ Done | 4 kolom + social links |

---

## ⚙️ Konfigurasi

### 1. Update Informasi Kontak

Edit [`apps/web/src/lib/constants.ts`](./apps/web/src/lib/constants.ts):

```typescript
export const SITE = {
  name: "Webifylab",
  email: "hello@webifylab.com",        // ← Update email
  whatsapp: "6281234567890",           // ← Update nomor WA (format: 628xxx)
  url: "https://webifylab.com",        // ← Update domain
  social: {
    linkedin: "https://linkedin.com/company/webifylab",
    github:   "https://github.com/webifylab",
    instagram: "https://instagram.com/webifylab",
  },
};
```

### 2. Setup Formspree (Form Kontak)

1. Daftar di [formspree.io](https://formspree.io) (gratis, 50 submissions/bulan)
2. Buat form baru, salin Form ID
3. Edit [`apps/web/src/components/sections/Contact.astro`](./apps/web/src/components/sections/Contact.astro) baris 7:

```typescript
const FORMSPREE_ENDPOINT = "https://formspree.io/f/YOUR_FORM_ID"; // ← Ganti ini
```

### 3. Environment Variables

Salin `.env.example` menjadi `.env`:

```bash
cp .env.example .env
```

Isi variabel yang diperlukan sesuai komentar di dalam file.

### 4. Design Tokens

Seluruh design token (warna, font, spacing) dikonfigurasi via CSS `@theme` di [`apps/web/src/styles/global.css`](./apps/web/src/styles/global.css) mengikuti Tailwind CSS v4 CSS-first config:

```css
@theme {
  --color-primary:       #1E3A5F;   /* Deep Blue */
  --color-accent-indigo: #4F46E5;   /* Indigo */
  --color-accent-cyan:   #06B6D4;   /* Cyan */
  --font-sans:           "Inter", system-ui, sans-serif;
}
```

---

## 🚀 Deployment

### Arsitektur Hosting (VPS)

```
Internet (HTTPS)
      │
      ▼
┌─────────────────────────────────┐
│   VPS — 1 vCPU, 1GB RAM        │
│   ┌───────────────────────────┐ │
│   │   Nginx (Reverse Proxy)   │ │
│   │   • SSL Termination       │ │
│   │   • Serve static files    │ │
│   │   • Gzip/Brotli           │ │
│   │   • Security Headers      │ │
│   └─────────┬─────────────────┘ │
│             │                   │
│     ┌───────┴──────────┐       │
│     ▼                  ▼       │
│  Static Files      Go API      │
│  /var/www/         :8080       │
│  webifylab/        (V1.5)      │
└─────────────────────────────────┘
```

### Deploy Manual

```bash
# 1. Build frontend
cd apps/web && npm run build

# 2. Upload ke VPS via rsync
rsync -avz --delete dist/ deploy@your-vps-ip:/var/www/webifylab/dist/

# 3. Atau gunakan script
./infra/scripts/deploy.sh
```

### Deploy via GitHub Actions

Lihat [`.github/workflows/deploy.yml`](./.github/workflows/deploy.yml) untuk konfigurasi CI/CD otomatis pada push ke branch `main`.

> 📖 Lihat [`docs/deployment-runbook.md`](./docs/deployment-runbook.md) untuk panduan lengkap setup VPS, Nginx, SSL, dan deployment step-by-step.

---

## 🗺️ Roadmap

### ✅ V1 — Landing Page (Selesai)

- [x] Dokumentasi lengkap (PRD, TDD, Design, Wireframe, Content, SEO, Deploy)
- [x] Landing page statis Astro + Tailwind CSS v4
- [x] 9 section: Navbar, Hero, Problem, Services, Why, Approach, Portfolio, Contact, Footer
- [x] Form kontak via Formspree
- [x] WhatsApp CTA integration
- [x] Full SEO meta tags (OG, Twitter Card)
- [x] Custom 404 page
- [ ] Go-live ke VPS production
- [ ] Google Search Console setup
- [ ] Analytics (Plausible / GA4)

### 🔜 V1.5 — Backend API (Golang)

- [ ] Golang API (`apps/api`) dengan Chi router
- [ ] Endpoint `/api/contact` — terima form, validasi, notifikasi Telegram/Email
- [ ] Endpoint `/api/health` — health check
- [ ] Migrasi form dari Formspree ke Golang API
- [ ] PostgreSQL connection + migrations
- [ ] Systemd service + auto-restart

### 🔮 V2 — Ekosistem Digital

- [ ] Blog section (Astro Content Collections)
- [ ] Portfolio dengan case studies nyata
- [ ] Testimonial section aktif
- [ ] Admin panel (dashboard leads)
- [ ] Integrasi AI & Data features
- [ ] Privacy Policy & Terms of Service pages

---

## 📚 Dokumentasi

Semua dokumentasi proyek tersedia di folder [`docs/`](./docs/):

| File | Deskripsi |
|------|-----------|
| [`prd.md`](./docs/prd.md) | Product Requirements — fitur, persona, KPIs |
| [`tdd.md`](./docs/tdd.md) | Technical Design — arsitektur, stack, infra |
| [`design.md`](./docs/design.md) | Design System — colors, typography, components |
| [`wireframe-ui.md`](./docs/wireframe-ui.md) | Wireframe & UI spec per section |
| [`content-copywriting.md`](./docs/content-copywriting.md) | Konten teks, headline, microcopy |
| [`database-schema.md`](./docs/database-schema.md) | PostgreSQL schema (V1.5 & V2 ready) |
| [`seo-strategy.md`](./docs/seo-strategy.md) | SEO keyword, meta tags, structured data |
| [`deployment-runbook.md`](./docs/deployment-runbook.md) | Panduan setup VPS, Nginx, CI/CD |
| [`progress-tracker.md`](./docs/progress-tracker.md) | Status kemajuan proyek per phase |

---

## 🚧 Blockers & Risiko

| ID | Deskripsi | Dampak | Mitigasi |
|----|-----------|--------|----------|
| B1 | VPS 1GB RAM — risiko OOM | High | Swap 1GB aktif, Astro static, matikan service tidak perlu |
| B2 | Portfolio belum ada proyek nyata | Medium | Placeholder "Coming Soon" profesional di S7 |
| B3 | Formspree free tier 50 sub/bulan | Low | Upgrade atau percepat Golang API (V1.5) |

---

## 👤 Author

**Rizal** — System Analyst / Software Engineer / AI & Data Enthusiast

- 🌐 Website: [webifylab.com](https://webifylab.com)
- 📧 Email: hello@webifylab.com
- 💼 LinkedIn: [linkedin.com/company/webifylab](https://linkedin.com/company/webifylab)

---

<div align="center">

Made with ❤️ using **Astro** & **Tailwind CSS**

*© 2026 Webifylab. All rights reserved.*

</div>
