# 🚀 Deployment & Operations Runbook

## Webifylab Landing Page — Version 1.0

| Metadata         | Detail                                 |
| ---------------- | -------------------------------------- |
| **Product**      | Webifylab Landing Page                 |
| **Version**      | 1.0                                    |
| **Author**       | Rizal                                  |
| **Created**      | 17 September 2026                      |
| **Status**       | Draft                                  |
| **Architecture** | Monorepo (Astro + Golang)              |
| **VPS Specs**    | 1 vCPU, 1 GB RAM, 20 GB SSD, Unmetered |

---

## 1. Prerequisites

### 1.1 What You Need Before Starting

| Item              | Status | Notes                                |
| ----------------- | ------ | ------------------------------------ |
| **VPS**           | ☐      | Ubuntu 22.04 LTS / Debian 12         |
| **Domain**        | ☐      | webifylab.com (atau domain pilihan)  |
| **SSH Access**    | ☐      | Root atau sudo user                  |
| **Local Machine** | ☐      | Laptop Advan Workpro (Linux/WSL/Mac) |
| **Git**           | ☐      | Installed di local & VPS             |
| **Node.js**       | ☐      | v20+ (local only, untuk build Astro) |
| **Go**            | ☐      | v1.22+ (VPS, untuk V1.5)             |
| **Code Editor**   | ☐      | VS Code                              |

### 1.2 Environment Variables

Buat file `.env` di root monorepo (JANGAN commit ke Git):

```bash
# .env.example (copy ke .env dan isi)

# Domain
DOMAIN=webifylab.com
WWW_DOMAIN=www.webifylab.com

# VPS
VPS_IP=xxx.xxx.xxx.xxx
VPS_USER=deploy
VPS_SSH_PORT=22

# Golang API (V1.5)
API_PORT=8080
API_HOST=127.0.0.1

# Formspree (V1)
FORMSPREE_ENDPOINT=https://formspree.io/f/xxxxx

# WhatsApp
WHATSAPP_NUMBER=628xxxxxxxxxx

# Notification (V1.5)
TELEGRAM_BOT_TOKEN=xxx
TELEGRAM_CHAT_ID=xxx

# Database (V1.5)
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=webifylab
DB_PASSWORD=xxx
DB_NAME=webifylab
DB_SSLMODE=disable
```

---

## 2. VPS Initial Setup

### 2.1 Connect to VPS

```bash
# Login sebagai root
ssh root@xxx.xxx.xxx.xxx

# Jika menggunakan custom port
ssh root@xxx.xxx.xxx.xxx -p 2222
```

### 2.2 Update System

```bash
# Update package list
sudo apt update && sudo apt upgrade -y

# Install essential packages
sudo apt install -y \
    curl \
    wget \
    git \
    make \
    htop \
    ufw \
    fail2ban \
    nginx \
    certbot python3-certbot-nginx \
    build-essential \
    software-properties-common \
    apt-transport-https \
    ca-certificates \
    gnupg \
    lsb-release \
    unzip \
    jq
```

### 2.3 Create Deploy User

**PENTING:** Jangan gunakan root untuk deployment. Buat user khusus.

```bash
# Create user
sudo adduser deploy

# Add to sudo group
sudo usermod -aG sudo deploy

# Switch to deploy user
su - deploy
```

### 2.4 Setup SSH Key Authentication

**Di laptop lokal:**

```bash
# Generate SSH key (jika belum ada)
ssh-keygen -t ed25519 -C "rizal@webifylab.com"

# Copy public key ke VPS
ssh-copy-id deploy@xxx.xxx.xxx.xxx

# Test login tanpa password
ssh deploy@xxx.xxx.xxx.xxx
```

**Di VPS (sebagai deploy user):**

```bash
# Verify SSH key
cat ~/.ssh/authorized_keys
```

### 2.5 Harden SSH

```bash
# Edit SSH config
sudo nano /etc/ssh/sshd_config
```

**Ubah baris berikut:**

```
# /etc/ssh/sshd_config

Port 22                          # Bisa diganti ke port lain (misal 2222)
PermitRootLogin no               # Disable root login
PasswordAuthentication no        # Disable password auth
PubkeyAuthentication yes         # Enable key auth
MaxAuthTries 3                   # Limit auth attempts
ClientAliveInterval 300          # Timeout after 5 min idle
ClientAliveCountMax 2            # Disconnect after 2 timeouts
AllowUsers deploy                # Only allow deploy user
```

```bash
# Restart SSH
sudo systemctl restart sshd

# Test SSH dari terminal BARU (jangan close terminal saat ini!)
ssh deploy@xxx.xxx.xxx.xxx
```

### 2.6 Configure Firewall (UFW)

```bash
# Allow SSH
sudo ufw allow 22/tcp

# Allow HTTP & HTTPS
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp

# Enable firewall
sudo ufw enable

# Check status
sudo ufw status verbose
```

**Expected output:**

