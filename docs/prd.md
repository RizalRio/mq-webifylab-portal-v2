# 📄 Product Requirements Document (PRD)

## Webifylab Landing Page — Version 1.0

| Metadata       | Detail                             |
| -------------- | ---------------------------------- |
| **Product**    | Webifylab Landing Page             |
| **Version**    | 1.0                                |
| **Author**     | Rizal                              |
| **Role**       | System Analyst / Software Engineer |
| **Created**    | 17 September 2026                  |
| **Status**     | Draft                              |
| **Tech Stack** | Astro, Golang, Tailwind CSS        |
| **Hosting**    | VPS 1 vCPU, 1 GB RAM, 20 GB SSD    |

---

## 1. Product Overview & Objectives

### 1.1 Product Name

Webifylab Landing Page V2

### 1.2 Background

Webifylab adalah software house yang didirikan oleh Rizal, seorang lulusan Informatika dengan keahlian sebagai System Analyst, Software Engineer, serta enthusiast di bidang AI dan Data. Webifylab menawarkan tiga pilar layanan utama: pengembangan aplikasi, desain grafis & web, serta pengembangan ekosistem SaaS. Landing page ini merupakan aset digital pertama yang berfungsi sebagai wajah dan pintu masuk utama bagi calon klien.

### 1.3 Problem Statement

Calon klien (UMKM, startup, maupun perusahaan) kesulitan menemukan partner teknologi yang tidak hanya mampu membangun website atau aplikasi, tetapi juga memiliki pemahaman mendalam tentang arsitektur sistem, ekosistem SaaS, serta integrasi AI dan Data untuk kebutuhan jangka panjang. Banyak software house hanya menawarkan jasa "sekali buat", tanpa visi skalabilitas dan ekosistem.

### 1.4 Vision

Menjadi landing page yang secara efektif mengkomunikasikan positioning Webifylab sebagai software house modern yang berorientasi pada ekosistem digital, serta menghasilkan leads berkualitas untuk dikonversi menjadi klien.

### 1.5 Objectives

| No  | Objective                                         | Target                                                          |
| --- | ------------------------------------------------- | --------------------------------------------------------------- |
| O1  | Mendapatkan leads (kontak masuk) dari calon klien | Minimal 3 leads/bulan dalam 3 bulan pertama                     |
| O2  | Mengkomunikasikan 3 pilar layanan dengan jelas    | Pengunjung memahami layanan dalam < 30 detik pertama            |
| O3  | Membangun kredibilitas dan trust                  | Portfolio dan value proposition tersampaikan dengan profesional |
| O4  | Optimal untuk mesin pencari (SEO)                 | Ter-index di Google untuk keyword target dalam 2 bulan          |

### 1.6 Success Metrics (KPIs)

| Metric                           | Target                     | Tools Pengukuran                |
| -------------------------------- | -------------------------- | ------------------------------- |
| Page Load Time                   | < 2 detik                  | Lighthouse / PageSpeed Insights |
| Lighthouse Performance Score     | > 90                       | Chrome DevTools                 |
| Bounce Rate                      | < 60%                      | Plausible / Google Analytics    |
| CTA Click-through Rate           | > 5% dari total pengunjung | Analytics event tracking        |
| Form Submission / WhatsApp Click | > 3/bulan                  | Formspree / WA API log          |
| Mobile Responsiveness            | 100% layout tidak broken   | Manual test + BrowserStack      |

---

## 2. Target Audience (User Personas)

### 2.1 Persona 1: UMKM Owner — "Pak Budi"

| Atribut                            | Detail                                                                                                              |
| ---------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| **Usia**                           | 35–55 tahun                                                                                                         |
| **Pekerjaan**                      | Pemilik usaha lokal (toko, restoran, jasa)                                                                          |
| **Tech Literacy**                  | Rendah – Menengah                                                                                                   |
| **Pain Points**                    | Belum punya website, kalah saing dengan kompetitor yang sudah online, tidak paham istilah teknis                    |
| **Kebutuhan**                      | Website profil perusahaan atau katalog produk sederhana, harga terjangkau, proses cepat, bahasa yang mudah dipahami |
| **Decision Factor**                | Harga, kecepatan pengerjaan, ada yang menghandle teknis sepenuhnya                                                  |
| **Layanan Webifylab yang Relevan** | Pembuatan Website, Desain Grafis                                                                                    |

