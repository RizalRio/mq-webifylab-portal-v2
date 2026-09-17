# 🧪 Testing Strategy Document

## Webifylab Landing Page — Version 1.0

| Metadata         | Detail                                      |
|------------------|---------------------------------------------|
| **Product**      | Webifylab Landing Page                      |
| **Version**      | 1.0                                         |
| **Author**       | Rizal                                       |
| **Created**      | 17 September 2026                           |
| **Status**       | Draft                                       |
| **Scope**        | Frontend (Astro) + Backend API (V1.5)       |

---

## 1. Testing Strategy Overview

### 1.1 Testing Pyramid

```
         ╱╲
        ╱  ╲        E2E Tests (Playwright)
       ╱ E2E╲       — 5-10 critical flows
      ╱──────╲      — Slow, expensive
     ╱        ╲
    ╱Integration╲   Integration Tests (Vitest)
   ╱────────────╲   — API endpoints (V1.5)
  ╱              ╲  — Database queries
 ╱   Unit Tests   ╲ Unit Tests (Vitest)
╱──────────────────╲— Utility functions
                    — Component logic
                    — Fast, cheap
```

### 1.2 Testing Priorities (V1)

| Priority | Testing Type | Tools | Effort | Impact |
|----------|-------------|-------|--------|--------|
| 🔴 **Critical** | Manual Functional | Browser DevTools | Low | High |
| 🔴 **Critical** | Performance | Lighthouse | Low | High |
| 🔴 **Critical** | SEO Validation | Lighthouse + Manual | Low | High |
| 🟡 **High** | Cross-Browser | Manual + BrowserStack | Medium | High |
| 🟡 **High** | Accessibility | axe DevTools + Pa11y | Low | Medium |
| 🟡 **High** | Responsive | Manual + Responsinator | Low | High |
| 🟢 **Medium** | Automated E2E | Playwright | High | Medium |
| 🟢 **Medium** | Visual Regression | Playwright screenshots | High | Medium |
| 🔵 **Low** | Unit Tests | Vitest | Medium | Low (V1) |
| 🔵 **Low** | Security Scan | npm audit + manual | Low | Medium |

---

## 2. Manual Testing Checklist (Paling Praktis)

### 2.1 Functional Testing

#### ✅ Navigation & Links
```powershell
# Checklist manual
[ ] Navbar sticky saat scroll
[ ] Semua anchor link berfungsi (smooth scroll)
  - #layanan → scroll ke Services
  - #pendekatan → scroll ke Approach
  - #portfolio → scroll ke Portfolio
  - #kontak → scroll ke Contact
[ ] Hamburger menu open/close di mobile
[ ] Menu mobile close setelah klik link
[ ] Logo clickable → scroll ke top
[ ] CTA button "Konsultasi Gratis" berfungsi
```

#### ✅ Form Kontak (S9)
```powershell
[ ] Semua field required berfungsi
[ ] Validasi email (format harus valid)
[ ] Validasi pesan (min 10 karakter)
[ ] Dropdown "Jenis Layanan" berfungsi
[ ] Submit button menampilkan loading state
[ ] Success message muncul setelah submit
[ ] Error message muncul jika gagal
[ ] Form data terkirim ke Formspree/Web3Forms
[ ] Notifikasi masuk ke email/Telegram
```

#### ✅ WhatsApp CTA
```powershell
[ ] Tombol WhatsApp clickable
[ ] Membuka WhatsApp Web/App
[ ] Pre-filled message terisi dengan benar:
    "Halo Webifylab, saya tertarik untuk konsultasi mengenai: ..."
[ ] Nomor WhatsApp valid
```

#### ✅ Content Display
```powershell
[ ] Hero headline & sub-headline tampil
[ ] 3 Problem cards tampil dengan icon
[ ] 3 Services cards tampil dengan features list
[ ] 4 Why Webifylab items tampil
[ ] 4 Approach steps tampil
[ ] 3 Portfolio placeholder cards tampil
[ ] FAQ accordion expand/collapse
[ ] Footer multi-column tampil
[ ] Copyright tahun dinamis (2026)
```

---

### 2.2 Responsive Testing

#### 📱 Breakpoints yang Harus Dites

