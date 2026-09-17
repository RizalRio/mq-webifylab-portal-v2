# 🎨 Design System Document

## Webifylab Landing Page — Version 1.0

| Metadata      | Detail                            |
| ------------- | --------------------------------- |
| **Product**   | Webifylab Landing Page            |
| **Version**   | 1.0                               |
| **Author**    | Rizal                             |
| **Created**   | 17 September 2026                 |
| **Status**    | Draft                             |
| **Framework** | Tailwind CSS                      |
| **Tools**     | Figma (optional), Tailwind config |

---

## 1. Design Principles

### 1.1 Core Values

| Principle          | Deskripsi                                | Implementasi                                            |
| ------------------ | ---------------------------------------- | ------------------------------------------------------- |
| **Clarity First**  | Informasi harus jelas dan mudah dipahami | Typography hierarkis, spacing generous, kontras tinggi  |
| **Modern & Clean** | Tampilan modern tanpa berlebihan         | Minimalis, banyak white space, tidak ramai              |
| **Tech-Forward**   | Menunjukkan inovasi tanpa intimidatif    | Gradient halus, ikon konsisten, animasi subtle          |
| **Accessible**     | Bisa diakses semua orang                 | Kontras WCAG AA, font size minimal 16px, touch-friendly |
| **Consistent**     | Setiap elemen terasa "Webifylab"         | Reusable components, design tokens, strict guidelines   |

### 1.2 Design Philosophy

```
✅ DO:
- Gunakan white space generously
- Fokus pada satu CTA per section
- Gunakan ikon untuk visual hierarchy
- Animasi subtle (fade-in, slide-up)
- Konsisten di semua breakpoint

❌ DON'T:
- Jangan gunakan terlalu banyak warna
- Jangan gunakan animasi yang mengganggu
- Jangan gunakan foto stock generik
- Jangan gunakan teks terlalu kecil
- Jangan gunakan border radius berbeda-beda
```

---

## 2. Color Palette

### 2.1 Primary Colors

| Name          | Hex       | RGB         | Usage                                          |
| ------------- | --------- | ----------- | ---------------------------------------------- |
| **Deep Blue** | `#1E3A5F` | 30, 58, 95  | Primary brand color, headings, primary buttons |
| **Indigo**    | `#4F46E5` | 79, 70, 229 | Accent color, links, interactive elements      |
| **Cyan**      | `#06B6D4` | 6, 182, 212 | Secondary accent, highlights, badges           |

### 2.2 Neutral Colors

| Name          | Hex       | RGB           | Usage                       |
| ------------- | --------- | ------------- | --------------------------- |
| **Slate 900** | `#0F172A` | 15, 23, 42    | Body text, dark backgrounds |
| **Slate 700** | `#334155` | 51, 65, 85    | Secondary text              |
| **Slate 500** | `#64748B` | 100, 116, 139 | Muted text, placeholders    |
| **Slate 300** | `#CBD5E1` | 203, 213, 225 | Borders, dividers           |
| **Slate 100** | `#F1F5F9` | 241, 245, 249 | Light backgrounds, cards    |
| **White**     | `#FFFFFF` | 255, 255, 255 | Backgrounds, text on dark   |

### 2.3 Semantic Colors

| Name        | Hex       | Usage                        |
| ----------- | --------- | ---------------------------- |
| **Success** | `#10B981` | Success messages, checkmarks |
| **Warning** | `#F59E0B` | Warning messages, alerts     |
| **Error**   | `#EF4444` | Error messages, validation   |
| **Info**    | `#3B82F6` | Informational messages       |

### 2.4 Color Usage Guidelines

```css
/* Tailwind CSS Configuration */
module.exports = {
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: '#1E3A5F',
          light: '#2D5A8C',
          dark: '#152A45',
        },
        accent: {
          indigo: '#4F46E5',
          cyan: '#06B6D4',
        },
        neutral: {
          900: '#0F172A',
          700: '#334155',
          500: '#64748B',
          300: '#CBD5E1',
          100: '#F1F5F9',
        },
        semantic: {
          success: '#10B981',
          warning: '#F59E0B',
          error: '#EF4444',
          info: '#3B82F6',
        },
      },
    },
  },
}
```

### 2.5 Color Combinations