### 2.2 Persona 2: Startup Founder — "Sarah"

| Atribut                            | Detail                                                                                                           |
| ---------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| **Usia**                           | 22–35 tahun                                                                                                      |
| **Pekerjaan**                      | Founder / Co-founder startup tahap awal                                                                          |
| **Tech Literacy**                  | Menengah – Tinggi                                                                                                |
| **Pain Points**                    | Butuh MVP cepat tapi scalable, budget terbatas, perlu partner yang paham arsitektur sistem dan bisa grow bersama |
| **Kebutuhan**                      | MVP aplikasi web/mobile, sistem yang siap di-scale, pemahaman tentang SaaS dan data                              |
| **Decision Factor**                | Kualitas arsitektur, pemahaman bisnis, kemampuan scaling, ekosistem                                              |
| **Layanan Webifylab yang Relevan** | Pembuatan Aplikasi, SaaS Ecosystem                                                                               |

### 2.3 Persona 3: Corporate IT Manager — "Mas Andi"

| Atribut                            | Detail                                                                                                                      |
| ---------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| **Usia**                           | 30–45 tahun                                                                                                                 |
| **Pekerjaan**                      | IT Manager / Head of Digital di perusahaan menengah                                                                         |
| **Tech Literacy**                  | Tinggi                                                                                                                      |
| **Pain Points**                    | Vendor sebelumnya tidak deliver sesuai spec, sulit cari partner yang paham AI/Data integration, butuh dokumentasi yang rapi |
| **Kebutuhan**                      | Sistem internal, integrasi AI/ML, dashboard data, partner yang profesional dan terstruktur                                  |
| **Decision Factor**                | Portofolio, metodologi kerja, dokumentasi teknis, kemampuan AI/Data                                                         |
| **Layanan Webifylab yang Relevan** | Pembuatan Aplikasi, AI & Data Solutions                                                                                     |

---

## 3. Information Architecture & User Flow

### 3.1 Page Structure (Single Page — Scroll Flow)

Landing page ini menggunakan pendekatan **single-page** dengan scroll vertikal. Setiap section dirancang untuk memandu pengunjung dari kesadaran masalah hingga aksi (konversi).

| No  | Section             | Tujuan                               | Konten Utama                                                |
| --- | ------------------- | ------------------------------------ | ----------------------------------------------------------- |
| S1  | **Navbar**          | Navigasi cepat                       | Logo, menu anchor link, tombol CTA "Hubungi Kami"           |
| S2  | **Hero**            | First impression & value proposition | Headline kuat, sub-headline, 2 tombol CTA, visual/ilustrasi |
| S3  | **Problem**         | Agitasi masalah klien                | 3 pain point utama yang dihadapi calon klien                |
| S4  | **Services**        | Solusi yang ditawarkan               | 3 pilar layanan (App Dev, Design, SaaS Ecosystem)           |
| S5  | **Why Webifylab**   | Differentiator & trust               | Keunggulan kompetitif, pendekatan ekosistem                 |
| S6  | **Tech & Approach** | Kredibilitas teknis                  | Tech stack yang digunakan, metodologi kerja                 |
| S7  | **Portfolio**       | Bukti kemampuan                      | 3–6 showcase proyek (bisa placeholder di V1)                |
| S8  | **Testimonial**     | Social proof                         | Kutipan klien (bisa placeholder di V1)                      |
| S9  | **CTA / Contact**   | Konversi                             | Form kontak atau tombol WhatsApp                            |
| S10 | **Footer**          | Informasi penutup                    | Copyright, social links, quick links                        |

