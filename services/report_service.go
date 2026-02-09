package services

import (
	"backend-api/models"
	"backend-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

// GetTransactionReportService - Mengambil laporan transaksi dengan filter tanggal opsional
func (rs *ReportService) GetTransactionReportService(startDate, endDate string) (*models.TransactionReport, error) {
	return rs.repo.GetTransactionReportRepo(startDate, endDate)
}
