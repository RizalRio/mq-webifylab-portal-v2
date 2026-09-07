# DOKUMEN: CATATAN PENGEMBANGAN SISTEM (TIKES)
**Periode:** Minggu ke-2, September 2026 (Rencana Kerja Mingguan)  
**Penanggung Jawab:** Rizal Rio Andrian – Product & Technology Lead  
**Fokus Utama:** Pematangan Modul Administrasi, Perbaikan Logika Pembayaran, dan Penyempurnaan UX.

---

## 1. RENCANA KERJA MINGGUAN (SPRINT BACKLOG)

### A. Peningkatan Pengalaman Pengguna (UX) & Antarmuka
*Fokus: Meningkatkan interaktivitas aplikasi dan kejelasan visual bagi pengguna akhir.*

1.  **Implementasi *State Management* pada Elemen Tombol**
    *   **Deskripsi:** Menerapkan logika *disable* dan indikator *loading* (spinner/skeleton) pada seluruh tombol aksi kritis (simpan, bayar, kirim) saat proses berjalan.
    *   **Tujuan:** Mencegah *double-submit* (klik ganda) dan memberikan umpan balik visual yang jelas kepada pengguna bahwa sistem sedang memproses data.
2.  **Redesain Navigasi Utama (Menu *Insight*)**
    *   **Deskripsi:** Mengubah ikon menu *Insight* pada *bottom navigation bar* menjadi bentuk lingkaran yang lebih menonjol (*center floating action* atau *highlighted circle*) untuk menarik perhatian.
    *   **Tindakan:** Menghilangkan elemen visual tanda tambah (+) yang membingungkan dan menyederhanakan hierarki visual.
3.  **Penyesuaian UI Status Langganan (*Subscription*)**
    *   **Deskripsi:** Memperbarui tampilan profil/langganan. Jika pengguna sudah berlangganan Premium, tampilkan tombol "Perpanjang" dan informasi "Tanggal Berakhir Langganan" secara eksplisit.

### B. Pengembangan Modul Administrasi (*Admin Dashboard*)
*Fokus: Membangun kendali penuh bagi Administrator untuk mengelola operasional platform.*

1.  **Manajemen Asesmen & Umpan Balik**
    *   Menambahkan antarmuka untuk memantau hasil asesmen pengguna dan mengelola masuknya umpan balik (*feedback*) dari aplikasi.
2.  **Manajemen Pembayaran, Langganan, & Pemesanan**
    *   **Logika Bisnis:** Membedakan secara jelas antara transaksi yang menggunakan "Benefit Premium" (misal: 1x konsultasi gratis/bulan) dengan "Booking Biasa" (berbayar).
    *   **Fitur:** Tabel riwayat transaksi dengan filter status dan jenis benefit.
3.  **Detail Pengguna (*User Management*)**
    *   Memisahkan tampilan detail profil berdasarkan peran: Tampilan khusus untuk **Client** (riwayat mood, jurnal, langganan) dan **Konselor** (jadwal praktik, rating, verifikasi STR/SIP).

### C. Perbaikan Sistem, Logika Bisnis & Infrastruktur
*Fokus: Menjamin integritas data, keakuratan perhitungan, dan stabilitas lingkungan.*

1.  **Integrasi Data Riil pada Laporan PDF**
    *   Mengganti data *placeholder/mock* pada fitur ekspor PDF bulanan dengan data agregat riil dari database (Mood, Habit, dan Skor SPK).
2.  **Validasi & Koreksi Algoritma SPK**
    *   Melakukan *audit* ulang pada modul *Insight*. Memastikan perhitungan AHP-TOPSIS dan normalisasi skor sudah sesuai dengan spesifikasi SRS (presisi FLOAT8) dan menampilkan rekomendasi yang akurat.
3.  **Perbaikan Skema Migrasi Database**
    *   Mengatasi isu migrasi tabel asesmen yang tereksekusi ulang (*re-run*) setiap kali perintah `docker-compose up` dijalankan.
4.  **Manajemen Sesi Autentikasi Admin**
    *   Memperbaiki logika *middleware* dan *state* UI pada Dashboard Admin. Memastikan dashboard terkunci/redirect ke login secara otomatis saat token sesi habis (*expired*).