### 3.2 Primary User Flow

```
[Pengunjung mendarat di Hero]
        │
        ▼
[Membaca Headline & Value Proposition]
        │
        ▼
[Scroll ke Problem — "Ini masalah saya!"]
        │
        ▼
[Scroll ke Services — "Oh, mereka bisa bantu"]
        │
        ▼
[Scroll ke Why Webifylab & Portfolio — "Mereka kredibel"]
        │
        ▼
[Klik CTA "Konsultasi Gratis" / "Hubungi Kami"]
        │
        ├──► [Opsi A] Diarahkan ke WhatsApp (wa.me/...)
        │
        └──► [Opsi B] Mengisi Form Kontak (Formspree/Web3Forms)
                    │
                    ▼
            [Notifikasi masuk ke Email/Telegram Rizal]
```

### 3.3 Secondary User Flow

```
[Pengunjung dari search engine / social media]
        │
        ▼
[Langsung scroll ke Portfolio atau Services]
        │
        ▼
[Klik anchor link di Navbar untuk navigasi cepat]
        │
        ▼
[Klik CTA]
```

---

## 4. Functional Requirements

### 4.1 Navbar (S1)

| ID     | Requirement                                                | Priority |
| ------ | ---------------------------------------------------------- | -------- |
| FR-1.1 | Navbar sticky (fixed position) saat di-scroll              | Must     |
| FR-1.2 | Logo Webifylab di kiri, menu navigasi di kanan             | Must     |
| FR-1.3 | Menu navigasi berupa anchor link ke masing-masing section  | Must     |
| FR-1.4 | Tombol CTA "Hubungi Kami" di ujung kanan navbar            | Must     |
| FR-1.5 | Hamburger menu untuk tampilan mobile (< 768px)             | Must     |
| FR-1.6 | Navbar berubah background (transparan → solid) saat scroll | Should   |

### 4.2 Hero Section (S2)

| ID     | Requirement                                                                 | Priority |
| ------ | --------------------------------------------------------------------------- | -------- |
| FR-2.1 | Headline utama (H1) yang jelas dan kuat                                     | Must     |
| FR-2.2 | Sub-headline yang menjelaskan value proposition                             | Must     |
| FR-2.3 | 2 tombol CTA: "Konsultasi Gratis" (primary) dan "Lihat Layanan" (secondary) | Must     |
| FR-2.4 | Visual pendukung (ilustrasi/gradient/abstract shape, bukan foto berat)      | Should   |
| FR-2.5 | Animasi entrance halus (fade-in/slide-up) tanpa library berat               | Could    |

### 4.3 Problem Section (S3)

| ID     | Requirement                                             | Priority |
| ------ | ------------------------------------------------------- | -------- |
| FR-3.1 | Menampilkan 3 pain point dalam format card              | Must     |
| FR-3.2 | Setiap card memiliki ikon, judul, dan deskripsi singkat | Must     |
| FR-3.3 | Layout responsif: 3 kolom (desktop), 1 kolom (mobile)   | Must     |

### 4.4 Services Section (S4)

| ID     | Requirement                                                               | Priority |
| ------ | ------------------------------------------------------------------------- | -------- |
| FR-4.1 | Menampilkan 3 pilar layanan dalam format card                             | Must     |
| FR-4.2 | Setiap card: ikon, judul layanan, deskripsi, daftar fitur (bullet points) | Must     |
| FR-4.3 | Pilar 1: Pengembangan Aplikasi (Web & Mobile)                             | Must     |
| FR-4.4 | Pilar 2: Desain Grafis & Web Design                                       | Must     |
| FR-4.5 | Pilar 3: SaaS & Ekosistem Digital                                         | Must     |
| FR-4.6 | Setiap card memiliki tombol "Pelajari Lebih Lanjut" (anchor ke CTA)       | Should   |

### 4.5 Why Webifylab Section (S5)

