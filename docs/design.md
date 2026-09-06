# 📄 UI/UX DESIGN SPECIFICATION DOCUMENT
## WebifyLab Portal — Phase 1 (MVP)

---

### 1. DOCUMENT INFORMATION

| Field | Detail |
|---|---|
| **Product** | WebifyLab Portal |
| **Document Type** | UI/UX Design Specification |
| **Version** | 1.0.0 |
| **Date** | 05 September 2026 |
| **Target Pages** | Public Pages + Admin Dashboard |
| **Design System** | Tailwind CSS + shadcn/ui |
| **Author** | Rizal (System Analyst & Fullstack Developer) |
| **Status** | Approved — Ready for Frontend Development |

---

### 2. DESIGN PRINCIPLES

#### 2.1 Brand-Aligned Design Philosophy

Berdasarkan brand philosophy **"We Experiment, You Grow"** di PRD, desain harus mencerminkan:

| Principle | Design Implementation |
|---|---|
| **Experimental** | Gunakan aksen visual yang playful (orange accent, subtle animations), tapi tetap profesional |
| **Educative** | Layout yang mudah dipindai, typography yang readable, clear hierarchy |
| **Collaborative** | Tone visual yang approachable, bukan intimidating corporate |
| **Iterative** | Design system yang modular, mudah di-adjust seiring feedback |

#### 2.2 Design Principles for WebifyLab

| No | Principle | Implementation |
|---|---|---|
| 1 | **Clarity over Cleverness** | User harus bisa find what they need dalam < 3 detik |
| 2 | **Content-First** | Konten (case study, blog) adalah hero, bukan decoration |
| 3 | **Progressive Disclosure** | Tampilkan info penting dulu, detail on-demand |
| 4 | **Mobile-First** | Design untuk mobile dulu, lalu enhance untuk desktop |
| 5 | **Accessibility** | Contrast ratio ≥ 4.5:1, focus states, alt text |
| 6 | **Performance** | Minimal JavaScript, lazy load images, optimize fonts |

#### 2.3 Visual Tone

| Aspect | Direction |
|---|---|
| **Overall Feel** | Modern, clean, tech-forward tapi approachable |
| **Color Usage** | Lab Blue dominan, Experiment Orange sebagai accent |
| **Imagery** | Screenshots project real, abstract tech illustrations |
| **Iconography** | Lucide Icons — outlined style, consistent stroke |
| **Animations** | Subtle, purposeful (fade-in on scroll, hover states) |
| **Whitespace** | Generous, jangan cram content |

---

### 3. DESIGN SYSTEM

#### 3.1 Color Palette

**Primary Colors:**
```
Lab Blue:          #2563EB  → Primary CTA, links, brand accent
Lab Blue Hover:    #1D4ED8  → Hover state
Lab Blue Dark:     #1E40AF  → Active state

Experiment Orange: #F97316  → Secondary CTA, highlights, badges
Experiment Orange Hover: #EA580C
```

**Secondary Colors:**
```
Growth Green:      #10B981  → Success states, metrics, results
Deep Slate:        #0F172A  → Primary text, headings
Slate Gray:        #64748B  → Secondary text
Light Gray:        #F1F5F9  → Backgrounds, cards
Border Gray:       #E2E8F0  → Borders, dividers
```

**Semantic Colors:**
```
Error Red:         #EF4444  → Error states
Warning Amber:     #F59E0B  → Warning states
Info Blue:         #3B82F6  → Info states
```

**Tailwind Config:**
```typescript
// frontend/tailwind.config.ts
import type { Config } from 'tailwindcss'

const config: Config = {
  content: ['./src/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        // Brand colors
        'lab-blue': {
          DEFAULT: '#2563EB',
          hover: '#1D4ED8',
          dark: '#1E40AF',
        },
        'experiment-orange': {
          DEFAULT: '#F97316',
          hover: '#EA580C',
        },
        'growth-green': {
          DEFAULT: '#10B981',
        },
        'deep-slate': '#0F172A',
        'slate-gray': '#64748B',
        'light-gray': '#F1F5F9',
        'border-gray': '#E2E8F0',
        // Semantic
        'error': '#EF4444',
        'warning': '#F59E0B',
        'info': '#3B82F6',
      },
      fontFamily: {
        sans: ['Inter', 'Plus Jakarta Sans', 'sans-serif'],
        mono: ['JetBrains Mono', 'monospace'],
      },
    },
  },
  plugins: [],
}

export default config
```

#### 3.2 Typography Scale

```
Display XL:    4rem    (64px)  / 700  → Hero headline
Display LG:    3rem    (48px)  / 700  → Section headings
Heading XL:    2.25rem (36px)  / 700  → Page titles
Heading LG:    1.875rem (30px) / 600  → Card titles
Heading MD:    1.5rem  (24px)  / 600  → Sub-sections
Heading SM:    1.25rem (20px)  / 600  → Small headings
Body LG:       1.125rem (18px) / 400  → Lead paragraphs
Body MD:       1rem    (16px)  / 400  → Default body text
Body SM:       0.875rem (14px) / 400  → Captions, metadata
Body XS:       0.75rem (12px)  / 400  → Labels, badges
```

**Line Heights:**
- Headings: 1.2
- Body: 1.6
- Captions: 1.4

#### 3.3 Spacing Scale

```
4px   → xs   (gap between inline elements)
8px   → sm   (padding small, gap between related items)
16px  → md   (padding medium, gap between sections)
24px  → lg   (padding large)
32px  → xl   (section spacing)
48px  → 2xl  (major section spacing)
64px  → 3xl  (page-level spacing)
96px  → 4xl  (hero spacing)
```

#### 3.4 Border Radius

```
4px   → sm   (badges, small buttons)
8px   → md   (buttons, inputs, cards)
12px  → lg   (large cards, modals)
16px  → xl   (featured cards)
9999px → full (pills, avatars)
```

#### 3.5 Shadows

```
shadow-sm:  0 1px 2px rgba(0,0,0,0.05)  → Cards default
shadow-md:  0 4px 6px rgba(0,0,0,0.07)  → Cards hover
shadow-lg:  0 10px 15px rgba(0,0,0,0.1) → Modals, dropdowns
shadow-xl:  0 20px 25px rgba(0,0,0,0.1) → Hero elements
```

#### 3.6 Breakpoints

| Name | Min Width | Usage |
|---|---|---|
| `mobile` | 320px | Mobile phones (default) |
| `sm` | 640px | Small phones landscape |
| `md` | 768px | Tablets |
| `lg` | 1024px | Desktops |
| `xl` | 1280px | Large desktops |
| `2xl` | 1536px | Extra large screens |

---

### 4. COMPONENT LIBRARY

#### 4.1 Buttons

**Primary Button (Lab Blue):**
```
┌─────────────────────────┐
│   Mulai Eksperimen      │
└─────────────────────────┘
Background: #2563EB
Text: white
Padding: 12px 24px
Border-radius: 8px
Font: 600, 16px
Hover: #1D4ED8, slight scale(1.02)
Focus: ring 2px #2563EB
```

