package repositories

import (
	"backend-api/models"
	"database/sql"
	"fmt"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// GetTransactionReportRepo - Mengambil laporan transaksi dengan filter tanggal opsional
func (repo *ReportRepository) GetTransactionReportRepo(startDate, endDate string) (*models.TransactionReport, error) {
	report := &models.TransactionReport{}

	// Susun query dengan filter tanggal opsional
	baseQuery := "FROM transactions WHERE 1=1"
	args := []interface{}{}
	argIndex := 1

	if startDate != "" {
		baseQuery += fmt.Sprintf(" AND DATE(created_at) >= $%d", argIndex)
		args = append(args, startDate)
		argIndex++
	}

	if endDate != "" {
		baseQuery += fmt.Sprintf(" AND DATE(created_at) <= $%d", argIndex)
		args = append(args, endDate)
		argIndex++
	}

	// Ambil total pendapatan dan jumlah transaksi
	query := "SELECT COALESCE(SUM(total_amount), 0), COUNT(*) " + baseQuery
	err := repo.db.QueryRow(query, args...).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, fmt.Errorf("error getting report summary: %v", err)
	}

	// Ambil produk terlaris - semua produk dengan jumlah terbanyak
	// Menggunakan CTE (Common Table Expression) untuk mencari semua produk dengan qty maksimal
	bestSellingQuery := `
		WITH product_totals AS (
			SELECT td.product_name, COALESCE(SUM(td.quantity), 0) as total_qty
			FROM transaction_details td
			JOIN transactions t ON td.transaction_id = t.id
			WHERE 1=1
	`

	if startDate != "" {
		bestSellingQuery += " AND DATE(t.created_at) >= $1"
	}
	if endDate != "" {
		if startDate != "" {
			bestSellingQuery += " AND DATE(t.created_at) <= $2"
		} else {
			bestSellingQuery += " AND DATE(t.created_at) <= $1"
		}
	}

	bestSellingQuery += `
			GROUP BY td.product_name
		)
		SELECT product_name, total_qty
		FROM product_totals
		WHERE total_qty = (SELECT MAX(total_qty) FROM product_totals)
		ORDER BY product_name
	`

	rows, err := repo.db.Query(bestSellingQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("error getting best selling products: %v", err)
	}
	defer rows.Close()

	var bestSellingProducts []models.BestSellingProduct
	for rows.Next() {
		var productName string
		var totalQty int
		err := rows.Scan(&productName, &totalQty)
		if err != nil {
			return nil, fmt.Errorf("error scanning best selling product: %v", err)
		}
		bestSellingProducts = append(bestSellingProducts, models.BestSellingProduct{
			Nama:       productName,
			QtyTerjual: totalQty,
		})
	}

	report.ProdukTerlaris = bestSellingProducts

	return report, nil
}