| ID     | Requirement                                           | Priority |
| ------ | ----------------------------------------------------- | -------- |
| FR-5.1 | Menampilkan 4 keunggulan kompetitif dalam grid 2x2    | Must     |
| FR-5.2 | Setiap item: ikon, judul, deskripsi singkat           | Must     |
| FR-5.3 | Menyampaikan konsep "Ekosistem, bukan sekadar vendor" | Must     |

### 4.6 Tech & Approach Section (S6)

| ID     | Requirement                                                                             | Priority |
| ------ | --------------------------------------------------------------------------------------- | -------- |
| FR-6.1 | Menampilkan logo/ikon tech stack yang digunakan                                         | Should   |
| FR-6.2 | Menampilkan 3–4 langkah metodologi kerja (misal: Discovery → Design → Develop → Deploy) | Should   |

### 4.7 Portfolio Section (S7)

| ID     | Requirement                                                                 | Priority |
| ------ | --------------------------------------------------------------------------- | -------- |
| FR-7.1 | Menampilkan 3–6 item portofolio dalam grid                                  | Must     |
| FR-7.2 | Setiap item: thumbnail gambar, judul proyek, kategori, deskripsi singkat    | Must     |
| FR-7.3 | Jika belum ada proyek nyata, gunakan placeholder dengan label "Coming Soon" | Must     |
| FR-7.4 | Gambar dioptimalkan (WebP, lazy loading)                                    | Must     |

### 4.8 Testimonial Section (S8)

| ID     | Requirement                                                            | Priority |
| ------ | ---------------------------------------------------------------------- | -------- |
| FR-8.1 | Menampilkan 2–3 testimonial dalam card/carousel sederhana              | Should   |
| FR-8.2 | Setiap testimonial: kutipan, nama, jabatan/perusahaan, foto avatar     | Should   |
| FR-8.3 | Jika belum ada, gunakan placeholder atau sembunyikan section ini di V1 | Could    |

### 4.9 CTA / Contact Section (S9)

| ID     | Requirement                                                                 | Priority |
| ------ | --------------------------------------------------------------------------- | -------- |
| FR-9.1 | Form kontak dengan field: Nama, Email, Jenis Layanan (dropdown), Pesan      | Must     |
| FR-9.2 | Validasi client-side untuk semua field                                      | Must     |
| FR-9.3 | Form submission ditangani oleh layanan pihak ketiga (Formspree / Web3Forms) | Must     |
| FR-9.4 | Tombol alternatif "Chat via WhatsApp" dengan link wa.me                     | Must     |
| FR-9.5 | Success message setelah form berhasil dikirim                               | Must     |
| FR-9.6 | Headline CTA yang persuasif                                                 | Must     |

### 4.10 Footer (S10)

| ID      | Requirement                                        | Priority |
| ------- | -------------------------------------------------- | -------- |
| FR-10.1 | Copyright notice dengan tahun dinamis              | Must     |
| FR-10.2 | Link ke social media (LinkedIn, GitHub, Instagram) | Must     |
| FR-10.3 | Quick links ke section utama                       | Should   |
| FR-10.4 | Email kontak                                       | Must     |

---

## 5. Non-Functional Requirements

### 5.1 Performance

| ID      | Requirement                                             | Target      |
| ------- | ------------------------------------------------------- | ----------- |
| NFR-1.1 | Page load time (LCP)                                    | < 2 detik   |
| NFR-1.2 | Total page size (HTML + CSS + JS + Images)              | < 2 MB      |
| NFR-1.3 | Lighthouse Performance Score                            | > 90        |
| NFR-1.4 | First Contentful Paint (FCP)                            | < 1.5 detik |
| NFR-1.5 | Cumulative Layout Shift (CLS)                           | < 0.1       |
| NFR-1.6 | Gambar menggunakan format WebP/AVIF dengan lazy loading | Wajib       |
| NFR-1.7 | CSS & JS di-minify dan di-bundle                        | Wajib       |

### 5.2 Responsiveness