| Device | Resolution | Type |
|--------|-----------|------|
| **iPhone SE** | 375 x 667 | Mobile Small |
| **iPhone 14** | 390 x 844 | Mobile |
| **iPad Mini** | 768 x 1024 | Tablet |
| **iPad Pro** | 1024 x 1366 | Tablet Large |
| **MacBook Air** | 1440 x 900 | Laptop |
| **Full HD** | 1920 x 1080 | Desktop |
| **4K** | 2560 x 1440 | Large Desktop |

#### 🧪 Cara Test di Chrome DevTools

```powershell
# Buka DevTools (F12)
# Klik icon "Toggle device toolbar" (Ctrl+Shift+M)
# Pilih device dari dropdown atau custom resolution

# Test checklist per breakpoint:
[ ] Tidak ada horizontal scroll
[ ] Text tidak terpotong
[ ] Gambar tidak overflow
[ ] Grid layout adaptif (1 → 2 → 3 kolom)
[ ] Font size readable (min 14px)
[ ] Touch target min 44x44px (mobile)
[ ] Padding & spacing proporsional
[ ] Hero section layout benar (stacked di mobile)
```

#### 🌐 Tools Online untuk Responsive Testing

```
1. Responsinator: https://www.responsinator.com/
   → Masukkan URL, lihat preview di multiple devices

2. BrowserStack: https://www.browserstack.com/
   → Test di real devices (free trial 30 menit)

3. LambdaTest: https://www.lambdatest.com/
   → Alternative BrowserStack (free tier available)
```

---

### 2.3 Cross-Browser Testing

#### 🌍 Browser yang Harus Dites

| Browser | Desktop | Mobile | Priority |
|---------|---------|--------|----------|
| **Chrome** | ✅ | ✅ | 🔴 Critical |
| **Safari** | ✅ | ✅ (iOS) | 🔴 Critical |
| **Firefox** | ✅ | ✅ | 🟡 High |
| **Edge** | ✅ | ✅ | 🟡 High |
| **Samsung Internet** | - | ✅ | 🟢 Medium |
| **Opera** | ✅ | ✅ | 🔵 Low |

#### 🧪 Testing Matrix

```powershell
# Test checklist per browser:
[ ] Layout tidak broken
[ ] Font rendering benar (Inter)
[ ] Animasi smooth
[ ] Form submission berfungsi
[ ] Images load dengan benar
[ ] Video/embed berfungsi (jika ada)
[ ] Console tidak ada error
```

---

## 3. Performance Testing (Lighthouse)

### 3.1 Cara Menjalankan Lighthouse

#### 🖥️ Method 1: Chrome DevTools (Paling Mudah)

```powershell
# 1. Buka website di Chrome
# 2. Tekan F12 → tab "Lighthouse"
# 3. Pilih kategori:
#    ☑ Performance
#    ☑ Accessibility
#    ☑ Best Practices
#    ☑ SEO
# 4. Device: Mobile (lebih ketat)
# 5. Klik "Analyze page load"
# 6. Tunggu hasil (30-60 detik)
```

#### 💻 Method 2: CLI (Untuk CI/CD)

```powershell
# Install Lighthouse CLI
npm install -g lighthouse

# Run audit
lighthouse https://webifylab.com --view

# Save report as JSON
lighthouse https://webifylab.com --output=json --output-path=./report.json

# Mobile emulation
lighthouse https://webifylab.com --preset=perf
```

#### 🌐 Method 3: PageSpeed Insights (Online)

```
URL: https://pagespeed.web.dev/
→ Masukkan URL website
→ Lihat skor Performance, Accessibility, Best Practices, SEO
→ Lihat Core Web Vitals (LCP, FID, CLS)
```

### 3.2 Performance Budget (Target)

| Metric | Target | Tool |
|--------|--------|------|
| **Performance Score** | > 95 | Lighthouse |
| **Accessibility Score** | > 90 | Lighthouse |
| **Best Practices Score** | > 95 | Lighthouse |
| **SEO Score** | > 95 | Lighthouse |
| **LCP** (Largest Contentful Paint) | < 2.5s | Lighthouse |
| **FID** (First Input Delay) | < 100ms | Lighthouse |
| **CLS** (Cumulative Layout Shift) | < 0.1 | Lighthouse |
| **TTFB** (Time to First Byte) | < 800ms | WebPageTest |
| **Total Page Size** | < 500KB | `du -sh dist/` |
| **Total Requests** | < 50 | Chrome DevTools |

### 3.3 Cara Optimasi Jika Skor Rendah

#### 🐢 Jika Performance < 90

