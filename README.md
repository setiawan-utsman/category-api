# 📦 Golang Category, Product & Transaction API (Clean Architecture)

Backend REST API menggunakan **Go (net/http)** dengan pola design **Service & Repository Pattern**. Project ini mendemonstrasikan integrasi database PostgreSQL (Supabase) dan pengelolaan data yang terstruktur.

---

## 📌 Ketentuan Cabang (Branches)

*   **`crud-no-database`**: Versi stabil dasar. Data masih disimpan di **memory (slice)** dan belum terkoneksi ke database.
*   **`crud-database`**: Versi lanjutan yang sudah **terkoneksi database Supabase**. Menggunakan UUID untuk ID dan mendukung relasi antar tabel.
*   **`report-transaction`**: Fitur terbaru dengan sistem laporan transaksi (Latest Update).


---

## 🚀 Fitur Saat Ini

### 📁 Category Management
*   **Get All Categories**: Mengambil daftar kategori lengkap.
*   **Create Category**: Membuat kategori baru (Auto-generate UUID & CreatedAt).
*   **Update Category**: Memperbarui informasi kategori.
*   **Delete Category**: Menghapus kategori.

### 🍱 Product Management
*   **Get All Products**: Mengambil daftar produk lengkap dengan **JOIN** detail kategori.
*   **Get Products by Category ID**: Filter produk berdasarkan ID kategori tertentu.
*   **Create Product**: Membuat produk baru terintegrasi dengan `category_id`.
*   **Update Product**: Memperbarui data produk.
*   **Delete Product**: Menghapus produk.

### 💰 Transaction Management
*   **Create Transaction**: Membuat transaksi baru dengan multiple items.
*   **Get All Transactions**: Mengambil semua transaksi dengan detail items.
*   **Get Transaction by ID**: Mengambil detail transaksi tertentu.
*   **Auto Stock Deduction**: Otomatis mengurangi stok produk saat transaksi.
*   **Validation**: Validasi ketersediaan stok sebelum transaksi.

### 📊 Report & Analytics
*   **Transaction Report**: Laporan total revenue, jumlah transaksi, dan produk terlaris.
*   **Date Filtering**: Filter laporan berdasarkan rentang tanggal.
*   **Multiple Best Sellers**: Menampilkan semua produk dengan penjualan tertinggi (jika tied).

---

## 🛠️ Tech Stack

*   **Go (Golang)**: Core language & net/http standard library.
*   **PostgreSQL (Supabase)**: Database Cloud untuk penyimpanan data.
*   **Viper**: Management konfigurasi dan environment variables (`.env`).
*   **JSON Encoding / Decoding**: Pertukaran data antar client-server.

---

## 📁 Struktur Folder

```
category-api/
├── database/     # Inisialisasi DB
├── handlers/     # Request handling & Validation
├── services/     # Business logic
├── repositories/ # Query database
├── models/       # Struct definitions
├── untils/       # Helpers (Response JSON)
├── main.go       # Entry point
└── .env          # Database credentials
```

---

## 🌐 API Endpoints

### 📁 Category Endpoints

#### 1. Get All Categories
```
GET /api/categories
```
**Response:**
```json
{
  "status": 200,
  "message": "Categories retrieved successfully",
  "data": [
    {
      "id": "uuid-123",
      "name": "Food",
      "created_at": "2026-01-15T10:00:00Z"
    }
  ]
}
```

#### 2. Create Category
```
POST /api/categories
```
**Request Body:**
```json
{
  "name": "Beverages"
}
```

#### 3. Update Category
```
PUT /api/categories/
```
**Request Body:**
```json
{
  "id": "uuid-123",
  "name": "Fast Food"
}
```

#### 4. Delete Category
```
DELETE /api/categories/{id}
```
**Example:** `DELETE /api/categories/uuid-123`

---

### � Product Endpoints

#### 1. Get All Products
```
GET /api/products
```
**Optional Query Param:**
- `name` - Filter by product name (partial match)

**Example:** `GET /api/products?name=indomie`

