# 📦 Golang Category & Product API (Clean Architecture)

Backend REST API menggunakan **Go (net/http)** dengan pola design **Service & Repository Pattern**. Project ini mendemonstrasikan integrasi database PostgreSQL (Supabase) dan pengelolaan data yang terstruktur.

---

## 📌 Ketentuan Cabang (Branches)

*   **`crud-no-database`**: Versi stabil dasar. Data masih disimpan di **memory (slice)** dan belum terkoneksi ke database.
*   **`crud-database`**: Versi lanjutan yang sudah **terkoneksi database Supabase**. Menggunakan UUID untuk ID dan mendukung relasi antar tabel.

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

## 📦 Model Utama

### Product Request Pattern
```go
type Product struct {
	Id         string    `json:"id"`
	CategoryId string    `json:"category_id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	Category   Category  `json:"category"` // Populated via JOIN
}
```

---

## 🌐 Endpoint API Utama

### 🍱 Product
*   `GET /api/products` - List all products (+ details kategori)
*   `GET /api/products/{category_id}` - Filter produk by kategori
*   `POST /api/products` - Create new product
*   `PUT /api/products/` - Update product
*   `DELETE /api/products/{id}` - Delete product

### � Category
*   `GET /api/categories` - List all categories
*   `POST /api/categories` - Create new category
*   `PUT /api/categories/` - Update category
*   `DELETE /api/categories/{id}` - Delete category

---

## ▶️ Cara Menjalankan Project

1.  Clone project dan masuk ke branch `crud-database`.
2.  Siapkan file `.env` dengan format:
    ```env
    PORT=8181
    DB_CONN="postgres://user:pass@host:port/dbname"
    ```
3.  Jalankan perintah:
    ```bash
    go run main.go
    ```

---

## 🧪 Testing API

Disarankan menggunakan **Postman** untuk menguji endpoint. Jangan lupa untuk menyertakan `id` kategori yang benar saat membuat produk demi keamanan relasi data (Foreign Key).

---

Happy coding 🚀
Belajar pelan-pelan tapi konsisten 💪