```
Status: active

To                         Action      From
--                         ------      ----
22/tcp                     ALLOW       Anywhere
80/tcp                     ALLOW       Anywhere
443/tcp                    ALLOW       Anywhere
22/tcp (v6)                ALLOW       Anywhere (v6)
80/tcp (v6)                ALLOW       Anywhere (v6)
443/tcp (v6)               ALLOW       Anywhere (v6)
```

### 2.7 Configure Fail2Ban

```bash
# Enable fail2ban
sudo systemctl enable fail2ban
sudo systemctl start fail2ban

# Create local config
sudo nano /etc/fail2ban/jail.local
```

```ini
# /etc/fail2ban/jail.local

[DEFAULT]
bantime = 1h
findtime = 10m
maxretry = 3
backend = systemd

[sshd]
enabled = true
port = 22
filter = sshd
logpath = /var/log/auth.log

[nginx-http-auth]
enabled = true
port = http,https
filter = nginx-http-auth
logpath = /var/log/nginx/error.log

[nginx-botsearch]
enabled = true
port = http,https
filter = nginx-botsearch
logpath = /var/log/nginx/access.log
maxretry = 4
```

```bash
# Restart fail2ban
sudo systemctl restart fail2ban

# Check status
sudo fail2ban-client status
sudo fail2ban-client status sshd
```

### 2.8 Setup Swap (Safety Net untuk RAM 1GB)

```bash
# Check current swap
sudo swapon --show

# Create 1GB swap file
sudo fallocate -l 1G /swapfile

# Set permissions
sudo chmod 600 /swapfile

# Setup swap
sudo mkswap /swapfile
sudo swapon /swapfile

# Make permanent
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab

# Optimize swappiness (use swap only when necessary)
sudo sysctl vm.swappiness=10
echo 'vm.swappiness=10' | sudo tee -a /etc/sysctl.conf

# Verify
free -h
```

**Expected output:**

```
              total        used        free      shared  buff/cache   available
Mem:          987Mi       234Mi       456Mi       12Mi       297Mi       612Mi
Swap:         1.0Gi          0B       1.0Gi
```

### 2.9 Optimize Kernel Parameters

```bash
sudo nano /etc/sysctl.conf
```

```ini
# /etc/sysctl.conf — Web server optimization

# Network
net.core.somaxconn = 1024
net.ipv4.tcp_max_syn_backlog = 1024
net.ipv4.tcp_tw_reuse = 1
net.ipv4.ip_local_port_range = 1024 65535

# Memory
vm.swappiness = 10
vm.dirty_ratio = 15
vm.dirty_background_ratio = 5

# File descriptors
fs.file-max = 65536
```

```bash
# Apply changes
sudo sysctl -p
```

### 2.10 Create Web Directory

```bash
# Create directory structure
sudo mkdir -p /var/www/webifylab
sudo mkdir -p /var/www/webifylab/dist
sudo mkdir -p /var/log/webifylab
sudo mkdir -p /backup/webifylab

# Set ownership
sudo chown -R deploy:deploy /var/www/webifylab
sudo chown -R deploy:deploy /var/log/webifylab
sudo chown -R deploy:deploy /backup/webifylab

# Set permissions
sudo chmod -R 755 /var/www/webifylab
```

---

## 3. Domain & DNS Setup

### 3.1 DNS Records

**Di domain registrar (Niagahoster, Cloudflare, Namecheap, dll):**

| Type     | Name  | Value                                 | TTL  |
| -------- | ----- | ------------------------------------- | ---- |
| **A**    | `@`   | `xxx.xxx.xxx.xxx`                     | 3600 |
| **A**    | `www` | `xxx.xxx.xxx.xxx`                     | 3600 |
| **AAAA** | `@`   | (IPv6 jika ada)                       | 3600 |
| **TXT**  | `@`   | `v=spf1 include:_spf.google.com ~all` | 3600 |

### 3.2 Verify DNS Propagation

```bash
# Check A record
dig webifylab.com A +short

# Check www
dig www.webifylab.com A +short

# Check dari multiple DNS servers
dig @8.8.8.8 webifylab.com A +short
dig @1.1.1.1 webifylab.com A +short
```

**Tunggu hingga DNS propagate (biasanya 5-30 menit, maksimal 24 jam).**

---

## 4. SSL Certificate (Let's Encrypt)

### 4.1 Install Certbot

```bash
# Certbot sudah diinstall di step 2.2
# Verify
certbot --version
```

### 4.2 Obtain SSL Certificate

**PENTING:** Pastikan DNS sudah propagate dan port 80 terbuka sebelum menjalankan ini.

```bash
# Obtain certificate (standalone mode, stop nginx dulu)
sudo systemctl stop nginx

sudo certbot certonly --standalone \
    -d webifylab.com \
    -d www.webifylab.com \
    --email rizal@webifylab.com \
    --agree-tos \
    --no-eff-email

# Start nginx kembali
sudo systemctl start nginx
```