```powershell
# Cek di Lighthouse "Opportunities" section
# Biasanya masalah:
[ ] Gambar tidak di-optimize → Convert ke WebP/AVIF
[ ] CSS/JS tidak di-minify → Pastikan build production
[ ] Render-blocking resources → Defer non-critical CSS/JS
[ ] Unused JavaScript → Tree-shaking Tailwind
[ ] Server response time lambat → Enable gzip/brotli
```

#### ♿ Jika Accessibility < 90

```powershell
# Cek di Lighthouse "Accessibility" section
# Biasanya masalah:
[ ] Image tidak ada alt text → Tambahkan alt
[ ] Color contrast rendah → Ubah warna
[ ] Form input tidak ada label → Tambahkan <label>
[ ] Heading tidak hierarkis → Perbaiki H1-H6
[ ] Link tidak deskriptif → Ubah "klik di sini" → "lihat layanan"
```

---

## 4. SEO Testing

### 4.1 SEO Validation Checklist

#### 🔍 On-Page SEO

```powershell
[ ] Title tag (50-60 karakter, mengandung keyword)
[ ] Meta description (150-160 karakter)
[ ] H1 hanya 1 per halaman
[ ] Heading hierarchy (H1 → H2 → H3)
[ ] URL structure (lowercase, hyphens)
[ ] Canonical URL terpasang
[ ] Open Graph tags (og:title, og:description, og:image)
[ ] Twitter Card tags
[ ] Favicon terpasang
[ ] Robots.txt accessible
[ ] Sitemap.xml ter-generate
```

#### 🧪 Cara Test SEO

```powershell
# 1. Lighthouse SEO Audit
#    → F12 → Lighthouse → SEO → Analyze

# 2. Cek Meta Tags Manual
#    → View Page Source (Ctrl+U)
#    → Cari <title>, <meta name="description">, <meta property="og:...">

# 3. Cek Sitemap
#    → Buka https://webifylab.com/sitemap-index.xml
#    → Pastikan semua halaman ter-list

# 4. Cek Robots.txt
#    → Buka https://webifylab.com/robots.txt
#    → Pastikan tidak ada disallow yang salah

# 5. Google Rich Results Test
#    → https://search.google.com/test/rich-results
#    → Masukkan URL, cek structured data

# 6. Schema Validator
#    → https://validator.schema.org/
#    → Validasi JSON-LD schema
```

### 4.2 SEO Tools

| Tool | Purpose | URL |
|------|---------|-----|
| **Google Search Console** | Monitor indexing & keywords | https://search.google.com/search-console |
| **Google PageSpeed Insights** | Performance & SEO audit | https://pagespeed.web.dev/ |
| **Rich Results Test** | Validate structured data | https://search.google.com/test/rich-results |
| **Mobile-Friendly Test** | Check mobile optimization | https://search.google.com/test/mobile-friendly |
| **Screaming Frog** | Crawl website (free up to 500 URLs) | https://www.screamingfrog.co.uk/ |
| **Ahrefs Webmaster Tools** | Backlink & keyword monitoring | https://ahrefs.com/webmaster-tools |

---

## 5. Accessibility Testing

### 5.1 WCAG 2.1 AA Checklist

```powershell
# Perceivable
[ ] Semua gambar punya alt text
[ ] Video punya captions
[ ] Color contrast ratio ≥ 4.5:1 (normal text)
[ ] Color contrast ratio ≥ 3:1 (large text)
[ ] Text bisa di-resize sampai 200%
[ ] Content tidak bergantung pada color alone

# Operable
[ ] Semua fungsi bisa diakses via keyboard
[ ] Tidak ada keyboard traps
[ ] Focus indicator visible
[ ] Skip to content link tersedia
[ ] Page has descriptive title
[ ] Link text is descriptive

# Understandable
[ ] Language of page is identified
[ ] Form inputs have labels
[ ] Error messages are clear
[ ] Consistent navigation

# Robust
[ ] Valid HTML (no errors)
[ ] Status messages programmatically determined
```

### 5.2 Accessibility Testing Tools

#### 🛠️ Tools yang Direkomendasikan

