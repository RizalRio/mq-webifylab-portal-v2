# 🖼️ Wireframe & UI Specification Document
## Webifylab Landing Page — Version 1.0

| Metadata         | Detail                                      |
|------------------|---------------------------------------------|
| **Product**      | Webifylab Landing Page                      |
| **Version**      | 1.0                                         |
| **Author**       | Rizal                                       |
| **Created**      | 17 September 2026                           |
| **Status**       | Draft                                       |
| **Tools**        | Figma (optional), ASCII Wireframes          |

---

## 1. Wireframe Overview

### 1.1 Page Structure

Landing page menggunakan pendekatan **single-page scroll** dengan 10 section utama:

```
┌─────────────────────────────────────┐
│ S1: Navbar (Sticky)                 │ ← Fixed di atas
├─────────────────────────────────────┤
│ S2: Hero Section                    │ ← Full viewport height
├─────────────────────────────────────┤
│ S3: Problem Section                 │ ← 3 cards
├─────────────────────────────────────┤
│ S4: Services Section                │ ← 3 cards
├─────────────────────────────────────┤
│ S5: Why Webifylab Section           │ ← 2x2 grid
├─────────────────────────────────────┤
│ S6: Approach Section                │ ← 4 steps horizontal
├─────────────────────────────────────┤
│ S7: Portfolio Section               │ ← 3 cards
├─────────────────────────────────────┤
│ S8: Testimonial Section             │ ← Optional / Placeholder
├─────────────────────────────────────┤
│ S9: CTA / Contact Section           │ ← Form + WhatsApp
├─────────────────────────────────────┤
│ S10: Footer                         │ ← Multi-column
└─────────────────────────────────────┘
```

### 1.2 Global Layout Rules

| Rule | Value |
|------|-------|
| **Container Max Width** | 1280px (7xl) |
| **Container Padding** | 16px (mobile), 24px (tablet), 32px (desktop) |
| **Section Padding** | 60px (mobile), 80px (tablet), 120px (desktop) |
| **Grid Gap** | 16px (mobile), 24px (tablet), 32px (desktop) |
| **Background Pattern** | Alternating white & slate-100 |

---

## 2. Section-by-Section Wireframe

### S1: Navbar (Sticky)

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│  [Logo Webifylab]          Layanan  Pendekatan  Portfolio  Kontak  │
│                                              [Konsultasi Gratis]   │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Height:** 72px
- **Background:** White (transparent saat di atas, solid saat scroll)
- **Position:** Fixed top, z-index: 50
- **Logo:** Left aligned, height 40px
- **Menu Items:** Center-right, gap 32px, font-weight 500
- **CTA Button:** Right aligned, Primary Button variant
- **Shadow:** Subtle shadow saat scroll (shadow-sm)

#### Mobile View (<768px)

```
┌─────────────────────────────────────┐
│  [Logo Webifylab]              [☰] │
└─────────────────────────────────────┘

[Slide-down menu saat hamburger diklik]
┌─────────────────────────────────────┐
│  Layanan                            │
│  Pendekatan                         │
│  Portfolio                          │
│  Kontak                             │
│  [Konsultasi Gratis]                │
└─────────────────────────────────────┘
```

**Specifications:**
- **Height:** 64px
- **Hamburger Icon:** Right aligned, 24x24px
- **Mobile Menu:** Full-width dropdown, background white, shadow-lg
- **Menu Items:** Stacked vertical, padding 16px, border-bottom slate-300
- **CTA Button:** Full-width, margin-top 16px

**Interactions:**
- Scroll > 50px → Background berubah dari transparent ke white + shadow
- Hamburger click → Toggle mobile menu dengan animasi slide-down (300ms)
- Menu item click → Smooth scroll ke section + close mobile menu

---