**Secondary Button (Experiment Orange):**
```
┌─────────────────────────┐
│   Lihat Portfolio       │
└─────────────────────────┘
Background: transparent
Border: 2px solid #F97316
Text: #F97316
Padding: 12px 24px
Border-radius: 8px
Hover: bg #FFF7ED
```

**Ghost Button:**
```
┌─────────────────────────┐
│   Pelajari Lebih Lanjut │
└─────────────────────────┘
Background: transparent
Text: #2563EB
Padding: 8px 16px
Hover: underline or bg #EFF6FF
```

**Disabled State:**
```
Opacity: 0.5
Cursor: not-allowed
No hover effects
```

#### 4.2 Cards

**Portfolio/Case Study Card:**
```
┌─────────────────────────────────┐
│  ┌─────────────────────────┐   │
│  │                         │   │
│  │     Thumbnail Image     │   │
│  │     (16:9 ratio)        │   │
│  │                         │   │
│  └─────────────────────────┘   │
│                                 │
│  [Web Development]  [SPK]      │ ← Category badges
│                                 │
│  Optimasi Stok UMKM            │ ← Title (Heading LG)
│  dengan SPK                    │
│                                 │
│  Lorem ipsum dolor sit amet,   │ ← Excerpt (Body SM)
│  consectetur adipiscing...     │
│                                 │
│  ───────────────────────────   │
│                                 │
│  🏷️ Retail    📅 Aug 2026     │ ← Meta info
│                                 │
└─────────────────────────────────┘
```

**Blog Card:**
```
┌─────────────────────────────────┐
│  ┌─────────────────────────┐   │
│  │                         │   │
│  │     Cover Image         │   │
│  │     (16:9 ratio)        │   │
│  │                         │   │
│  └─────────────────────────┘   │
│                                 │
│  [Tutorial]                    │ ← Category badge
│                                 │
│  Membuat REST API dengan       │ ← Title (Heading MD)
│  Golang dan Gin                │
│                                 │
│  Panduan lengkap membuat       │ ← Excerpt (Body SM)
│  REST API dari nol...          │
│                                 │
│  ───────────────────────────   │
│                                 │
│  👤 Rizal  📅 Sep 5  ⏱️ 8 min │ ← Author, date, read time
│                                 │
└─────────────────────────────────┘
```

**Service Card:**
```
┌─────────────────────────────────┐
│                                 │
│        🌐                      │ ← Icon (Lucide)
│                                 │
│   Web Development              │ ← Title (Heading LG)
│                                 │
│   Pembuatan website            │ ← Description (Body MD)
│   profesional untuk UMKM       │
│   dan Startup                  │
│                                 │
│   ─────────────────────────    │
│                                 │
│   ✓ Company profile            │ ← Use cases list
│   ✓ Landing page               │
│   ✓ Web app                    │
│   ✓ E-commerce                 │
│                                 │
│   ─────────────────────────    │
│                                 │
│   ⏱️ 2-6 minggu                │ ← Timeline
│   🛠️ Next.js, Golang           │ ← Tech stack
│                                 │
│   [Diskusikan Eksperimen Ini →]│ ← CTA (Link)
│                                 │
└─────────────────────────────────┘
```

#### 4.3 Forms

**Input Field:**
```
┌─────────────────────────────────┐
│  Nama Lengkap *                 │ ← Label (Body SM, 600)
│  ┌─────────────────────────┐   │
│  │                         │   │
│  │  Budi Santoso           │   │ ← Input (Body MD)
│  │                         │   │
│  └─────────────────────────┘   │
│                                 │
│  Border: #E2E8F0               │
│  Border-radius: 8px            │
│  Padding: 12px 16px            │
│  Focus: border #2563EB,        │
│         ring 2px               │
│  Error: border #EF4444,        │
│         message below          │
└─────────────────────────────────┘
```

**Textarea:**
```
┌─────────────────────────────────┐
│  Deskripsi Project *            │
│  ┌─────────────────────────┐   │
│  │                         │   │
│  │  Saya ingin membuat     │   │
│  │  website company        │   │
│  │  profile untuk...       │   │
│  │                         │   │
│  │                         │   │
│  │                         │   │
│  └─────────────────────────┘   │
│                                 │
│  Min-height: 120px             │
│  Resize: vertical              │
└─────────────────────────────────┘
```

**Select/Dropdown:**
```
┌─────────────────────────────────┐
│  Jenis Layanan *                │
│  ┌─────────────────────────┐   │
│  │ Web Development      ▼  │   │ ← Custom select
│  └─────────────────────────┘   │
│                                 │
│  Options:                       │
│  - Web Development              │
│  - Mobile App                   │
│  - UI/UX Design                 │
│  - AI Integration               │
│  - Other                        │
└─────────────────────────────────┘
```

**Error State:**
```
┌─────────────────────────────────┐
│  Email *                        │
│  ┌─────────────────────────┐   │
│  │  invalid-email          │   │ ← Red border
│  └─────────────────────────┘   │
│  ⚠️ Email tidak valid           │ ← Error message (Error Red)
└─────────────────────────────────┘
```

#### 4.4 Badges

**Category Badge:**
```
┌──────────────┐
│ Web Development │
└──────────────┘
Background: #EFF6FF (light blue)
Text: #2563EB
Padding: 4px 12px
Border-radius: 9999px (pill)
Font: 500, 12px
```

**Status Badge:**
```
┌─────────┐
│ Published │
└─────────┘
Published: bg #D1FAE5, text #065F46
Draft:     bg #FEF3C7, text #92400E
Archived:  bg #F1F5F9, text #64748B
```

**AI Type Badge:**
```
┌─────┐  ┌────────────┐  ┌─────────────┐
│ SPK │  │ Sistem Pakar │  │ Data Analytics │
└─────┘  └────────────┘  └─────────────┘
Background: #FFF7ED (light orange)
Text: #F97316
```

#### 4.5 Navigation

**Desktop Header:**
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  🔬 WebifyLab                    Home  About  Services  Portfolio│
│                                 Case Studies  Blog  Products    │
│                                                                 │
│                              [Mulai Eksperimen]  [Login]        │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
Height: 72px
Background: white
Border-bottom: 1px solid #E2E8F0
Sticky: top-0
```

**Mobile Header:**
```
┌─────────────────────────────────┐
│                                 │
│  🔬 WebifyLab           [☰]    │ ← Hamburger menu
│                                 │
└─────────────────────────────────┘
Height: 64px
```

**Mobile Menu (Expanded):**
```
┌─────────────────────────────────┐
│                                 │
│  🔬 WebifyLab           [✕]    │
│                                 │
│  ───────────────────────────   │
│                                 │
│  Home                          │
│  About                         │
│  Services                      │
│  Portfolio                     │
│  Case Studies                  │
│  Blog                          │
│  Products                      │
│  Contact                       │
│                                 │
│  ───────────────────────────   │
│                                 │
│  [Mulai Eksperimen]            │
│                                 │
└─────────────────────────────────┘
```

**Footer:**
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  🔬 WebifyLab                                                   │
│  We Experiment, You Grow                                        │
│                                                                 │
│  ─────────────────────────────────────────────────────────────  │
│                                                                 │
│  Services          Company         Resources       Contact     │
│  ─────────         ─────────       ─────────       ─────────   │
│  Web Development   About           Blog            Email       │
│  Mobile App        Portfolio       Case Studies    WhatsApp    │
│  UI/UX Design      Team            Tutorials       Social      │
│  AI Integration                    Lessons                    │
│                                                                 │
│  ─────────────────────────────────────────────────────────────  │
│                                                                 │
│  © 2026 WebifyLab. All rights reserved.                        │
│  Built with 🧪 in Bogor                                        │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
Background: #0F172A (Deep Slate)
Text: white
Padding: 64px 0 32px
```