**Atau jika nginx sudah running (webroot mode):**

```bash
sudo certbot --nginx \
    -d webifylab.com \
    -d www.webifylab.com \
    --email rizal@webifylab.com \
    --agree-tos \
    --no-eff-email
```

### 4.3 Verify Certificate

```bash
# Check certificate
sudo certbot certificates

# Check expiry
sudo openssl x509 -dates -noout -in /etc/letsencrypt/live/webifylab.com/fullchain.pem
```

### 4.4 Auto-Renewal

```bash
# Test renewal
sudo certbot renew --dry-run

# Certbot otomatis setup cron/timer
# Verify
sudo systemctl status certbot.timer
```

**Expected:** Certbot akan auto-renew 30 hari sebelum expiry.

---

## 5. Nginx Configuration

### 5.1 Create Site Config

```bash
sudo nano /etc/nginx/sites-available/webifylab
```

```nginx
# /etc/nginx/sites-available/webifylab

# Rate limiting
limit_req_zone $binary_remote_addr zone=general:10m rate=30r/s;
limit_req_zone $binary_remote_addr zone=api:10m rate=10r/s;

# Redirect HTTP to HTTPS
server {
    listen 80;
    listen [::]:80;
    server_name webifylab.com www.webifylab.com;
    return 301 https://$host$request_uri;
}

# Main HTTPS server
server {
    listen 443 ssl http2;
    listen [::]:443 ssl http2;
    server_name webifylab.com www.webifylab.com;

    # ─── SSL Configuration ───
    ssl_certificate /etc/letsencrypt/live/webifylab.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/webifylab.com/privkey.pem;
    ssl_trusted_certificate /etc/letsencrypt/live/webifylab.com/chain.pem;

    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384;
    ssl_prefer_server_ciphers off;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 1d;
    ssl_session_tickets off;

    # OCSP Stapling
    ssl_stapling on;
    ssl_stapling_verify on;
    resolver 8.8.8.8 8.8.4.4 valid=300s;
    resolver_timeout 5s;

    # ─── Security Headers ───
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;
    add_header Permissions-Policy "camera=(), microphone=(), geolocation=()" always;
    add_header Strict-Transport-Security "max-age=63072000; includeSubDomains; preload" always;

    add_header Content-Security-Policy "
        default-src 'self';
        script-src 'self' 'unsafe-inline' https://plausible.io;
        style-src 'self' 'unsafe-inline' https://fonts.googleapis.com;
        img-src 'self' data: https: blob:;
        font-src 'self' https://fonts.gstatic.com;
        connect-src 'self' https://formspree.io https://plausible.io;
        frame-ancestors 'self';
        base-uri 'self';
        form-action 'self' https://formspree.io;
    " always;

    # ─── Root Directory ───
    root /var/www/webifylab/dist;
    index index.html;

    # ─── Logging ───
    access_log /var/log/webifylab/access.log;
    error_log /var/log/webifylab/error.log warn;

    # ─── Gzip Compression ───
    gzip on;
    gzip_vary on;
    gzip_proxied any;
    gzip_comp_level 6;
    gzip_min_length 256;
    gzip_types
        text/plain
        text/css
        text/xml
        text/javascript
        application/json
        application/javascript
        application/x-javascript
        application/xml
        application/xml+rss
        application/vnd.ms-fontobject
        application/x-font-ttf
        font/opentype
        image/svg+xml
        image/x-icon;

    # ─── Static Assets (Long Cache) ───
    location ~* \.(jpg|jpeg|png|gif|ico|webp|avif|svg)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
        access_log off;
    }

    location ~* \.(css|js)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
        access_log off;
    }

    location ~* \.(woff|woff2|ttf|eot|otf)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
        add_header Access-Control-Allow-Origin "*";
        access_log off;
    }

    # ─── HTML (Short Cache) ───
    location ~* \.html$ {
        expires 1h;
        add_header Cache-Control "public, must-revalidate";
    }

    # ─── API Proxy (V1.5 — Golang Backend) ───
    location /api/ {
        limit_req zone=api burst=20 nodelay;

        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Connection "";

        # Timeouts
        proxy_connect_timeout 10s;
        proxy_send_timeout 30s;
        proxy_read_timeout 30s;

        # Buffering
        proxy_buffering on;
        proxy_buffer_size 4k;
        proxy_buffers 8 4k;
    }

    # ─── Health Check ───
    location = /health {
        access_log off;
        return 200 "OK\n";
        add_header Content-Type text/plain;
    }

    # ─── Main Location ───
    location / {
        limit_req zone=general burst=50 nodelay;
        try_files $uri $uri/ $uri.html =404;
    }

    # ─── Deny Hidden Files ───
    location ~ /\. {
        deny all;
        access_log off;
        log_not_found off;
    }

    # ─── Deny Sensitive Files ───
    location ~* \.(env|git|gitignore|htaccess|htpasswd|ini|log|sql|bak)$ {
        deny all;
        access_log off;
        log_not_found off;
    }
}
```