### S2: Hero Section

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│                                                                     │
│   Dari Ide Menjadi                  [Illustration/Gradient]         │
│   Ekosistem Digital                 ┌─────────────────────┐        │
│                                                                     │
│   Webifylab membantu bisnis Anda    │   Abstract shapes     │        │
│   membangun aplikasi, desain,       │   Gradient: Indigo    │        │
│   dan sistem SaaS yang siap         │   → Cyan              │        │
│   berkembang — didukung oleh        │   Network nodes       │        │
│   keahlian AI & Data.               └─────────────────────┘        │
│                                                                     │
│   [Konsultasi Gratis →]  [Lihat Layanan]                           │
│                                                                     │
│   ✓ Dipercaya oleh UMKM & Startup                                  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Min Height:** 100vh (full viewport)
- **Layout:** 2 columns (60% text, 40% visual)
- **Background:** Gradient (white → slate-100) atau solid white
- **H1:** 48px desktop, 32px mobile, font-weight 800, color slate-900
- **Sub-headline:** 18px, font-weight 400, color slate-700, max-width 600px
- **CTA Buttons:** 
  - Primary: "Konsultasi Gratis →" (margin-right 16px)
  - Secondary: "Lihat Layanan"
- **Trust Badge:** Below CTA, icon Check + text, color slate-500, font-size 14px
- **Visual:** Right side, abstract illustration atau gradient blob, max-width 500px

#### Mobile View (<768px)

```
┌─────────────────────────────────────┐
│                                     │
│   Dari Ide Menjadi                  │
│   Ekosistem Digital                 │
│                                     │
│   Webifylab membantu bisnis Anda    │
│   membangun aplikasi, desain,       │
│   dan sistem SaaS yang siap         │
│   berkembang.                       │
│                                     │
│   [Konsultasi Gratis →]             │
│   [Lihat Layanan]                   │
│                                     │
│   [Illustration]                    │
│   ┌───────────────────────┐        │
│   │   Abstract shapes     │        │
│   └───────────────────────┘        │
│                                     │
└─────────────────────────────────────┘
```

**Specifications:**
- **Layout:** Single column, text top, visual bottom
- **H1:** 32px, centered atau left-aligned
- **CTA Buttons:** Full-width, stacked vertical, gap 12px
- **Visual:** Below text, max-width 100%, margin-top 32px

**Animations:**
- Text: Fade-in + slide-up (delay 0ms, duration 600ms)
- Visual: Fade-in (delay 200ms, duration 600ms)
- CTA Buttons: Fade-in (delay 400ms, duration 600ms)

---

### S3: Problem Section

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   Mengapa Bisnis Anda Butuh Partner Teknologi yang Tepat?          │
│   Banyak bisnis kehilangan kesempatan karena partner teknologi     │
│   yang tidak memahami kebutuhan jangka panjang.                    │
│                                                                     │
│   ┌──────────────┐  ┌──────────────┐  ┌──────────────┐           │
│   │   [Icon]     │  │   [Icon]     │  │   [Icon]     │           │
│   │   Globe      │  │  Handshake   │  │    Brain     │           │
│   │              │  │              │  │              │           │
│   │  Website     │  │   Vendor     │  │  Tidak Siap  │           │
│   │  Saja Tidak  │  │  yang Hilang │  │  Era AI &    │           │
│   │  Cukup       │  │              │  │  Data        │           │
│   │              │  │              │  │              │           │
│   │  Bisnis Anda │  │  Anda butuh  │  │  Kompetitor  │           │
│   │  butuh sistem│  │  partner     │  │  Anda sudah  │           │
│   │  yang        │  │  jangka      │  │  mulai       │           │
│   │  terintegrasi│  │  panjang...  │  │  otomatisasi │           │
│   └──────────────┘  └──────────────┘  └──────────────┘           │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Background:** Slate-100 (alternating dari Hero)
- **Section Title:** 36px desktop, 28px mobile, centered, margin-bottom 16px
- **Section Subtitle:** 18px, color slate-700, centered, max-width 700px, margin-bottom 48px
- **Grid:** 3 columns desktop, 2 columns tablet, 1 column mobile
- **Cards:** 
  - Background: White
  - Padding: 32px
  - Border-radius: 16px
  - Shadow: sm (default), lg (hover)
  - Border: 1px slate-300 (default), indigo (hover)
- **Icon:** 48x48px, background indigo 10%, color indigo, padding 12px, border-radius 12px
- **Card Title:** 20px, font-weight 600, margin-top 16px
- **Card Description:** 16px, color slate-700, margin-top 8px

#### Mobile View (<768px)