#### 4.6 Modal

```
┌─────────────────────────────────────────────────┐
│                                                 │
│  ┌─────────────────────────────────────────┐   │
│  │                                         │   │
│  │  ⚠️ Konfirmasi Hapus                    │   │ ← Title
│  │                                         │   │
│  │  Apakah Anda yakin ingin menghapus      │   │ ← Body
│  │  blog post ini? Tindakan ini tidak      │   │
│  │  dapat dibatalkan.                      │   │
│  │                                         │   │
│  │  ─────────────────────────────────────  │   │
│  │                                         │   │
│  │  [Batal]              [Hapus]           │   │ ← Actions
│  │                                         │   │
│  └─────────────────────────────────────────┘   │
│                                                 │
│  Overlay: rgba(0,0,0,0.5)                      │
│  Backdrop blur                                 │
│                                                 │
└─────────────────────────────────────────────────┘
```

#### 4.7 Toast/Notification

```
┌─────────────────────────────────┐
│  ✓ Blog post berhasil disimpan │
└─────────────────────────────────┘
Position: top-right
Background: white
Border-left: 4px solid #10B981 (success)
Shadow: shadow-lg
Auto-dismiss: 5 seconds
```

**Types:**
- Success: border-left #10B981
- Error: border-left #EF4444
- Warning: border-left #F59E0B
- Info: border-left #3B82F6

#### 4.8 Loading States

**Skeleton Loading:**
```
┌─────────────────────────────────┐
│  ┌─────────────────────────┐   │
│  │░░░░░░░░░░░░░░░░░░░░░░░░░│   │ ← Image placeholder
│  └─────────────────────────┘   │
│                                 │
│  ┌──────────────────────┐      │
│  │░░░░░░░░░░░░░░░░░░░░░░│      │ ← Title placeholder
│  └──────────────────────┘      │
│                                 │
│  ┌───────────────────────────┐ │
│  │░░░░░░░░░░░░░░░░░░░░░░░░░░░│ │ ← Text placeholder
│  └───────────────────────────┘ │
│  ┌───────────────────────────┐ │
│  │░░░░░░░░░░░░░░░░░░░░░░░░░░░│ │
│  └───────────────────────────┘ │
└─────────────────────────────────┘
```

**Button Loading:**
```
┌─────────────────────────┐
│  ⏳ Menyimpan...        │
└─────────────────────────┘
Disabled state + spinner icon
```

---

### 5. PAGE DESIGNS — PUBLIC PAGES

#### 5.1 Home Page

**Layout Overview:**
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │              HERO SECTION                               │   │
│  │                                                         │   │
│  │     We Experiment, You Grow                            │   │ ← Display XL
│  │                                                         │   │
│  │     Kami adalah lab digital yang membantu UMKM         │   │ ← Body LG
│  │     dan Startup tumbuh melalui eksperimen web,         │   │
│  │     mobile, dan AI.                                    │   │
│  │                                                         │   │
│  │     [Mulai Eksperimen]  [Lihat Portfolio]              │   │ ← CTAs
│  │                                                         │   │
│  │     ────────────────────────────────────────────       │   │
│  │                                                         │   │
│  │     🏢 10+ UMKM dibantu  📱 5+ App dibangun           │   │ ← Stats
│  │     🤖 3+ AI diintegrasikan                            │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     LAYANAN UNGGULAN                                    │   │ ← Section Heading
│  │                                                         │   │
│  │     Apa yang bisa kami bantu?                          │   │ ← Subheading
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │
│  │  │  🌐      │  │  📱      │  │  🧠      │             │   │ ← Service Cards
│  │  │  Web Dev │  │  Mobile  │  │  AI      │             │   │    (GET /services)
│  │  │          │  │  App     │  │  Integr. │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  │              [Lihat Semua Layanan →]                    │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     KENAPA WEBIFYLAB?                                   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🧪 Kami Eksperimen, Bukan Janji               │   │   │
│  │  │  Setiap project adalah eksperimen. Kami         │   │   │
│  │  │  dokumentasikan proses, kegagalan, dan          │   │   │
│  │  │  pembelajaran secara transparan.                │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🤖 AI-Enhanced Development                    │   │   │
│  │  │  Integrasikan SPK, Sistem Pakar, dan Data       │   │   │
│  │  │  Analytics untuk keputusan bisnis lebih cerdas. │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🤝 Partner, Bukan Sekadar Vendor              │   │   │
│  │  │  Kami belajar bareng kamu. Setiap teknologi     │   │   │
│  │  │  yang dipakai, kami jelaskan dengan bahasa      │   │   │
│  │  │  yang kamu pahami.                              │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     EKSPERIMEN TERBARU                                  │   │
│  │                                                         │   │
│  │     Lihat apa yang sedang kami kerjakan                │   │
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │
│  │  │  Case    │  │  Case    │  │  Case    │             │   │ ← Case Study Cards
│  │  │  Study 1 │  │  Study 2 │  │  Study 3 │             │   │    (GET /case-studies?limit=3)
│  │  │          │  │          │  │          │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  │              [Lihat Semua Case Study →]                 │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     CATATAN BELAJAR                                     │   │
│  │                                                         │   │
│  │     Sharing yang kami pelajari dari setiap eksperimen  │   │
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │
│  │  │  Blog 1  │  │  Blog 2  │  │  Blog 3  │             │   │ ← Blog Cards
│  │  │          │  │          │  │          │             │   │    (GET /blog?limit=3)
│  │  │          │  │          │  │          │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  │              [Lihat Semua Artikel →]                    │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     CLIENT LOGOS / TESTIMONIALS                         │   │
│  │                                                         │   │
│  │     ┌────┐  ┌────┐  ┌────┐  ┌────┐  ┌────┐            │   │
│  │     │ C1 │  │ C2 │  │ C3 │  │ C4 │  │ C5 │            │   │ ← Logo Grid
│  │     └────┘  └────┘  └────┘  └────┘  └────┘            │   │
│  │                                                         │   │
│  │     "WebifyLab membantu kami..."                        │   │ ← Testimonial
│  │     — Budi, Pemilik Toko Sejahtera                     │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     SIAP EKSPERIMEN BARENG?                             │   │ ← CTA Section
│  │                                                         │   │
│  │     Ceritakan project kamu, mari kita wujudkan         │   │
│  │     bersama. Konsultasi awal gratis.                   │   │
│  │                                                         │   │
│  │     [Mulai Konsultasi]                                  │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

