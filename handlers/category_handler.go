package handlers

import (
	"backend-api/models"
	"backend-api/services"
	"backend-api/untils"
	"encoding/json"
	"net/http"
	"strings"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

// ResponseSuccess - Response untuk sukses dengan data
func (ph *CategoryHandler) ResponseSuccess(w http.ResponseWriter, data any, message string) {
	untils.JSONRespon(w, http.StatusOK, data, message)
}

// ResponseError - Response untuk error
func (ph *CategoryHandler) ResponseError(w http.ResponseWriter, message string, statusCode int) {
	untils.JSONRespon(w, statusCode, nil, message)
}

// ResponseNull - Response untuk data null/tidak ada
func (ph *CategoryHandler) ResponseNull(w http.ResponseWriter, message string) {
	untils.JSONRespon(w, http.StatusOK, nil, message)
}

func (ph *CategoryHandler) CategoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ph.getAllCategory(w, r)
	case http.MethodPost:
		ph.createCategory(w, r)
	default:
		ph.ResponseError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ph *CategoryHandler) CategoryIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		ph.updateCategory(w, r)
	case http.MethodDelete:
		ph.deleteCategory(w, r)
	default:
		ph.ResponseError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (ph *CategoryHandler) getAllCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	categories := ph.service.GetAllCategoryService()
	if len(categories) == 0 {
		ph.ResponseNull(w, "No categories found")
		return
	}

	ph.ResponseSuccess(w, categories, "Categories retrieved successfully")
}

func (ph *CategoryHandler) createCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		ph.ResponseError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validation
	if category.Name == "" {
		ph.ResponseError(w, "Name is required", http.StatusBadRequest)
		return
	}

	if err := ph.service.CreateCategoryService(&category); err != nil {
		ph.ResponseError(w, "Failed to create category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ph.ResponseSuccess(w, category, "Category created successfully")
}

func (ph *CategoryHandler) updateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		ph.ResponseError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validation
	if category.Id == "" {
		ph.ResponseError(w, "Category ID is required", http.StatusBadRequest)
		return
	}
	if category.Name == "" {
		ph.ResponseError(w, "Name is required", http.StatusBadRequest)
		return
	}

	if err := ph.service.UpdateCategoryService(&category); err != nil {
		ph.ResponseError(w, "Failed to update category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ph.ResponseSuccess(w, category, "Category updated successfully")
}

func (ph *CategoryHandler) deleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		ph.ResponseError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	if id == "" {
		ph.ResponseError(w, "ID is required", http.StatusBadRequest)
		return
	}

	if err := ph.service.DeleteCategoryService(id); err != nil {
		ph.ResponseError(w, "Failed to delete category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ph.ResponseSuccess(w, nil, "Category deleted successfully")
}