```
┌─────────────────────────────────────┐
│   Mengapa Bisnis Anda Butuh         │
│   Partner Teknologi yang Tepat?     │
│                                     │
│   Banyak bisnis kehilangan          │
│   kesempatan karena partner         │
│   teknologi yang tidak memahami     │
│   kebutuhan jangka panjang.         │
│                                     │
│   ┌───────────────────────────┐    │
│   │   [Icon]                  │    │
│   │   Website Saja Tidak      │    │
│   │   Cukup                   │    │
│   │                           │    │
│   │   Bisnis Anda butuh       │    │
│   │   sistem yang             │    │
│   │   terintegrasi...         │    │
│   └───────────────────────────┘    │
│                                     │
│   ┌───────────────────────────┐    │
│   │   [Icon]                  │    │
│   │   Vendor yang Hilang      │    │
│   │   ...                     │    │
│   └───────────────────────────┘    │
│                                     │
│   ┌───────────────────────────┐    │
│   │   [Icon]                  │    │
│   │   Tidak Siap Era AI &     │    │
│   │   Data                    │    │
│   │   ...                     │    │
│   └───────────────────────────┘    │
│                                     │
└─────────────────────────────────────┘
```

**Animations:**
- Cards: Fade-in + slide-up (staggered delay: 0ms, 100ms, 200ms)

---

### S4: Services Section

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   Layanan Kami                                                      │
│   Tiga pilar layanan yang saling terintegrasi untuk membangun      │
│   ekosistem digital bisnis Anda.                                   │
│                                                                     │
│   ┌──────────────┐  ┌──────────────┐  ┌──────────────┐           │
│   │   [Icon]     │  │   [Icon]     │  │   [Icon]     │           │
│   │    Code      │  │   Palette    │  │   Network    │           │
│   │              │  │              │  │              │           │
│   │ Pengembangan │  │  Desain      │  │  SaaS &      │           │
│   │ Aplikasi     │  │  Grafis &    │  │  Ekosistem   │           │
│   │              │  │  Web Design  │  │  Digital     │           │
│   │ Kami membangun│  │              │  │              │           │
│   │ aplikasi web │  │ Visual yang  │  │ Produk SaaS  │           │
│   │ dan mobile   │  │ konsisten    │  │ dari         │           │
│   │ yang         │  │ dan          │  │ Webifylab    │           │
│   │ scalable...  │  │ profesional..│  │ yang saling  │           │
│   │              │  │              │  │ terintegrasi │           │
│   │ • Web App    │  │ • UI/UX      │  │ • SaaS siap  │           │
│   │ • Mobile App │  │ • Brand      │  │ • API-first  │           │
│   │ • API Dev    │  │ • Marketing  │  │ • AI & Data  │           │
│   │ • Legacy     │  │ • Design Sys │  │ • Automation │           │
│   │              │  │              │  │              │           │
│   │ [Pelajari →] │  │ [Pelajari →] │  │ [Pelajari →] │           │
│   └──────────────┘  └──────────────┘  └──────────────┘           │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Background:** White
- **Layout:** Same as Problem Section (3 columns grid)
- **Cards:** Same styling as Problem cards
- **Features List:** 
  - Bullet points dengan icon Check (small)
  - Font-size: 14px, color slate-700
  - Margin-top: 16px
  - Gap: 8px antar item
- **CTA Link:** 
  - Ghost Button variant
  - "Pelajari Lebih Lanjut →"
  - Margin-top: 24px
  - Icon ArrowRight (16x16px)

---

