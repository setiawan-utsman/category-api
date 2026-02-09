package handlers

import (
	"backend-api/services"
	"backend-api/untils"
	"net/http"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

// ResponseSuccess - Response untuk sukses dengan data
func (rh *ReportHandler) ResponseSuccess(w http.ResponseWriter, data any, message string) {
	untils.JSONRespon(w, http.StatusOK, data, message)
}

// ResponseError - Response untuk error
func (rh *ReportHandler) ResponseError(w http.ResponseWriter, message string, statusCode int) {
	untils.JSONRespon(w, statusCode, nil, message)
}

// GetReportHandler - Menangani permintaan GET untuk mengambil laporan transaksi
func (rh *ReportHandler) GetReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rh.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ambil filter tanggal opsional dari query parameter
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	// Ambil laporan
	report, err := rh.service.GetTransactionReportService(startDate, endDate)
	if err != nil {
		rh.ResponseError(w, "Failed to get report: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Response sukses dengan data
	rh.ResponseSuccess(w, report, "Report retrieved successfully")
}
