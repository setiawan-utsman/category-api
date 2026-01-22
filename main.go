package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var categories = []Category{
	{Id: 1, Name: "Category 1", Description: "Description 1"},
	{Id: 2, Name: "Category 2", Description: "Description 2"},
	{Id: 3, Name: "Category 3", Description: "Description 3"},
}

func main() {
	// Localhost:8181/api/categories
	// getAllCategory()

	// Localhost:8181/api/categories
	// createCategory()
	http.HandleFunc("/api/categories", categoryHandler)
	http.HandleFunc("/api/categories/", categoryIdHandler)

	fmt.Println("Server Running di localhost:8181")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Gagal Running Server")
	}
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getAllCategory(w, r)
	case http.MethodPost:
		createCategory(w, r)
	}
}

func categoryIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getCategoryById(w, r)
	case http.MethodPut:
		updateCategory(w, r)
	case http.MethodDelete:
		deleteCategory(w, r)
	}
}

func getAllCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var category Category
	json.NewDecoder(r.Body).Decode(&category)

	category.Id = len(categories) + 1
	categories = append(categories, category)

	response := Response{
		Status:  "success",
		Message: "Category created successfully",
		Data:    category,
	}

	json.NewEncoder(w).Encode(response)
}

func getCategoryById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, category := range categories {
		if category.Id == id {
			json.NewEncoder(w).Encode(category)
			return
		}
	}
	http.Error(w, "Category Not Found", http.StatusNotFound)
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var category Category
	json.NewDecoder(r.Body).Decode(&category)

	for i, c := range categories {
		if c.Id == id {
			category.Id = c.Id
			categories[i] = category

			response := Response{
				Status:  "success",
				Message: "Category updated successfully",
				Data:    category,
			}

			json.NewEncoder(w).Encode(response)
			return
		}
	}
	http.Error(w, "Category Not Found", http.StatusNotFound)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, category := range categories {
		if category.Id == id {
			categories = append(categories[:i], categories[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted successfully"})
			return
		}
	}
	http.Error(w, "Category Not Found", http.StatusNotFound)
}