### S5: Why Webifylab Section

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   Mengapa Webifylab?                                               │
│   Kami bukan sekadar vendor — kami adalah partner teknologi yang   │
│   memahami visi jangka panjang bisnis Anda.                        │
│                                                                     │
│   ┌──────────────────────────┐  ┌──────────────────────────┐      │
│   │   [Icon]                 │  │   [Icon]                 │      │
│   │   Blueprint              │  │   Ecosystem              │      │
│   │                          │  │                          │      │
│   │   System-First Thinking  │  │   Ekosistem, Bukan       │      │
│   │                          │  │   Sekadar Proyek         │      │
│   │                          │  │                          │      │
│   │   Setiap proyek dimulai  │  │   Kami membangun fondasi │      │
│   │   dari analisis sistem   │  │   yang bisa berkembang.  │      │
│   │   yang mendalam...       │  │   Website hari ini bisa  │      │
│   │                          │  │   menjadi bagian dari    │      │
│   │                          │  │   ekosistem SaaS besok.  │      │
│   └──────────────────────────┘  └──────────────────────────┘      │
│                                                                     │
│   ┌──────────────────────────┐  ┌──────────────────────────┐      │
│   │   [Icon]                 │  │   [Icon]                 │      │
│   │   Brain                  │  │   Chat                   │      │
│   │                          │  │                          │      │
│   │   AI & Data Ready        │  │   Transparent & Agile    │      │
│   │                          │  │                          │      │
│   │   Dengan keahlian di     │  │   Komunikasi terbuka,    │      │
│   │   bidang AI dan Data,    │  │   iterasi cepat,         │      │
│   │   kami membantu bisnis   │  │   dokumentasi rapi.      │      │
│   │   Anda siap menghadapi   │  │   Anda selalu tahu       │      │
│   │   era otomatisasi...     │  │   progress proyek...     │      │
│   └──────────────────────────┘  └──────────────────────────┘      │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Background:** Slate-100
- **Layout:** 2x2 grid (desktop), 1 column (mobile)
- **Cards:** 
  - Background: White
  - Padding: 32px
  - Border-radius: 16px
  - No border (clean look)
  - Shadow: sm
- **Icon:** 40x40px, color indigo
- **Card Title:** 20px, font-weight 600, margin-top 12px
- **Card Description:** 16px, color slate-700, margin-top 8px, line-height 1.6

---

### S6: Approach Section

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   Pendekatan Kami                                                  │
│   Metodologi kerja yang terstruktur untuk memastikan proyek        │
│   berjalan lancar dan hasil maksimal.                              │
│                                                                     │
│   ┌─────────┐    ┌─────────┐    ┌─────────┐    ┌─────────┐       │
│   │   01    │───►│   02    │───►│   03    │───►│   04    │       │
│   │         │    │         │    │         │    │         │       │
│   │Discovery│    │ Design  │    │Develop- │    │ Deploy  │       │
│   │         │    │&Planning│    │  ment   │    │& Support│       │
│   │         │    │         │    │         │    │         │       │
│   │Kami     │    │Kami     │    │Kami     │    │Kami     │       │
│   │mendeng- │    │merancang│    │membangun│    │membantu │       │
│   │arkan    │    │arsitektur│   │solusi   │    │deploy-  │       │
│   │kebutuhan│    │sistem,  │    │dengan   │    │ment,    │       │
│   │bisnis   │    │wireframe│    │tech     │    │memberi- │       │
│   │Anda...  │    │dan      │    │stack    │    │kan      │       │
│   │         │    │roadmap..│    │modern...│    │dokumen- │       │
│   │         │    │         │    │         │    │tasi...  │       │
│   └─────────┘    └─────────┘    └─────────┘    └─────────┘       │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Background:** White
- **Layout:** Horizontal flow (desktop), vertical stack (mobile)
- **Steps:** 4 items dengan connector line/arrow
- **Number Badge:** 
  - 48x48px circle
  - Background: indigo
  - Text: white, font-weight 700, font-size 20px
- **Step Title:** 20px, font-weight 600, margin-top 16px
- **Step Description:** 16px, color slate-700, margin-top 8px
- **Connector:** Line atau arrow antara steps, color slate-300

#### Mobile View (<768px)

```
┌─────────────────────────────────────┐
│   Pendekatan Kami                   │
│                                     │
│   Metodologi kerja yang terstruktur │
│   untuk memastikan proyek berjalan  │
│   lancar dan hasil maksimal.        │
│                                     │
│   ┌───────────────────────────┐    │
│   │   01                      │    │
│   │   Discovery               │    │
│   │                           │    │
│   │   Kami mendengarkan       │    │
│   │   kebutuhan bisnis Anda... │    │
│   └───────────────────────────┘    │
│              │                      │
│              ▼                      │
│   ┌───────────────────────────┐    │
│   │   02                      │    │
│   │   Design & Planning       │    │
│   │   ...                     │    │
│   └───────────────────────────┘    │
│              │                      │
│              ▼                      │
│   ┌───────────────────────────┐    │
│   │   03                      │    │
│   │   Development             │    │
│   │   ...                     │    │
│   └───────────────────────────┘    │
│              │                      │
│              ▼                      │
│   ┌───────────────────────────┐    │
│   │   04                      │    │
│   │   Deploy & Support        │    │
│   │   ...                     │    │
│   └───────────────────────────┘    │
│                                     │
└─────────────────────────────────────┘
```

