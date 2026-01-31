# 📦 Golang Category API (Beginner Friendly)

Simple backend REST API menggunakan **Go (net/http)** untuk CRUD **Category**.
Project ini cocok untuk **pemula yang baru belajar Golang backend** tanpa framework tambahan.

---

## 📌 Ketentuan Cabang (Branches)

*   **`main`**: Versi stabil dasar. Data masih disimpan di **memory (slice)** dan belum terkoneksi ke database.
*   **`crud-database`**: Versi lanjutan yang sudah **terkoneksi database Supabase**. Menggunakan UUID untuk ID dan mendukung relasi antar tabel.


## 🚀 Fitur

* Get All Categories
* Get Category by ID
* Create Category
* Update Category
* Delete Category
* Response JSON standar (`status`, `message`, `data`)

---

## 🛠️ Tech Stack

* Go (Golang)
* net/http (standard library)
* JSON Encoding / Decoding

---

## 📁 Struktur Folder

```
category-api/
├── main.go
├── go.mod
└── README.md
```

---

## 📦 Model Category

```go
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
```

---

## 🔁 Format Response API

Semua response menggunakan format yang konsisten:

```json
{
  "status": "success",
  "message": "Category created successfully",
  "data": {
    "id": 1,
    "name": "Makanan",
    "description": "Produk makanan"
  }
}
```

---

## 🌐 Endpoint API

### 🔹 Get All Categories

```
GET /api/categories
```

### 🔹 Get Category by ID

```
GET /api/categories?id=1
```

### 🔹 Create Category

```
POST /api/categories
```

Body JSON:

```json
{
  "name": "Minuman",
  "description": "Produk minuman"
}
```

### 🔹 Update Category

```
PUT /api/categories?id=1
```

Body JSON:

```json
{
  "name": "Snack",
  "description": "Makanan ringan"
}
```

### 🔹 Delete Category

```
DELETE /api/categories?id=1
```

---

## ▶️ Cara Menjalankan Project

### 1️⃣ Pastikan Go Terinstall

```bash
go version
```

### 2️⃣ Masuk ke Folder Project

```bash
cd category-api
```

### 3️⃣ Jalankan Server

```bash
go run main.go
```

Jika berhasil:

```
Server running at http://localhost:8181
```

---

## 🧪 Testing API

Gunakan:

* Postman
* Thunder Client (VS Code)
* curl

---

## 🎯 Tujuan Project Ini

* Memahami dasar **routing di Golang**
* Memahami **HTTP Method (GET, POST, PUT, DELETE)**
* Belajar **struktur backend sederhana**
* Menyiapkan pondasi sebelum memakai framework (Gin / Fiber)

---

## 📌 Catatan

* Data masih disimpan di memory (slice)
* Belum menggunakan database
* Cocok untuk belajar konsep dasar backend

---

## 📚 Next Step (Recommended)

* Gunakan **Gin / Fiber**
* Tambahkan **Database (MySQL / PostgreSQL)**
* Implement **Service & Repository Pattern**
* Tambahkan **Middleware**

---

Happy coding 🚀
Belajar pelan-pelan tapi konsisten 💪