**Data Mapping:**
| Section | API Endpoint |
|---|---|
| Hero | Static content |
| Services | `GET /services` |
| Why WebifyLab | Static content |
| Case Studies | `GET /case-studies?limit=3` |
| Blog | `GET /blog?limit=3` |
| Testimonials | Static or from case studies |
| Settings | `GET /settings` |

**Responsive Behavior:**
- Hero: Stack vertically on mobile, side-by-side on desktop
- Services: 1 column mobile, 2 columns tablet, 3 columns desktop
- Case Studies: 1 column mobile, 2 columns tablet, 3 columns desktop
- Blog: 1 column mobile, 2 columns tablet, 3 columns desktop

---

#### 5.2 About Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     TENTANG WEBIFYLAB                                   │   │
│  │                                                         │   │
│  │     Kami adalah lab digital yang percaya bahwa          │   │
│  │     eksperimen adalah cara terbaik untuk tumbuh.       │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     FILOSOFI KAMI                                       │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🧪 Transparan                                 │   │   │
│  │  │  Kami dokumentasikan proses, termasuk           │   │   │
│  │  │  kegagalan dan pembelajaran.                    │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🤝 Kolaboratif                                │   │   │
│  │  │  Kamu bukan klien, kamu partner. Kami belajar   │   │   │
│  │  │  bareng di setiap project.                      │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🔄 Iteratif                                   │   │   │
│  │  │  MVP dulu, lalu improve berdasarkan data        │   │   │
│  │  │  dan feedback.                                  │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  📚 Edukatif                                   │   │   │
│  │  │  Setiap konten bertujuan membagi ilmu,          │   │   │
│  │  │  bukan sekadar jualan.                          │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     TIM KAMI                                            │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  👤 Rizal                                      │   │   │
│  │  │  Founder & Fullstack Developer                 │   │   │
│  │  │                                                 │   │   │
│  │  │  System Analyst, Software Engineer, AI &        │   │   │
│  │  │  Data Analyst Enthusiast                       │   │   │
│  │  │                                                 │   │   │
│  │  │  Sedang belajar: Golang advanced, ML           │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     TECH STACK KAMI                                     │   │
│  │                                                         │   │
│  │     ┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐ ┌────┐        │   │
│  │     │ ⚛️ │ │ 🐹 │ │ 🐘 │ │ 🐳 │ │ 🎨 │ │ 📧 │        │   │ ← Tech badges
│  │     │Next│ │Go  │ │PG  │ │Docker│ │Tail│ │Resend│      │   │
│  │     └────┘ └────┘ └────┘ └────┘ └────┘ └────┘        │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [CTA Section - Same as Home]                                   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.3 Services Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     LAYANAN KAMI                                        │   │
│  │                                                         │   │
│  │     Dari konsep hingga deployment, kami bantu kamu      │   │
│  │     di setiap tahap.                                    │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🌐 Web Development                             │   │   │ ← Service Card 1
│  │  │                                                 │   │   │    (GET /services/:slug)
│  │  │  Pembuatan website profesional untuk UMKM      │   │   │
│  │  │  dan Startup.                                  │   │   │
│  │  │                                                 │   │   │
│  │  │  Use Cases:                                    │   │   │
│  │  │  • Company profile                             │   │   │
│  │  │  • Landing page                                │   │   │
│  │  │  • Web app                                     │   │   │
│  │  │  • E-commerce                                  │   │   │
│  │  │                                                 │   │   │
│  │  │  ⏱️ 2-6 minggu  🛠️ Next.js, Golang, PG        │   │   │
│  │  │                                                 │   │   │
│  │  │  💡 Yang akan kita pelajari bareng:            │   │   │
│  │  │  Bagaimana website bisa jadi motor             │   │   │
│  │  │  pertumbuhan bisnis kamu.                      │   │   │
│  │  │                                                 │   │   │
│  │  │  [Diskusikan Eksperimen Ini →]                 │   │   │ ← CTA → Contact
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │  📱 Mobile App Development                      │   │   │ ← Service Card 2
│  │  │  ...                                            │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │  🎨 UI/UX Design                               │   │   │ ← Service Card 3
│  │  │  ...                                            │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │  🧠 AI Integration                             │   │   │ ← Service Card 4
│  │  │  ...                                            │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     BAGAIMANA KAMI BEKERJA                              │   │
│  │                                                         │   │
│  │     ┌────┐    ┌────┐    ┌────┐    ┌────┐    ┌────┐    │   │
│  │     │ 1  │ →  │ 2  │ →  │ 3  │ →  │ 4  │ →  │ 5  │    │   │ ← Process Steps
│  │     │Disc│    │Des │    │Dev │    │Test│    │Dep │    │   │
│  │     │overy│   │ign │    │elop│    │ing │    │loy │    │   │
│  │     └────┘    └────┘    └────┘    └────┘    └────┘    │   │
│  │                                                         │   │
│  │     + Learn & Iterate (ongoing)                        │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [CTA Section]                                                  │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.4 Portfolio Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     PORTFOLIO                                           │   │
│  │                                                         │   │
│  │     Eksperimen yang sudah kami wujudkan                 │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  Filter: [Semua] [Web] [Mobile] [AI] [E-commerce]      │   │ ← Filter Tabs
│  │                                                         │   │
│  │  ───────────────────────────────────────────────────   │   │
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │
│  │  │  Port 1  │  │  Port 2  │  │  Port 3  │             │   │ ← Portfolio Cards
│  │  │          │  │          │  │          │             │   │    (GET /portfolio)
│  │  │          │  │          │  │          │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │
│  │  │  Port 4  │  │  Port 5  │  │  Port 6  │             │   │
│  │  │          │  │          │  │          │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  │  ───────────────────────────────────────────────────   │   │
│  │                                                         │   │
│  │  [Load More] atau Pagination                            │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [CTA Section]                                                  │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.5 Portfolio Detail Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  ← Kembali ke Portfolio                                 │   │ ← Back link
│  │                                                         │   │
│  │     E-Commerce Platform untuk UMKM                      │   │ ← Title (Display LG)
│  │                                                         │   │
│  │     [E-commerce]  [Next.js]  [Golang]  [PostgreSQL]    │   │ ← Category + Tech Stack badges
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │              Hero Image                         │   │   │ ← Thumbnail
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │     🔗 Live: https://example.com                       │   │ ← Live URL
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     OVERVIEW                                            │   │
│  │                                                         │   │
│  │     Platform e-commerce lengkap untuk membantu         │   │
│  │     UMKM go digital...                                 │   │
│  │                                                         │   │
│  │     ───────────────────────────────────────────────    │   │
│  │                                                         │   │
│  │     GALLERY                                             │   │
│  │                                                         │   │
│  │     ┌────┐  ┌────┐  ┌────┐                            │   │ ← Image Gallery
│  │     │ 1  │  │ 2  │  │ 3  │                            │   │
│  │     └────┘  └────┘  └────┘                            │   │
│  │                                                         │   │
│  │     ───────────────────────────────────────────────    │   │
│  │                                                         │   │
│  │     💡 LESSONS LEARNED                                  │   │
│  │                                                         │   │
│  │     Pelajaran utama dalam project ini adalah           │   │
│  │     pentingnya memahami kebutuhan user...              │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     PROJECT SERUPA                                      │   │
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │ ← Related Portfolios
│  │  │  Rel 1   │  │  Rel 2   │  │  Rel 3   │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [CTA Section]                                                  │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.6 Case Studies Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     CASE STUDIES                                        │   │
│  │                                                         │   │
│  │     Dokumentasi eksperimen kami bersama klien          │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  Filter by Industry: [Semua] [Retail] [F&B] [Tech]     │   │ ← Industry Filter
│  │                                                         │   │
│  │  Filter by AI: [Semua] [SPK] [Sistem Pakar] [ML]       │   │ ← AI Type Filter
│  │                                                         │   │
│  │  ───────────────────────────────────────────────────   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  ┌─────────────────────────────────────────┐   │   │   │ ← Case Study Card 1
│  │  │  │  ┌─────────────┐                        │   │   │   │
│  │  │  │  │             │  Optimasi Stok UMKM   │   │   │   │
│  │  │  │  │  Cover Img  │  dengan SPK           │   │   │   │
│  │  │  │  │             │                        │   │   │   │
│  │  │  │  └─────────────┘  🏢 Toko Sejahtera    │   │   │   │
│  │  │  │                    🏭 Retail            │   │   │   │
│  │  │  │                                         │   │   │   │
│  │  │  │  [SPK] [Data Analytics]                │   │   │   │ ← AI Type badges
│  │  │  │                                         │   │   │   │
│  │  │  │  📊 Hasil:                             │   │   │   │
│  │  │  │  • Stock accuracy: 95%                 │   │   │   │ ← Key Metrics
│  │  │  │  • Waste reduction: 40%                │   │   │   │
│  │  │  │                                         │   │   │   │
│  │  │  │  [Baca Case Study →]                   │   │   │   │ ← CTA
│  │  │  │                                         │   │   │   │
│  │  │  └─────────────────────────────────────────┘   │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │  Case Study Card 2                              │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [CTA Section]                                                  │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.7 Case Study Detail Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  ← Kembali ke Case Studies                              │   │
│  │                                                         │   │
│  │     Optimasi Stok UMKM dengan SPK                       │   │ ← Title (Display LG)
│  │                                                         │   │
│  │     🏢 Toko Sejahtera  •  🏭 Retail  •  📅 Aug 2026   │   │ ← Meta
│  │                                                         │   │
│  │     [SPK] [Data Analytics]                             │   │ ← AI Type badges
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🎯 CHALLENGE                                   │   │   │ ← Section 1
│  │  │                                                 │   │   │
│  │  │  Toko Sejahtera mengalami kesulitan mengelola   │   │   │
│  │  │  stok barang...                                 │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🧭 APPROACH                                    │   │   │ ← Section 2
│  │  │                                                 │   │   │
│  │  │  Kami menggunakan metode SPK dengan...          │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  💡 SOLUTION                                    │   │   │ ← Section 3
│  │  │                                                 │   │   │
│  │  │  Sistem yang kami bangun menganalisis...        │   │   │
│  │  │                                                 │   │   │
│  │  │  ┌────┐  ┌────┐  ┌────┐                       │   │   │ ← Solution Images
│  │  │  │ 1  │  │ 2  │  │ 3  │                       │   │   │
│  │  │  └────┘  └────┘  └────┘                       │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  📊 RESULTS                                     │   │   │ ← Section 4
│  │  │                                                 │   │   │
│  │  │  Setelah 3 bulan implementasi:                 │   │   │
│  │  │                                                 │   │   │
│  │  │  ┌────────────┐  ┌────────────┐               │   │   │ ← Metric Cards
│  │  │  │    95%     │  │    40%     │               │   │   │
│  │  │  │  Stock     │  │  Waste     │               │   │   │
│  │  │  │  Accuracy  │  │  Reduction │               │   │   │
│  │  │  └────────────┘  └────────────┘               │   │   │
│  │  │                                                 │   │   │
│  │  │  ┌────────────┐                                │   │   │
│  │  │  │  10 hrs    │                                │   │   │
│  │  │  │  Time      │                                │   │   │
│  │  │  │  Saved/wk  │                                │   │   │
│  │  │  └────────────┘                                │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  📚 LESSONS LEARNED                             │   │   │ ← Section 5
│  │  │                                                 │   │   │
│  │  │  Pelajaran utama adalah pentingnya...           │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  💬 TESTIMONIAL                                 │   │   │ ← Section 6
│  │  │                                                 │   │   │
│  │  │  "WebifyLab membantu kami..."                  │   │   │
│  │  │                                                 │   │   │
│  │  │  — Budi, Pemilik Toko Sejahtera                │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     INGIN EKSPERIMEN SERUPA?                            │   │ ← CTA
│  │                                                         │   │
│  │     [Diskusikan Project Kamu →]                         │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.8 Blog Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     CATATAN BELAJAR                                     │   │
│  │                                                         │   │
│  │     Sharing yang kami pelajari dari setiap eksperimen  │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  🔍 [Search artikel...]                                 │   │ ← Search Bar
│  │                                                         │   │
│  │  Kategori: [Semua] [Tutorial] [Business] [AI & Data]   │   │ ← Category Filter
│  │            [Eksperimen] [Lessons Learned]              │   │
│  │                                                         │   │
│  │  ───────────────────────────────────────────────────   │   │
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │
│  │  │  Blog 1  │  │  Blog 2  │  │  Blog 3  │             │   │ ← Blog Cards Grid
│  │  │          │  │          │  │          │             │   │    (GET /blog)
│  │  │          │  │          │  │          │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │
│  │  │  Blog 4  │  │  Blog 5  │  │  Blog 6  │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  │  ───────────────────────────────────────────────────   │   │
│  │                                                         │   │
│  │  Pagination: [← Prev]  1  2  3  [Next →]               │   │ ← Pagination
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.9 Blog Detail Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  ← Kembali ke Blog                                      │   │
│  │                                                         │   │
│  │     Membuat REST API dengan Golang dan Gin              │   │ ← Title (Display LG)
│  │                                                         │   │
│  │     👤 Rizal  •  📅 Sep 5, 2026  •  ⏱️ 8 min read     │   │ ← Meta
│  │                                                         │   │
│  │     [Tutorial]  [Golang]  [Next.js]                    │   │ ← Category + Tags
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │              Cover Image                        │   │   │ ← Cover Image
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  ┌─────────────┐  ┌──────────────────────────────┐    │   │
│  │  │             │  │                              │    │   │
│  │  │  Table of   │  │  Article Content             │    │   │ ← 2-column layout
│  │  │  Contents   │  │                              │    │   │    (desktop only)
│  │  │             │  │  ## Pendahuluan              │    │   │
│  │  │  1. Intro   │  │                              │    │   │
│  │  │  2. Setup   │  │  Lorem ipsum dolor sit       │    │   │
│  │  │  3. Routes  │  │  amet...                     │    │   │
│  │  │  4. Deploy  │  │                              │    │   │
│  │  │             │  │  ## Setup                    │    │   │
│  │  │  (sticky)   │  │                              │    │   │
│  │  │             │  │  ```go                       │    │   │ ← Code blocks
│  │  │             │  │  func main() {               │    │   │    with syntax
│  │  │             │  │      // code                 │    │   │    highlighting
│  │  │             │  │  }                           │    │   │
│  │  │             │  │  ```                         │    │   │
│  │  │             │  │                              │    │   │
│  │  └─────────────┘  └──────────────────────────────┘    │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     📚 ARTIKEL TERKAIT                                  │   │
│  │                                                         │   │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐             │   │ ← Related Articles
│  │  │  Rel 1   │  │  Rel 2   │  │  Rel 3   │             │   │
│  │  └──────────┘  └──────────┘  └──────────┘             │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     📤 SHARE ARTIKEL INI                                │   │
│  │                                                         │   │
│  │     [LinkedIn]  [Twitter/X]  [WhatsApp]  [Copy Link]   │   │ ← Share Buttons
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.10 Contact Page

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     MULAI EKSPERIMEN                                    │   │
│  │                                                         │   │
│  │     Ceritakan project kamu, mari kita wujudkan         │   │
│  │     bersama. Konsultasi awal gratis.                   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  📋 Form Inquiry                                │   │   │ ← Contact Form
│  │  │                                                 │   │   │    (POST /leads)
│  │  │  Nama Lengkap *                                 │   │   │
│  │  │  ┌─────────────────────────────────────────┐   │   │   │
│  │  │  │                                         │   │   │   │
│  │  │  └─────────────────────────────────────────┘   │   │   │
│  │  │                                                 │   │   │
│  │  │  Email *                                        │   │   │
│  │  │  ┌─────────────────────────────────────────┐   │   │   │
│  │  │  │                                         │   │   │   │
│  │  │  └─────────────────────────────────────────┘   │   │   │
│  │  │                                                 │   │   │
│  │  │  No. WhatsApp *                                 │   │   │
│  │  │  ┌─────────────────────────────────────────┐   │   │   │
│  │  │  │                                         │   │   │   │
│  │  │  └─────────────────────────────────────────┘   │   │   │
│  │  │                                                 │   │   │
│  │  │  Jenis Layanan *                                │   │   │
│  │  │  ┌─────────────────────────────────────────┐   │   │   │
│  │  │  │  [Web Development              ▼]      │   │   │   │
│  │  │  └─────────────────────────────────────────┘   │   │   │
│  │  │                                                 │   │   │
│  │  │  Budget Range *                                 │   │   │
│  │  │  ┌─────────────────────────────────────────┐   │   │   │
│  │  │  │  [5-15jt                       ▼]      │   │   │   │
│  │  │  └─────────────────────────────────────────┘   │   │   │
│  │  │                                                 │   │   │
│  │  │  Deskripsi Project *                            │   │   │
│  │  │  ┌─────────────────────────────────────────┐   │   │   │
│  │  │  │                                         │   │   │   │
│  │  │  │                                         │   │   │   │
│  │  │  │                                         │   │   │   │
│  │  │  └─────────────────────────────────────────┘   │   │   │
│  │  │                                                 │   │   │
│  │  │  [Kirim Inquiry]                                │   │   │ ← Submit Button
│  │  │                                                 │   │   │
│  │  │  ───────────────────────────────────────────   │   │   │
│  │  │                                                 │   │   │
│  │  │  ✅ Success message:                            │   │   │ ← Success State
│  │  │  "Inquiry submitted! Kami akan menghubungi     │   │   │
│  │  │  kamu via WhatsApp dalam 24 jam."              │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  📞 Kontak Langsung                             │   │   │ ← Contact Info
│  │  │                                                 │   │   │    (GET /settings)
│  │  │  📧 hello@webifylab.my.id                      │   │   │
│  │  │  📱 +62 812-3456-7890                          │   │   │
│  │  │                                                 │   │   │
│  │  │  🌐 Social Media:                              │   │   │
│  │  │  [Instagram] [LinkedIn] [GitHub]               │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