```powershell
# 1. axe DevTools (Chrome Extension)
#    → Install dari Chrome Web Store
#    → Klik icon axe → "Scan all of page"
#    → Lihat violations & recommendations

# 2. WAVE (Chrome Extension)
#    → https://wave.webaim.org/extension/
#    → Visual feedback untuk accessibility issues

# 3. Lighthouse Accessibility Audit
#    → F12 → Lighthouse → Accessibility

# 4. Pa11y (CLI)
#    → npm install -g pa11y
#    → pa11y https://webifylab.com

# 5. Color Contrast Checker
#    → https://webaim.org/resources/contrastchecker/
#    → Cek contrast ratio antara text & background
```

### 5.3 Keyboard Navigation Test

```powershell
# Test manual dengan keyboard saja (tanpa mouse):
[ ] Tab → Navigate forward
[ ] Shift+Tab → Navigate backward
[ ] Enter → Activate link/button
[ ] Space → Activate button
[ ] Esc → Close modal/menu
[ ] Arrow keys → Navigate within components

# Checklist:
[ ] Focus order logical (top → bottom, left → right)
[ ] Focus indicator visible (outline/ring)
[ ] No keyboard traps (can't get stuck)
[ ] Skip to content link works
[ ] Modal traps focus correctly
```

---

## 6. Automated E2E Testing (Playwright)

### 6.1 Setup Playwright

```powershell
# Install Playwright
cd apps/web
npm install -D @playwright/test

# Initialize Playwright
npx playwright install

# Create config file
npx playwright init
```

### 6.2 Contoh Test Cases

#### 📝 `tests/landing-page.spec.ts`

```typescript
import { test, expect } from '@playwright/test';

test.describe('Landing Page - Critical Flows', () => {
  
  test('homepage loads successfully', async ({ page }) => {
    await page.goto('http://localhost:4321');
    
    // Check page title
    await expect(page).toHaveTitle(/Webifylab/);
    
    // Check hero section visible
    await expect(page.locator('h1')).toContainText('Dari Ide Menjadi Ekosistem Digital');
  });

  test('navigation links work', async ({ page }) => {
    await page.goto('http://localhost:4321');
    
    // Click "Layanan" link
    await page.click('a[href="#layanan"]');
    await expect(page.locator('#layanan')).toBeVisible();
    
    // Click "Kontak" link
    await page.click('a[href="#kontak"]');
    await expect(page.locator('#kontak')).toBeVisible();
  });

  test('contact form submission', async ({ page }) => {
    await page.goto('http://localhost:4321#kontak');
    
    // Fill form
    await page.fill('input[name="name"]', 'Test User');
    await page.fill('input[name="email"]', 'test@example.com');
    await page.selectOption('select[name="service"]', 'app_development');
    await page.fill('textarea[name="message"]', 'This is a test message for validation.');
    
    // Submit form
    await page.click('button[type="submit"]');
    
    // Check success message
    await expect(page.locator('.success-message')).toBeVisible();
  });

  test('mobile menu toggle', async ({ page }) => {
    // Set mobile viewport
    await page.setViewportSize({ width: 375, height: 667 });
    await page.goto('http://localhost:4321');
    
    // Click hamburger menu
    await page.click('[data-testid="hamburger-menu"]');
    
    // Check menu is visible
    await expect(page.locator('[data-testid="mobile-menu"]')).toBeVisible();
    
    // Click a link
    await page.click('[data-testid="mobile-menu"] a:first-child');
    
    // Check menu closes
    await expect(page.locator('[data-testid="mobile-menu"]')).not.toBeVisible();
  });

  test('WhatsApp CTA works', async ({ page }) => {
    await page.goto('http://localhost:4321');
    
    // Click WhatsApp button
    const whatsappLink = page.locator('a[href*="wa.me"]');
    await expect(whatsappLink).toBeVisible();
    
    // Check href contains phone number
    const href = await whatsappLink.getAttribute('href');
    expect(href).toContain('wa.me');
  });

  test('all sections are visible', async ({ page }) => {
    await page.goto('http://localhost:4321');
    
    const sections = [
      '#hero',
      '#problem',
      '#services',
      '#why-webifylab',
      '#approach',
      '#portfolio',
      '#contact',
      '#footer'
    ];
    
    for (const section of sections) {
      await expect(page.locator(section)).toBeVisible();
    }
  });
});
```

### 6.3 Menjalankan Playwright Tests

```powershell
# Run all tests
npx playwright test

# Run with UI mode (visual feedback)
npx playwright test --ui

# Run specific test file
npx playwright test tests/landing-page.spec.ts

# Run in headed mode (see browser)
npx playwright test --headed

# Generate HTML report
npx playwright test --reporter=html

# View report
npx playwright show-report
```