| Combination               | Usage                     | Example                                |
| ------------------------- | ------------------------- | -------------------------------------- |
| **Deep Blue + White**     | Hero section, primary CTA | Background: Deep Blue, Text: White     |
| **White + Slate 900**     | Body content, cards       | Background: White, Text: Slate 900     |
| **Slate 100 + Slate 900** | Section backgrounds       | Background: Slate 100, Text: Slate 900 |
| **Indigo + White**        | Links, buttons            | Button: Indigo, Text: White            |
| **Cyan + Deep Blue**      | Badges, highlights        | Badge: Cyan, Text: Deep Blue           |

---

## 3. Typography

### 3.1 Font Family

| Type            | Font              | Fallback                             | Usage               |
| --------------- | ----------------- | ------------------------------------ | ------------------- |
| **Primary**     | Inter             | system-ui, -apple-system, sans-serif | All text            |
| **Alternative** | Plus Jakarta Sans | Inter, sans-serif                    | Headings (optional) |

### 3.2 Type Scale

| Level          | Size            | Line Height | Weight           | Usage                    |
| -------------- | --------------- | ----------- | ---------------- | ------------------------ |
| **H1**         | 48px (3rem)     | 1.2         | 800 (Extra Bold) | Hero headline            |
| **H2**         | 36px (2.25rem)  | 1.3         | 700 (Bold)       | Section titles           |
| **H3**         | 24px (1.5rem)   | 1.4         | 600 (Semibold)   | Card titles, subsections |
| **H4**         | 20px (1.25rem)  | 1.4         | 600 (Semibold)   | Small headings           |
| **Body Large** | 18px (1.125rem) | 1.6         | 400 (Regular)    | Lead paragraphs          |
| **Body**       | 16px (1rem)     | 1.6         | 400 (Regular)    | Main body text           |
| **Body Small** | 14px (0.875rem) | 1.5         | 400 (Regular)    | Secondary text, captions |
| **Caption**    | 12px (0.75rem)  | 1.4         | 400 (Regular)    | Labels, metadata         |

### 3.3 Typography Guidelines

```css
/* Tailwind CSS Configuration */
module.exports = {
  theme: {
    extend: {
      fontFamily: {
        sans: ['Inter', 'system-ui', '-apple-system', 'sans-serif'],
      },
      fontSize: {
        'h1': ['3rem', { lineHeight: '1.2', fontWeight: '800' }],
        'h2': ['2.25rem', { lineHeight: '1.3', fontWeight: '700' }],
        'h3': ['1.5rem', { lineHeight: '1.4', fontWeight: '600' }],
        'h4': ['1.25rem', { lineHeight: '1.4', fontWeight: '600' }],
        'body-lg': ['1.125rem', { lineHeight: '1.6', fontWeight: '400' }],
        'body': ['1rem', { lineHeight: '1.6', fontWeight: '400' }],
        'body-sm': ['0.875rem', { lineHeight: '1.5', fontWeight: '400' }],
        'caption': ['0.75rem', { lineHeight: '1.4', fontWeight: '400' }],
      },
    },
  },
}
```

### 3.4 Typography Rules

| Rule            | Do ✅                                   | Don't ❌                                   |
| --------------- | --------------------------------------- | ------------------------------------------ |
| **Hierarchy**   | Gunakan H1 → H2 → H3 secara berurutan   | Jangan skip heading level (H1 → H3)        |
| **Contrast**    | Headings: Slate 900, Body: Slate 700    | Jangan gunakan teks abu-abu terlalu terang |
| **Line Length** | Maksimal 65-75 karakter per baris       | Jangan biarkan paragraf terlalu panjang    |
| **Alignment**   | Left-align untuk body text              | Jangan center-align paragraf panjang       |
| **Font Weight** | Bold untuk headings, regular untuk body | Jangan gunakan terlalu banyak bold         |

---

## 4. Spacing System

### 4.1 Base Unit

**Base unit: 4px (0.25rem)**

Semua spacing harus kelipatan dari 4px untuk konsistensi visual.

### 4.2 Spacing Scale

| Token   | Value          | Usage                       |
| ------- | -------------- | --------------------------- |
| **xs**  | 4px (0.25rem)  | Icon spacing, tight gaps    |
| **sm**  | 8px (0.5rem)   | Small padding, icon margins |
| **md**  | 16px (1rem)    | Default padding, card gaps  |
| **lg**  | 24px (1.5rem)  | Section padding, large gaps |
| **xl**  | 32px (2rem)    | Component spacing           |
| **2xl** | 48px (3rem)    | Section margins             |
| **3xl** | 64px (4rem)    | Large section spacing       |
| **4xl** | 80px (5rem)    | Section padding vertical    |
| **5xl** | 120px (7.5rem) | Hero section padding        |

