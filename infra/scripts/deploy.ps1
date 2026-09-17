# ========================================
# Webifylab Deploy Script (PowerShell)
# ========================================

param(
    [string]$VPSHost = $(Read-Host "Enter VPS IP"),
    [string]$VPSUser = "deploy",
    [string]$TargetPath = "/var/www/webifylab/dist"
)

Write-Host "🚀 Starting deployment..." -ForegroundColor Cyan

# Build Astro
Write-Host "📦 Building Astro..." -ForegroundColor Yellow
Set-Location apps/web
npm run build
Set-Location ../..

# Upload via rsync (requires rsync installed, or use scp)
Write-Host "📤 Uploading to VPS..." -ForegroundColor Yellow
$distPath = "apps/web/dist/"

# Option 1: Using scp (simpler, Windows native)
scp -r $distPath* ${VPSUser}@${VPSHost}:${TargetPath}

# Option 2: Using rsync (if installed via Git Bash/WSL)
# rsync -avz --delete $distPath ${VPSUser}@${VPSHost}:${TargetPath}

Write-Host "✅ Deployment complete!" -ForegroundColor Green