### 6.4 Visual Regression Testing

```typescript
// tests/visual-regression.spec.ts
import { test, expect } from '@playwright/test';

test('homepage visual regression', async ({ page }) => {
  await page.goto('http://localhost:4321');
  
  // Take screenshot
  await expect(page).toHaveScreenshot('homepage.png', {
    maxDiffPixels: 100,
  });
});

test('mobile view visual regression', async ({ page }) => {
  await page.setViewportSize({ width: 375, height: 667 });
  await page.goto('http://localhost:4321');
  
  await expect(page).toHaveScreenshot('homepage-mobile.png', {
    maxDiffPixels: 100,
  });
});
```

---

## 7. Security Testing (Basic)

### 7.1 Security Checklist

```powershell
# Dependencies
[ ] Run `npm audit` → No critical vulnerabilities
[ ] Run `npm audit fix` → Auto-fix vulnerabilities
[ ] Check for outdated packages → `npm outdated`

# HTTPS & SSL
[ ] Website accessible via HTTPS
[ ] SSL certificate valid (not expired)
[ ] HTTP → HTTPS redirect works
[ ] HSTS header present

# Security Headers
[ ] X-Frame-Options: SAMEORIGIN
[ ] X-Content-Type-Options: nosniff
[ ] X-XSS-Protection: 1; mode=block
[ ] Content-Security-Policy configured
[ ] Referrer-Policy configured

# Form Security
[ ] Form validation (client + server side)
[ ] CSRF protection (jika ada backend)
[ ] Rate limiting pada form submission
[ ] No sensitive data in URL parameters

# Code Security
[ ] No API keys in frontend code
[ ] No .env files committed to Git
[ ] No console.log in production
[ ] No debug mode enabled
```

### 7.2 Security Testing Tools

```powershell
# 1. npm audit
npm audit
npm audit fix

# 2. Snyk (free for open source)
npm install -g snyk
snyk test

# 3. OWASP ZAP (advanced)
# → Download dari https://www.zaproxy.org/
# → Run automated scan

# 4. Security Headers Test
# → https://securityheaders.com/
# → Masukkan URL, lihat security headers score

# 5. SSL Labs Test
# → https://www.ssllabs.com/ssltest/
# → Cek SSL configuration
```

---

## 8. Testing Timeline

### 📅 Pre-Development Testing (Setup)

```powershell
[ ] Setup testing tools (Playwright, Lighthouse CLI)
[ ] Create test data (dummy form submissions)
[ ] Setup test environment (localhost:4321)
[ ] Create testing checklist document
```

### 📅 During Development Testing

```powershell
# Setiap selesai 1 section:
[ ] Manual functional test
[ ] Responsive test (3 breakpoints)
[ ] Accessibility check (axe DevTools)

# Setiap selesai 3 sections:
[ ] Cross-browser test (Chrome, Firefox, Safari)
[ ] Performance test (Lighthouse)
[ ] SEO validation
```

### 📅 Pre-Launch Testing (Final)

```powershell
# 1 hari sebelum launch:
[ ] Full manual testing checklist
[ ] Playwright E2E tests (all pass)
[ ] Lighthouse audit (all scores > 90)
[ ] Cross-browser testing (5 browsers)
[ ] Accessibility audit (WCAG AA)
[ ] Security scan (npm audit)
[ ] Performance optimization check

# Hari launch:
[ ] Final smoke test (production URL)
[ ] Form submission test
[ ] WhatsApp CTA test
[ ] Submit sitemap to Google Search Console
```

### 📅 Post-Launch Testing

```powershell
# 1 minggu setelah launch:
[ ] Monitor Google Search Console (indexing)
[ ] Check analytics (traffic, bounce rate)
[ ] Monitor uptime (UptimeRobot)
[ ] Check form submissions (Formspree)

# 1 bulan setelah launch:
[ ] Performance audit (Lighthouse)
[ ] SEO audit (rankings, backlinks)
[ ] Accessibility re-check
[ ] User feedback collection
```

---

## 9. Testing Automation (CI/CD)

### 9.1 GitHub Actions Workflow

