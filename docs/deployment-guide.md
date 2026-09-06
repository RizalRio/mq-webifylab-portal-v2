# 📄 DEPLOYMENT GUIDE
## WebifyLab Portal — Phase 1 (MVP)

---

### 1. INFORMASI DOKUMEN

| Field | Detail |
|---|---|
| **Produk** | WebifyLab Portal |
| **Tipe Dokumen** | Deployment Guide |
| **Versi** | 1.0.0 |
| **Tanggal** | 05 September 2026 |
| **Target Server** | VPS 1 vCPU, 1GB RAM, 20GB SSD, Bogor T3 |
| **Domain Frontend** | webifylab.my.id |
| **Domain API** | api.webifylab.my.id |
| **Penulis** | Rizal (System Analyst & Fullstack Developer) |
| **Status** | Approved — Ready for Execution |

---

### 2. OVERVIEW DEPLOYMENT

#### 2.1 Arsitektur Deployment

```
┌──────────────────────────────────────────────────────────────────┐
│                     LAPTOP (16GB RAM)                            │
│                                                                  │
│  1. Development (Next.js + Golang)                              │
│  2. Build Docker images                                         │
│  3. Save images ke .tar.gz                                      │
│  4. Upload ke VPS via SCP                                       │
│                                                                  │
└────────────────────────┬─────────────────────────────────────────┘
                         │
                         │ SCP / rsync
                         │
                         ▼
┌──────────────────────────────────────────────────────────────────┐
│                  VPS (1GB RAM, Bogor T3)                         │
│                                                                  │
│  ┌────────────────────────────────────────────────────────────┐ │
│  │                   Nginx (Host)                             │ │
│  │  • SSL termination (Let's Encrypt)                        │ │
│  │  • Reverse proxy                                           │ │
│  │  • webifylab.my.id → localhost:3000                       │ │
│  │  • api.webifylab.my.id → localhost:8080                   │ │
│  │  • Serve static files (uploads)                            │ │
│  └──────────────┬───────────────────────┬────────────────────┘ │
│                 │                       │                       │
│                 ▼                       ▼                       │
│  ┌────────────────────────┐  ┌────────────────────────┐        │
│  │  Docker: Next.js       │  │  Docker: Golang (Gin)  │        │
│  │  • Port 3000           │  │  • Port 8080           │        │
│  │  • SSR/SSG             │  │  • REST API            │        │
│  │  • Admin CMS           │  │  • GORM + PostgreSQL   │        │
│  │  • Limit: 350MB RAM    │  │  • Resend email        │        │
│  │                        │  │  • Limit: 150MB RAM    │        │
│  └────────────────────────┘  └────────────────────────┘        │
│                                                                  │
│  ┌────────────────────────────────────────────────────────────┐ │
│  │              Docker: PostgreSQL 15                          │ │
│  │  • Port 5432 (internal only)                               │ │
│  │  • Tuned untuk low memory                                  │ │
│  │  • Limit: 200MB RAM                                        │ │
│  └────────────────────────────────────────────────────────────┘ │
│                                                                  │
│  Total RAM containers: ~530MB / 1GB                             │
│  Buffer OS + spike: ~470MB                                      │
│  Swap: 2GB (fallback)                                           │
└──────────────────────────────────────────────────────────────────┘
```

#### 2.2 Alur Deployment

```
┌─────────┐     ┌─────────┐     ┌─────────┐     ┌─────────┐
│  Build  │ ──► │  Save   │ ──► │ Upload  │ ──► │  Load   │
│  Docker │     │  Images │     │  ke VPS │     │ & Start │
│  Images │     │  .tar.gz│     │  via SCP│     │Containers│
└─────────┘     └─────────┘     └─────────┘     └─────────┘
  (Laptop)        (Laptop)        (Laptop→VPS)      (VPS)
```

#### 2.3 Direktori di VPS

```
/var/www/webifylab/
├── docker-compose.yml          # Orchestration
├── .env                        # Environment variables (RAHASIA)
├── frontend.tar.gz             # Docker image frontend (sementara)
├── backend.tar.gz              # Docker image backend (sementara)
├── infra/
│   ├── nginx/
│   │   ├── webifylab.conf      # Nginx config frontend
│   │   └── api.conf            # Nginx config API
│   ├── postgres/
│   │   └── postgresql.conf     # PostgreSQL tuning
│   └── scripts/
│       ├── deploy.sh           # Deployment script
│       ├── backup.sh           # Backup script
│       └── restore.sh          # Restore script
└── uploads/                    # Uploaded files (volume mount)
    └── media/
```

---

### 3. PRASYARAT (Sebelum Deployment)

#### 3.1 Checklist Prasyarat

| No | Item | Status | Cara Verifikasi |
|---|---|---|---|
| P1 | Domain webifylab.my.id sudah aktif | ⬜ | `nslookup webifylab.my.id` |
| P2 | DNS A record api.webifylab.my.id → IP VPS | ⬜ | `nslookup api.webifylab.my.id` |
| P3 | Akses SSH ke VPS | ⬜ | `ssh user@IP_VPS` |
| P4 | Docker terinstall di VPS | ⬜ | `docker --version` |
| P5 | Docker Compose terinstall di VPS | ⬜ | `docker compose version` |
| P6 | Nginx terinstall di VPS | ⬜ | `nginx -v` |
| P7 | Previous deployment sudah dihapus | ⬜ | `docker ps` (kosong) |
| P8 | Swap space 2GB sudah aktif | ⬜ | `free -h` |
| P9 | Node.js terinstall di laptop | ⬜ | `node --version` (≥ v20) |
| P10 | Go terinstall di laptop | ⬜ | `go version` (≥ 1.22) |
| P11 | Docker terinstall di laptop | ⬜ | `docker --version` |
| P12 | Akun Resend sudah verifikasi domain | ⬜ | Cek dashboard Resend |
| P13 | API key Resend sudah di-generate | ⬜ | Cek dashboard Resend |
| P14 | Kode sudah di-build dan di-test lokal | ⬜ | `make dev` berjalan |

#### 3.2 Setup DNS (Jika Belum)

Login ke DNS provider (registrar domain), tambahkan:

| Type | Name | Value | TTL |
|---|---|---|---|
| A | `@` | `IP_VPS_KAMU` | 3600 |
| A | `api` | `IP_VPS_KAMU` | 3600 |

Verifikasi propagation:
```bash
# Tunggu 5-30 menit setelah setup DNS, lalu cek:
nslookup webifylab.my.id
nslookup api.webifylab.my.id

# Atau pakai dig:
dig webifylab.my.id +short
dig api.webifylab.my.id +short
```

