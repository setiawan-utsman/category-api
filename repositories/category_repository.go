package repositories

import (
	"backend-api/models"
	"database/sql"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAllCategoryRepo() []models.Category {
	var categories []models.Category
	rows, err := repo.db.Query("SELECT id, name, description, created_at FROM categories")
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.Description, &category.CreatedAt); err != nil {
			return nil
		}
		categories = append(categories, category)
	}
	return categories
}

func (repo *CategoryRepository) CreateCategoryRepo(category *models.Category) error {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id, created_at"
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.Id, &category.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CategoryRepository) UpdateCategoryRepo(category *models.Category) error {
	query := "UPDATE categories SET name = $1, description = $2 WHERE id = $3"
	_, err := repo.db.Exec(query, category.Name, category.Description, category.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CategoryRepository) DeleteCategoryRepo(id string) error {
	query := "DELETE FROM categories WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
