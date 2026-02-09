package models

// TransactionReport - Model respons laporan
type TransactionReport struct {
	TotalRevenue   int                  `json:"total_revenue"`
	TotalTransaksi int                  `json:"total_transaksi"`
	ProdukTerlaris []BestSellingProduct `json:"produk_terlaris,omitempty"`
}

// BestSellingProduct - Model produk terlaris
type BestSellingProduct struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