### 4.3 Spacing Guidelines

```css
/* Tailwind CSS Configuration */
module.exports = {
  theme: {
    extend: {
      spacing: {
        '18': '4.5rem',   // 72px
        '88': '22rem',    // 352px
        '128': '32rem',   // 512px
      },
    },
  },
}
```

### 4.4 Spacing Rules

| Context                          | Spacing                              | Example              |
| -------------------------------- | ------------------------------------ | -------------------- |
| **Section padding (vertical)**   | 80-120px                             | `py-20 md:py-32`     |
| **Section padding (horizontal)** | 16-24px                              | `px-4 md:px-6`       |
| **Card padding**                 | 24-32px                              | `p-6 md:p-8`         |
| **Button padding**               | 12-16px vertical, 24-32px horizontal | `px-6 py-3`          |
| **Icon spacing**                 | 8-12px                               | `gap-2` atau `gap-3` |
| **Text spacing**                 | 8-16px                               | `mb-2` atau `mb-4`   |
| **Grid gap**                     | 16-32px                              | `gap-4 md:gap-8`     |

---

## 5. Component Library

### 5.1 Buttons

#### Primary Button

```jsx
// Usage
<Button variant="primary">Konsultasi Gratis</Button>

// Tailwind Classes
className="bg-accent-indigo text-white font-semibold px-6 py-3 rounded-lg
           hover:bg-opacity-90 transition-all duration-200
           focus:outline-none focus:ring-2 focus:ring-accent-indigo focus:ring-offset-2"
```