---

### S7: Portfolio Section

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   Portfolio                                                        │
│   Beberapa proyek yang telah kami kerjakan.                        │
│                                                                     │
│   ┌──────────────┐  ┌──────────────┐  ┌──────────────┐           │
│   │              │  │              │  │              │           │
│   │  [Thumbnail] │  │  [Thumbnail] │  │  [Thumbnail] │           │
│   │              │  │              │  │              │           │
│   │  ┌────────┐  │  │  ┌────────┐  │  │  ┌────────┐  │           │
│   │  │Coming  │  │  │  │Coming  │  │  │  │Coming  │  │           │
│   │  │Soon    │  │  │  │Soon    │  │  │  │Soon    │  │           │
│   │  └────────┘  │  │  └────────┘  │  │  └────────┘  │           │
│   │              │  │              │  │              │           │
│   ├──────────────┤  ├──────────────┤  ├──────────────┤           │
│   │ Proyek 1     │  │ Proyek 2     │  │ Proyek 3     │           │
│   │ Web App      │  │ UI/UX Design │  │ SaaS Dev     │           │
│   │              │  │              │  │              │           │
│   │ Detail proyek│  │ Detail proyek│  │ Detail proyek│           │
│   │ akan segera  │  │ akan segera  │  │ akan segera  │           │
│   │ ditampilkan. │  │ ditampilkan. │  │ ditampilkan. │           │
│   └──────────────┘  └──────────────┘  └──────────────┘           │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Background:** Slate-100
- **Layout:** 3 columns grid (desktop), 2 columns (tablet), 1 column (mobile)
- **Cards:**
  - Background: White
  - Border-radius: 16px
  - Overflow: hidden
  - Shadow: sm
- **Thumbnail:**
  - Aspect ratio: 16:9
  - Background: gradient (indigo → cyan) atau placeholder image
  - Overlay: "Coming Soon" badge (centered, background white 90%, padding 8px 16px)
- **Card Content:**
  - Padding: 24px
  - Title: 20px, font-weight 600
  - Category Badge: Small badge (cyan), margin-top 8px
  - Description: 14px, color slate-700, margin-top 12px

---

### S8: Testimonial Section (Optional)

**Status:** Hidden di V1 jika belum ada testimonial

**Jika ada testimonial:**

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   Apa Kata Klien Kami                                             │
│   Kepuasan klien adalah prioritas utama kami.                      │
│                                                                     │
│   ┌──────────────────────────────────────────────────────────┐    │
│   │  "Webifylab membantu kami membangun sistem yang scalable │    │
│   │   dan mudah dikembangkan. Tim mereka sangat profesional  │    │
│   │   dan responsif."                                        │    │
│   │                                                          │    │
│   │  [Avatar]  Nama Klien                                    │    │
│   │            CEO, PT Contoh                                │    │
│   └──────────────────────────────────────────────────────────┘    │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Background:** White
- **Layout:** Carousel atau grid (jika > 1 testimonial)
- **Card:**
  - Background: slate-100
  - Padding: 32px
  - Border-radius: 16px
  - Quote icon: Large, color indigo, opacity 20%
- **Quote Text:** 18px, font-style italic, color slate-900
- **Author:** 
  - Avatar: 48x48px circle
  - Name: 16px, font-weight 600
  - Position: 14px, color slate-700

---

