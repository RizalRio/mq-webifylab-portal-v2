# 🏗️ Technical Design Document (TDD)

## Webifylab Landing Page — Version 1.0

| Metadata       | Detail                                        |
| -------------- | --------------------------------------------- |
| **Product**    | Webifylab Landing Page                        |
| **Version**    | 1.0                                           |
| **Author**     | Rizal                                         |
| **Role**       | System Analyst / Software Engineer            |
| **Created**    | 17 September 2026                             |
| **Status**     | Draft                                         |
| **Tech Stack** | Astro (Frontend), Golang (Backend API - V1.5) |
| **Hosting**    | VPS 1 vCPU, 1 GB RAM, 20 GB SSD               |

---

## 1. Architecture Overview

### 1.1 Design Principles

| Principle                   | Deskripsi                                                                                |
| --------------------------- | ---------------------------------------------------------------------------------------- |
| **Static-First**            | Landing page V1 100% static (HTML/CSS/JS) untuk performa maksimal dan hemat resource VPS |
| **Progressive Enhancement** | Backend (Golang) ditambahkan bertahap saat dibutuhkan                                    |
| **Resource Efficiency**     | Setiap komponen harus ringan — VPS hanya 1GB RAM                                         |
| **Future-Proof**            | Struktur siap untuk ekspansi ke SaaS & AI ecosystem                                      |
| **Security by Default**     | HTTPS, CSP headers, no inline scripts                                                    |
| **Observable**              | Logging & monitoring sejak hari pertama                                                  |

### 1.2 High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         INTERNET                                │
└───────────────────────────┬─────────────────────────────────────┘
                            │ HTTPS (443)
                            ▼