#### 5.11 Products Page (Placeholder)

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Header Navigation]                                            │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     PRODUK & EKOSISTEM                                  │   │
│  │                                                         │   │
│  │     SaaS yang lahir dari eksperimen kami               │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🚀 SaaS Product 1                             │   │   │ ← Product Card 1
│  │  │                                                 │   │   │
│  │  │  Deskripsi singkat produk...                   │   │   │
│  │  │                                                 │   │   │
│  │  │  Status: [Live]                                │   │   │ ← Status Badge
│  │  │                                                 │   │   │
│  │  │  [Kunjungi Produk →]                           │   │   │ ← CTA to subdomain
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  │  ┌─────────────────────────────────────────────────┐   │   │
│  │  │                                                 │   │   │
│  │  │  🧪 SaaS Product 2                             │   │   │ ← Product Card 2
│  │  │                                                 │   │   │
│  │  │  Deskripsi singkat produk...                   │   │   │
│  │  │                                                 │   │   │
│  │  │  Status: [Experimenting]                       │   │   │ ← Status Badge
│  │  │                                                 │   │   │
│  │  │  Coming Soon                                   │   │   │
│  │  │                                                 │   │   │
│  │  └─────────────────────────────────────────────────┘   │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                                                         │   │
│  │     EKSPERIMEN SELANJUTNYA                              │   │ ← Coming Soon
│  │                                                         │   │
│  │     Kami sedang mengeksplorasi ide-ide baru...         │   │
│  │                                                         │   │
│  └─────────────────────────────────────────────────────────┘   │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  [Footer]                                                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