### S9: CTA / Contact Section

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   Siap Membangun Sesuatu yang Luar Biasa?                         │
│   Ceritakan ide atau tantangan Anda. Konsultasi pertama gratis,   │
│   tanpa komitmen.                                                  │
│                                                                     │
│   ┌──────────────────────────────────────────────────────────┐    │
│   │                                                          │    │
│   │   Nama Lengkap                                           │    │
│   │   ┌────────────────────────────────────────────────┐    │    │
│   │   │ Masukkan nama Anda                             │    │    │
│   │   └────────────────────────────────────────────────┘    │    │
│   │                                                          │    │
│   │   Email Aktif                                            │    │
│   │   ┌────────────────────────────────────────────────┐    │    │
│   │   │ nama@email.com                                 │    │    │
│   │   └────────────────────────────────────────────────┘    │    │
│   │                                                          │    │
│   │   Jenis Layanan yang Dibutuhkan                         │    │
│   │   ┌────────────────────────────────────────────────┐    │    │
│   │   │ Pilih layanan                              ▼   │    │    │
│   │   └────────────────────────────────────────────────┘    │    │
│   │                                                          │    │
│   │   Ceritakan Kebutuhan Anda                               │    │
│   │   ┌────────────────────────────────────────────────┐    │    │
│   │   │                                                │    │    │
│   │   │ Deskripsikan proyek atau tantangan Anda...     │    │    │
│   │   │                                                │    │    │
│   │   │                                                │    │    │
│   │   └────────────────────────────────────────────────┘    │    │
│   │                                                          │    │
│   │   [Kirim Pesan →]                                        │    │
│   │                                                          │    │
│   │   Atau Chat via WhatsApp 💬                              │    │
│   │                                                          │    │
│   └──────────────────────────────────────────────────────────┘    │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Background:** Gradient (indigo → deep blue) atau solid deep blue
- **Text Color:** White
- **Layout:** Centered, max-width 600px
- **Section Title:** 36px, font-weight 700, centered
- **Section Subtitle:** 18px, opacity 90%, centered, margin-bottom 48px
- **Form Container:**
  - Background: White
  - Padding: 48px
  - Border-radius: 24px
  - Shadow: 2xl
- **Form Fields:**
  - Label: 14px, font-weight 500, color slate-900, margin-bottom 8px
  - Input: Full-width, padding 12px 16px, border 1px slate-300, border-radius 8px
  - Focus: Ring 2px indigo, border transparent
  - Gap antar fields: 24px
- **Submit Button:** 
  - Primary variant, full-width
  - Text: "Kirim Pesan →"
  - Loading state: Spinner + "Mengirim..."
- **WhatsApp Link:**
  - Below form, margin-top 24px
  - Text: "Atau Chat via WhatsApp 💬"
  - Style: Ghost button atau text link
  - Icon: WhatsApp (optional)

#### Mobile View (<768px)

```
┌─────────────────────────────────────┐
│   Siap Membangun Sesuatu yang       │
│   Luar Biasa?                       │
│                                     │
│   Ceritakan ide atau tantangan      │
│   Anda. Konsultasi pertama gratis.  │
│                                     │
│   ┌───────────────────────────┐    │
│   │                           │    │
│   │   Nama Lengkap            │    │
│   │   ┌─────────────────┐    │    │
│   │   │                 │    │    │
│   │   └─────────────────┘    │    │
│   │                           │    │
│   │   Email Aktif             │    │
│   │   ┌─────────────────┐    │    │
│   │   │                 │    │    │
│   │   └─────────────────┘    │    │
│   │                           │    │
│   │   Jenis Layanan           │    │
│   │   ┌─────────────────┐    │    │
│   │   │             ▼   │    │    │
│   │   └─────────────────┘    │    │
│   │                           │    │
│   │   Ceritakan Kebutuhan     │    │
│   │   ┌─────────────────┐    │    │
│   │   │                 │    │    │
│   │   │                 │    │    │
│   │   └─────────────────┘    │    │
│   │                           │    │
│   │   [Kirim Pesan →]         │    │
│   │                           │    │
│   │   Atau Chat via WhatsApp  │    │
│   │                           │    │
│   └───────────────────────────┘    │
│                                     │
└─────────────────────────────────────┘
```

**Specifications (Mobile):**
- Form container padding: 24px
- Form fields: Same as desktop
- Submit button: Full-width

**Success State:**

```
┌─────────────────────────────────────┐
│   ✓                                 │
│                                     │
│   Pesan Terkirim!                   │
│                                     │
│   Terima kasih telah menghubungi   │
│   kami. Kami akan membalas dalam    │
│   1x24 jam kerja.                   │
│                                     │
│   [Kembali ke Beranda]              │
│                                     │
└─────────────────────────────────────┘
```

**Error State:**

