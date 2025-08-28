# Redikru - Intern Backend Engineer (Golang) Test Submission

Ini adalah proyek submission untuk technical test posisi Backend Engineer Intern di Redikru. Proyek ini dibangun dengan fokus pada kualitas kode, performa, keamanan, dan kemudahan penggunaan.

## Deskripsi Singkat

REST API untuk platform lowongan pekerjaan yang memungkinkan pengguna untuk membuat dan melihat daftar pekerjaan. Aplikasi ini dibangun dengan Go (Gin) dan dirancang dengan arsitektur berlapis yang bersih (Handler, Service, Repository) untuk memastikan kode yang *scalable* dan mudah dikelola.

## Fitur Unggulan

- **API Lengkap**: Menyediakan endpoint untuk `POST` dan `GET` data pekerjaan.
- **Pencarian & Filter**: Mendukung pencarian berdasarkan kata kunci dan filter berdasarkan nama perusahaan.
- **Performa Tinggi**: Menggunakan **Redis** sebagai *cache* untuk mempercepat respons pada endpoint `GET /jobs`, mengurangi beban pada database secara signifikan.
- **Keamanan**:
  - **Pencegahan Stored XSS**: Input dari pengguna dibersihkan (*sanitized*) menggunakan Bluemonday untuk mencegah serangan XSS.
  - **Validasi Input**: Memastikan integritas data dengan validasi format UUID pada *request payload*.
- **Dokumentasi API**: Dokumentasi interaktif yang lengkap disediakan menggunakan **Swagger**.
- **Siap Jalan dengan Docker**: Seluruh aplikasi dan layanannya (PostgreSQL, Redis) sudah dibungkus dalam **Docker**, memungkinkan setup yang instan dengan satu perintah.
- **Unit Test**: *Service layer* sudah dicakup oleh *unit test* yang solid menggunakan Testify dan Miniredis.

## Tech Stack

- **Bahasa**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Caching**: Redis
- **Kontainerisasi**: Docker & Docker Compose
- **Dokumentasi**: Swaggo


## Cara Menjalankan

Ada dua cara untuk menjalankan proyek ini.

### Cara 1: Menjalankan dengan Docker 


**Langkah-langkah:**

1.  **Clone repository ini.**

2.  **Jalankan dengan Docker Compose:**
    Dari direktori utama proyek, jalankan satu perintah berikut:
    ```bash
    docker-compose up --build
    ```

### Cara 2: Menjalankan Secara Lokal (Untuk Development)

**Prasyarat:**
- Go (versi 1.24 atau lebih baru)
- PostgreSQL 
- Redis 

**Langkah-langkah:**

1.  **Clone repository ini.**

2.  **Buat dan Konfigurasi File `.env`:**
    Salin file `.env.example` menjadi `.env`.
    ```bash
    cp .env.example .env
    ```
    Kemudian, sesuaikan isi file `.env` dengan konfigurasi PostgreSQL dan Redis lokal anda. Contohnya:
    ```
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASS=root
    DB_NAME=test_BE
    DB_PORT=5432
    DB_SSLMODE=disable

    REDIS_ADDR=localhost:6379
    REDIS_PASSWORD=
    REDIS_DB=0
    ```

3.  **Install Dependensi:**
    ```bash
    go mod tidy
    ```

4.  **Jalankan Aplikasi:**
    ```bash
    go run main.go
    ```

---

## Akses Aplikasi

- **API Server**: `http://localhost:8080`
- **Dokumentasi API (Swagger)**: `http://localhost:8080/swagger/index.html`

---