**Catatan:** DNS propagation bisa memakan waktu 1-48 jam. Biasanya 5-30 menit untuk provider umum.

#### 3.3 Setup Swap Space di VPS (Jika Belum)

```bash
# SSH ke VPS
ssh user@IP_VPS

# Cek apakah swap sudah ada
free -h

# Jika belum ada swap, buat 2GB:
sudo fallocate -l 2G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile

# Persist setelah reboot
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab

# Set swappiness rendah (gunakan swap hanya saat perlu)
sudo sysctl vm.swappiness=10
echo 'vm.swappiness=10' | sudo tee -a /etc/sysctl.conf

# Verifikasi
free -h
# Output harus menunjukkan Swap: 2.0G
```

#### 3.4 Hapus Previous Deployment

```bash
# SSH ke VPS
ssh user@IP_VPS

# Stop semua container yang berjalan
docker stop $(docker ps -q) 2>/dev/null

# Hapus semua container
docker rm $(docker ps -aq) 2>/dev/null

# Hapus semua image (HATI-HATI: ini hapus SEMUA image)
docker rmi $(docker images -q) 2>/dev/null

# Hapus volume yang tidak dipakai
docker volume prune -f

# Hapus network yang tidak dipakai
docker network prune -f

# Hapus file deployment lama
sudo rm -rf /var/www/old-webifylab/*

# Verifikasi bersih
docker ps          # Harus kosong
docker images      # Harus kosong
docker volume ls   # Harus kosong (atau hanya volume sistem)
df -h              # Cek storage tersedia
free -h            # Cek RAM tersedia
```

---

### 4. PERSIAPAN FILE DI LAPTOP

#### 4.1 Struktur File yang Akan Di-upload

Sebelum deployment, pastikan kamu punya file-file berikut di laptop:

```
webifylab-portal/                    # Root monorepo di laptop
├── docker-compose.yml               # ✅ Akan di-upload
├── .env.production                  # ✅ Akan di-upload sebagai .env
├── frontend/
│   └── Dockerfile                   # Dipakai untuk build di laptop
├── backend/
│   └── Dockerfile                   # Dipakai untuk build di laptop
└── infra/
    ├── nginx/
    │   ├── webifylab.conf           # ✅ Akan di-upload
    │   └── api.conf                 # ✅ Akan di-upload
    ├── postgres/
    │   └── postgresql.conf          # ✅ Akan di-upload
    └── scripts/
        ├── deploy.sh                # ✅ Akan di-upload
        ├── backup.sh                # ✅ Akan di-upload
        └── restore.sh               # ✅ Akan di-upload
```

#### 4.2 Buat File .env.production

Buat file `.env.production` di root monorepo (jangan commit ke git!):

```bash
# ===================================
# WebifyLab Portal - Production Environment
# ===================================

# Application
APP_ENV=production
SITE_URL=https://webifylab.my.id
API_URL=https://api.webifylab.my.id

# Database
DB_USER=webifylab
DB_PASSWORD=GantiDenganPasswordKuatMinimal16Karakter!
DB_NAME=webifylab

# JWT (generate dengan: openssl rand -base64 32)
JWT_SECRET=HasilGenerateOpenSSLDisiniMinimal32Karakter!

# Resend Email
RESEND_API_KEY=re_GantiDenganAPIKeyResendKamu
ADMIN_EMAIL=admin@webifylab.my.id
```

**Cara generate JWT secret:**
```bash
openssl rand -base64 32
# Output contoh: K7gNU3sdo+OL0wNhqoVWhr3g6s1xYv72ol/pe/Unols=
```

#### 4.3 Buat File docker-compose.yml

File ini sudah ada di monorepo. Pastikan isinya sesuai:

```yaml
# docker-compose.yml
version: '3.8'

services:
  frontend:
    image: webifylab-frontend:latest
    container_name: webifylab-frontend
    restart: unless-stopped
    environment:
      - NODE_ENV=production
      - NEXT_PUBLIC_API_URL=${API_URL}
      - NEXT_PUBLIC_SITE_URL=${SITE_URL}
    ports:
      - "3000:3000"
    networks:
      - webifylab-network
    deploy:
      resources:
        limits:
          memory: 350M

  backend:
    image: webifylab-backend:latest
    container_name: webifylab-backend
    restart: unless-stopped
    env_file:
      - .env
    environment:
      - APP_ENV=${APP_ENV}
      - APP_PORT=8080
      - APP_HOST=0.0.0.0
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=${DB_USER}
      - DB_PASSWORD=${DB_PASSWORD}
      - DB_NAME=${DB_NAME}
      - DB_SSLMODE=disable
      - JWT_SECRET=${JWT_SECRET}
      - JWT_EXPIRY=24h
      - REFRESH_TOKEN_EXPIRY=168h
      - RESEND_API_KEY=${RESEND_API_KEY}
      - ADMIN_EMAIL=${ADMIN_EMAIL}
      - FRONTEND_URL=${SITE_URL}
      - CORS_ORIGIN=${SITE_URL}
      - UPLOAD_DIR=/app/uploads
      - MAX_UPLOAD_SIZE=5242880
    depends_on:
      postgres:
        condition: service_healthy
    ports:
      - "8080:8080"
    volumes:
      - uploads:/app/uploads
    networks:
      - webifylab-network
    deploy:
      resources:
        limits:
          memory: 150M

  postgres:
    image: postgres:15-alpine
    container_name: webifylab-postgres
    restart: unless-stopped
    environment:
      - POSTGRES_DB=${DB_NAME}
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
    volumes:
      - pgdata:/var/lib/postgresql/data
      - ./infra/postgres/postgresql.conf:/etc/postgresql/postgresql.conf
    command: postgres -c config_file=/etc/postgresql/postgresql.conf
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U ${DB_USER}"]
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

**PENTING:** Perhatikan bahwa di production, ports di-expose ke host (`3000:3000` dan `8080:8080`) karena Nginx di host akan proxy ke port ini. PostgreSQL TIDAK di-expose ke host (aman).

#### 4.4 Buat File Nginx Configurations

**File: `infra/nginx/webifylab.conf`**

```nginx
# /etc/nginx/sites-available/webifylab

# ============================================
# Redirect HTTP ke HTTPS
# ============================================
server {
    listen 80;
    server_name webifylab.my.id api.webifylab.my.id;
    return 301 https://$server_name$request_uri;
}