### 6. PAGE DESIGNS — ADMIN DASHBOARD

#### 6.1 Admin Layout

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  ┌──────────┐  ┌──────────────────────────────────────────┐   │
│  │          │  │                                          │   │
│  │  🔬      │  │  [Admin Header]                          │   │
│  │  Webify  │  │                                          │   │
│  │  Lab     │  │  🔍 Search...     👤 Rizal  [Logout]   │   │
│  │          │  │                                          │   │
│  │  ──────  │  ├──────────────────────────────────────────┤   │
│  │          │  │                                          │   │
│  │  📊      │  │  [Content Area]                          │   │
│  │  Dash    │  │                                          │   │
│  │          │  │                                          │   │
│  │  📝      │  │                                          │   │
│  │  Blog    │  │                                          │   │
│  │          │  │                                          │   │
│  │  📋      │  │                                          │   │
│  │  Case    │  │                                          │   │
│  │  Studies │  │                                          │   │
│  │          │  │                                          │   │
│  │  💼      │  │                                          │   │
│  │  Port    │  │                                          │   │
│  │  folio   │  │                                          │   │
│  │          │  │                                          │   │
│  │  📬      │  │                                          │   │
│  │  Leads   │  │                                          │   │
│  │          │  │                                          │   │
│  │  🖼️      │  │                                          │   │
│  │  Media   │  │                                          │   │
│  │          │  │                                          │   │
│  │  ⚙️      │  │                                          │   │
│  │  Settings│  │                                          │   │
│  │          │  │                                          │   │
│  └──────────┘  └──────────────────────────────────────────┘   │
│                                                                 │
│  Sidebar: 256px wide, collapsible                              │
│  Background: #0F172A (Deep Slate)                              │
│  Text: white                                                   │
│  Active item: bg #2563EB, text white                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

