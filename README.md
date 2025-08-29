# Redikru - Intern Backend Engineer (Golang) Test Submission

Ini adalah proyek submission untuk technical test posisi Backend Engineer Intern di Redikru. Proyek ini dibangun dengan fokus pada kualitas kode, performa, keamanan, dan kemudahan penggunaan.

## Deskripsi Singkat

REST API untuk platform lowongan pekerjaan yang memungkinkan pengguna untuk membuat dan melihat daftar pekerjaan. Aplikasi ini dibangun dengan Go (Gin) dan dirancang dengan arsitektur berlapis yang bersih (Handler, Service, Repository) untuk memastikan kode yang *scalable* dan mudah dikelola.

## Solusi & Fitur Utama

* **API & Pencarian Cerdas**: Menyediakan endpoint `POST` dan `GET`. Endpoint `GET` mendukung **filter `keyword` dan `companyName`** untuk pencarian data yang cepat dan relevan.
* **Performa untuk Skala Besar**:
    * **Pagination**: Mengatasi "ribuan data" dengan membatasi jumlah data yang diambil dari database dalam satu waktu (`page` & `limit`).
    * **Redis Caching**: Mengatasi "banyak pengguna" dengan menyimpan data di *cache*, sehingga respons API menjadi super cepat dan beban database berkurang drastis.
* **Keamanan Anti XSS**: Mencegah serangan *XSS* dengan **membersihkan (sanitasi)** semua input dari pengguna sebelum disimpan ke database menggunakan Bluemonday.
* **Membuat Unit Test**: Mencakup logika bisnis dan *caching* pada *service layer*.
* **Dokumentasi API (Swagger)**: Dilengkapi dokumentasi API interaktif menggunakan **Swagger** untuk memudahkan penggunaan dan pemahaman endpoint yang ada.
* **Setup Docker**: Memungkinkan seluruh aplikasi dan layanannya dijalankan dengan satu perintah mudah.

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