```yaml
# .github/workflows/test.yml
name: Test & Lint

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '20'
          cache: 'npm'
          cache-dependency-path: apps/web/package-lock.json
      
      - name: Install dependencies
        working-directory: ./apps/web
        run: npm ci
      
      - name: Run linter
        working-directory: ./apps/web
        run: npm run lint
      
      - name: Build Astro
        working-directory: ./apps/web
        run: npm run build
      
      - name: Install Playwright browsers
        working-directory: ./apps/web
        run: npx playwright install --with-deps
      
      - name: Run Playwright tests
        working-directory: ./apps/web
        run: npx playwright test
      
      - name: Run Lighthouse audit
        uses: treosh/lighthouse-ci-action@v10
        with:
          urls: |
            http://localhost:4321
          uploadArtifacts: true
      
      - name: Upload test results
        if: always()
        uses: actions/upload-artifact@v4
        with:
          name: playwright-report
          path: apps/web/playwright-report/
```

### 9.2 Lighthouse CI Configuration

```javascript
// lighthouserc.js
module.exports = {
  ci: {
    collect: {
      startServerCommand: 'npm run preview',
      urls: [
        'http://localhost:4321/',
        'http://localhost:4321/kontak',
      ],
    },
    assert: {
      assertions: {
        'categories:performance': ['error', { minScore: 0.9 }],
        'categories:accessibility': ['error', { minScore: 0.9 }],
        'categories:best-practices': ['error', { minScore: 0.9 }],
        'categories:seo': ['error', { minScore: 0.9 }],
      },
    },
    upload: {
      target: 'temporary-public-storage',
    },
  },
};
```

---

## 10. Testing Tools Summary

### 🛠️ Essential Tools (Wajib)

| Tool | Purpose | Cost | URL |
|------|---------|------|-----|
| **Chrome DevTools** | Manual testing, Lighthouse | Free | Built-in |
| **Playwright** | E2E automation | Free | https://playwright.dev |
| **axe DevTools** | Accessibility testing | Free | Chrome Extension |
| **Google Search Console** | SEO monitoring | Free | https://search.google.com/search-console |
| **UptimeRobot** | Uptime monitoring | Free | https://uptimerobot.com |

### 🔧 Advanced Tools (Opsional)

| Tool | Purpose | Cost | URL |
|------|---------|------|-----|
| **BrowserStack** | Cross-browser testing | $29/mo | https://www.browserstack.com |
| **Ahrefs** | SEO & backlink analysis | $99/mo | https://ahrefs.com |
| **Sentry** | Error tracking | Free tier | https://sentry.io |
| **Plausible** | Privacy-friendly analytics | $9/mo | https://plausible.io |

---

## 11. Testing Checklist (Quick Reference)

### ✅ Pre-Launch Checklist

```powershell
# Functional
[ ] All navigation links work
[ ] Form submission works
[ ] WhatsApp CTA works
[ ] Mobile menu works
[ ] All sections visible

# Responsive
[ ] Mobile (375px) ✓
[ ] Tablet (768px) ✓
[ ] Desktop (1440px) ✓
[ ] No horizontal scroll

# Performance
[ ] Lighthouse Performance > 90
[ ] Page load < 2s
[ ] Total size < 500KB

# SEO
[ ] Title tag optimized
[ ] Meta description optimized
[ ] H1 hierarchy correct
[ ] Sitemap generated
[ ] robots.txt configured

# Accessibility
[ ] WCAG AA compliant
[ ] Keyboard navigation works
[ ] Color contrast sufficient
[ ] Alt text on images

# Security
[ ] HTTPS enabled
[ ] Security headers present
[ ] npm audit clean
[ ] No sensitive data exposed

# Cross-Browser
[ ] Chrome ✓
[ ] Firefox ✓
[ ] Safari ✓
[ ] Edge ✓
```

---

## 12. Open Questions

| No | Pertanyaan | Status |
|----|-----------|--------|
| Q1 | Apakah ingin setup Playwright sekarang atau nanti saja? | Pending |
| Q2 | Apakah ada budget untuk BrowserStack/Ahrefs? | Pending |
| Q3 | Apakah ingin automated testing di CI/CD sejak V1? | Pending |
| Q4 | Siapa yang akan melakukan testing? (Solo atau tim?) | Pending |
| Q5 | Apakah ada tools testing yang sudah familiar? | Pending |

---

*Dokumen ini adalah living document. Versi akan diperbarui seiring perkembangan proyek.*

**Last Updated:** 17 September 2026
**Next Step:** Jalankan manual testing checklist, setup Playwright untuk automated testing.