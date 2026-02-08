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
- `GET /api/produk?name={keyword}` - Search produk by name
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

### Checkout
- `POST /api/checkout` - Proses transaksi (auto update stock)

### Reports
- `GET /api/report/hari-ini` - Laporan penjualan hari ini
- `GET /api/report?start_date={YYYY-MM-DD}&end_date={YYYY-MM-DD}` - Laporan by date range

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

## Checkout Example

Request body untuk `POST /api/checkout`:

```json
{
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 4,
      "quantity": 1
    }
  ]
}
```

Response:

```json
{
  "id": 1,
  "total_amount": 10500,
  "created_at": "2026-02-08T16:05:00Z",
  "details": [
    {
      "id": 1,
      "transaction_id": 1,
      "product_id": 1,
      "product_name": "Teh Pucuk",
      "quantity": 2,
      "subtotal": 7000
    }
  ]
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
