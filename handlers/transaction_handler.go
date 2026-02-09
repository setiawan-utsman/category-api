package handlers

import (
	"backend-api/models"
	"backend-api/services"
	"backend-api/untils"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// ResponseSuccess - Response untuk sukses dengan data
func (th *TransactionHandler) ResponseSuccess(w http.ResponseWriter, data any, message string) {
	untils.JSONRespon(w, http.StatusOK, data, message)
}

// ResponseError - Response untuk error
func (th *TransactionHandler) ResponseError(w http.ResponseWriter, message string, statusCode int) {
	untils.JSONRespon(w, statusCode, nil, message)
}

// ResponseNull - Response untuk data null/tidak ada
func (th *TransactionHandler) ResponseNull(w http.ResponseWriter, message string) {
	untils.JSONRespon(w, http.StatusOK, nil, message)
}

// TransactionHandler - Handler utama untuk /api/transactions
func (th *TransactionHandler) TransactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		th.GetAllTransactionHandler(w, r)
	case http.MethodPost:
		th.CreateTransactionHandler(w, r)
	default:
		th.ResponseError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// TransactionIdHandler - Handler utama untuk /api/transactions/ dengan ID
func (th *TransactionHandler) TransactionIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		th.GetTransactionByIDHandler(w, r)
	default:
		th.ResponseError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// CreateTransactionHandler - Menangani permintaan POST untuk membuat transaksi
func (th *TransactionHandler) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var items []models.TransactionDetail

	// Decode body permintaan
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		th.ResponseError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi item
	if len(items) == 0 {
		th.ResponseError(w, "Items cannot be empty", http.StatusBadRequest)
		return
	}

	// Buat transaksi
	transaction, err := th.service.CreateTransactionService(items)
	if err != nil {
		th.ResponseError(w, "Failed to create transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Kembalikan respons sukses
	th.ResponseSuccess(w, transaction, "Transaction created successfully")
}

// GetAllTransactionHandler - Menangani permintaan GET untuk mengambil semua transaksi
func (th *TransactionHandler) GetAllTransactionHandler(w http.ResponseWriter, r *http.Request) {
	transactions, err := th.service.GetAllTransactionService()
	if err != nil {
		th.ResponseError(w, "Failed to get transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Cek jika data kosong/null
	if len(transactions) == 0 {
		th.ResponseNull(w, "No transactions found")
		return
	}

	// Response sukses dengan data
	th.ResponseSuccess(w, transactions, "Transactions retrieved successfully")
}

// GetTransactionByIDHandler - Menangani permintaan GET untuk mengambil transaksi berdasarkan ID
func (th *TransactionHandler) GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari URL path (contoh: /api/transactions/1)
	idStr := strings.TrimPrefix(r.URL.Path, "/api/transactions/")
	if idStr == "" || idStr == "/api/transactions" {
		th.ResponseError(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	// Konversi ID ke integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		th.ResponseError(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	// Ambil transaksi
	transaction, err := th.service.GetTransactionByIDService(id)
	if err != nil {
		th.ResponseError(w, "Transaction not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Response sukses dengan data
	th.ResponseSuccess(w, transaction, "Transaction retrieved successfully")
}
