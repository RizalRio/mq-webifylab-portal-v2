// Site constants — update these before going live

export const SITE = {
  name: "Webifylab",
  tagline: "Dari Ide Menjadi Ekosistem Digital",
  description:
    "Webifylab membantu bisnis membangun aplikasi, desain, dan sistem SaaS yang scalable. Partner teknologi jangka panjang dengan keahlian AI & Data. Konsultasi gratis!",
  url: "https://webifylab.com",
  email: "hello@webifylab.com",
  // Replace with your real WhatsApp number (format: 628xxxxxxxxxx)
  whatsapp: "6281234567890",
  social: {
    linkedin: "https://linkedin.com/company/webifylab",
    github:   "https://github.com/webifylab",
    instagram: "https://instagram.com/webifylab",
  },
} as const;

export const NAV_LINKS = [
  { label: "Layanan",    href: "#layanan" },
  { label: "Pendekatan", href: "#pendekatan" },
  { label: "Portfolio",  href: "#portfolio" },
  { label: "Kontak",     href: "#kontak" },
] as const;

export const SERVICES = [
  {
    icon: "code",
    title: "Pengembangan Aplikasi",
    description:
      "Kami membangun aplikasi web dan mobile yang scalable, dengan arsitektur yang siap berkembang bersama bisnis Anda.",
    features: [
      "Web Application (SaaS, Dashboard, Internal Tools)",
      "Mobile Application (Android & iOS)",
      "API Development & Integration",
      "Sistem Legacy Modernization",
    ],
    badge: "Most Popular",
  },
  {
    icon: "palette",
    title: "Desain Grafis & Web Design",
    description:
      "Visual yang konsisten dan profesional untuk membangun brand yang kuat dan pengalaman pengguna yang optimal.",
    features: [
      "UI/UX Design (Website & Aplikasi)",
      "Brand Identity (Logo, Guidelines)",
      "Marketing Collateral (Banner, Social Media)",
      "Design System & Component Library",
    ],
    badge: null,
  },
  {
    icon: "network",
    title: "SaaS & Ekosistem Digital",
    description:
      "Produk SaaS dari Webifylab yang saling terintegrasi, siap di-scale dengan AI & Data untuk efisiensi maksimal.",
    features: [
      "Produk SaaS siap pakai",
      "Integrasi antar sistem (API-first)",
      "AI & Data Analytics",
      "Automation & Workflow Optimization",
    ],
    badge: "Future-Ready",
  },
] as const;

export const WHY_ITEMS = [
  {
    icon: "layout-template",
    title: "System-First Thinking",
    description:
      "Setiap proyek dimulai dari analisis sistem yang mendalam. Kami tidak langsung coding — kami merancang fondasi yang kuat terlebih dahulu.",
  },
  {
    icon: "share-2",
    title: "Ekosistem, Bukan Sekadar Proyek",
    description:
      "Kami membangun fondasi yang bisa berkembang. Website hari ini bisa menjadi bagian dari ekosistem SaaS besok.",
  },
  {
    icon: "brain",
    title: "AI & Data Ready",
    description:
      "Dengan keahlian di bidang AI dan Data, kami membantu bisnis Anda siap menghadapi era otomatisasi dan data-driven decision.",
  },
  {
    icon: "message-square",
    title: "Transparent & Agile",
    description:
      "Komunikasi terbuka, iterasi cepat, dokumentasi rapi. Anda selalu tahu progress proyek dan bisa memberikan feedback kapan saja.",
  },
] as const;

export const APPROACH_STEPS = [
  {
    number: "01",
    title: "Discovery",
    description:
      "Kami mendengarkan kebutuhan bisnis Anda, menganalisis masalah, dan mendefinisikan tujuan yang jelas.",
  },
  {
    number: "02",
    title: "Design & Planning",
    description:
      "Kami merancang arsitektur sistem, wireframe, dan roadmap pengembangan yang detail dan terukur.",
  },
  {
    number: "03",
    title: "Development",
    description:
      "Kami membangun solusi dengan tech stack modern, mengikuti best practices, dan melakukan testing ketat.",
  },
  {
    number: "04",
    title: "Deploy & Support",
    description:
      "Kami membantu deployment, memberikan dokumentasi lengkap, dan siap mendukung Anda pasca-launch.",
  },
] as const;

export const PORTFOLIO_ITEMS = [
  {
    title: "Proyek 1 — Coming Soon",
    category: "Web Application",
    description: "Detail proyek akan segera ditampilkan.",
    gradient: "from-accent-indigo to-accent-cyan",
    comingSoon: true,
  },
  {
    title: "Proyek 2 — Coming Soon",
    category: "UI/UX Design",
    description: "Detail proyek akan segera ditampilkan.",
    gradient: "from-primary to-accent-indigo",
    comingSoon: true,
  },
  {
    title: "Proyek 3 — Coming Soon",
    category: "SaaS Development",
    description: "Detail proyek akan segera ditampilkan.",
    gradient: "from-accent-cyan to-accent-indigo",
    comingSoon: true,
  },
] as const;

export const FAQ_ITEMS = [
  {
    question: "Berapa lama waktu pengerjaan proyek?",
    answer:
      "Tergantung kompleksitas proyek. Website sederhana bisa selesai dalam 2–4 minggu, sedangkan aplikasi kompleks bisa memakan waktu 2–6 bulan. Kami akan memberikan estimasi timeline yang jelas setelah sesi discovery.",
  },
  {
    question: "Apakah ada garansi setelah proyek selesai?",
    answer:
      "Ya, kami memberikan garansi bug fix selama 30 hari setelah launch. Setelah itu, kami menawarkan paket maintenance bulanan jika Anda membutuhkan dukungan berkelanjutan.",
  },
  {
    question: "Bagaimana sistem pembayaran?",
    answer:
      "Kami menggunakan sistem milestone-based payment. Biasanya 40% di awal, 30% di tengah proyek, dan 30% setelah proyek selesai. Detail akan dibahas di kontrak.",
  },
  {
    question: "Apakah saya perlu paham teknis untuk bekerja dengan Webifylab?",
    answer:
      "Tidak perlu. Kami akan menjelaskan semua hal teknis dengan bahasa yang mudah dipahami. Yang penting, Anda punya visi bisnis yang jelas — kami yang akan menerjemahkannya menjadi solusi teknis.",
  },
] as const;