| ID      | Requirement                                     | Breakpoint       |
| ------- | ----------------------------------------------- | ---------------- |
| NFR-2.1 | Tampilan optimal di Mobile                      | 375px – 767px    |
| NFR-2.2 | Tampilan optimal di Tablet                      | 768px – 1023px   |
| NFR-2.3 | Tampilan optimal di Desktop                     | 1024px – 1440px+ |
| NFR-2.4 | Tidak ada horizontal scroll di semua breakpoint | Wajib            |
| NFR-2.5 | Touch-friendly: tombol & link min 44x44px       | Wajib            |

### 5.3 SEO

| ID      | Requirement                                          | Detail                                  |
| ------- | ---------------------------------------------------- | --------------------------------------- |
| NFR-3.1 | Meta title & description yang optimal                | Title < 60 char, Desc < 160 char        |
| NFR-3.2 | Open Graph tags (og:title, og:description, og:image) | Untuk preview di WhatsApp, Twitter, dll |
| NFR-3.3 | Struktur heading hierarkis (H1 → H2 → H3)            | Wajib                                   |
| NFR-3.4 | Alt text untuk semua gambar                          | Wajib                                   |
| NFR-3.5 | Semantic HTML (header, main, section, footer, nav)   | Wajib                                   |
| NFR-3.6 | Sitemap.xml dan robots.txt                           | Should                                  |
| NFR-3.7 | Canonical URL                                        | Wajib                                   |

### 5.4 Accessibility

| ID      | Requirement                         | Detail                         |
| ------- | ----------------------------------- | ------------------------------ |
| NFR-4.1 | Kontras warna memenuhi WCAG AA      | Rasio minimal 4.5:1 untuk teks |
| NFR-4.2 | Navigasi keyboard (tab order)       | Wajib                          |
| NFR-4.3 | ARIA labels untuk elemen interaktif | Should                         |

### 5.5 Infrastructure & Hosting

