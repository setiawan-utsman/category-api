package handlers

import (
	"backend-api/models"
	"backend-api/services"
	"backend-api/untils"
	"encoding/json"
	"net/http"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// ResponseSuccess - Response untuk sukses dengan data
func (ph *ProductHandler) ResponseSuccess(w http.ResponseWriter, data any, message string) {
	untils.JSONRespon(w, http.StatusOK, data, message)
}

// ResponseError - Response untuk error
func (ph *ProductHandler) ResponseError(w http.ResponseWriter, message string, statusCode int) {
	untils.JSONRespon(w, statusCode, nil, message)
}

// ResponseNull - Response untuk data null/tidak ada
func (ph *ProductHandler) ResponseNull(w http.ResponseWriter, message string) {
	untils.JSONRespon(w, http.StatusOK, nil, message)
}

func (ph *ProductHandler) ProductHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ph.getAllProduct(w, r)
	case http.MethodPost:
		ph.createProduct(w, r)
	default:
		ph.ResponseError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ph *ProductHandler) ProductIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ph.getAllProductByCategoryId(w, r)
	case http.MethodPut:
		ph.updateProduct(w, r)
	case http.MethodDelete:
		ph.deleteProduct(w, r)
	default:
		ph.ResponseError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ph *ProductHandler) getAllProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	products := ph.service.GetAllProductService(name)

	// Cek jika data kosong/null
	if len(products) == 0 {
		ph.ResponseNull(w, "No products found")
		return
	}

	// Response sukses dengan data
	ph.ResponseSuccess(w, products, "Products retrieved successfully")
}

func (ph *ProductHandler) createProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var product models.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		ph.ResponseError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi
	if product.CategoryId == "" {
		ph.ResponseError(w, "Category ID is required", http.StatusBadRequest)
		return
	}
	if product.Name == "" {
		ph.ResponseError(w, "Name is required", http.StatusBadRequest)
		return
	}
	// End

	err = ph.service.CreateProductService(&product)
	if err != nil {
		ph.ResponseError(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ph.ResponseSuccess(w, product, "Product created successfully")
}

func (ph *ProductHandler) getAllProductByCategoryId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ambil categoryId dari URL path (e.g., /api/products/uuid)
	categoryId := strings.TrimPrefix(r.URL.Path, "/api/products/")
	if categoryId == "" {
		ph.ResponseError(w, "Category ID is required", http.StatusBadRequest)
		return
	}

	products := ph.service.GetProductByCategoryIdService(categoryId)

	// Cek jika data kosong/null
	if len(products) == 0 {
		ph.ResponseNull(w, "No products found for this category")
		return
	}

	// Response sukses dengan data
	ph.ResponseSuccess(w, products, "Products retrieved successfully for category")
}

func (ph *ProductHandler) updateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// // Ambil categoryId dari URL path (e.g., /api/products/uuid)
	// categoryId := strings.TrimPrefix(r.URL.Path, "/api/products/")
	// if categoryId == "" {
	// 	ph.ResponseError(w, "Category ID is required", http.StatusBadRequest)
	// 	return
	// }

	var product models.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		ph.ResponseError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi
	if product.CategoryId == "" {
		ph.ResponseError(w, "Category ID is required", http.StatusBadRequest)
		return
	}
	if product.Name == "" {
		ph.ResponseError(w, "Name is required", http.StatusBadRequest)
		return
	}
	// End

	err = ph.service.UpdateProductService(&product)
	if err != nil {
		ph.ResponseError(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ph.ResponseSuccess(w, product, "Product updated successfully")
}

func (ph *ProductHandler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ambil categoryId dari URL path (e.g., /api/products/uuid)
	id := strings.TrimPrefix(r.URL.Path, "/api/products/")

	if id == "" {
		ph.ResponseError(w, "ID is required", http.StatusBadRequest)
		return
	}

	err := ph.service.DeleteProductService(id)
	if err != nil {
		ph.ResponseError(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ph.ResponseSuccess(w, nil, "Product deleted successfully")
}