```
┌─────────────────────────────────────┐
│   ✗                                 │
│                                     │
│   Oops! Terjadi Kesalahan           │
│                                     │
│   Gagal mengirim pesan. Silakan     │
│   coba lagi atau hubungi kami via   │
│   WhatsApp.                         │
│                                     │
│   [Coba Lagi]  [Chat WhatsApp]      │
│                                     │
└─────────────────────────────────────┘
```

---

### S10: Footer

#### Desktop View (≥1024px)

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   [Logo Webifylab]                                                  │
│   Dari Ide Menjadi Ekosistem Digital                               │
│                                                                     │
│   ┌──────────────┐  ┌──────────────┐  ┌──────────────┐           │
│   │ Layanan      │  │ Perusahaan   │  │ Kontak       │           │
│   │              │  │              │  │              │           │
│   │ • Pengembangan│  │ • Tentang    │  │ • Email      │           │
│   │   Aplikasi   │  │   Kami       │  │   hello@...  │           │
│   │ • Desain     │  │ • Portfolio  │  │              │           │
│   │   Grafis     │  │ • Pendekatan │  │ • WhatsApp   │           │
│   │ • SaaS       │  │ • Kontak     │  │   +62...     │           │
│   │              │  │              │  │              │           │
│   │              │  │ Legal        │  │ Social       │           │
│   │              │  │ • Privacy    │  │ • LinkedIn   │           │
│   │              │  │ • Terms      │  │ • GitHub     │           │
│   │              │  │              │  │ • Instagram  │           │
│   └──────────────┘  └──────────────┘  └──────────────┘           │
│                                                                     │
│   ─────────────────────────────────────────────────────────────── │
│                                                                     │
│   © 2026 Webifylab. All rights reserved.                         │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

**Specifications:**
- **Background:** Slate-900 (dark)
- **Text Color:** White (headings), slate-300 (links)
- **Padding:** 80px vertical, 32px horizontal
- **Layout:** 
  - Top: Logo + tagline (left aligned)
  - Middle: 3-4 columns grid (links)
  - Bottom: Copyright (centered, border-top slate-700, padding-top 32px)
- **Logo:** White version, height 40px
- **Tagline:** 16px, color slate-300, margin-top 12px
- **Column Titles:** 14px, font-weight 600, color white, margin-bottom 16px
- **Links:** 14px, color slate-300, hover: white, gap 12px
- **Social Icons:** 24x24px, color slate-300, hover: white, gap 16px
- **Copyright:** 14px, color slate-500, centered

#### Mobile View (<768px)

```
┌─────────────────────────────────────┐
│   [Logo Webifylab]                  │
│   Dari Ide Menjadi Ekosistem Digital│
│                                     │
│   Layanan                           │
│   • Pengembangan Aplikasi           │
│   • Desain Grafis                   │
│   • SaaS                            │
│                                     │
│   Perusahaan                        │
│   • Tentang Kami                    │
│   • Portfolio                       │
│   • Pendekatan                      │
│   • Kontak                          │
│                                     │
│   Kontak                            │
│   • hello@webifylab.com             │
│   • +62...                          │
│                                     │
│   Social                            │
│   [LinkedIn] [GitHub] [Instagram]   │
│                                     │
│   ──────────────────────────────── │
│                                     │
│   © 2026 Webifylab. All rights      │
│   reserved.                         │
│                                     │
└─────────────────────────────────────┘
```

**Specifications (Mobile):**
- Columns: Stacked vertical
- Gap antar columns: 32px

---

## 3. Responsive Breakpoints Summary

| Section | Mobile (<768px) | Tablet (768-1023px) | Desktop (≥1024px) |
|---------|-----------------|---------------------|-------------------|
| **Navbar** | Hamburger menu | Hamburger menu | Full menu |
| **Hero** | Stacked (text + visual) | 2 columns (50/50) | 2 columns (60/40) |
| **Problem** | 1 column | 2 columns | 3 columns |
| **Services** | 1 column | 2 columns | 3 columns |
| **Why** | 1 column | 2 columns | 2x2 grid |
| **Approach** | Vertical stack | Horizontal | Horizontal |
| **Portfolio** | 1 column | 2 columns | 3 columns |
| **Testimonial** | 1 column | 1 column | Carousel/Grid |
| **Contact** | Form full-width | Form max-600px | Form max-600px |
| **Footer** | Stacked columns | 2 columns | 3-4 columns |