┌─────────────────────────────────────────────────────────────────┐
│                    VPS (1 vCPU, 1GB RAM, 20GB SSD)              │
│                                                                 │
│  ┌───────────────────────────────────────────────────────────┐  │
│  │                    NGINX (Reverse Proxy)                  │  │
│  │  • SSL Termination (Let's Encrypt)                        │  │
│  │  • Serve static files (Astro build output)                │  │
│  │  • Gzip/Brotli compression                                │  │
│  │  • Security headers (CSP, X-Frame-Options, etc)           │  │
│  │  • Rate limiting                                          │  │
│  └───────────────────────────┬───────────────────────────────┘  │
│                              │                                  │
│           ┌──────────────────┴──────────────────┐               │
│           ▼                                     ▼               │
│  ┌────────────────────┐              ┌────────────────────┐    │
│  │ Static Files       │              │ Golang API (V1.5+) │    │
│  │ /var/www/webifylab │              │ :8080              │    │
│  │ (Astro build)      │              │ • /api/contact     │    │
│  │                    │              │ • /api/health      │    │
│  │ ~500KB total       │              │ • /api/analytics   │    │
│  └────────────────────┘              └────────────────────┘    │
│                                                                 │
│  Resource Usage (Estimated):                                   │
│  • Nginx: ~5MB RAM                                            │
│  • Golang API: ~20MB RAM (saat idle)                          │
│  • OS + System: ~300MB RAM                                    │
│  • Available: ~675MB RAM (buffer aman)                        │
└─────────────────────────────────────────────────────────────────┘
                            │
                            │ External Services (V1)
                            ▼
┌─────────────────────────────────────────────────────────────────┐
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────┐     │
│  │  Formspree   │  │  Plausible / │  │  WhatsApp API    │     │
│  │  (Form)      │  │  GA (Track)  │  │  (wa.me)         │     │
│  └──────────────┘  └──────────────┘  └──────────────────┘     │
└─────────────────────────────────────────────────────────────────┘
```

### 1.3 Technology Stack Details

| Layer                      | Technology                | Version   | Alasan                                                      |
| -------------------------- | ------------------------- | --------- | ----------------------------------------------------------- |
| **Frontend Framework**     | Astro                     | 4.x       | SSG terbaik untuk content-focused sites, zero JS by default |
| **Styling**                | Tailwind CSS              | 3.x       | Utility-first, tree-shaking, rapid development              |
| **Icons**                  | Lucide Astro              | Latest    | Ringan, konsisten, tree-shakeable                           |
| **Fonts**                  | Inter / Plus Jakarta Sans | -         | Modern, readable, variable font                             |
| **Backend API (V1.5)**     | Golang                    | 1.22+     | Efficient, low memory, cocok untuk VPS 1GB                  |
| **HTTP Router (Go)**       | Chi atau Fiber            | Latest    | Lightweight, fast, idiomatic                                |
| **Web Server**             | Nginx                     | 1.24+     | Stabil, ringan, fitur lengkap                               |
| **SSL**                    | Let's Encrypt + Certbot   | -         | Gratis, auto-renew                                          |
| **Form Handling (V1)**     | Formspree                 | Free tier | Tanpa backend, cepat setup                                  |
| **Analytics (V1)**         | Plausible / GA4           | -         | Plausible lebih ringan & privacy-friendly                   |
| **Process Manager (V1.5)** | Systemd                   | -         | Native Linux, auto-restart                                  |
| **Version Control**        | Git + GitHub              | -         | Standard industry                                           |
| **CI/CD (Opsional)**       | GitHub Actions            | -         | Auto deploy on push                                         |
| **Monitoring**             | UptimeRobot (free) + htop | -         | External + internal monitoring                              |

---

## 2. Project Structure

### 2.1 Repository Structure (Monorepo)

```
webifylab/
├── .github/
│   └── workflows/
│       └── deploy.yml              # CI/CD pipeline
├── apps/
│   ├── web/                        # Astro frontend (V1)
│   │   ├── src/
│   │   │   ├── components/         # Reusable UI components
│   │   │   │   ├── ui/             # Base components (Button, Card, Input)
│   │   │   │   ├── sections/       # Section components (Hero, Services, etc)
│   │   │   │   └── layout/         # Header, Footer, Navbar
│   │   │   ├── layouts/            # Astro layouts
│   │   │   ├── pages/              # Astro pages
│   │   │   │   └── index.astro     # Landing page
│   │   │   ├── styles/             # Global styles
│   │   │   ├── assets/             # Images, fonts, icons
│   │   │   ├── lib/                # Utility functions
│   │   │   └── content/            # Content collections (blog, portfolio)
│   │   ├── public/                 # Static assets
│   │   ├── astro.config.mjs
│   │   ├── tailwind.config.js
│   │   ├── package.json
│   │   └── tsconfig.json
│   │
│   └── api/                        # Golang backend (V1.5)
│       ├── cmd/
│       │   └── server/
│       │       └── main.go         # Entry point
│       ├── internal/
│       │   ├── handlers/           # HTTP handlers
│       │   │   ├── contact.go
│       │   │   ├── health.go
│       │   │   └── analytics.go
│       │   ├── models/             # Data models
│       │   ├── services/           # Business logic
│       │   ├── repositories/       # Data access
│       │   ├── middleware/         # Auth, CORS, logging
│       │   └── config/             # Configuration
│       ├── pkg/                    # Shared packages
│       │   ├── logger/
│       │   ├── validator/
│       │   └── response/
│       ├── migrations/             # DB migrations (V2)
│       ├── go.mod
│       ├── go.sum
│       ├── Makefile
│       └── Dockerfile
│
├── infra/                          # Infrastructure as Code
│   ├── nginx/
│   │   └── webifylab.conf          # Nginx config
│   ├── systemd/
│   │   └── webifylab-api.service   # Systemd service (V1.5)
│   └── scripts/
│       ├── deploy.sh               # Deployment script
│       ├── backup.sh               # Backup script
│       └── setup-vps.sh            # VPS initial setup
│
├── docs/                           # Documentation
│   ├── PRD.md
│   ├── TDD.md                      # This document
│   ├── CONTENT.md
│   └── DESIGN_SYSTEM.md
│
├── .env.example
├── .gitignore
├── Makefile                        # Root Makefile
└── README.md
```

### 2.2 Astro Frontend Structure Detail

```
apps/web/src/
├── components/
│   ├── ui/
│   │   ├── Button.astro            # Reusable button component
│   │   ├── Card.astro              # Card component
│   │   ├── Input.astro             # Form input
│   │   ├── Badge.astro             # Badge/tag
│   │   └── Container.astro         # Max-width container
│   ├── sections/
│   │   ├── Hero.astro              # S2
│   │   ├── Problem.astro           # S3
│   │   ├── Services.astro          # S4
│   │   ├── WhyWebifylab.astro      # S5
│   │   ├── Approach.astro          # S6
│   │   ├── Portfolio.astro         # S7
│   │   ├── Testimonial.astro       # S8
│   │   ├── Contact.astro           # S9
│   │   └── FAQ.astro               # Bonus
│   └── layout/
│       ├── Header.astro            # Navbar
│       ├── Footer.astro            # Footer
│       └── BaseLayout.astro        # HTML skeleton
├── layouts/
│   └── Layout.astro                # Main layout wrapper
├── pages/
│   ├── index.astro                 # Landing page
│   ├── 404.astro                   # Not found page
│   └── privacy.astro               # Privacy policy (optional)
├── styles/
│   └── global.css                  # Tailwind imports + custom styles
├── content/
│   ├── portfolio/                  # Content collection
│   │   ├── project-1.md
│   │   └── project-2.md
│   └── config.ts                   # Content collection schema
└── lib/
    ├── utils.ts                    # Utility functions
    └── constants.ts                # Site constants
```

### 2.3 Golang Backend Structure Detail (V1.5)

```
apps/api/
├── cmd/server/
│   └── main.go                     # Entry point, setup router
├── internal/
│   ├── handlers/
│   │   ├── contact.go              # POST /api/contact
│   │   ├── health.go               # GET /api/health
│   │   └── analytics.go            # POST /api/analytics/event
│   ├── models/
│   │   ├── contact.go              # ContactRequest struct
│   │   └── response.go             # Standard API response
│   ├── services/
│   │   ├── contact_service.go      # Business logic for contact
│   │   └── notification.go         # Send email/Telegram notification
│   ├── middleware/
│   │   ├── cors.go                 # CORS middleware
│   │   ├── ratelimit.go            # Rate limiting
│   │   ├── logger.go               # Request logging
│   │   └── recovery.go             # Panic recovery
│   └── config/
│       └── config.go               # Load env vars
├── pkg/
│   ├── logger/
│   │   └── logger.go               # Structured logging (zerolog)
│   ├── validator/
│   │   └── validator.go            # Input validation
│   └── response/
│       └── response.go             # JSON response helpers
├── go.mod
├── go.sum
├── Makefile                        # Build, test, run commands
└── Dockerfile                      # Multi-stage build
```

---

## 3. VPS Configuration Plan

### 3.1 VPS Specifications Reminder

| Resource  | Value                        | Notes                                  |
| --------- | ---------------------------- | -------------------------------------- |
| CPU       | 1 vCPU                       | Shared, cukup untuk static + light API |
| RAM       | 1 GB                         | Sangat terbatas, harus efisien         |
| Storage   | 20 GB SSD                    | Cukup untuk static files + logs        |
| Bandwidth | Unmetered                    | Tidak perlu khawatir traffic           |
| OS        | Ubuntu 22.04 LTS / Debian 12 | Stabil, well-supported                 |

### 3.2 Memory Budget Plan

```
Total RAM: 1024 MB
├── OS + System Services:        ~300 MB
├── Nginx:                        ~5 MB
├── Golang API (V1.5, idle):     ~20 MB
├── Golang API (V1.5, peak):     ~50 MB
├── Certbot (cron):               ~0 MB (only during renewal)
├── Logging buffer:              ~20 MB
├── Buffer/Cache:                ~100 MB
└── AVAILABLE (safety margin):   ~529 MB ✅
```

**Kesimpulan:** Konfigurasi ini **sangat aman** dengan margin ~50% RAM tersedia.

### 3.3 Initial VPS Setup Script (infra/scripts/setup-vps.sh)

```bash
#!/bin/bash
# Initial VPS setup for Webifylab

# 1. Update system
sudo apt update && sudo apt upgrade -y

# 2. Install essential packages
sudo apt install -y \
    nginx \
    certbot python3-certbot-nginx \
    ufw \
    fail2ban \
    htop \
    curl \
    git \
    make

# 3. Install Golang (for V1.5)
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# 4. Configure UFW (firewall)
sudo ufw allow OpenSSH
sudo ufw allow 'Nginx Full'
sudo ufw enable

# 5. Configure fail2ban (brute force protection)
sudo systemctl enable fail2ban

# 6. Create web directory
sudo mkdir -p /var/www/webifylab
sudo chown -R $USER:$USER /var/www/webifylab

# 7. Setup swap (safety net for RAM)
sudo fallocate -l 1G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab

# 8. Optimize sysctl for web server
cat << EOF | sudo tee -a /etc/sysctl.conf
# Web server optimization
net.core.somaxconn = 1024
net.ipv4.tcp_max_syn_backlog = 1024
vm.swappiness = 10
EOF
sudo sysctl -p

echo "✅ VPS setup complete!"
```

### 3.4 Nginx Configuration (infra/nginx/webifylab.conf)

```nginx
# /etc/nginx/sites-available/webifylab

# Rate limiting zone
limit_req_zone $binary_remote_addr zone=api_limit:10m rate=10r/s;

server {
    listen 80;
    server_name webifylab.com www.webifylab.com;
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name webifylab.com www.webifylab.com;

    # SSL (managed by Certbot)
    ssl_certificate /etc/letsencrypt/live/webifylab.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/webifylab.com/privkey.pem;
    include /etc/letsencrypt/options-ssl-nginx.conf;
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

    # Security headers
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;
    add_header Content-Security-Policy "default-src 'self'; script-src 'self' 'unsafe-inline' https://plausible.io; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self'; connect-src 'self' https://formspree.io https://plausible.io;" always;

    # Root directory (Astro build output)
    root /var/www/webifylab/dist;
    index index.html;

    # Gzip compression
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml text/javascript image/svg+xml;

    # Cache static assets
    location ~* \.(jpg|jpeg|png|gif|ico|css|js|svg|woff|woff2|ttf|eot)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
        access_log off;
    }

    # API proxy (V1.5)
    location /api/ {
        limit_req zone=api_limit burst=20 nodelay;
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # SPA fallback (for future routes)
    location / {
        try_files $uri $uri/ $uri.html =404;
    }

    # Health check endpoint
    location = /health {
        access_log off;
        return 200 "OK";
        add_header Content-Type text/plain;
    }

    # Deny access to hidden files
    location ~ /\. {
        deny all;
        access_log off;
        log_not_found off;
    }
}
```

---

## 4. Deployment Strategy

### 4.1 Deployment Flow

```
┌──────────────┐      ┌──────────────┐      ┌──────────────┐
│   Local Dev  │ ───► │   GitHub     │ ───► │     VPS      │
│  (Laptop)    │      │  Repository  │      │  (Production)│
└──────────────┘      └──────────────┘      └──────────────┘
      │                       │                     │
      │ 1. git push           │ 2. GitHub Actions   │ 3. Deploy
      │                       │    triggers         │
      │                       │                     │
      ▼                       ▼                     ▼
┌──────────────┐      ┌──────────────┐      ┌──────────────┐
│ Build Astro  │      │ Run tests    │      │ Pull & Build │
│ locally      │      │ Lint, format │      │ Restart nginx│
│ Test         │      │ Build        │      │              │
└──────────────┘      └──────────────┘      └──────────────┘
```

### 4.2 Deployment Options

#### Option A: Manual Deployment (Recommended for V1)

**Pros:** Simple, no CI/CD setup needed, full control
**Cons:** Manual steps, prone to human error

```bash
# Local: Build Astro
cd apps/web
npm run build

# Local: Upload to VPS
rsync -avz --delete dist/ user@vps-ip:/var/www/webifylab/dist/

# VPS: Reload nginx (if needed)
ssh user@vps-ip "sudo nginx -t && sudo systemctl reload nginx"
```

#### Option B: GitHub Actions CI/CD (Recommended for V1.5+)

**.github/workflows/deploy.yml:**

```yaml
name: Deploy to VPS

on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Node
        uses: actions/setup-node@v4
        with:
          node-version: "20"

      - name: Install dependencies
        working-directory: ./apps/web
        run: npm ci

      - name: Build Astro
        working-directory: ./apps/web
        run: npm run build

      - name: Deploy to VPS
        uses: appleboy/scp-action@master
        with:
          host: ${{ secrets.VPS_HOST }}
          username: ${{ secrets.VPS_USER }}
          key: ${{ secrets.SSH_PRIVATE_KEY }}
          source: "./apps/web/dist/*"
          target: "/var/www/webifylab/dist"
          strip_components: 3

      - name: Reload Nginx
        uses: appleboy/ssh-action@master
        with:
          host: ${{ secrets.VPS_HOST }}
          username: ${{ secrets.VPS_USER }}
          key: ${{ secrets.SSH_PRIVATE_KEY }}
          script: |
            sudo nginx -t
            sudo systemctl reload nginx
```

### 4.3 Deployment Checklist

- [ ] SSH key configured (no password login)
- [ ] GitHub secrets set (VPS_HOST, VPS_USER, SSH_PRIVATE_KEY)
- [ ] Domain DNS pointed to VPS IP
- [ ] SSL certificate issued (Let's Encrypt)
- [ ] Nginx config tested (`nginx -t`)
- [ ] First deployment successful
- [ ] HTTPS working
- [ ] Form submission working
- [ ] Analytics tracking

---

## 5. Security Considerations

### 5.1 Security Layers

| Layer           | Measure                   | Implementation                                        |
| --------------- | ------------------------- | ----------------------------------------------------- |
| **Network**     | Firewall                  | UFW: only allow 22 (SSH), 80 (HTTP), 443 (HTTPS)      |
| **Network**     | SSH hardening             | Key-based auth only, disable root login, change port  |
| **Network**     | Fail2ban                  | Block IPs after failed login attempts                 |
| **Transport**   | HTTPS                     | Let's Encrypt, auto-renew, HSTS enabled               |
| **Application** | CSP Headers               | Restrict script/style sources                         |
| **Application** | No inline scripts         | All JS in separate files                              |
| **Application** | Rate limiting             | Nginx limit_req for API endpoints                     |
| **Application** | Input validation          | Server-side validation for all forms                  |
| **Data**        | No sensitive data in repo | .env for secrets, .gitignore for .env                 |
| **Monitoring**  | Log monitoring            | Check /var/log/nginx/access.log & error.log regularly |

### 5.2 Secrets Management

```bash
# .env.example (NEVER commit this file)
# Formspree
FORMSPREE_ENDPOINT=https://formspree.io/f/xxxxx

# WhatsApp
WHATSAPP_NUMBER=628xxxxxxxxxx

# Golang API (V1.5)
API_PORT=8080
API_HOST=127.0.0.1
TELEGRAM_BOT_TOKEN=xxx
TELEGRAM_CHAT_ID=xxx

# Analytics
PLAUSIBLE_DOMAIN=webifylab.com
```

### 5.3 Security Audit Checklist

- [ ] SSH password authentication disabled
- [ ] Root login disabled
- [ ] UFW firewall active
- [ ] Fail2ban running
- [ ] SSL certificate valid
- [ ] HSTS header enabled
- [ ] CSP header configured
- [ ] No sensitive data in git history
- [ ] npm audit passed (no vulnerabilities)
- [ ] gosec passed (no Go vulnerabilities)

---

## 6. Monitoring & Logging

### 6.1 Monitoring Stack (Lightweight)

| Tool                        | Purpose                      | Resource Usage |
| --------------------------- | ---------------------------- | -------------- |
| **htop**                    | Real-time process monitoring | ~2MB RAM       |
| **UptimeRobot**             | External uptime monitoring   | Free, external |
| **Nginx access/error logs** | Request & error tracking     | Disk only      |
| **Golang structured logs**  | API request logging (V1.5)   | Disk only      |
| **Systemd journal**         | System logs                  | Disk only      |

### 6.2 Log Rotation

```bash
# /etc/logrotate.d/webifylab
/var/log/nginx/webifylab/*.log {
    daily
    rotate 14
    compress
    delaycompress
    missingok
    notifempty
    create 0640 www-data adm
    sharedscripts
    postrotate
        [ -f /var/run/nginx.pid ] && kill -USR1 `cat /var/run/nginx.pid`
    endscript
}
```

### 6.3 Monitoring Commands

```bash
# Check RAM usage
free -h

# Check disk usage
df -h

# Check running processes
htop

# Check Nginx status
systemctl status nginx

# Check recent errors
tail -f /var/log/nginx/error.log

# Check API logs (V1.5)
journalctl -u webifylab-api -f
```

---

## 7. Backup & Disaster Recovery

### 7.1 Backup Strategy

| What             | Frequency  | Retention             | Method                         |
| ---------------- | ---------- | --------------------- | ------------------------------ |
| Website files    | Daily      | 7 days                | rsync to local / cloud storage |
| Nginx config     | On change  | Git versioned         | Git repository                 |
| SSL certificates | Auto-renew | Let's Encrypt manages | Certbot                        |
| Database (V2+)   | Daily      | 30 days               | pg_dump + offsite backup       |
| Logs             | Daily      | 14 days               | Logrotate                      |

### 7.2 Backup Script (infra/scripts/backup.sh)

```bash
#!/bin/bash
BACKUP_DIR="/backup/webifylab"
DATE=$(date +%Y%m%d)

mkdir -p $BACKUP_DIR

# Backup website files
tar -czf $BACKUP_DIR/web-$DATE.tar.gz /var/www/webifylab/

# Backup nginx config
tar -czf $BACKUP_DIR/nginx-$DATE.tar.gz /etc/nginx/

# Delete backups older than 7 days
find $BACKUP_DIR -type f -mtime +7 -delete

echo "✅ Backup completed: $DATE"
```

### 7.3 Disaster Recovery Plan

| Scenario        | Recovery Steps                 | RTO        | RPO      |
| --------------- | ------------------------------ | ---------- | -------- |
| VPS down        | Restore from backup to new VPS | 1 hour     | 24 hours |
| SSL expired     | Run certbot renew              | 5 minutes  | 0        |
| Nginx misconfig | Restore config from Git        | 10 minutes | 0        |
| Disk full       | Clean logs, expand disk        | 30 minutes | 0        |
| RAM OOM         | Restart services, add swap     | 5 minutes  | 0        |

---

## 8. Performance Optimization

### 8.1 Frontend Performance

| Technique              | Implementation                    | Impact                     |
| ---------------------- | --------------------------------- | -------------------------- |
| **Static generation**  | Astro SSG                         | Zero server-side rendering |
| **Zero JS by default** | Astro islands only where needed   | Minimal JS payload         |
| **Image optimization** | WebP/AVIF, responsive sizes       | 50-80% smaller images      |
| **Lazy loading**       | `loading="lazy"` for images       | Faster initial load        |
| **Font optimization**  | Variable fonts, preconnect        | Faster font loading        |
| **CSS purging**        | Tailwind tree-shaking             | Minimal CSS                |
| **Compression**        | Gzip/Brotli via Nginx             | 60-80% smaller transfers   |
| **Caching**            | Long-term cache for static assets | Repeat visits instant      |

### 8.2 Performance Budget

| Metric                         | Budget   | Tool           |
| ------------------------------ | -------- | -------------- |
| Total page size                | < 500 KB | `du -sh dist/` |
| LCP (Largest Contentful Paint) | < 1.5s   | Lighthouse     |
| FCP (First Contentful Paint)   | < 1.0s   | Lighthouse     |
| CLS (Cumulative Layout Shift)  | < 0.05   | Lighthouse     |
| TBT (Total Blocking Time)      | < 100ms  | Lighthouse     |
| Lighthouse Performance         | > 95     | Lighthouse     |
| Lighthouse SEO                 | > 95     | Lighthouse     |
| Lighthouse Accessibility       | > 90     | Lighthouse     |
| Lighthouse Best Practices      | > 95     | Lighthouse     |

---

## 9. Golang Backend Design (V1.5 Preview)

### 9.1 API Endpoints

| Method | Endpoint               | Description         | Auth         |
| ------ | ---------------------- | ------------------- | ------------ |
| GET    | `/api/health`          | Health check        | None         |
| POST   | `/api/contact`         | Submit contact form | Rate limited |
| POST   | `/api/analytics/event` | Track custom events | Rate limited |

### 9.2 Contact API Design

```go
// POST /api/contact
// Request:
{
  "name": "string (required, max 100)",
  "email": "string (required, valid email)",
  "service": "string (required, enum)",
  "message": "string (required, min 10, max 2000)"
}

// Response (200 OK):
{
  "success": true,
  "message": "Pesan terkirim",
  "data": {
    "id": "uuid"
  }
}

// Response (400 Bad Request):
{
  "success": false,
  "message": "Validasi gagal",
  "errors": [
    {"field": "email", "message": "Email tidak valid"}
  ]
}
```

### 9.3 Golang Tech Stack Detail (V1.5)

| Component      | Choice                   | Reason                                             |
| -------------- | ------------------------ | -------------------------------------------------- |
| HTTP Framework | Chi                      | Lightweight, stdlib-compatible, middleware support |
| Logger         | Zerolog                  | Zero-allocation, fast, structured                  |
| Validator      | Go-playground/validator  | Industry standard                                  |
| Config         | Viper                    | Env + file config                                  |
| Notification   | Telegram Bot API         | Free, instant notification                         |
| Testing        | stdlib testing + testify | Standard + assertions                              |

---

## 10. Development Workflow

### 10.1 Local Development Setup

```bash
# 1. Clone repository
git clone https://github.com/yourusername/webifylab.git
cd webifylab

# 2. Install root dependencies (if any)
make setup

# 3. Setup frontend
cd apps/web
npm install
npm run dev          # Start Astro dev server at localhost:4321

# 4. Setup backend (V1.5)
cd ../api
go mod download
make run             # Start Go server at localhost:8080
```

### 10.2 Makefile Commands

```makefile
# Root Makefile
.PHONY: setup dev build deploy test

setup:
	cd apps/web && npm install
	cd apps/api && go mod download

dev-web:
	cd apps/web && npm run dev

dev-api:
	cd apps/api && make run

build-web:
	cd apps/web && npm run build

build-api:
	cd apps/api && make build

test-web:
	cd apps/web && npm run test

test-api:
	cd apps/api && make test

deploy:
	./infra/scripts/deploy.sh
```

### 10.3 Git Branching Strategy

```
main (production)
 │
 ├── develop (staging)
 │    │
 │    ├── feature/hero-section
 │    ├── feature/services-section
 │    └── fix/navbar-mobile
 │
 └── release/v1.0
```

- **main:** Production-ready code
- **develop:** Integration branch
- **feature/\*:** New features
- **fix/\*:** Bug fixes
- **release/\*:** Release preparation

---

## 11. Risks & Mitigation (Technical)

| Risk                     | Impact          | Likelihood | Mitigation                                   |
| ------------------------ | --------------- | ---------- | -------------------------------------------- |
| VPS 1GB OOM under load   | Website down    | Low        | Static files + swap + monitoring             |
| SSL certificate expired  | Browser warning | Low        | Certbot auto-renew + monitoring              |
| Disk full (logs)         | Service crash   | Medium     | Logrotate + monitoring                       |
| DDoS attack              | Website down    | Low        | Cloudflare (free tier) + Nginx rate limit    |
| Dependency vulnerability | Security breach | Medium     | Regular `npm audit` & `gosec`                |
| Deployment failure       | Downtime        | Low        | Test deploy script locally first             |
| VPS provider downtime    | Website down    | Low        | External monitoring + backup plan            |
| Golang learning curve    | Delay V1.5      | Medium     | Start with simple endpoints, learn gradually |

---

## 12. Future Roadmap (Technical)

### V1.0 (Now) — Landing Page

- [x] Astro static site
- [x] Nginx + SSL
- [x] Formspree integration
- [x] Basic analytics

### V1.5 (Month 2-3) — Backend Foundation

- [ ] Golang API (Chi framework)
- [ ] Contact form handler (replace Formspree)
- [ ] Self-hosted analytics endpoint
- [ ] Telegram notification for new leads
- [ ] Systemd service management

### V2.0 (Month 6+) — SaaS Ecosystem

- [ ] Authentication service (JWT)
- [ ] Multi-tenant architecture
- [ ] PostgreSQL database
- [ ] Redis cache
- [ ] Docker containerization
- [ ] Kubernetes-ready design

### V3.0 (Year 2+) — AI & Data

- [ ] Data pipeline (ETL) in Golang
- [ ] ML model serving (ONNX runtime)
- [ ] Vector database (pgvector)
- [ ] RAG system for AI features
- [ ] Real-time analytics dashboard

---

## 13. Open Questions (Technical)

| No  | Pertanyaan                                                                 | Status  |
| --- | -------------------------------------------------------------------------- | ------- |
| Q1  | Apakah VPS sudah di-setup dengan OS Ubuntu/Debian?                         | Pending |
| Q2  | Apakah domain sudah dibeli dan DNS pointed ke VPS?                         | Pending |
| Q3  | Preferensi framework Golang: Chi atau Fiber?                               | Pending |
| Q4  | Apakah ingin self-hosted analytics (Plausible) atau GA4?                   | Pending |
| Q5  | Apakah perlu CI/CD sejak V1 atau manual deploy dulu?                       | Pending |
| Q6  | Apakah ada preference untuk notification channel (Email/Telegram/Discord)? | Pending |

---

_Dokumen ini adalah living document. Versi akan diperbarui seiring perkembangan arsitektur._

**Last Updated:** 17 September 2026
**Next Step:** Setup VPS, initialize repository, mulai development Astro frontend.