### 5.2 Enable Site

```bash
# Create symlink
sudo ln -s /etc/nginx/sites-available/webifylab /etc/nginx/sites-enabled/

# Remove default site
sudo rm -f /etc/nginx/sites-enabled/default

# Test config
sudo nginx -t
```

**Expected output:**

```
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
```

```bash
# Reload nginx
sudo systemctl reload nginx
sudo systemctl enable nginx
```

### 5.3 Log Rotation

```bash
sudo nano /etc/logrotate.d/webifylab
```

```
/var/log/webifylab/*.log {
    daily
    rotate 14
    compress
    delaycompress
    missingok
    notifempty
    create 0640 deploy deploy
    sharedscripts
    postrotate
        [ -f /var/run/nginx.pid ] && kill -USR1 `cat /var/run/nginx.pid`
    endscript
}
```

---

## 6. Deployment Flow

### 6.1 Monorepo Structure Reminder

```
webifylab/
├── apps/
│   ├── web/          # Astro frontend
│   └── api/          # Golang backend (V1.5)
├── infra/
│   ├── nginx/
│   ├── systemd/
│   └── scripts/
├── docs/
├── .env
├── .gitignore
├── Makefile
└── README.md
```

### 6.2 Manual Deployment (V1 — Recommended)

#### Step 1: Build Astro Locally

```bash
# Di laptop lokal
cd webifylab/apps/web

# Install dependencies
npm install

# Build static site
npm run build

# Verify output
ls -la dist/
du -sh dist/
```

**Expected:** `dist/` folder berisi HTML, CSS, JS, images. Total size < 2MB.

#### Step 2: Upload to VPS

```bash
# Option A: rsync (recommended)
rsync -avz --delete \
    --exclude='.git' \
    --exclude='node_modules' \
    apps/web/dist/ \
    deploy@xxx.xxx.xxx.xxx:/var/www/webifylab/dist/

# Option B: scp
scp -r apps/web/dist/* deploy@xxx.xxx.xxx.xxx:/var/www/webifylab/dist/
```

#### Step 3: Verify on VPS

```bash
# SSH ke VPS
ssh deploy@xxx.xxx.xxx.xxx

# Check files
ls -la /var/www/webifylab/dist/
du -sh /var/www/webifylab/dist/

# Check nginx
sudo nginx -t
curl -I https://webifylab.com
```

**Expected response:**

```
HTTP/2 200
server: nginx
content-type: text/html
content-encoding: gzip
strict-transport-security: max-age=63072000; includeSubDomains; preload
x-frame-options: SAMEORIGIN
x-content-type-options: nosniff
```

### 6.3 Automated Deployment (V1.5 — GitHub Actions)

#### Step 1: Generate SSH Deploy Key

```bash
# Di laptop lokal
ssh-keygen -t ed25519 -f ~/.ssh/webifylab_deploy -N ""

# Copy public key ke VPS
ssh-copy-id -i ~/.ssh/webifylab_deploy.pub deploy@xxx.xxx.xxx.xxx
```

#### Step 2: Add GitHub Secrets

Di GitHub repository → Settings → Secrets and variables → Actions:

| Secret Name    | Value                                |
| -------------- | ------------------------------------ |
| `VPS_HOST`     | `xxx.xxx.xxx.xxx`                    |
| `VPS_USER`     | `deploy`                             |
| `VPS_SSH_KEY`  | (isi dari `~/.ssh/webifylab_deploy`) |
| `VPS_SSH_PORT` | `22`                                 |

#### Step 3: Create GitHub Actions Workflow

```yaml
# .github/workflows/deploy.yml
name: Deploy Webifylab

on:
  push:
    branches: [main]
    paths:
      - "apps/web/**"
      - ".github/workflows/deploy.yml"

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: "20"
          cache: "npm"
          cache-dependency-path: apps/web/package-lock.json

      - name: Install dependencies
        working-directory: ./apps/web
        run: npm ci

      - name: Build Astro
        working-directory: ./apps/web
        run: npm run build
        env:
          PUBLIC_ANALYTICS_ID: ${{ secrets.ANALYTICS_ID }}

      - name: Deploy to VPS
        uses: appleboy/scp-action@v0.1.7
        with:
          host: ${{ secrets.VPS_HOST }}
          username: ${{ secrets.VPS_USER }}
          key: ${{ secrets.VPS_SSH_KEY }}
          port: ${{ secrets.VPS_SSH_PORT }}
          source: "apps/web/dist/*"
          target: "/var/www/webifylab/dist"
          strip_components: 3
          overwrite: true

      - name: Verify deployment
        uses: appleboy/ssh-action@v1.0.3
        with:
          host: ${{ secrets.VPS_HOST }}
          username: ${{ secrets.VPS_USER }}
          key: ${{ secrets.VPS_SSH_KEY }}
          port: ${{ secrets.VPS_SSH_PORT }}
          script: |
            echo "=== Deployment Verification ==="
            echo "Files:"
            ls -la /var/www/webifylab/dist/ | head -10
            echo ""
            echo "Size:"
            du -sh /var/www/webifylab/dist/
            echo ""
            echo "Nginx status:"
            sudo nginx -t
            echo ""
            echo "Health check:"
            curl -s -o /dev/null -w "%{http_code}" https://webifylab.com/health
            echo ""
            echo "✅ Deployment complete!"
```