#### 6.2 Admin Dashboard

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Sidebar]  │  Dashboard                                        │
│             │                                                   │
│             │  ┌──────────────────────────────────────────┐   │
│             │  │                                          │   │
│             │  │  ┌────┐  ┌────┐  ┌────┐  ┌────┐        │   │ ← Stats Cards
│             │  │  │ 25 │  │ 12 │  │ 45 │  │ 150│        │   │
│             │  │  │Blog│  │Case│  │Lead│  │Med │        │   │
│             │  │  │    │  │    │  │    │  │    │        │   │
│             │  │  └────┘  └────┘  └────┘  └────┘        │   │
│             │  │                                          │   │
│             │  └──────────────────────────────────────────┘   │
│             │                                                   │
│             │  ┌──────────────────────────────────────────┐   │
│             │  │                                          │   │
│             │  │  📈 Leads per Month                      │   │ ← Chart
│             │  │                                          │   │    (GET /dashboard/leads-chart)
│             │  │  ┌──────────────────────────────────┐   │   │
│             │  │  │                                  │   │   │
│             │  │  │  ▁ ▂ ▃ ▄ ▅ ▆ ▇ █ ▇ ▆ ▅ ▄      │   │   │
│             │  │  │                                  │   │   │
│             │  │  └──────────────────────────────────┘   │   │
│             │  │                                          │   │
│             │  └──────────────────────────────────────────┘   │
│             │                                                   │
│             │  ┌─────────────────────┐  ┌──────────────────┐ │
│             │  │                     │  │                  │ │
│             │  │  📬 Recent Leads    │  │  📝 Recent Posts │ │ ← Recent Lists
│             │  │                     │  │                  │ │
│             │  │  • Budi Santoso     │  │  • Blog Post 1  │ │
│             │  │  • Sari Dewi        │  │  • Blog Post 2  │ │
│             │  │  • Andi Pratama     │  │  • Blog Post 3  │ │
│             │  │                     │  │                  │ │
│             │  │  [Lihat Semua →]    │  │  [Lihat Semua →] │ │
│             │  │                     │  │                  │ │
│             │  └─────────────────────┘  └──────────────────┘ │
│             │                                                   │
└─────────────────────────────────────────────────────────────────┘
```

#### 6.3 Blog Management

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Sidebar]  │  Blog Management                                  │
│             │                                                   │
│             │  ┌──────────────────────────────────────────┐   │
│             │  │                                          │   │
│             │  │  [+ Tulis Artikel Baru]                  │   │ ← Create Button
│             │  │                                          │   │
│             │  │  ─────────────────────────────────────  │   │
│             │  │                                          │   │
│             │  │  Filter: [Semua] [Published] [Draft]    │   │ ← Status Filter
│             │  │                                          │   │
│             │  │  ┌──────────────────────────────────┐   │   │
│             │  │  │                                  │   │   │ ← Blog Table
│             │  │  │ Title        │ Status  │ Date    │   │   │
│             │  │  │──────────────│─────────│─────────│   │   │
│             │  │  │ REST API...  │Published│Sep 5   │   │   │
│             │  │  │              │         │         │   │   │
│             │  │  │ [Edit] [Delete]                  │   │   │ ← Actions
│             │  │  │                                  │   │   │
│             │  │  │──────────────│─────────│─────────│   │   │
│             │  │  │ Docker...    │ Draft   │ Sep 4   │   │   │
│             │  │  │              │         │         │   │   │
│             │  │  │ [Edit] [Delete]                  │   │   │
│             │  │  │                                  │   │   │
│             │  │  └──────────────────────────────────┘   │   │
│             │  │                                          │   │
│             │  │  [← Prev]  1  2  3  [Next →]            │   │ ← Pagination
│             │  │                                          │   │
│             │  └──────────────────────────────────────────┘   │
│             │                                                   │
└─────────────────────────────────────────────────────────────────┘
```