5.  **Setup Awal Integrasi Eksternal**
    *   Melakukan konfigurasi dasar dan *environment setup* untuk **Push Notification (FCM)** dan **Google Login** (sebagai persiapan integrasi penuh).

### D. Perbaikan Bug (*Bug Fixes*)
*Fokus: Menyelesaikan hambatan fungsional kritis sebelum uji coba internal.*

1.  **Alur Pembayaran (*Payment Flow*)**
    *   **Masalah:** Klik "Bayar" pada jadwal *pending* langsung mengubah status lunas tanpa menampilkan UI pembayaran.
    *   **Perbaikan:** Memaksa *routing* ke halaman *Payment Gateway* (Mock) sebelum status berubah.
    *   **Visibilitas:** Menyembunyikan jadwal berstatus "Menunggu Pembayaran" dari tampilan depan utama. Menampilkan lencana notifikasi (*badge*) pada menu "Jadwalku" sebagai pengingat.
2.  **Sinkronisasi Profil & *Home Screen***
    *   Memperbaiki data profil pada Beranda yang tidak *update* saat pertama kali dibuka.
    *   Menambahkan fitur *Pull-to-Refresh* pada *Profile Screen* untuk memuat data terbaru secara manual.
    *   Memperbaiki tampilan status "Premium" yang tidak muncul dengan benar di *Profile Screen*.
3.  **Logika Streak & Notifikasi**
    *   Memperbaiki status "Streak hari ini" pada menu Jurnal yang tidak berubah sesuai aktivitas.
    *   Investigasi izin notifikasi yang tidak muncul pada perangkat produksi spesifik (Redmi Note 12 / MIUI).

### E. Riset & Eksplorasi Teknis (*Research*)
*Fokus: Persiapan untuk skalabilitas dan kustomisasi aset di masa depan.*

1.  **Integrasi API Wilayah & Pekerjaan**
    *   Mengeksplorasi penggunaan API publik (misal: API Wilayah Indonesia) untuk input "Domisili Kota" dan "Pekerjaan" guna menstandarisasi data dan mengurangi *typo* pengguna.
2.  **Kustomisasi Aset Android**
    *   Mempelajari kemungkinan teknis implementasi fitur pemilihan ikon badge dan tantangan khusus untuk manajemen Badge dan Tantangan pada dashboard admin.

---

## 2. CATATAN TEKNIS & MITIGASI RISIKO

*   **Prioritas Kritis:** Perbaikan **Alur Pembayaran** dan **Manajemen Sesi Admin** harus diselesaikan di hari pertama (Senin-Selasa) karena menyangkut integritas transaksi dan keamanan akses.
*   **Pendekatan Migrasi:** Untuk isu migrasi yang terulang, tim backend wajib memeriksa *checksum* pada tabel `schema_migrations` dan memastikan file `.sql` tidak dimodifikasi setelah *commit*.
*   **Kompatibilitas Android:** Untuk isu notifikasi di Redmi, tim mobile perlu menambahkan penanganan khusus untuk *Battery Saver* dan *Auto-start permission* yang sering memblokir FCM pada MIUI.

---

## 3. TARGET LUARAN (DELIVERABLES) MINGGU INI

1.  [ ] Modul Admin Dashboard (Asesmen, Feedback, User Details) dapat diakses dan berfungsi.
2.  [ ] Alur pembayaran berjalan sesuai urutan: Pilih Jadwal -> UI Pembayaran -> Status Berubah.
3.  [ ] Laporan PDF berhasil diunduh dengan data riil pengguna.
4.  [ ] Seluruh tombol aksi utama memiliki state *loading* dan *disable*.
5.  [ ] Daftar *bug* pada bagian D terselesaikan 100%.

***

### 💡 Saran Implementasi untuk Tim Developer:
*   **Untuk Admin Payment Logic:** Gunakan *enum* atau *flag* pada tabel `payment_transactions` atau `consultation_schedules` untuk menandai `is_premium_benefit = true/false`. Ini akan memudahkan admin membedakan mana sesi yang dipotong kuota premium dan mana yang uang masuk.
*   **Untuk UI Insight Navbar:** Pastikan perubahan bentuk lingkaran tidak melanggar *Human Interface Guidelines* (iOS) atau *Material Design* (Android) terkait area sentuh (*touch target*) yang minimal 48x48dp.