**Response:**
```json
{
  "status": 200,
  "message": "Products retrieved successfully",
  "data": [
    {
      "id": "uuid-456",
      "category_id": "uuid-123",
      "name": "Indomie Goreng",
      "price": 3000,
      "stock": 100,
      "category": {
        "id": "uuid-123",
        "name": "Food"
      }
    }
  ]
}
```

#### 2. Get Products by Category
```
GET /api/products/{category_id}
```
**Example:** `GET /api/products/uuid-123`

#### 3. Create Product
```
POST /api/products
```
**Request Body:**
```json
{
  "category_id": "uuid-123",
  "name": "Indomie Goreng",
  "price": 3000,
  "stock": 100
}
```

#### 4. Update Product
```
PUT /api/products/
```
**Request Body:**
```json
{
  "id": "uuid-456",
  "category_id": "uuid-123",
  "name": "Indomie Soto",
  "price": 3500,
  "stock": 80
}
```

#### 5. Delete Product
```
DELETE /api/products/{id}
```
**Example:** `DELETE /api/products/uuid-456`

---

### 💰 Transaction Endpoints

#### 1. Create Transaction
```
POST /api/transactions
```
**Request Body:**
```json
[
  {
    "product_id": "uuid-456",
    "quantity": 5
  },
  {
    "product_id": "uuid-789",
    "quantity": 3
  }
]
```

**Response:**
```json
{
  "status": 200,
  "message": "Transaction created successfully",
  "data": {
    "id": 1,
    "total_amount": 25000,
    "created_at": "2026-02-08T10:00:00Z",
    "details": [
      {
        "id": 1,
        "transaction_id": 1,
        "product_id": "uuid-456",
        "product_name": "Indomie Goreng",
        "quantity": 5,
        "subtotal": 15000
      }
    ]
  }
}
```

#### 2. Get All Transactions
```
GET /api/transactions
```

#### 3. Get Transaction by ID
```
GET /api/transactions/{id}
```
**Example:** `GET /api/transactions/1`

---

### 📊 Report Endpoint

#### Get Transaction Report
```
GET /api/report
```

**Optional Query Params:**
- `start_date` - Filter dari tanggal (format: YYYY-MM-DD)
- `end_date` - Filter sampai tanggal (format: YYYY-MM-DD)

**Example:** 
- All time: `GET /api/report`
- With filter: `GET /api/report?start_date=2026-01-01&end_date=2026-02-01`

**Response:**
```json
{
  "status": 200,
  "message": "Report retrieved successfully",
  "data": {
    "total_revenue": 150000,
    "total_transaksi": 10,
    "produk_terlaris": [
      {
        "nama": "Indomie Goreng",
        "qty_terjual": 25
      },
      {
        "nama": "Mie Sedaap",
        "qty_terjual": 25
      }
    ]
  }
}
```

**Note:** Jika ada lebih dari 1 produk dengan qty tertinggi yang sama, semua produk akan ditampilkan dalam array `produk_terlaris`.

---

## ▶️ Cara Menjalankan Project

1.  Clone project dan masuk ke branch `report-transaction`.
2.  Siapkan file `.env` dengan format:
    ```env
    PORT=8181
    DB_CONN="postgres://user:pass@host:port/dbname"
    ```
3.  Jalankan perintah:
    ```bash
    go run main.go
    ```

Server akan berjalan di `http://localhost:8181`

---

## 🧪 Testing API

Disarankan menggunakan **Postman** atau **Thunder Client** untuk menguji endpoint. 

### Contoh Flow Testing:
1. Buat category terlebih dahulu
2. Buat product dengan `category_id` dari category yang sudah dibuat
3. Buat transaction dengan `product_id` yang valid
4. Cek report untuk melihat analytics

---

## 📋 Response Format

Semua response mengikuti format standar:
```json
{
  "status": 200,
  "message": "Success message",
  "data": { }
}
```

---

Happy coding 🚀  
Belajar pelan-pelan tapi konsisten 💪