#### 6.4 Lead Management

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  [Sidebar]  │  Lead Management                                  │
│             │                                                   │
│             │  ┌──────────────────────────────────────────┐   │
│             │  │                                          │   │
│             │  │  Filter: [Semua] [New] [Contacted]      │   │ ← Status Filter
│             │  │          [Negotiation] [Won] [Lost]     │   │
│             │  │                                          │   │
│             │  │  [Export CSV]                            │   │ ← Export Button
│             │  │                                          │   │
│             │  │  ┌──────────────────────────────────┐   │   │
│             │  │  │                                  │   │   │ ← Leads Table
│             │  │  │ Name  │ Email │ Service │ Status│   │   │
│             │  │  │───────│───────│─────────│───────│   │   │
│             │  │  │ Budi  │ b@... │ Web Dev │ New   │   │   │
│             │  │  │       │       │         │       │   │   │
│             │  │  │ [View] [📋 Copy WA]             │   │   │ ← Actions
│             │  │  │                                  │   │   │
│             │  │  │──────────────│─────────│─────────│   │   │
│             │  │  │ Sari  │ s@... │ AI Int  │ Contact│   │   │
│             │  │  │       │       │         │       │   │   │
│             │  │  │ [View] [📋 Copy WA]             │   │   │
│             │  │  │                                  │   │   │
│             │  │  └──────────────────────────────────┘   │   │
│             │  │                                          │   │
│             │  └──────────────────────────────────────────┘   │
│             │                                                   │
│             │  ┌──────────────────────────────────────────┐   │
│             │  │                                          │   │
│             │  │  📋 Lead Detail Modal                    │   │ ← Detail Modal
│             │  │                                          │   │
│             │  │  Name: Budi Santoso                     │   │
│             │  │  Email: budi@example.com                │   │
│             │  │  WhatsApp: +62 812-3456-7890 [📋]      │   │
│             │  │  Service: Web Development               │   │
│             │  │  Budget: 5-15jt                         │   │
│             │  │                                          │   │
│             │  │  Description:                           │   │
│             │  │  "Saya ingin membuat website..."        │   │
│             │  │                                          │   │
│             │  │  Status: [New ▼]                        │   │ ← Status Update
│             │  │                                          │   │
│             │  │  Notes:                                 │   │
│             │  │  ┌──────────────────────────────────┐   │   │
│             │  │  │                                  │   │   │
│             │  │  └──────────────────────────────────┘   │   │
│             │  │                                          │   │
│             │  │  [Save] [Close]                         │   │
│             │  │                                          │   │
│             │  └──────────────────────────────────────────┘   │
│             │                                                   │
└─────────────────────────────────────────────────────────────────┘
```

---

### 7. RESPONSIVE DESIGN

#### 7.1 Breakpoint Strategy

| Component | Mobile (< 768px) | Tablet (768-1024px) | Desktop (> 1024px) |
|---|---|---|---|
| **Header** | Hamburger menu | Hamburger menu | Full nav |
| **Hero** | Stack vertical | Stack vertical | Side-by-side |
| **Services Grid** | 1 column | 2 columns | 3-4 columns |
| **Portfolio Grid** | 1 column | 2 columns | 3 columns |
| **Blog Grid** | 1 column | 2 columns | 3 columns |
| **Case Studies** | 1 column | 1 column | 2 columns |
| **Blog Detail** | Single column | Single column | TOC + Content |
| **Contact Form** | Full width | Max-w-md | Max-w-lg |
| **Admin Sidebar** | Hidden, hamburger | Collapsible | Fixed |

#### 7.2 Mobile-Specific Patterns

**Sticky CTA on Mobile:**
```
┌─────────────────────────────────┐
│                                 │
│  [Content...]                   │
│                                 │
│  ───────────────────────────   │
│                                 │
│  ┌─────────────────────────┐   │
│  │  [Mulai Eksperimen]     │   │ ← Sticky bottom CTA
│  └─────────────────────────┘   │
│                                 │
└─────────────────────────────────┘
```

**Mobile Navigation Drawer:**
- Slide-in from right
- Backdrop overlay with blur
- Close on backdrop click or ESC

---

### 8. INTERACTION DESIGN

#### 8.1 Animations

| Element | Animation | Duration | Easing |
|---|---|---|---|
| Page load | Fade in | 300ms | ease-out |
| Scroll reveal | Fade up + slide | 500ms | ease-out |
| Button hover | Scale 1.02 | 150ms | ease-in-out |
| Card hover | Shadow elevation | 200ms | ease-in-out |
| Modal open | Scale 0.95→1 + fade | 200ms | ease-out |
| Modal close | Scale 1→0.95 + fade | 150ms | ease-in |
| Toast | Slide in from right | 300ms | ease-out |
| Skeleton | Pulse | 1.5s | ease-in-out |

#### 8.2 Scroll Behaviors

- **Smooth scroll** untuk anchor links
- **Sticky header** dengan backdrop blur saat scroll
- **Scroll reveal** untuk sections (Framer Motion `whileInView`)
- **Reading progress bar** untuk blog detail

#### 8.3 Form Interactions

- **Inline validation** on blur
- **Real-time validation** untuk email dan phone
- **Loading state** pada submit button
- **Success/error toast** setelah submit
- **Auto-save** untuk blog editor (draft)

---

### 9. ACCESSIBILITY GUIDELINES

#### 9.1 Color Contrast

| Element | Foreground | Background | Ratio |
|---|---|---|---|
| Body text | #0F172A | #FFFFFF | 15.4:1 ✅ |
| Secondary text | #64748B | #FFFFFF | 4.6:1 ✅ |
| Primary button | #FFFFFF | #2563EB | 4.6:1 ✅ |
| Link | #2563EB | #FFFFFF | 4.6:1 ✅ |
| Error text | #EF4444 | #FFFFFF | 4.5:1 ✅ |

#### 9.2 Keyboard Navigation

- Semua interactive elements harus focusable
- Visible focus ring: `outline: 2px solid #2563EB`
- Tab order logical (follow DOM order)
- Modal traps focus, returns focus on close

#### 9.3 ARIA Labels

```html
<!-- Navigation -->
<nav aria-label="Main navigation">...</nav>

<!-- Buttons -->
<button aria-label="Close menu">✕</button>

<!-- Images -->
<img src="..." alt="Screenshot of e-commerce dashboard" />

<!-- Form fields -->
<label for="email">Email</label>
<input id="email" type="email" aria-required="true" />

<!-- Status messages -->
<div role="alert">Form berhasil dikirim</div>
```

#### 9.4 Form Accessibility

- Semua input punya `<label>` yang terhubung
- Error messages linked dengan `aria-describedby`
- Required fields marked with `aria-required`
- Success/error states announced with `role="status"`

---

### 10. ASSETS & RESOURCES

#### 10.1 Icons

**Library:** Lucide Icons (https://lucide.dev)

**Usage:**
```tsx
import { Globe, Smartphone, Brain, Palette } from 'lucide-react'

<Globe className="w-6 h-6 text-lab-blue" />
```

**Icon Mapping:**
| Service | Icon |
|---|---|
| Web Development | `Globe` |
| Mobile App | `Smartphone` |
| UI/UX Design | `Palette` |
| AI Integration | `Brain` |

#### 10.2 Fonts

```css
/* Google Fonts */
@import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Plus+Jakarta+Sans:wght@600;700&family=JetBrains+Mono:wght@400;500&display=swap');
```

#### 10.3 Images

| Type | Aspect Ratio | Usage |
|---|---|---|
| Blog cover | 16:9 | Blog cards, blog detail |
| Portfolio thumbnail | 16:9 | Portfolio cards |
| Case study cover | 16:9 | Case study cards |
| OG image | 1200x630 | Social sharing |

---

### 11. IMPLEMENTATION CHECKLIST

#### 11.1 Frontend Setup

- [ ] Install Next.js 15+ dengan App Router
- [ ] Install Tailwind CSS + shadcn/ui
- [ ] Install Framer Motion
- [ ] Install Lucide Icons
- [ ] Setup font optimization
- [ ] Setup color palette di tailwind.config.ts

#### 11.2 Public Pages

- [ ] Home page (all sections)
- [ ] About page
- [ ] Services page
- [ ] Portfolio page + detail
- [ ] Case Studies page + detail
- [ ] Blog page + detail
- [ ] Contact page
- [ ] Products page (placeholder)

#### 11.3 Admin Dashboard

- [ ] Admin layout (sidebar + header)
- [ ] Login page
- [ ] Dashboard overview
- [ ] Blog management (list, create, edit, delete)
- [ ] Case study management
- [ ] Portfolio management
- [ ] Lead management
- [ ] Media manager
- [ ] Settings page

#### 11.4 Components

- [ ] Button (primary, secondary, ghost, disabled)
- [ ] Card (portfolio, blog, service)
- [ ] Input (text, textarea, select)
- [ ] Badge (category, status, AI type)
- [ ] Modal
- [ ] Toast/Notification
- [ ] Skeleton loader
- [ ] Pagination
- [ ] Navigation (header, footer, mobile menu)

#### 11.5 Responsive & Accessibility

- [ ] Mobile breakpoints tested
- [ ] Keyboard navigation tested
- [ ] Screen reader tested
- [ ] Color contrast verified
- [ ] Alt text for all images

---

### 12. DESIGN DELIVERABLES SUMMARY

| Deliverable | Status | Location |
|---|---|---|
| Design System | ✅ Complete | Section 3 |
| Component Library | ✅ Complete | Section 4 |
| Public Page Wireframes | ✅ Complete | Section 5 |
| Admin Page Wireframes | ✅ Complete | Section 6 |
| Responsive Guidelines | ✅ Complete | Section 7 |
| Interaction Design | ✅ Complete | Section 8 |
| Accessibility Guidelines | ✅ Complete | Section 9 |
| Implementation Checklist | ✅ Complete | Section 11 |