#### Step 4: Test CI/CD

```bash
# Di laptop lokal
git add .
git commit -m "feat: setup CI/CD pipeline"
git push origin main

# Monitor di GitHub Actions tab
```

---

## 7. Golang Backend Deployment (V1.5)

### 7.1 Install Go on VPS

```bash
# SSH ke VPS
ssh deploy@xxx.xxx.xxx.xxx

# Download Go
wget https://go.dev/dl/go1.22.5.linux-amd64.tar.gz

# Extract
sudo tar -C /usr/local -xzf go1.22.5.linux-amd64.tar.gz

# Add to PATH
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc

# Verify
go version
```

### 7.2 Build Go Binary

**Di laptop lokal (cross-compile untuk Linux):**

```bash
cd webifylab/apps/api

# Build untuk Linux AMD64
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build \
    -ldflags="-s -w" \
    -o webifylab-api \
    ./cmd/server/main.go

# Check binary size
ls -lh webifylab-api
```

**Expected:** Binary size ~10-20MB.

### 7.3 Upload Binary to VPS

```bash
# Upload binary
scp apps/api/webifylab-api deploy@xxx.xxx.xxx.xxx:/var/www/webifylab/

# Upload config
scp apps/api/.env deploy@xxx.xxx.xxx.xxx:/var/www/webifylab/.env
```

### 7.4 Create Systemd Service

```bash
# Di VPS
sudo nano /etc/systemd/system/webifylab-api.service
```

```ini
# /etc/systemd/system/webifylab-api.service

[Unit]
Description=Webifylab API Server
Documentation=https://webifylab.com
After=network.target postgresql.service
Wants=postgresql.service

[Service]
Type=simple
User=deploy
Group=deploy
WorkingDirectory=/var/www/webifylab
ExecStart=/var/www/webifylab/webifylab-api
Restart=on-failure
RestartSec=5s
StandardOutput=journal
StandardError=journal
SyslogIdentifier=webifylab-api

# Environment
EnvironmentFile=/var/www/webifylab/.env

# Security hardening
NoNewPrivileges=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/var/www/webifylab /var/log/webifylab
PrivateTmp=true

# Resource limits
MemoryMax=200M
CPUQuota=50%

[Install]
WantedBy=multi-user.target
```

### 7.5 Start API Service

```bash
# Reload systemd
sudo systemctl daemon-reload

# Enable & start
sudo systemctl enable webifylab-api
sudo systemctl start webifylab-api

# Check status
sudo systemctl status webifylab-api

# Check logs
sudo journalctl -u webifylab-api -f
```

### 7.6 Verify API

```bash
# Health check
curl http://127.0.0.1:8080/api/health

# Via Nginx (HTTPS)
curl https://webifylab.com/api/health
```

**Expected response:**

```json
{ "status": "ok", "version": "1.0.0", "timestamp": "2026-09-17T10:00:00Z" }
```

---

## 8. Monitoring & Alerting

### 8.1 Uptime Monitoring

**UptimeRobot (Free):**

- [ ] Daftar di https://uptimerobot.com
- [ ] Add monitor: `https://webifylab.com`
- [ ] Interval: 5 menit
- [ ] Alert: Email + Telegram

**Health Check Endpoint:**

```bash
# Test dari luar
curl -I https://webifylab.com/health
```

### 8.2 Server Monitoring

```bash
# Real-time resource usage
htop

# Memory usage
free -h

# Disk usage
df -h

# Network connections
ss -tulnp

# Nginx status
sudo systemctl status nginx

# API status (V1.5)
sudo systemctl status webifylab-api

# Recent errors
tail -50 /var/log/webifylab/error.log

# Nginx access log (real-time)
tail -f /var/log/webifylab/access.log
```

### 8.3 Monitoring Script