# ============================================
# Frontend (Next.js) - webifylab.my.id
# ============================================
server {
    listen 443 ssl http2;
    server_name webifylab.my.id;

    # SSL Certificate (akan di-generate oleh Certbot)
    ssl_certificate /etc/letsencrypt/live/webifylab.my.id/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/webifylab.my.id/privkey.pem;

    # SSL Settings
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;

    # Security Headers
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;

    # Gzip Compression
    gzip on;
    gzip_vary on;
    gzip_proxied any;
    gzip_comp_level 6;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript image/svg+xml;

    # Client body size (untuk upload)
    client_max_body_size 10M;

    # Uploaded media files (serve langsung dari Nginx, bukan lewat container)
    location /uploads/ {
        alias /var/www/webifylab/uploads/;
        expires 30d;
        add_header Cache-Control "public, immutable";
        access_log off;
    }

    # Next.js static files (cache agresif)
    location /_next/static/ {
        proxy_pass http://127.0.0.1:3000;
        proxy_cache_valid 200 365d;
        add_header Cache-Control "public, max-age=31536000, immutable";
        access_log off;
    }

    # Semua request lainnya ke Next.js
    location / {
        proxy_pass http://127.0.0.1:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_cache_bypass $http_upgrade;

        # Timeouts
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }

    # Logging
    access_log /var/log/nginx/webifylab_access.log;
    error_log /var/log/nginx/webifylab_error.log;
}
```

**File: `infra/nginx/api.conf`**

```nginx
# /etc/nginx/sites-available/api

