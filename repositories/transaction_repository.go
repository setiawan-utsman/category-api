package repositories

import (
	"backend-api/models"
	"database/sql"
	"fmt"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// GetProductByIDRepo - Mengambil detail produk berdasarkan ID
func (repo *TransactionRepository) GetProductByIDRepo(productID string) (*models.Product, error) {
	var product models.Product
	query := `
		SELECT id, name, price, stock 
		FROM products 
		WHERE id = $1
	`
	err := repo.db.QueryRow(query, productID).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Stock,
	)
	if err != nil {
		return nil, fmt.Errorf("error getting product details: %v", err)
	}
	return &product, nil
}

// UpdateProductStockRepo - Mengurangi stok produk
func (repo *TransactionRepository) UpdateProductStockRepo(productID string, quantity int) error {
	query := `
		UPDATE products 
		SET stock = stock - $1 
		WHERE id = $2 AND stock >= $1
	`
	result, err := repo.db.Exec(query, quantity, productID)
	if err != nil {
		return fmt.Errorf("error updating stock: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking affected rows: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("insufficient stock for product %s", productID)
	}

	return nil
}

// CreateTransactionRepo - Memasukkan transaksi dan detail transaksi
func (repo *TransactionRepository) CreateTransactionRepo(transaction *models.Transaction) error {
	// Mulai transaksi
	tx, err := repo.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}
	defer tx.Rollback()

	// Masukkan ke tabel transactions
	var transactionID int
	query := `
		INSERT INTO transactions (total_amount) 
		VALUES ($1) 
		RETURNING id
	`
	err = tx.QueryRow(query, transaction.TotalAmount).Scan(&transactionID)
	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	// Masukkan detail transaksi
	detailQuery := `
		INSERT INTO transaction_details (transaction_id, product_id, product_name, quantity, subtotal) 
		VALUES ($1, $2, $3, $4, $5)
	`
	for _, detail := range transaction.Details {
		_, err = tx.Exec(
			detailQuery,
			transactionID,
			detail.ProductID,
			detail.ProductName,
			detail.Quantity,
			detail.Subtotal,
		)
		if err != nil {
			return fmt.Errorf("error creating transaction detail: %v", err)
		}
	}

	// Commit transaksi
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	transaction.ID = transactionID
	fmt.Printf("Transaction created successfully with ID: %d\n", transactionID)
	return nil
}

// GetAllTransactionRepo - Mengambil semua transaksi beserta detailnya
func (repo *TransactionRepository) GetAllTransactionRepo() ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := `
		SELECT id, total_amount, created_at 
		FROM transactions 
		ORDER BY created_at DESC
	`
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error getting transactions: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var txn models.Transaction
		err := rows.Scan(&txn.ID, &txn.TotalAmount, &txn.CreatedAt)
		if err != nil {
			fmt.Printf("Error scanning transaction: %v\n", err)
			continue
		}

		// Ambil detail untuk transaksi ini
		details, err := repo.GetTransactionDetailsRepo(txn.ID)
		if err != nil {
			fmt.Printf("Error getting transaction details: %v\n", err)
			continue
		}
		txn.Details = details

		transactions = append(transactions, txn)
	}

	return transactions, nil
}

// GetTransactionDetailsRepo - Mengambil detail transaksi berdasarkan ID transaksi
func (repo *TransactionRepository) GetTransactionDetailsRepo(transactionID int) ([]models.TransactionDetail, error) {
	var details []models.TransactionDetail

	query := `
		SELECT id, transaction_id, product_id, product_name, quantity, subtotal 
		FROM transaction_details 
		WHERE transaction_id = $1
	`
	rows, err := repo.db.Query(query, transactionID)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction details: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var detail models.TransactionDetail
		err := rows.Scan(
			&detail.ID,
			&detail.TransactionID,
			&detail.ProductID,
			&detail.ProductName,
			&detail.Quantity,
			&detail.Subtotal,
		)
		if err != nil {
			fmt.Printf("Error scanning transaction detail: %v\n", err)
			continue
		}
		details = append(details, detail)
	}

	return details, nil
}

// GetTransactionByIDRepo - Mengambil transaksi berdasarkan ID
func (repo *TransactionRepository) GetTransactionByIDRepo(id int) (*models.Transaction, error) {
	var txn models.Transaction

	query := `
		SELECT id, total_amount, created_at 
		FROM transactions 
		WHERE id = $1
	`
	err := repo.db.QueryRow(query, id).Scan(&txn.ID, &txn.TotalAmount, &txn.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, fmt.Errorf("error getting transaction: %v", err)
	}

	// Ambil detail
	details, err := repo.GetTransactionDetailsRepo(txn.ID)
	if err != nil {
		return nil, err
	}
	txn.Details = details

	return &txn, nil
}