```bash
#!/bin/bash
# /opt/scripts/monitor.sh — Run via cron every 5 minutes

DOMAIN="webifylab.com"
LOG_FILE="/var/log/webifylab/monitor.log"
ALERT_THRESHOLD_RAM=80
ALERT_THRESHOLD_DISK=85

# Check HTTP status
HTTP_STATUS=$(curl -s -o /dev/null -w "%{http_code}" https://$DOMAIN/health)

if [ "$HTTP_STATUS" != "200" ]; then
    echo "$(date): ⚠️ ALERT — $DOMAIN returned HTTP $HTTP_STATUS" >> $LOG_FILE
    # Send Telegram alert (V1.5)
    # curl -s "https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/sendMessage?chat_id=$TELEGRAM_CHAT_ID&text=⚠️ Webifylab DOWN (HTTP $HTTP_STATUS)"
fi

# Check RAM usage
RAM_USAGE=$(free | awk '/Mem/{printf("%.0f"), $3/$2*100}')
if [ "$RAM_USAGE" -gt "$ALERT_THRESHOLD_RAM" ]; then
    echo "$(date): ⚠️ ALERT — RAM usage at ${RAM_USAGE}%" >> $LOG_FILE
fi

# Check disk usage
DISK_USAGE=$(df / | awk 'NR==2{print $5}' | sed 's/%//')
if [ "$DISK_USAGE" -gt "$ALERT_THRESHOLD_DISK" ]; then
    echo "$(date): ⚠️ ALERT — Disk usage at ${DISK_USAGE}%" >> $LOG_FILE
fi

echo "$(date): ✅ OK — HTTP $HTTP_STATUS, RAM ${RAM_USAGE}%, Disk ${DISK_USAGE}%" >> $LOG_FILE
```

```bash
# Make executable
chmod +x /opt/scripts/monitor.sh

# Add to cron (every 5 minutes)
crontab -e
*/5 * * * * /opt/scripts/monitor.sh
```

### 8.4 Log Analysis

```bash
# Top 10 most visited pages
awk '{print $7}' /var/log/webifylab/access.log | sort | uniq -c | sort -rn | head -10

# Top 10 IP addresses
awk '{print $1}' /var/log/webifylab/access.log | sort | uniq -c | sort -rn | head -10

# HTTP status code distribution
awk '{print $9}' /var/log/webifylab/access.log | sort | uniq -c | sort -rn

# 404 errors
grep " 404 " /var/log/webifylab/access.log | tail -20

# Requests per hour
awk '{print $4}' /var/log/webifylab/access.log | cut -d: -f2 | sort | uniq -c | sort -rn
```

---

## 9. Backup & Disaster Recovery

### 9.1 Backup Strategy

| What                 | Frequency | Retention     | Method                |
| -------------------- | --------- | ------------- | --------------------- |
| **Website files**    | Daily     | 7 days        | tar + local           |
| **Nginx config**     | On change | Git versioned | Git                   |
| **SSL certificates** | Auto      | Let's Encrypt | Certbot               |
| **Database (V1.5)**  | Daily     | 30 days       | pg_dump               |
| **Logs**             | Daily     | 14 days       | Logrotate             |
| **Full VPS**         | Weekly    | 4 weeks       | VPS provider snapshot |

### 9.2 Backup Script

```bash
#!/bin/bash
# /opt/scripts/backup.sh

BACKUP_DIR="/backup/webifylab"
DATE=$(date +%Y%m%d_%H%M%S)
RETENTION_DAYS=7

mkdir -p $BACKUP_DIR

echo "🔄 Starting backup at $(date)"

# 1. Backup website files
echo "📦 Backing up website files..."
tar -czf "$BACKUP_DIR/web-$DATE.tar.gz" \
    -C /var/www/webifylab dist/

# 2. Backup nginx config
echo "📦 Backing up nginx config..."
tar -czf "$BACKUP_DIR/nginx-$DATE.tar.gz" \
    -C /etc/nginx sites-available/ sites-enabled/ nginx.conf

# 3. Backup systemd services
echo "📦 Backing up systemd services..."
tar -czf "$BACKUP_DIR/systemd-$DATE.tar.gz" \
    -C /etc/systemd/system webifylab-api.service 2>/dev/null

# 4. Backup database (V1.5)
# echo "📦 Backing up database..."
# pg_dump -U webifylab -h localhost -F c -b -v \
#     -f "$BACKUP_DIR/db-$DATE.backup" webifylab

# 5. Delete old backups
echo "🗑️ Cleaning old backups..."
find $BACKUP_DIR -type f -mtime +$RETENTION_DAYS -delete

# 6. Summary
echo "✅ Backup completed at $(date)"
echo "📊 Backup size:"
du -sh "$BACKUP_DIR"/*-$DATE* 2>/dev/null
echo "📊 Total backup directory:"
du -sh $BACKUP_DIR
```

```bash
# Make executable
chmod +x /opt/scripts/backup.sh

# Add to cron (daily at 2 AM)
crontab -e
0 2 * * * /opt/scripts/backup.sh >> /var/log/webifylab/backup.log 2>&1
```

### 9.3 Restore Procedure

#### Restore Website Files

```bash
# List available backups
ls -la /backup/webifylab/web-*.tar.gz

# Restore specific backup
tar -xzf /backup/webifylab/web-20260917_020000.tar.gz \
    -C /var/www/webifylab/

# Verify
ls -la /var/www/webifylab/dist/
curl -I https://webifylab.com
```

#### Restore Nginx Config

```bash
tar -xzf /backup/webifylab/nginx-20260917_020000.tar.gz \
    -C /etc/nginx/

sudo nginx -t
sudo systemctl reload nginx
```