# ============================================
# Backend API (Golang + Gin) - api.webifylab.my.id
# ============================================
server {
    listen 443 ssl http2;
    server_name api.webifylab.my.id;

    # SSL Certificate
    ssl_certificate /etc/letsencrypt/live/api.webifylab.my.id/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.webifylab.my.id/privkey.pem;

    # SSL Settings
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;

    # Security Headers
    add_header X-Frame-Options "DENY" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;

    # Client body size (untuk upload media)
    client_max_body_size 10M;

    # Rate Limiting
    limit_req_zone $binary_remote_addr zone=api_general:10m rate=30r/s;
    limit_req_zone $binary_remote_addr zone=api_login:10m rate=5r/m;
    limit_req_zone $binary_remote_addr zone=api_leads:10m rate=5r/h;

    # Health check (no rate limit)
    location = /api/v1/health {
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        access_log off;
    }

    # Login endpoint (strict rate limit)
    location = /api/v1/auth/login {
        limit_req zone=api_login burst=3 nodelay;
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Leads submit endpoint (strict rate limit)
    location = /api/v1/leads {
        limit_req zone=api_leads burst=3 nodelay;
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # Uploaded files (serve langsung dari Nginx)
    location /uploads/ {
        alias /var/www/webifylab/uploads/;
        expires 30d;
        add_header Cache-Control "public, immutable";
        access_log off;
    }

    # Semua API request lainnya
    location / {
        limit_req zone=api_general burst=20 nodelay;
        proxy_pass http://127.0.0.1:8080;
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

    # Logging
    access_log /var/log/nginx/api_access.log;
    error_log /var/log/nginx/api_error.log;
}
```

#### 4.5 Buat File PostgreSQL Tuning

**File: `infra/postgres/postgresql.conf`**

```ini
# PostgreSQL Tuning untuk VPS 1GB RAM
# ============================================

# Memory Settings
shared_buffers = 128MB
effective_cache_size = 256MB
work_mem = 4MB
maintenance_work_mem = 64MB

# WAL Settings
wal_buffers = 4MB
checkpoint_completion_target = 0.9

# Connection Settings
max_connections = 20

# Logging
log_min_duration_statement = 1000
log_line_prefix = '%t [%p]: [%l-1] user=%u,db=%d '

# Locale
lc_messages = 'en_US.UTF-8'
```

#### 4.6 Buat Deployment Scripts

**File: `infra/scripts/deploy.sh`**

```bash
#!/bin/bash
# ============================================
# WebifyLab Portal - Deployment Script
# ============================================
# Jalankan dari laptop: ./infra/scripts/deploy.sh
# ============================================

set -e

# Konfigurasi
VPS_USER="user"                    # Ganti dengan username VPS kamu
VPS_IP="IP_VPS_KAMU"              # Ganti dengan IP VPS kamu
VPS_DIR="/var/www/webifylab"
PROJECT_DIR="$(cd "$(dirname "$0")/../.." && pwd)"

echo "🚀 ============================================"
echo "🚀 WebifyLab Portal - Deployment"
echo "🚀 ============================================"
echo ""

# Step 1: Build Docker images
echo "📦 Step 1/6: Building Docker images..."
cd "$PROJECT_DIR"
docker build -t webifylab-frontend:latest -f frontend/Dockerfile frontend/
docker build -t webifylab-backend:latest -f backend/Dockerfile backend/
echo "✅ Docker images berhasil di-build"
echo ""

# Step 2: Save images ke tar.gz
echo "💾 Step 2/6: Saving Docker images..."
docker save webifylab-frontend:latest | gzip > frontend.tar.gz
docker save webifylab-backend:latest | gzip > backend.tar.gz
echo "✅ Images berhasil di-save"
ls -lh frontend.tar.gz backend.tar.gz
echo ""

# Step 3: Upload ke VPS
echo "📤 Step 3/6: Uploading ke VPS..."
ssh $VPS_USER@$VPS_IP "mkdir -p $VPS_DIR/infra/nginx $VPS_DIR/infra/postgres $VPS_DIR/infra/scripts $VPS_DIR/uploads"

scp frontend.tar.gz $VPS_USER@$VPS_IP:$VPS_DIR/
scp backend.tar.gz $VPS_USER@$VPS_IP:$VPS_DIR/
scp docker-compose.yml $VPS_USER@$VPS_IP:$VPS_DIR/
scp .env.production $VPS_USER@$VPS_IP:$VPS_DIR/.env
scp infra/nginx/webifylab.conf $VPS_USER@$VPS_IP:$VPS_DIR/infra/nginx/
scp infra/nginx/api.conf $VPS_USER@$VPS_IP:$VPS_DIR/infra/nginx/
scp infra/postgres/postgresql.conf $VPS_USER@$VPS_IP:$VPS_DIR/infra/postgres/
scp infra/scripts/backup.sh $VPS_USER@$VPS_IP:$VPS_DIR/infra/scripts/
scp infra/scripts/restore.sh $VPS_USER@$VPS_IP:$VPS_DIR/infra/scripts/
echo "✅ Upload berhasil"
echo ""

# Step 4: Load images di VPS
echo "📥 Step 4/6: Loading Docker images di VPS..."
ssh $VPS_USER@$VPS_IP "cd $VPS_DIR && docker load -i frontend.tar.gz && docker load -i backend.tar.gz"
echo "✅ Images berhasil di-load"
echo ""

# Step 5: Start containers
echo "▶️  Step 5/6: Starting containers..."
ssh $VPS_USER@$VPS_IP "cd $VPS_DIR && docker compose down 2>/dev/null; docker compose up -d"
echo "✅ Containers berhasil di-start"
echo ""

# Step 6: Run migrations & seed
echo "🗄️  Step 6/6: Running migrations..."
sleep 10  # Tunggu postgres siap
ssh $VPS_USER@$VPS_IP "cd $VPS_DIR && docker compose exec backend ./server migrate" || echo "⚠️  Migration skipped (mungkin sudah up-to-date)"
ssh $VPS_USER@$VPS_IP "cd $VPS_DIR && docker compose exec backend ./server seed" || echo "⚠️  Seed skipped (mungkin sudah di-seed)"
echo ""

# Cleanup
echo "🧹 Cleaning up..."
rm -f frontend.tar.gz backend.tar.gz
ssh $VPS_USER@$VPS_IP "rm -f $VPS_DIR/frontend.tar.gz $VPS_DIR/backend.tar.gz"
echo ""

# Verify
echo "🔍 Verifying deployment..."
echo "   Health check:"
curl -s https://api.webifylab.my.id/api/v1/health | head -c 200
echo ""
echo "   Frontend:"
curl -s -o /dev/null -w "   HTTP Status: %{http_code}\n" https://webifylab.my.id
echo ""

echo "🎉 ============================================"
echo "🎉 Deployment selesai!"
echo "🎉 Frontend: https://webifylab.my.id"
echo "🎉 API:      https://api.webifylab.my.id"
echo "🎉 Admin:    https://webifylab.my.id/admin"
echo "🎉 ============================================"
```

**File: `infra/scripts/backup.sh`**

```bash
#!/bin/bash
# ============================================
# WebifyLab Portal - Backup Script
# ============================================
# Setup di cron: 0 2 * * * /var/www/webifylab/infra/scripts/backup.sh
# ============================================

BACKUP_DIR="/var/backups/webifylab"
DATE=$(date +%Y%m%d_%H%M%S)
COMPOSE_DIR="/var/www/webifylab"

mkdir -p $BACKUP_DIR

echo "[$(date)] Memulai backup..."

# Backup database
docker compose -f $COMPOSE_DIR/docker-compose.yml exec -T postgres \
  pg_dump -U webifylab webifylab | gzip > $BACKUP_DIR/db_$DATE.sql.gz

if [ $? -eq 0 ]; then
    echo "[$(date)] ✅ Database backup berhasil: db_$DATE.sql.gz"
else
    echo "[$(date)] ❌ Database backup GAGAL!"
fi

# Backup uploads
if [ -d "$COMPOSE_DIR/uploads" ]; then
    tar -czf $BACKUP_DIR/uploads_$DATE.tar.gz -C $COMPOSE_DIR uploads/
    echo "[$(date)] ✅ Uploads backup berhasil: uploads_$DATE.tar.gz"
fi

# Hapus backup lama (simpan 7 hari terakhir)
find $BACKUP_DIR -name "db_*.sql.gz" -type f -mtime +7 -delete
find $BACKUP_DIR -name "uploads_*.tar.gz" -type f -mtime +7 -delete

# Hitung total size backup
TOTAL_SIZE=$(du -sh $BACKUP_DIR | cut -f1)
echo "[$(date)] Backup selesai. Total size: $TOTAL_SIZE"
```

**File: `infra/scripts/restore.sh`**

```bash
#!/bin/bash
# ============================================
# WebifyLab Portal - Restore Script
# ============================================
# Cara pakai: ./restore.sh /path/to/db_YYYYMMDD_HHMMSS.sql.gz
# ============================================

if [ -z "$1" ]; then
    echo "Penggunaan: $0 <file_backup.sql.gz>"
    echo "Contoh: $0 /var/backups/webifylab/db_20260905_020000.sql.gz"
    echo ""
    echo "Backup yang tersedia:"
    ls -lh /var/backups/webifylab/db_*.sql.gz 2>/dev/null
    exit 1
fi

BACKUP_FILE=$1
COMPOSE_DIR="/var/www/webifylab"

if [ ! -f "$BACKUP_FILE" ]; then
    echo "❌ File tidak ditemukan: $BACKUP_FILE"
    exit 1
fi

echo "⚠️  PERINGATAN: Ini akan menimpa database saat ini!"
read -p "Apakah Anda yakin? (ketik 'YA' untuk konfirmasi): " CONFIRM

if [ "$CONFIRM" != "YA" ]; then
    echo "Dibatalkan."
    exit 0
fi

echo "🔄 Menghentikan backend..."
docker compose -f $COMPOSE_DIR/docker-compose.yml stop backend

echo "🔄 Merestore database dari $BACKUP_FILE..."
gunzip -c "$BACKUP_FILE" | docker compose -f $COMPOSE_DIR/docker-compose.yml exec -T postgres \
  psql -U webifylab webifylab

if [ $? -eq 0 ]; then
    echo "✅ Database berhasil di-restore!"
else
    echo "❌ Restore GAGAL!"
fi

echo "🔄 Menyalakan backend..."
docker compose -f $COMPOSE_DIR/docker-compose.yml start backend

echo "✅ Selesai!"
```

Buat semua script executable:
```bash
chmod +x infra/scripts/deploy.sh
chmod +x infra/scripts/backup.sh
chmod +x infra/scripts/restore.sh
```

---

### 5. SETUP SSL CERTIFICATE DI VPS

**Langkah ini dilakukan SEKALI di awal, sebelum deployment pertama.**

#### 5.1 Install Certbot

```bash
# SSH ke VPS
ssh user@IP_VPS

# Update sistem
sudo apt update && sudo apt upgrade -y

# Install Certbot
sudo apt install certbot -y
```

#### 5.2 Generate SSL Certificate

**PENTING:** Stop Nginx dulu karena kita pakai mode standalone.

```bash
# Stop Nginx sementara
sudo systemctl stop nginx

# Generate SSL untuk webifylab.my.id
sudo certbot certonly --standalone -d webifylab.my.id

# Generate SSL untuk api.webifylab.my.id
sudo certbot certonly --standalone -d api.webifylab.my.id

# Start Nginx kembali
sudo systemctl start nginx
```

**Jika berhasil, output akan menunjukkan:**
```
Successfully received certificate.
Certificate is saved at: /etc/letsencrypt/live/webifylab.my.id/fullchain.pem
Key is saved at:         /etc/letsencrypt/live/webifylab.my.id/privkey.pem
```

#### 5.3 Verifikasi Certificate

```bash
ls -la /etc/letsencrypt/live/webifylab.my.id/
ls -la /etc/letsencrypt/live/api.webifylab.my.id/
```

#### 5.4 Setup Auto-Renewal

```bash
# Test renewal (dry run)
sudo certbot renew --dry-run

# Setup cron untuk auto-renewal
sudo crontab -e

# Tambahkan baris ini:
0 3 * * * certbot renew --quiet --pre-hook "systemctl stop nginx" --post-hook "systemctl start nginx"
```

---

### 6. SETUP NGINX DI VPS

#### 6.1 Copy Config Files

Setelah deployment script dijalankan, config files sudah ada di `/var/www/webifylab/infra/nginx/`. Sekarang symlink ke Nginx:

```bash
# SSH ke VPS
ssh user@IP_VPS

# Symlink config files
sudo ln -sf /var/www/webifylab/infra/nginx/webifylab.conf /etc/nginx/sites-available/webifylab
sudo ln -sf /var/www/webifylab/infra/nginx/api.conf /etc/nginx/sites-available/api

# Enable sites
sudo ln -sf /etc/nginx/sites-available/webifylab /etc/nginx/sites-enabled/
sudo ln -sf /etc/nginx/sites-available/api /etc/nginx/sites-enabled/

# Hapus default site (opsional)
sudo rm -f /etc/nginx/sites-enabled/default
```

#### 6.2 Setup Rate Limiting Zone

Tambahkan rate limiting zone di config utama Nginx:

```bash
sudo nano /etc/nginx/nginx.conf
```

Tambahkan di dalam block `http { }`:

```nginx
http {
    # ... existing config ...

    # Rate Limiting Zones
    limit_req_zone $binary_remote_addr zone=api_general:10m rate=30r/s;
    limit_req_zone $binary_remote_addr zone=api_login:10m rate=5r/m;
    limit_req_zone $binary_remote_addr zone=api_leads:10m rate=5r/h;

    # ... existing config ...
}
```

#### 6.3 Buat Direktori Uploads

```bash
sudo mkdir -p /var/www/webifylab/uploads/media
sudo chmod -R 755 /var/www/webifylab/uploads
```

#### 6.4 Setup Log Rotation

```bash
sudo nano /etc/logrotate.d/webifylab
```

Isi dengan:
```
/var/log/nginx/webifylab_*.log
/var/log/nginx/api_*.log {
    daily
    missingok
    rotate 14
    compress
    delaycompress
    notifempty
    create 0640 www-data adm
    sharedscripts
    postrotate
        [ -f /var/run/nginx.pid ] && kill -USR1 $(cat /var/run/nginx.pid)
    endscript
}
```

#### 6.5 Test & Reload Nginx

```bash
# Test konfigurasi
sudo nginx -t

# Output yang diharapkan:
# nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
# nginx: configuration file /etc/nginx/nginx.conf test is successful

# Reload Nginx
sudo systemctl reload nginx

# Verifikasi Nginx berjalan
sudo systemctl status nginx
```

---

### 7. DEPLOYMENT STEP-BY-STEP

#### 7.1 Deployment Pertama (Fresh Install)

**Dari Laptop:**

```bash
# 1. Pastikan semua kode sudah di-commit dan di-test
cd webifylab-portal
git status
make dev  # Test lokal dulu

# 2. Jalankan deployment script
./infra/scripts/deploy.sh
```

**Atau manual step-by-step:**

```bash
# Step 1: Build Docker images di laptop
cd webifylab-portal
docker build -t webifylab-frontend:latest -f frontend/Dockerfile frontend/
docker build -t webifylab-backend:latest -f backend/Dockerfile backend/

# Step 2: Save images
docker save webifylab-frontend:latest | gzip > frontend.tar.gz
docker save webifylab-backend:latest | gzip > backend.tar.gz

# Step 3: Buat direktori di VPS
ssh user@IP_VPS "mkdir -p /var/www/webifylab/infra/nginx /var/www/webifylab/infra/postgres /var/www/webifylab/infra/scripts /var/www/webifylab/uploads"

# Step 4: Upload semua file
scp frontend.tar.gz user@IP_VPS:/var/www/webifylab/
scp backend.tar.gz user@IP_VPS:/var/www/webifylab/
scp docker-compose.yml user@IP_VPS:/var/www/webifylab/
scp .env.production user@IP_VPS:/var/www/webifylab/.env
scp infra/nginx/*.conf user@IP_VPS:/var/www/webifylab/infra/nginx/
scp infra/postgres/postgresql.conf user@IP_VPS:/var/www/webifylab/infra/postgres/
scp infra/scripts/*.sh user@IP_VPS:/var/www/webifylab/infra/scripts/

# Step 5: Di VPS, load images dan start containers
ssh user@IP_VPS
cd /var/www/webifylab
docker load -i frontend.tar.gz
docker load -i backend.tar.gz
docker compose up -d

# Step 6: Tunggu PostgreSQL siap, lalu migrate & seed
sleep 15
docker compose exec backend ./server migrate
docker compose exec backend ./server seed

# Step 7: Setup Nginx (lihat Section 6)
sudo ln -sf /var/www/webifylab/infra/nginx/webifylab.conf /etc/nginx/sites-available/webifylab
sudo ln -sf /var/www/webifylab/infra/nginx/api.conf /etc/nginx/sites-available/api
sudo ln -sf /etc/nginx/sites-available/webifylab /etc/nginx/sites-enabled/
sudo ln -sf /etc/nginx/sites-available/api /etc/nginx/sites-enabled/
sudo rm -f /etc/nginx/sites-enabled/default
sudo nginx -t
sudo systemctl reload nginx

# Step 8: Setup backup cron
chmod +x /var/www/webifylab/infra/scripts/backup.sh
sudo crontab -e
# Tambahkan: 0 2 * * * /var/www/webifylab/infra/scripts/backup.sh

# Step 9: Cleanup
rm -f frontend.tar.gz backend.tar.gz
exit  # Kembali ke laptop
rm -f frontend.tar.gz backend.tar.gz
```

#### 7.2 Deployment Selanjutnya (Update Kode)

Setelah deployment pertama, untuk update kode:

```bash
# Dari laptop, cukup jalankan:
./infra/scripts/deploy.sh
```

Script ini akan otomatis:
1. Build ulang Docker images
2. Upload ke VPS
3. Load images baru
4. Restart containers
5. Run migrations (jika ada)
6. Cleanup file sementara

**Catatan:** Nginx config dan .env TIDAK akan berubah kecuali kamu memodifikasi file tersebut.

#### 7.3 Deployment Darurat (Hotfix)

Jika perlu fix cepat tanpa full deployment:

```bash
# Build hanya image yang berubah (misal: backend)
docker build -t webifylab-backend:latest -f backend/Dockerfile backend/
docker save webifylab-backend:latest | gzip > backend.tar.gz

# Upload dan restart
scp backend.tar.gz user@IP_VPS:/var/www/webifylab/
ssh user@IP_VPS "cd /var/www/webifylab && docker load -i backend.tar.gz && docker compose restart backend && rm backend.tar.gz"

# Cleanup lokal
rm backend.tar.gz
```

---

### 8. VERIFIKASI DEPLOYMENT

#### 8.1 Checklist Verifikasi

Jalankan semua perintah ini setelah deployment:

```bash
# === DI VPS ===

# 1. Cek semua container berjalan
docker compose ps
# Output yang diharapkan:
# NAME                   STATUS          PORTS
# webifylab-frontend     Up (healthy)    0.0.0.0:3000->3000/tcp
# webifylab-backend      Up              0.0.0.0:8080->8080/tcp
# webifylab-postgres     Up (healthy)    5432/tcp

# 2. Cek penggunaan RAM
docker stats --no-stream
# Pastikan total < 800MB

# 3. Cek logs backend
docker compose logs backend --tail 20

# 4. Cek logs frontend
docker compose logs frontend --tail 20

# 5. Cek logs postgres
docker compose logs postgres --tail 10

# 6. Test health check dari VPS
curl http://localhost:8080/api/v1/health

# 7. Test frontend dari VPS
curl -s -o /dev/null -w "%{http_code}" http://localhost:3000

# 8. Cek Nginx
sudo nginx -t
sudo systemctl status nginx

# 9. Cek SSL
echo | openssl s_client -connect webifylab.my.id:443 -servername webifylab.my.id 2>/dev/null | openssl x509 -noout -dates
echo | openssl s_client -connect api.webifylab.my.id:443 -servername api.webifylab.my.id 2>/dev/null | openssl x509 -noout -dates

# 10. Cek disk space
df -h

# 11. Cek swap
free -h

# === DARI LAPTOP (atau browser) ===

# 12. Test API dari luar
curl https://api.webifylab.my.id/api/v1/health

# 13. Test frontend dari luar
curl -s -o /dev/null -w "%{http_code}" https://webifylab.my.id

# 14. Test admin page
curl -s -o /dev/null -w "%{http_code}" https://webifylab.my.id/admin/login

# 15. Test contact form (submit lead)
curl -X POST https://api.webifylab.my.id/api/v1/leads \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","email":"test@test.com","phone":"6281234567890","service_type":"web_development","budget_range":"5-15jt","description":"Test deployment"}'

# 16. Cek email auto-reply (cek inbox test@test.com)
# 17. Cek email notifikasi admin (cek inbox admin@webifylab.my.id)
```

#### 8.2 Expected Results

| Test | Expected Result |
|---|---|
| `docker compose ps` | 3 containers Up |
| `docker stats` | Total RAM < 800MB |
| Health check | `{"success":true,"data":{"status":"ok"}}` |
| Frontend curl | HTTP 200 |
| SSL check | Certificate valid, not expired |
| API dari luar | JSON response |
| Lead submit | HTTP 201 + email terkirim |
| Admin page | HTTP 200 (login form) |

---

### 9. KOMANDO MANAJEMEN HARIAN

#### 9.1 Komando yang Sering Dipakai

```bash
# === DI VPS ===

# Lihat status containers
cd /var/www/webifylab && docker compose ps

# Lihat logs
docker compose logs -f              # Semua services
docker compose logs -f backend      # Backend saja
docker compose logs -f frontend     # Frontend saja
docker compose logs --tail 100      # 100 baris terakhir

# Restart services
docker compose restart              # Semua
docker compose restart backend      # Backend saja
docker compose restart frontend     # Frontend saja

# Stop semua
docker compose down

# Start semua
docker compose up -d

# Masuk ke container (debugging)
docker compose exec backend sh
docker compose exec postgres psql -U webifylab

# Cek resource usage
docker stats

# Cek disk usage
docker system df
```

#### 9.2 Database Management

```bash
# Masuk ke PostgreSQL shell
docker compose exec postgres psql -U webifylab

# Di dalam psql:
\dt                    # List semua tabel
\d blog_posts          # Detail tabel blog_posts
SELECT count(*) FROM leads;  # Hitung leads
\q                     # Keluar

# Backup manual
docker compose exec -T postgres pg_dump -U webifylab webifylab > backup_manual.sql

# Restore manual
cat backup_manual.sql | docker compose exec -T postgres psql -U webifylab webifylab
```

#### 9.3 Update Deployment

```bash
# Dari laptop, setelah update kode:
./infra/scripts/deploy.sh

# Atau jika hanya update config Nginx:
scp infra/nginx/*.conf user@IP_VPS:/var/www/webifylab/infra/nginx/
ssh user@IP_VPS "sudo cp /var/www/webifylab/infra/nginx/*.conf /etc/nginx/sites-available/ && sudo nginx -t && sudo systemctl reload nginx"
```

---

### 10. TROUBLESHOOTING

#### 10.1 Container Tidak Mau Start

**Gejala:** `docker compose ps` menunjukkan container `Exited`

**Diagnosa:**
```bash
# Cek logs container yang gagal
docker compose logs backend --tail 50
docker compose logs frontend --tail 50
docker compose logs postgres --tail 50
```

**Penyebab umum & solusi:**

| Penyebab | Solusi |
|---|---|
| PostgreSQL belum siap | Tunggu healthcheck pass, atau restart: `docker compose restart backend` |
| Environment variable salah | Cek `.env`: `cat .env` |
| Port sudah dipakai | Cek: `sudo lsof -i :3000` atau `sudo lsof -i :8080` |
| Image corrupt | Rebuild: `docker compose build --no-cache` |
| RAM habis | Cek: `free -h`, kill process tidak perlu |

#### 10.2 OOM (Out of Memory)

**Gejala:** Container tiba-tiba restart, logs menunjukkan `Killed` atau `OOM`

**Diagnosa:**
```bash
# Cek apakah OOM terjadi
dmesg | grep -i "out of memory"
docker inspect webifylab-backend | grep -i oom
free -h
```

**Solusi:**
```bash
# 1. Pastikan swap aktif
free -h  # Harus ada swap 2GB

# 2. Kurangi memory limit PostgreSQL
# Edit docker-compose.yml, ubah postgres memory limit ke 150M

# 3. Restart containers
docker compose restart

# 4. Jika masih OOM, stop service yang tidak perlu
sudo systemctl stop snapd  # Jika ada, bisa makan RAM
sudo systemctl stop unattended-upgrades
```

#### 10.3 Nginx Error: 502 Bad Gateway

**Gejala:** Buka website dapat 502 Bad Gateway

**Diagnosa:**
```bash
# Cek apakah container berjalan
docker compose ps

# Cek Nginx error log
sudo tail -50 /var/log/nginx/webifylab_error.log
sudo tail -50 /var/log/nginx/api_error.log

# Cek apakah port 3000/8080 listening
sudo ss -tlnp | grep -E '3000|8080'
```

**Solusi:**
```bash
# Biasanya container belum siap atau crash
docker compose restart

# Jika container tidak mau start, cek logs
docker compose logs backend --tail 50
```

#### 10.4 SSL Certificate Error

**Gejala:** Browser menunjukkan "Not Secure" atau "Certificate Error"

**Diagnosa:**
```bash
# Cek certificate
sudo certbot certificates

# Cek expiry date
echo | openssl s_client -connect webifylab.my.id:443 2>/dev/null | openssl x509 -noout -dates
```

**Solusi:**
```bash
# Renew certificate manual
sudo systemctl stop nginx
sudo certbot renew --force-renewal
sudo systemctl start nginx

# Verifikasi
sudo certbot certificates
```

#### 10.5 Database Connection Error

**Gejala:** Backend logs menunjukkan `connection refused` atau `dial tcp`

**Diagnosa:**
```bash
# Cek apakah PostgreSQL container berjalan
docker compose ps postgres

# Cek PostgreSQL logs
docker compose logs postgres --tail 20

# Test koneksi dari backend container
docker compose exec backend sh -c "nc -zv postgres 5432"
```

**Solusi:**
```bash
# Restart PostgreSQL
docker compose restart postgres

# Tunggu healthcheck pass
sleep 15
docker compose ps postgres  # Harus "healthy"

# Restart backend
docker compose restart backend
```

#### 10.6 Email Tidak Terkirim

**Gejala:** Submit lead berhasil tapi tidak ada email

**Diagnosa:**
```bash
# Cek backend logs untuk error email
docker compose logs backend | grep -i "email\|resend\|mail"

# Cek apakah RESEND_API_KEY ter-set
docker compose exec backend env | grep RESEND
```

**Solusi:**
```bash
# Verifikasi API key di .env
cat .env | grep RESEND

# Test kirim email manual dari container
docker compose exec backend sh -c "curl -X POST https://api.resend.com/emails -H 'Authorization: Bearer $RESEND_API_KEY' -H 'Content-Type: application/json' -d '{\"from\":\"onboarding@resend.dev\",\"to\":\"admin@webifylab.my.id\",\"subject\":\"Test\",\"html\":\"<p>Test email</p>\"}'"
```

#### 10.7 Upload File Gagal

**Gejala:** Upload media return error

**Diagnosa:**
```bash
# Cek permission folder uploads
ls -la /var/www/webifylab/uploads/

# Cek apakah volume mount benar
docker compose exec backend ls -la /app/uploads/

# Cek Nginx client_max_body_size
grep client_max_body_size /etc/nginx/sites-available/api
```

**Solusi:**
```bash
# Fix permission
sudo chmod -R 755 /var/www/webifylab/uploads/

# Pastikan volume mount di docker-compose.yml benar
# volumes:
#   - uploads:/app/uploads

# Restart containers
docker compose restart backend
```

#### 10.8 Disk Space Penuh

**Gejala:** `df -h` menunjukkan disk > 90%

**Diagnosa:**
```bash
df -h
docker system df
du -sh /var/backups/webifylab/*
du -sh /var/www/webifylab/uploads/*
```

**Solusi:**
```bash
# Bersihkan Docker resources yang tidak dipakai
docker system prune -af

# Hapus backup lama
find /var/backups/webifylab -type f -mtime +7 -delete

# Hapus log lama
sudo journalctl --vacuum-time=7d
sudo find /var/log -name "*.gz" -mtime +7 -delete

# Cek Docker images yang tidak dipakai
docker images
docker rmi <image_id_yang_tidak_dipakai>
```

---

### 11. MONITORING & MAINTENANCE

#### 11.1 Monitoring Harian

```bash
# Cek status containers
docker compose ps

# Cek RAM usage
docker stats --no-stream
free -h

# Cek disk usage
df -h

# Cek error logs
docker compose logs backend --since 24h | grep -i "error\|panic\|fatal"
docker compose logs frontend --since 24h | grep -i "error"

# Cek Nginx error log
sudo tail -20 /var/log/nginx/webifylab_error.log
sudo tail -20 /var/log/nginx/api_error.log
```

#### 11.2 Maintenance Mingguan

```bash
# Update sistem
sudo apt update && sudo apt upgrade -y

# Bersihkan Docker
docker system prune -f

# Cek SSL certificate expiry
sudo certbot certificates

# Cek backup terakhir
ls -lh /var/backups/webifylab/

# Test restore backup (di environment test, bukan production!)
```

#### 11.3 Maintenance Bulanan

```bash
# Review logs untuk pattern error
docker compose logs backend --since 30d | grep -c "error"

# Cek database size
docker compose exec postgres psql -U webifylab -c "SELECT pg_size_pretty(pg_database_size('webifylab'));"

# Review disk usage trend
df -h

# Test backup & restore
./infra/scripts/backup.sh
# Lalu test restore di local

# Review dan rotate API keys (jika perlu)
```

#### 11.4 Setup Uptime Monitoring (Opsional tapi Recommended)

**Option A: Uptime Kuma (Self-hosted)**
```bash
# Install Uptime Kuma di VPS (pakai Docker)
docker run -d --restart=always -p 3001:3001 \
  -v uptime-kuma:/app/data \
  --name uptime-kuma \
  louislam/uptime-kuma:1

# Akses di: http://IP_VPS:3001
# Tambahkan monitors:
# - webifylab.my.id (HTTPS, interval 1 menit)
# - api.webifylab.my.id/api/v1/health (HTTPS, interval 1 menit)
```

**Option B: UptimeRobot (Free, External)**
- Daftar di https://uptimerobot.com
- Tambahkan monitor untuk webifylab.my.id dan api.webifylab.my.id
- Setup alert ke email

#### 11.5 Backup Verification

**PENTING:** Backup yang tidak pernah di-test restore = tidak ada backup.

```bash
# Test restore backup (lakukan di laptop, bukan di VPS production!)

# 1. Download backup terbaru dari VPS
scp user@IP_VPS:/var/backups/webifylab/db_YYYYMMDD_HHMMSS.sql.gz ./

# 2. Restore ke database lokal
gunzip -c db_YYYYMMDD_HHMMSS.sql.gz | psql -U webifylab_dev -d webifylab_dev

# 3. Verifikasi data
psql -U webifylab_dev -d webifylab_dev -c "SELECT count(*) FROM blog_posts;"
psql -U webifylab_dev -d webifylab_dev -c "SELECT count(*) FROM leads;"
```

Lakukan test ini **minimal sebulan sekali**.

---

### 12. ROLLBACK PROCEDURE

Jika deployment baru menyebabkan masalah, rollback ke versi sebelumnya:

#### 12.1 Rollback Backend

```bash
# Di VPS
cd /var/www/webifylab

# Lihat image yang tersedia
docker images webifylab-backend

# Jika image lama masih ada, tag ulang
docker tag webifylab-backend:<old_tag> webifylab-backend:latest

# Restart
docker compose restart backend
```

#### 12.2 Rollback Database

```bash
# Gunakan restore script
/var/www/webifylab/infra/scripts/restore.sh /var/backups/webifylab/db_YYYYMMDD_HHMMSS.sql.gz
```

#### 12.3 Rollback Full

```bash
# Di laptop, checkout commit sebelumnya
cd webifylab-portal
git log --oneline -10  # Lihat commit history
git checkout <commit_hash_yang_stabil>

# Deploy ulang
./infra/scripts/deploy.sh
```

---

### 13. SECURITY CHECKLIST POST-DEPLOYMENT

| No | Item | Cara Verifikasi | Status |
|---|---|---|---|
| S1 | HTTPS aktif untuk kedua domain | `curl -I https://webifylab.my.id` → 200 | ⬜ |
| S2 | HTTP redirect ke HTTPS | `curl -I http://webifylab.my.id` → 301 | ⬜ |
| S3 | Security headers ada | `curl -I https://webifylab.my.id` → cek X-Frame-Options dll | ⬜ |
| S4 | PostgreSQL tidak ter-expose ke luar | `nmap -p 5432 IP_VPS` → filtered/closed | ⬜ |
| S5 | Port 3000 & 8080 tidak ter-expose ke luar | `nmap -p 3000,8080 IP_VPS` → filtered/closed | ⬜ |
| S6 | Rate limiting berfungsi | Kirim 10 request cepat ke /auth/login → 429 | ⬜ |
| S7 | File .env tidak ter-commit ke git | `git ls-files .env` → kosong | ⬜ |
| S8 | SSL auto-renewal aktif | `sudo certbot renew --dry-run` → success | ⬜ |
| S9 | Backup cron aktif | `sudo crontab -l` → ada backup script | ⬜ |
| S10 | SSH key-based auth (bukan password) | Cek `/etc/ssh/sshd_config` | ⬜ |

**Catatan untuk S5:** Saat ini port 3000 dan 8080 di-expose ke host di docker-compose.yml. Untuk keamanan tambahan, bisa di-bind ke localhost saja:

```yaml
ports:
  - "127.0.0.1:3000:3000"   # Hanya accessible dari localhost
  - "127.0.0.1:8080:8080"   # Hanya accessible dari localhost
```

Ini memastikan hanya Nginx di host yang bisa akses container, bukan dari luar.

---

### 14. QUICK REFERENCE CARD

Cetak atau simpan ini untuk referensi cepat:

```
┌─────────────────────────────────────────────────────────────┐
│           WEBIFYLAB PORTAL - QUICK REFERENCE                │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  DEPLOYMENT (dari laptop):                                  │
│  ./infra/scripts/deploy.sh                                  │
│                                                             │
│  CEK STATUS (di VPS):                                       │
│  cd /var/www/webifylab && docker compose ps                 │
│  docker stats --no-stream                                   │
│  free -h                                                    │
│                                                             │
│  LIHAT LOGS:                                                │
│  docker compose logs -f backend                             │
│  docker compose logs -f frontend                            │
│  sudo tail -f /var/log/nginx/webifylab_error.log            │
│                                                             │
│  RESTART:                                                   │
│  docker compose restart                                     │
│  docker compose restart backend                             │
│                                                             │
│  DATABASE:                                                  │
│  docker compose exec postgres psql -U webifylab             │
│                                                             │
│  BACKUP:                                                    │
│  /var/www/webifylab/infra/scripts/backup.sh                 │
│                                                             │
│  RESTORE:                                                   │
│  /var/www/webifylab/infra/scripts/restore.sh <file.sql.gz>  │
│                                                             │
│  ROLLBACK:                                                  │
│  git checkout <commit_stabil> && ./infra/scripts/deploy.sh  │
│                                                             │
│  URLS:                                                      │
│  Frontend: https://webifylab.my.id                          │
│  API:      https://api.webifylab.my.id                      │
│  Admin:    https://webifylab.my.id/admin                    │
│  Health:   https://api.webifylab.my.id/api/v1/health        │
│                                                             │
│  FILE PENTING DI VPS:                                       │
│  /var/www/webifylab/docker-compose.yml                      │
│  /var/www/webifylab/.env                                    │
│  /etc/nginx/sites-available/webifylab                       │
│  /etc/nginx/sites-available/api                             │
│  /etc/letsencrypt/live/webifylab.my.id/                     │
│  /etc/letsencrypt/live/api.webifylab.my.id/                 │
│  /var/backups/webifylab/                                    │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

### 15. CHANGELOG

| Versi | Tanggal | Perubahan |
|---|---|---|
| 1.0.0 | 2026-09-05 | Initial release — Full deployment guide |