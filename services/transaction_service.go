package services

import (
	"backend-api/models"
	"backend-api/repositories"
	"fmt"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// CreateTransactionService - Logika bisnis untuk membuat transaksi
func (ts *TransactionService) CreateTransactionService(items []models.TransactionDetail) (*models.Transaction, error) {
	var totalAmount int
	var details []models.TransactionDetail

	// Loop melalui semua item
	for _, item := range items {
		// Ambil detail produk
		product, err := ts.repo.GetProductByIDRepo(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to get product %s: %v", item.ProductID, err)
		}

		// Periksa apakah stok mencukupi
		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s (available: %d, requested: %d)",
				product.Name, product.Stock, item.Quantity)
		}

		// Hitung subtotal
		subtotal := product.Price * item.Quantity

		// Tambahkan ke total jumlah
		totalAmount += subtotal

		// Kurangi stok
		err = ts.repo.UpdateProductStockRepo(item.ProductID, item.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to update stock for product %s: %v", product.Name, err)
		}

		// Tambahkan ke detail
		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// Buat transaksi
	transaction := &models.Transaction{
		TotalAmount: totalAmount,
		Details:     details,
	}

	// Masukkan transaksi dan detail transaksi
	err := ts.repo.CreateTransactionRepo(transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	return transaction, nil
}

// GetAllTransactionService - Mengambil semua transaksi
func (ts *TransactionService) GetAllTransactionService() ([]models.Transaction, error) {
	return ts.repo.GetAllTransactionRepo()
}

// GetTransactionByIDService - Mengambil transaksi berdasarkan ID
func (ts *TransactionService) GetTransactionByIDService(id int) (*models.Transaction, error) {
	return ts.repo.GetTransactionByIDRepo(id)
}