#### Restore Database (V1.5)

```bash
pg_restore -U webifylab -h localhost -d webifylab -v \
    /backup/webifylab/db-20260917_020000.backup
```

### 9.4 Disaster Recovery Plan

| Scenario                 | Steps                                  | RTO    | RPO |
| ------------------------ | -------------------------------------- | ------ | --- |
| **Website down (Nginx)** | `sudo systemctl restart nginx`         | 1 min  | 0   |
| **API down (V1.5)**      | `sudo systemctl restart webifylab-api` | 1 min  | 0   |
| **SSL expired**          | `sudo certbot renew`                   | 5 min  | 0   |
| **Disk full**            | Clean logs, expand disk                | 30 min | 0   |
| **RAM OOM**              | Restart services, check swap           | 5 min  | 0   |
| **Corrupted files**      | Restore from backup                    | 15 min | 24h |
| **VPS total failure**    | Provision new VPS, restore backups     | 1-2h   | 24h |
| **DDoS attack**          | Enable Cloudflare, rate limit          | 30 min | 0   |

---

## 10. Troubleshooting Guide

### 10.1 Common Issues

#### Issue: Website returns 502 Bad Gateway

```bash
# Check nginx error log
sudo tail -50 /var/log/webifylab/error.log

# Check if backend is running (V1.5)
sudo systemctl status webifylab-api

# Check if port 8080 is listening
ss -tulnp | grep 8080

# Restart services
sudo systemctl restart webifylab-api
sudo systemctl reload nginx
```

#### Issue: SSL Certificate Error

```bash
# Check certificate status
sudo certbot certificates

# Force renewal
sudo certbot renew --force-renewal

# Check nginx SSL config
sudo nginx -t

# Restart nginx
sudo systemctl restart nginx
```

#### Issue: 403 Forbidden

```bash
# Check file permissions
ls -la /var/www/webifylab/dist/

# Fix permissions
sudo chown -R deploy:deploy /var/www/webifylab/
sudo chmod -R 755 /var/www/webifylab/

# Check nginx config
sudo nginx -t
```

#### Issue: 404 Not Found

```bash
# Check if files exist
ls -la /var/www/webifylab/dist/

# Check nginx root directive
grep "root" /etc/nginx/sites-available/webifylab

# Check try_files directive
grep "try_files" /etc/nginx/sites-available/webifylab
```

#### Issue: Slow Page Load

```bash
# Check server resources
htop
free -h
df -h

# Check nginx access log for slow requests
awk '{print $10, $7}' /var/log/webifylab/access.log | sort -rn | head -20

# Test page speed
curl -o /dev/null -s -w "Time: %{time_total}s\nSize: %{size_download} bytes\n" https://webifylab.com

# Check gzip
curl -H "Accept-Encoding: gzip" -I https://webifylab.com | grep -i content-encoding
```

#### Issue: RAM Running Out

```bash
# Check memory usage
free -h

# Check top processes
ps aux --sort=-%mem | head -10

# Check swap usage
swapon --show

# Clear cache (safe)
sudo sync && echo 3 | sudo tee /proc/sys/vm/drop_caches

# Restart heavy services
sudo systemctl restart nginx
sudo systemctl restart webifylab-api  # V1.5
```

#### Issue: Disk Full

```bash
# Check disk usage
df -h

# Find large files
sudo du -sh /var/log/* | sort -rh | head -10
sudo du -sh /backup/* | sort -rh | head -10

# Clean old logs
sudo journalctl --vacuum-time=7d
sudo find /var/log -name "*.gz" -mtime +14 -delete

# Clean old backups
find /backup/webifylab -type f -mtime +7 -delete

# Clean apt cache
sudo apt clean
```

### 10.2 Emergency Contacts

| Service              | Contact                           | Notes                         |
| -------------------- | --------------------------------- | ----------------------------- |
| **VPS Provider**     | (isi sesuai provider)             | Untuk hardware/network issues |
| **Domain Registrar** | (isi sesuai registrar)            | Untuk DNS issues              |
| **Let's Encrypt**    | https://community.letsencrypt.org | Untuk SSL issues              |

---

## 11. Maintenance Checklist

### 11.1 Daily (Automated)

- [ ] Backup script runs at 2 AM
- [ ] Monitoring script runs every 5 minutes
- [ ] Log rotation runs
- [ ] SSL auto-renewal check

### 11.2 Weekly (Manual)

- [ ] Check server resources (`htop`, `free -h`, `df -h`)
- [ ] Review error logs (`/var/log/webifylab/error.log`)
- [ ] Check backup integrity
- [ ] Review monitoring alerts
- [ ] Check Google Search Console for errors

### 11.3 Monthly (Manual)

