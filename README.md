# Kasir API

Simple REST API buat sistem kasir pake Go dan PostgreSQL (Supabase).

## Setup

### 1. Clone repository
```bash
git clone https://github.com/YOUR_USERNAME/kasir-api.git
cd kasir-api
```

### 2. Buat file `.env`
```
PORT=8080
DB_CONN=your_supabase_connection_string_here
```

### 3. Install dependencies
```bash
go mod download
```

### 4. Jalankan
```bash
go run main.go
```

Server jalan di `http://localhost:8080`

## API Endpoints

### Products
- `GET /api/produk` - Semua produk
- `POST /api/produk` - Buat produk baru
- `GET /api/produk/{id}` - Detail produk (dengan category_name dari JOIN!)
- `PUT /api/produk/{id}` - Update produk
- `DELETE /api/produk/{id}` - Hapus produk

### Categories
- `GET /categories` - Semua kategori
- `POST /categories` - Buat kategori baru
- `GET /categories/{id}` - Detail kategori
- `PUT /categories/{id}` - Update kategori
- `DELETE /categories/{id}` - Hapus kategori

### Health Check
- `GET /health` - Cek status API

## JOIN Challenge ✅

Endpoint `GET /api/produk/{id}` mengembalikan `category_name` dari tabel categories menggunakan LEFT JOIN:

```json
{
  "id": 1,
  "name": "Teh Pucuk",
  "price": 3500,
  "stock": 10,
  "category_id": 1,
  "category_name": "Minuman"
}
```

## Struktur Project

```
kasir-api/
├── database/          # Koneksi database
├── models/           # Data structures
├── repositories/     # SQL queries
├── services/         # Business logic
├── handlers/         # HTTP handlers
├── main.go           # Entry point
├── schema.sql        # Database schema
└── .env             # Config (jangan di-push!)
```

## Deploy

Push ke GitHub, lalu deploy ke Railway/Zeabur. Set environment variables:
- `PORT=8080`
- `DB_CONN=your_supabase_connection_string`

## Tech Stack

- Go 1.21+
- PostgreSQL (Supabase)
- Viper (config)
- pgx/v5 (PostgreSQL driver)