| Property          | Value                          |
| ----------------- | ------------------------------ |
| **Background**    | Indigo (#4F46E5)               |
| **Text**          | White                          |
| **Padding**       | 12px vertical, 24px horizontal |
| **Border Radius** | 8px                            |
| **Font Weight**   | 600 (Semibold)                 |
| **Hover**         | Opacity 90%                    |
| **Focus**         | Ring 2px Indigo, offset 2px    |

#### Secondary Button

```jsx
// Usage
<Button variant="secondary">Lihat Layanan</Button>

// Tailwind Classes
className="bg-transparent text-primary border-2 border-primary font-semibold
           px-6 py-3 rounded-lg hover:bg-primary hover:text-white
           transition-all duration-200 focus:outline-none focus:ring-2
           focus:ring-primary focus:ring-offset-2"
```

| Property       | Value                            |
| -------------- | -------------------------------- |
| **Background** | Transparent                      |
| **Border**     | 2px Deep Blue                    |
| **Text**       | Deep Blue                        |
| **Hover**      | Background Deep Blue, Text White |

#### Ghost Button

```jsx
// Usage
<Button variant="ghost">Pelajari Lebih Lanjut →</Button>

// Tailwind Classes
className="bg-transparent text-accent-indigo font-semibold px-4 py-2
           hover:bg-accent-indigo hover:bg-opacity-10 transition-all duration-200
           focus:outline-none focus:ring-2 focus:ring-accent-indigo"
```

| Property       | Value                         |
| -------------- | ----------------------------- |
| **Background** | Transparent                   |
| **Text**       | Indigo                        |
| **Hover**      | Background Indigo 10% opacity |

### 5.2 Cards

#### Service Card

```jsx
// Usage
<ServiceCard
  icon={<CodeIcon />}
  title="Pengembangan Aplikasi"
  description="Kami membangun aplikasi web dan mobile yang scalable..."
  features={['Web Application', 'Mobile Application', 'API Development']}
/>

// Tailwind Classes
className="bg-white p-8 rounded-2xl shadow-sm border border-neutral-300
           hover:shadow-lg hover:border-accent-indigo transition-all duration-300"
```

| Property          | Value                          |
| ----------------- | ------------------------------ |
| **Background**    | White                          |
| **Padding**       | 32px                           |
| **Border Radius** | 16px                           |
| **Border**        | 1px Slate 300                  |
| **Shadow**        | Small (default), Large (hover) |
| **Hover**         | Border Indigo, shadow large    |

#### Problem Card

```jsx
// Usage
<ProblemCard
  icon={<GlobeIcon />}
  title="Website Saja Tidak Cukup"
  description="Bisnis Anda butuh sistem yang terintegrasi..."
/>;

// Tailwind Classes
className = "bg-slate-100 p-6 rounded-xl border border-neutral-300";
```

| Property          | Value         |
| ----------------- | ------------- |
| **Background**    | Slate 100     |
| **Padding**       | 24px          |
| **Border Radius** | 12px          |
| **Border**        | 1px Slate 300 |

### 5.3 Form Inputs

#### Text Input

```jsx
// Usage
<Input
  label="Nama Lengkap"
  placeholder="Masukkan nama Anda"
  type="text"
  required
/>

// Tailwind Classes
className="w-full px-4 py-3 border border-neutral-300 rounded-lg
           focus:outline-none focus:ring-2 focus:ring-accent-indigo
           focus:border-transparent placeholder:text-neutral-500"
```

| Property          | Value                               |
| ----------------- | ----------------------------------- |
| **Border**        | 1px Slate 300                       |
| **Padding**       | 12px horizontal, 12px vertical      |
| **Border Radius** | 8px                                 |
| **Focus**         | Ring 2px Indigo, border transparent |
| **Placeholder**   | Slate 500                           |

#### Textarea

```jsx
// Usage
<Textarea
  label="Pesan"
  placeholder="Ceritakan kebutuhan Anda..."
  rows={5}
  required
/>

// Tailwind Classes
className="w-full px-4 py-3 border border-neutral-300 rounded-lg
           focus:outline-none focus:ring-2 focus:ring-accent-indigo
           focus:border-transparent placeholder:text-neutral-500 resize-none"
```

### 5.4 Badges

#### Service Badge

```jsx
// Usage
<Badge variant="cyan">SaaS Ready</Badge>

// Tailwind Classes
className="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium
           bg-accent-cyan bg-opacity-10 text-accent-cyan"
```

| Property          | Value                         |
| ----------------- | ----------------------------- |
| **Background**    | Cyan 10% opacity              |
| **Text**          | Cyan                          |
| **Padding**       | 4px vertical, 12px horizontal |
| **Border Radius** | Full (9999px)                 |
| **Font Size**     | 14px                          |
| **Font Weight**   | 500 (Medium)                  |

### 5.5 Icons

#### Icon Guidelines

| Property         | Value                                      |
| ---------------- | ------------------------------------------ |
| **Library**      | Lucide Icons                               |
| **Style**        | Outline (stroke)                           |
| **Stroke Width** | 2px                                        |
| **Size**         | 24px (default), 20px (small), 32px (large) |
| **Color**        | Inherit from parent or use semantic colors |

```jsx
// Usage
import { Code, Palette, Network } from 'lucide-react';

<Code className="w-6 h-6 text-accent-indigo" />
<Palette className="w-6 h-6 text-accent-indigo" />
<Network className="w-6 h-6 text-accent-indigo" />
```

---

## 6. Layout System

### 6.1 Container

```jsx
// Usage
<Container>{/* Content */}</Container>;

// Tailwind Classes
className = "max-w-7xl mx-auto px-4 sm:px-6 lg:px-8";
```

| Property      | Value                                  |
| ------------- | -------------------------------------- |
| **Max Width** | 1280px (7xl)                           |
| **Padding**   | 16px mobile, 24px tablet, 32px desktop |
| **Margin**    | Auto (centered)                        |

### 6.2 Grid System

#### 3-Column Grid (Services, Problem)

```jsx
// Desktop: 3 columns
// Tablet: 2 columns
// Mobile: 1 column

className = "grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8";
```

#### 2-Column Grid (Why Webifylab)

```jsx
// Desktop: 2 columns
// Mobile: 1 column

className = "grid grid-cols-1 md:grid-cols-2 gap-8";
```

### 6.3 Section Layout

```jsx
// Standard Section
<section className="py-20 md:py-32 bg-white">
  <Container>
    <SectionHeader
      title="Layanan Kami"
      subtitle="Tiga pilar layanan yang saling terintegrasi..."
    />
    {/* Content */}
  </Container>
</section>
```

---

## 7. Animation & Motion

### 7.1 Animation Principles

| Principle      | Deskripsi                                                   |
| -------------- | ----------------------------------------------------------- |
| **Purposeful** | Animasi harus memiliki tujuan (guidance, feedback, delight) |
| **Subtle**     | Jangan berlebihan, cukup 200-300ms                          |
| **Consistent** | Gunakan easing dan duration yang sama                       |
| **Accessible** | Respect `prefers-reduced-motion`                            |

### 7.2 Animation Types

#### Fade In

```css
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.animate-fade-in {
  animation: fadeIn 0.3s ease-in-out;
}
```

#### Slide Up

```css
@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-slide-up {
  animation: slideUp 0.4s ease-out;
}
```

### 7.3 Animation Usage

| Context              | Animation       | Duration | Easing      |
| -------------------- | --------------- | -------- | ----------- |
| **Page load**        | Fade in         | 300ms    | ease-in-out |
| **Section entrance** | Slide up        | 400ms    | ease-out    |
| **Button hover**     | Scale + opacity | 200ms    | ease        |
| **Card hover**       | Shadow + border | 300ms    | ease        |
| **Modal open**       | Fade + scale    | 300ms    | ease-out    |

---

## 8. Responsive Design

### 8.1 Breakpoints

| Name              | Min Width | Usage                  |
| ----------------- | --------- | ---------------------- |
| **Mobile**        | 375px     | Default (mobile-first) |
| **Tablet**        | 768px     | `md:` prefix           |
| **Desktop**       | 1024px    | `lg:` prefix           |
| **Large Desktop** | 1280px    | `xl:` prefix           |

### 8.2 Responsive Guidelines

| Element               | Mobile | Tablet | Desktop |
| --------------------- | ------ | ------ | ------- |
| **Font Size (H1)**    | 32px   | 40px   | 48px    |
| **Font Size (H2)**    | 28px   | 32px   | 36px    |
| **Container Padding** | 16px   | 24px   | 32px    |
| **Grid Columns**      | 1      | 2      | 3       |
| **Section Padding**   | 60px   | 80px   | 120px   |
| **Button Width**      | 100%   | Auto   | Auto    |

### 8.3 Mobile-First Approach

```jsx
// Mobile-first: default styles for mobile
// Use md: and lg: for larger screens

<div
  className="
  text-3xl           /* Mobile: 30px */
  md:text-4xl        /* Tablet: 36px */
  lg:text-5xl        /* Desktop: 48px */
"
>
  Headline
</div>
```

---

## 9. Accessibility

### 9.1 Color Contrast

| Element                     | Minimum Ratio | WCAG Level |
| --------------------------- | ------------- | ---------- |
| **Normal text**             | 4.5:1         | AA         |
| **Large text (18px+ bold)** | 3:1           | AA         |
| **UI components**           | 3:1           | AA         |

### 9.2 Focus States

```css
/* All interactive elements must have visible focus */
button:focus,
a:focus,
input:focus {
  outline: 2px solid #4f46e5;
  outline-offset: 2px;
}
```

### 9.3 Touch Targets

| Element         | Minimum Size |
| --------------- | ------------ |
| **Buttons**     | 44x44px      |
| **Links**       | 44x44px      |
| **Form inputs** | 44px height  |

### 9.4 Accessibility Checklist

- [ ] All images have alt text
- [ ] All form inputs have labels
- [ ] Color contrast meets WCAG AA
- [ ] Focus states visible
- [ ] Keyboard navigation works
- [ ] ARIA labels for complex widgets
- [ ] Skip to content link
- [ ] Semantic HTML (header, main, section, footer)

---

## 10. Do's and Don'ts

### 10.1 Visual Do's ✅

- ✅ Gunakan white space generously
- ✅ Konsisten dengan border radius (8px untuk button, 12-16px untuk card)
- ✅ Gunakan ikon untuk visual hierarchy
- ✅ Animasi subtle dan purposeful
- ✅ Kontras tinggi untuk readability
- ✅ Mobile-first approach
- ✅ Semantic HTML
- ✅ Focus states visible

### 10.2 Visual Don'ts ❌

- ❌ Jangan gunakan lebih dari 3 warna utama
- ❌ Jangan gunakan animasi yang mengganggu
- ❌ Jangan gunakan foto stock generik
- ❌ Jangan gunakan teks terlalu kecil (< 14px)
- ❌ Jangan gunakan border radius berbeda-beda
- ❌ Jangan gunakan terlalu banyak shadow
- ❌ Jangan gunakan inline styles
- ❌ Jangan skip heading levels

---

## 11. Implementation Guide

### 11.1 Tailwind CSS Setup

```javascript
// tailwind.config.js
module.exports = {
  content: ["./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}"],
  theme: {
    extend: {
      colors: {
        primary: {
          DEFAULT: "#1E3A5F",
          light: "#2D5A8C",
          dark: "#152A45",
        },
        accent: {
          indigo: "#4F46E5",
          cyan: "#06B6D4",
        },
        neutral: {
          900: "#0F172A",
          700: "#334155",
          500: "#64748B",
          300: "#CBD5E1",
          100: "#F1F5F9",
        },
      },
      fontFamily: {
        sans: ["Inter", "system-ui", "-apple-system", "sans-serif"],
      },
      fontSize: {
        h1: ["3rem", { lineHeight: "1.2", fontWeight: "800" }],
        h2: ["2.25rem", { lineHeight: "1.3", fontWeight: "700" }],
        h3: ["1.5rem", { lineHeight: "1.4", fontWeight: "600" }],
      },
      borderRadius: {
        xl: "12px",
        "2xl": "16px",
      },
    },
  },
  plugins: [],
};
```

### 11.2 Global CSS

```css
/* src/styles/global.css */
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  html {
    font-family:
      "Inter",
      system-ui,
      -apple-system,
      sans-serif;
    scroll-behavior: smooth;
  }

  body {
    @apply text-neutral-900 bg-white;
  }

  h1,
  h2,
  h3,
  h4,
  h5,
  h6 {
    @apply font-bold text-neutral-900;
  }

  h1 {
    @apply text-h1;
  }
  h2 {
    @apply text-h2;
  }
  h3 {
    @apply text-h3;
  }
  h4 {
    @apply text-h4;
  }

  a {
    @apply text-accent-indigo hover:underline;
  }

  button:focus,
  a:focus,
  input:focus {
    @apply outline-none ring-2 ring-accent-indigo ring-offset-2;
  }
}

@layer components {
  .btn-primary {
    @apply bg-accent-indigo text-white font-semibold px-6 py-3 rounded-lg 
           hover:bg-opacity-90 transition-all duration-200;
  }

  .btn-secondary {
    @apply bg-transparent text-primary border-2 border-primary font-semibold 
           px-6 py-3 rounded-lg hover:bg-primary hover:text-white transition-all duration-200;
  }

  .card {
    @apply bg-white p-8 rounded-2xl shadow-sm border border-neutral-300 
           hover:shadow-lg hover:border-accent-indigo transition-all duration-300;
  }

  .input {
    @apply w-full px-4 py-3 border border-neutral-300 rounded-lg 
           focus:outline-none focus:ring-2 focus:ring-accent-indigo 
           focus:border-transparent placeholder:text-neutral-500;
  }
}
```

---

## 12. Design Assets

### 12.1 Figma Library (Optional)

Jika kamu menggunakan Figma, buat library dengan:

- Color styles (Primary, Accent, Neutral, Semantic)
- Text styles (H1-H4, Body, Caption)
- Component variants (Button, Card, Input, Badge)
- Auto layout untuk responsive design

### 12.2 Icon Library

**Lucide Icons:** https://lucide.dev

Recommended icons:

- Code, Palette, Network (Services)
- Globe, Handshake, Brain (Problem)
- Blueprint, Ecosystem, Brain, Chat (Why Webifylab)
- ArrowRight, Check, X (UI)

---

## 13. Design System Checklist

Sebelum development, pastikan:

- [ ] Color palette defined dan disetujui
- [ ] Typography scale finalized
- [ ] Spacing system understood
- [ ] Component library documented
- [ ] Responsive breakpoints defined
- [ ] Accessibility requirements met
- [ ] Tailwind config updated
- [ ] Global CSS created
- [ ] Icon library chosen
- [ ] Do's and Don'ts reviewed

---

## 14. Open Questions

| No  | Pertanyaan                                                                                   | Status  |
| --- | -------------------------------------------------------------------------------------------- | ------- |
| Q1  | Apakah color palette (Deep Blue, Indigo, Cyan) sudah sesuai? Atau ada preferensi warna lain? | Pending |
| Q2  | Apakah font Inter sudah sesuai? Atau preferensi Plus Jakarta Sans?                           | Pending |
| Q3  | Apakah ada brand logo yang sudah ada? Atau perlu dibuatkan?                                  | Pending |
| Q4  | Apakah ingin menggunakan Figma untuk design? Atau langsung ke code?                          | Pending |
| Q5  | Apakah ada referensi website yang disukai secara visual?                                     | Pending |

---

_Dokumen ini adalah living document. Versi akan diperbarui seiring feedback dan perkembangan desain._

**Last Updated:** 17 September 2026
**Next Step:** Review design system, finalisasi visual direction, lalu mulai wireframe atau langsung development.