- [ ] Update system packages (`sudo apt update && sudo apt upgrade`)
- [ ] Review SSL certificate expiry (`sudo certbot certificates`)
- [ ] Audit user access (`cat /etc/passwd | grep bash`)
- [ ] Review firewall rules (`sudo ufw status`)
- [ ] Test disaster recovery (restore from backup)
- [ ] Review performance metrics (Lighthouse, PageSpeed)
- [ ] Clean old backups and logs

### 11.4 Quarterly (Manual)

- [ ] Security audit (check for vulnerabilities)
- [ ] Review and update dependencies (`npm audit`, `go list -m -u all`)
- [ ] Review SEO performance
- [ ] Review backup strategy
- [ ] Update documentation

---

## 12. Quick Reference Commands

### 12.1 Service Management

```bash
# Nginx
sudo systemctl start nginx
sudo systemctl stop nginx
sudo systemctl restart nginx
sudo systemctl reload nginx
sudo systemctl status nginx
sudo nginx -t

# Golang API (V1.5)
sudo systemctl start webifylab-api
sudo systemctl stop webifylab-api
sudo systemctl restart webifylab-api
sudo systemctl status webifylab-api

# PostgreSQL (V1.5)
sudo systemctl start postgresql
sudo systemctl stop postgresql
sudo systemctl restart postgresql
sudo systemctl status postgresql
```

### 12.2 Logs

```bash
# Nginx access log
tail -f /var/log/webifylab/access.log

# Nginx error log
tail -f /var/log/webifylab/error.log

# API logs (V1.5)
sudo journalctl -u webifylab-api -f

# System logs
sudo journalctl -xe

# Auth logs (SSH)
sudo tail -f /var/log/auth.log
```

### 12.3 Network

```bash
# Check open ports
ss -tulnp

# Check connections
ss -s

# Test HTTP
curl -I https://webifylab.com

# Test API (V1.5)
curl https://webifylab.com/api/health

# DNS lookup
dig webifylab.com

# Ping
ping webifylab.com
```

### 12.4 Resources

```bash
# CPU & RAM
htop
free -h

# Disk
df -h
du -sh /var/www/webifylab/

# Processes
ps aux --sort=-%mem | head -10
ps aux --sort=-%cpu | head -10
```

---

## 13. Deployment Checklist (Pre-Launch)

### 13.1 VPS Setup

- [ ] VPS provisioned (Ubuntu 22.04 / Debian 12)
- [ ] System updated (`apt update && apt upgrade`)
- [ ] Deploy user created
- [ ] SSH key authentication configured
- [ ] Root login disabled
- [ ] Password authentication disabled
- [ ] UFW firewall configured (22, 80, 443)
- [ ] Fail2ban running
- [ ] Swap file created (1GB)
- [ ] Kernel parameters optimized
- [ ] Web directory created

### 13.2 Domain & SSL

- [ ] DNS A records configured
- [ ] DNS propagated (verified with `dig`)
- [ ] SSL certificate obtained (Let's Encrypt)
- [ ] SSL auto-renewal configured
- [ ] HTTPS redirect working

### 13.3 Nginx

- [ ] Site config created
- [ ] SSL configured
- [ ] Security headers configured
- [ ] Gzip compression enabled
- [ ] Caching configured
- [ ] Rate limiting configured
- [ ] Log rotation configured
- [ ] Config tested (`nginx -t`)
- [ ] Nginx running

### 13.4 Website

- [ ] Astro build successful
- [ ] Static files uploaded to VPS
- [ ] Website accessible via HTTPS
- [ ] All pages loading correctly
- [ ] Images loading correctly
- [ ] Form submission working
- [ ] WhatsApp CTA working
- [ ] Mobile responsive verified

### 13.5 Monitoring

- [ ] UptimeRobot configured
- [ ] Monitoring script running
- [ ] Alert notifications working
- [ ] Backup script running
- [ ] Cron jobs verified

### 13.6 Security

- [ ] SSH hardened
- [ ] Firewall active
- [ ] Fail2ban running
- [ ] SSL/TLS configured (TLS 1.2+)
- [ ] Security headers present
- [ ] CSP configured
- [ ] HSTS enabled
- [ ] Hidden files blocked

---

## 14. Open Questions

| No  | Pertanyaan                                                                | Status  |
| --- | ------------------------------------------------------------------------- | ------- |
| Q1  | VPS provider apa yang digunakan? (DigitalOcean, Vultr, IDCloudHost, dll)  | Pending |
| Q2  | Domain registrar apa yang digunakan? (Niagahoster, Cloudflare, Namecheap) | Pending |
| Q3  | Apakah ingin menggunakan Cloudflare sebagai CDN/proxy?                    | Pending |
| Q4  | Apakah ada preferensi untuk notification channel (Email/Telegram)?        | Pending |
| Q5  | Apakah VPS sudah di-provision atau masih perlu dibeli?                    | Pending |

---

_Dokumen ini adalah living document. Versi akan diperbarui seiring perubahan infrastruktur._

**Last Updated:** 17 September 2026
**Next Step:** Setup VPS, configure domain, deploy landing page.