---

## 4. Component Usage Matrix

| Component | S1 | S2 | S3 | S4 | S5 | S6 | S7 | S8 | S9 | S10 |
|-----------|----|----|----|----|----|----|----|----|----|----|
| **Button Primary** | ✓ | ✓ | | | | | | | ✓ | |
| **Button Secondary** | | ✓ | | | | | | | | |
| **Button Ghost** | | | | ✓ | | | | | | ✓ |
| **Card (Service)** | | | | ✓ | | | | | | |
| **Card (Problem)** | | | ✓ | | | | | | | |
| **Card (Why)** | | | | | ✓ | | | | | |
| **Card (Portfolio)** | | | | | | | ✓ | | | |
| **Card (Testimonial)**| | | | | | | | ✓ | | |
| **Input Text** | | | | | | | | | ✓ | |
| **Input Textarea** | | | | | | | | | ✓ | |
| **Input Select** | | | | | | | | | ✓ | |
| **Badge** | | | | ✓ | | | ✓ | | | |
| **Icon** | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | | ✓ |

---

## 5. Animation & Interaction Summary

| Element | Animation | Trigger | Duration | Easing |
|---------|-----------|---------|----------|--------|
| **Hero Text** | Fade-in + slide-up | Page load | 600ms | ease-out |
| **Hero Visual** | Fade-in | Page load (delay 200ms) | 600ms | ease-out |
| **Section Cards** | Fade-in + slide-up | Scroll into view (staggered) | 400ms | ease-out |
| **Button Hover** | Scale + opacity | Hover | 200ms | ease |
| **Card Hover** | Shadow + border | Hover | 300ms | ease |
| **Navbar** | Background change | Scroll > 50px | 300ms | ease |
| **Mobile Menu** | Slide-down | Hamburger click | 300ms | ease-out |
| **Form Success** | Fade-in | Form submit | 300ms | ease-out |

---

## 6. Accessibility Considerations

| Element | Requirement |
|---------|-------------|
| **Images** | Alt text untuk semua gambar (decorative images: alt="") |
| **Icons** | ARIA labels untuk icon-only buttons |
| **Forms** | Labels untuk semua inputs, error messages jelas |
| **Colors** | Kontras minimal 4.5:1 untuk teks, 3:1 untuk UI components |
| **Focus** | Visible focus states untuk semua interactive elements |
| **Keyboard** | Semua interaksi bisa diakses via keyboard (Tab, Enter, Esc) |
| **Motion** | Respect `prefers-reduced-motion` (disable animations) |
| **Skip Link** | "Skip to content" link di awal halaman |

---

## 7. Wireframe Checklist

Sebelum development, pastikan:

- [ ] Semua 10 section sudah di-wireframe
- [ ] Desktop view (≥1024px) sudah detail
- [ ] Mobile view (<768px) sudah detail
- [ ] Tablet view (768-1023px) sudah dipertimbangkan
- [ ] Component usage sudah jelas
- [ ] Animasi & interaksi sudah didefinisikan
- [ ] Accessibility requirements sudah di-note
- [ ] Responsive breakpoints sudah disetujui

---

## 8. Open Questions

| No | Pertanyaan | Status |
|----|-----------|--------|
| Q1 | Apakah layout Hero Section (2 columns desktop, stacked mobile) sudah sesuai? | Pending |
| Q2 | Apakah ingin menggunakan Figma untuk wireframe visual? Atau cukup dengan ASCII + spesifikasi ini? | Pending |
| Q3 | Apakah ada preferensi untuk visual Hero Section? (Ilustrasi abstrak, gradient blob, atau screenshot produk?) | Pending |
| Q4 | Apakah section Testimonial (S8) ingin ditampilkan di V1 dengan placeholder, atau disembunyikan? | Pending |
| Q5 | Apakah ada animasi khusus yang diinginkan? Atau cukup dengan fade-in + slide-up standar? | Pending |

---

*Dokumen ini adalah living document. Versi akan diperbarui seiring feedback dan perkembangan desain.*

**Last Updated:** 17 September 2026
**Next Step:** Review wireframe, finalisasi layout, lalu mulai development atau buat mockup visual di Figma.