| ID      | Requirement                                                   | Detail                                       |
| ------- | ------------------------------------------------------------- | -------------------------------------------- |
| NFR-5.1 | Website harus berupa static files (HTML/CSS/JS)               | Tidak ada server-side rendering saat runtime |
| NFR-5.2 | Di-serve melalui Nginx atau Caddy di VPS                      | Konfigurasi ringan                           |
| NFR-5.3 | HTTPS wajib (Let's Encrypt / SSL gratis)                      | Wajib                                        |
| NFR-5.4 | Gzip/Brotli compression aktif di web server                   | Wajib                                        |
| NFR-5.5 | Tidak ada database yang berjalan untuk landing page           | Wajib                                        |
| NFR-5.6 | Tidak ada backend server (Node/PHP/Python) untuk landing page | Wajib                                        |
| NFR-5.7 | Form handling via third-party service                         | Formspree / Web3Forms / Getform              |

### 5.6 Security

| ID      | Requirement                                 | Detail |
| ------- | ------------------------------------------- | ------ |
| NFR-6.1 | Content Security Policy (CSP) headers       | Should |
| NFR-6.2 | Tidak ada inline scripts yang tidak perlu   | Wajib  |
| NFR-6.3 | Dependency audit (npm audit) sebelum deploy | Wajib  |

---

## 6. Design Guidelines

### 6.1 Visual Direction

| Atribut           | Arahan                                                                                                                                 |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| **Style**         | Modern, clean, minimal, tech-forward                                                                                                   |
| **Tone**          | Profesional namun approachable, inovatif                                                                                               |
| **Color Palette** | Primary: Deep Blue / Indigo (#1E3A5F atau #4F46E5), Accent: Cyan / Teal (#06B6D4), Neutral: Slate grays, Background: White / Off-white |
| **Typography**    | Heading: Inter / Plus Jakarta Sans (bold), Body: Inter / Plus Jakarta Sans (regular)                                                   |
| **Iconography**   | Lucide Icons / Heroicons (outline style, konsisten)                                                                                    |
| **Spacing**       | Section padding: 80–120px vertikal, container max-width: 1200px                                                                        |
| **Border Radius** | Cards: 12–16px, Buttons: 8px                                                                                                           |
| **Shadows**       | Subtle, soft shadows untuk card elevation                                                                                              |

### 6.2 Tone of Voice

| Do ✅                                                            | Don't ❌                                                   |
| ---------------------------------------------------------------- | ---------------------------------------------------------- |
| "Kami membangun sistem yang siap berkembang bersama bisnis Anda" | "Kami adalah perusahaan IT terbaik dan paling profesional" |
| "Dari ide hingga ekosistem digital"                              | "Solusi IT one-stop terlengkap se-Indonesia"               |
| Bahasa yang jelas, langsung, dan berorientasi solusi             | Jargon teknis berlebihan tanpa konteks                     |

---

## 7. Tech Stack & Architecture

### 7.1 Technology Choices

| Layer                  | Technology                                                 | Alasan                                                                                    |
| ---------------------- | ---------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| **Framework**          | Astro (primary) atau Next.js (static export)               | Menghasilkan static HTML yang sangat ringan, cepat, dan SEO-friendly. Cocok untuk VPS 1GB |
| **Styling**            | Tailwind CSS                                               | Utility-first, cepat untuk development, output CSS minimal (tree-shaking)                 |
| **Icons**              | Lucide Icons                                               | Ringan, konsisten, tree-shakeable                                                         |
| **Form Handling**      | Formspree / Web3Forms                                      | Gratis untuk volume rendah, tidak perlu backend                                           |
| **WhatsApp CTA**       | wa.me API                                                  | Langsung, tanpa backend                                                                   |
| **Web Server**         | Nginx atau Caddy                                           | Ringan, efisien untuk static files                                                        |
| **SSL**                | Let's Encrypt (Certbot)                                    | Gratis, auto-renew                                                                        |
| **Analytics**          | Plausible (self-hosted) atau Google Analytics              | Plausible lebih ringan & privacy-friendly                                                 |
| **Deployment**         | Manual via Git + SSH ke VPS, atau CI/CD via GitHub Actions | Sederhana untuk V1                                                                        |
| **Image Optimization** | Astro built-in / Sharp                                     | Auto-convert ke WebP, responsive sizes                                                    |

### 7.2 Architecture Diagram (Simplified)

```
┌─────────────────────────────────────────────┐
│                  CLIENT                     │
│          (Browser / Mobile)                 │
└──────────────────┬──────────────────────────┘
                   │ HTTPS
                   ▼
┌─────────────────────────────────────────────┐
│           VPS (1 vCPU, 1GB RAM)            │
│                                             │
│  ┌─────────────────────────────────────┐    │
│  │        Nginx / Caddy              │    │
│  │   (Reverse Proxy + Static Serve)  │    │
│  │   + SSL (Let's Encrypt)           │    │
│  └──────────────┬──────────────────────┘    │
│                 │                           │
│  ┌──────────────▼──────────────────────┐    │
│  │     Static Files (HTML/CSS/JS)     │    │
│  │     /var/www/webifylab/            │    │
│  │     (Built with Astro/Next.js)     │    │
│  └─────────────────────────────────────┘    │
│                                             │
└─────────────────────────────────────────────┘
                   │
                   │ Form Submit
                   ▼
┌─────────────────────────────────────────────┐
│     External Services (Offloaded)          │
│                                             │
│  ┌──────────────┐  ┌───────────────────┐   │
│  │  Formspree / │  │  Plausible /      │   │
│  │  Web3Forms   │  │  Google Analytics │   │
│  └──────────────┘  └───────────────────┘   │
│                                             │
│  ┌──────────────┐                           │
│  │  WhatsApp    │                           │
│  │  (wa.me)     │                           │
│  └──────────────┘                           │
└─────────────────────────────────────────────┘
```

---

## 8. Content Structure (Per Section)

### S2: Hero Section

- **H1 (Headline):** _"Dari Ide Menjadi Ekosistem Digital"_
- **Sub-headline:** _"Webifylab membantu bisnis Anda membangun aplikasi, desain, dan sistem SaaS yang siap berkembang — didukung oleh keahlian AI & Data."_
- **CTA Primary:** "Konsultasi Gratis →" (link ke section Contact / WhatsApp)
- **CTA Secondary:** "Lihat Layanan" (anchor ke section Services)

### S3: Problem Section

- **Section Title:** _"Mengapa Bisnis Anda Butuh Partner Teknologi yang Tepat?"_
- **Card 1:** "Website saja tidak cukup" — Bisnis butuh sistem yang terintegrasi, bukan sekadar halaman statis.
- **Card 2:** "Vendor yang hilang setelah proyek selesai" — Butuh partner jangka panjang yang paham bisnis Anda.
- **Card 3:** "Tidak siap menghadapi era AI & Data" — Kompetitor sudah mulai otomatisasi, Anda belum.

### S4: Services Section

- **Section Title:** _"Layanan Kami"_
- **Card 1 — Pengembangan Aplikasi:** Web app, mobile app, sistem internal. Tech stack modern, arsitektur scalable.
- **Card 2 — Desain Grafis & Web Design:** UI/UX design, branding, desain marketing. Visual yang konsisten dan profesional.
- **Card 3 — SaaS & Ekosistem Digital:** Produk SaaS dari Webifylab yang saling terintegrasi. Siap di-scale dengan AI & Data.

### S5: Why Webifylab

- **Section Title:** _"Mengapa Webifylab?"_
- **Item 1:** "System-First Thinking" — Setiap proyek dimulai dari analisis sistem, bukan sekadar coding.
- **Item 2:** "Ekosistem, Bukan Sekadar Proyek" — Kami membangun fondasi yang bisa berkembang.
- **Item 3:** "AI & Data Ready" — Keahlian di bidang AI dan Data untuk solusi masa depan.
- **Item 4:** "Transparent & Agile" — Komunikasi terbuka, iterasi cepat, dokumentasi rapi.

### S9: CTA / Contact

- **Headline:** _"Siap Membangun Sesuatu yang Luar Biasa?"_
- **Sub-text:** _"Ceritakan ide atau tantangan Anda. Konsultasi pertama gratis, tanpa komitmen."_

---

## 9. Out of Scope (V1)

| ID   | Item                          | Alasan                                 | Kapan?          |
| ---- | ----------------------------- | -------------------------------------- | --------------- |
| OS-1 | Blog / Artikel                | Butuh CMS, menambah kompleksitas       | V2              |
| OS-2 | Sistem Login / Client Portal  | Butuh backend & database               | V2              |
| OS-3 | Animasi 3D (Spline, Three.js) | Terlalu berat untuk VPS 1GB dan mobile | V2 (jika perlu) |
| OS-4 | Payment Gateway               | Belum ada produk berbayar di V1        | V2              |
| OS-5 | Multi-language (EN/ID)        | Fokus pasar lokal dulu                 | V2              |
| OS-6 | Dark Mode                     | Nice-to-have, bukan prioritas          | V2              |
| OS-7 | Chatbot / AI Assistant        | Butuh backend AI                       | V3              |
| OS-8 | CMS (Headless CMS)            | Konten masih bisa di-hardcode di V1    | V2              |

---

## 10. Milestones & Timeline

| Phase       | Task                                             | Estimasi        | Output                        |
| ----------- | ------------------------------------------------ | --------------- | ----------------------------- |
| **Phase 0** | Finalisasi PRD & Content Outline                 | 1–2 hari        | Dokumen PRD final, copy draft |
| **Phase 1** | Wireframe & Design (Figma)                       | 2–3 hari        | Figma mockup desktop & mobile |
| **Phase 2** | Setup project (Astro/Next.js + Tailwind)         | 1 hari          | Repo, dev server running      |
| **Phase 3** | Development: Navbar, Hero, Problem               | 2 hari          | Section S1–S3 responsif       |
| **Phase 4** | Development: Services, Why, Tech                 | 2 hari          | Section S4–S6 responsif       |
| **Phase 5** | Development: Portfolio, Testimonial, CTA, Footer | 2 hari          | Section S7–S10 responsif      |
| **Phase 6** | Integrasi form (Formspree/Web3Forms) & WhatsApp  | 0.5 hari        | Form berfungsi                |
| **Phase 7** | SEO, Analytics, Performance optimization         | 1 hari          | Lighthouse > 90, meta tags    |
| **Phase 8** | Deploy ke VPS (Nginx + SSL)                      | 1 hari          | Live di domain                |
| **Phase 9** | Testing & bug fix                                | 1 hari          | Checklist DoD terpenuhi       |
| **Total**   |                                                  | **~13–15 hari** | Landing page live             |

---

## 11. Definition of Done (DoD)

Landing page V1 dianggap **selesai** jika:

- [ ] Semua 10 section (S1–S10) telah dibangun dan ditampilkan
- [ ] Responsif di Mobile (375px), Tablet (768px), Desktop (1440px)
- [ ] Lighthouse Performance score > 90
- [ ] Lighthouse SEO score > 90
- [ ] Form kontak berhasil mengirim data ke Formspree/Web3Forms
- [ ] Notifikasi form masuk ke email Rizal
- [ ] Tombol WhatsApp mengarah ke wa.me dengan pesan pre-filled
- [ ] HTTPS aktif (SSL certificate valid)
- [ ] Open Graph tags berfungsi (preview benar saat di-share di WA)
- [ ] Semua gambar menggunakan WebP dan lazy loading
- [ ] Total page size < 2 MB
- [ ] Tidak ada console error di browser
- [ ] Deployed dan accessible di domain/subdomain yang ditentukan
- [ ] Code di-push ke repository Git

---

## 12. Risks & Mitigation

| Risk                               | Impact                 | Likelihood                       | Mitigation                                                                      |
| ---------------------------------- | ---------------------- | -------------------------------- | ------------------------------------------------------------------------------- |
| VPS 1GB kehabisan RAM saat serving | Website down           | Low (static files sangat ringan) | Gunakan static files saja, monitoring RAM dengan `htop`                         |
| Belum ada portofolio nyata         | Kredibilitas rendah    | High (baru mulai)                | Gunakan placeholder "Coming Soon" + fokus pada value proposition dan pendekatan |
| Belum ada testimonial              | Social proof kurang    | High                             | Sembunyikan section atau gunakan "trusted by" logo placeholder                  |
| Scope creep (terus menambah fitur) | Delay timeline         | Medium                           | Patuhi Out of Scope, pindahkan ide ke backlog V2                                |
| SEO tidak ter-index cepat          | Traffic organik lambat | Medium                           | Submit sitemap ke Google Search Console, share di social media                  |

---

## 13. Open Questions

| No  | Pertanyaan                                                       | Status  |
| --- | ---------------------------------------------------------------- | ------- |
| Q1  | Apakah sudah memiliki domain untuk Webifylab?                    | Pending |
| Q2  | Apakah sudah memiliki logo dan brand assets?                     | Pending |
| Q3  | Nomor WhatsApp mana yang akan digunakan untuk CTA?               | Pending |
| Q4  | Apakah ada proyek sebelumnya yang bisa ditampilkan di portfolio? | Pending |
| Q5  | Preferensi framework: Astro atau Next.js?                        | Pending |

---

## Appendix A: Referensi & Inspirasi

- [Astro Documentation](https://docs.astro.build)
- [Tailwind CSS](https://tailwindcss.com)
- [Formspree](https://formspree.io)
- [Lucide Icons](https://lucide.dev)
- [Plausible Analytics](https://plausible.io)

---

_Dokumen ini adalah living document. Versi akan diperbarui seiring perkembangan proyek._

**Last Updated:** 17 September 2026
**Next Step:** Finalisasi content copywriting dan mulai wireframe di Figma.
