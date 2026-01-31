package repositories

import (
	"backend-api/models"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) GetAllProductRepo() []models.Product {
	var products []models.Product

	// Gunakan SELECT dengan join untuk mengambil detail kategori
	query := `
		SELECT 
			p.id, p.created_at, p.category_id, p.name, p.price, p.stock, 
			c.id, c.name, c.description 
		FROM products p 
		JOIN categories c ON c.id = p.category_id
	`
	rows, err := pr.db.Query(query)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		// Scan data dari products dan categories
		err := rows.Scan(
			&product.Id,
			&product.CreatedAt,
			&product.CategoryId,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Category.Id,
			&product.Category.Name,
			&product.Category.Description,
		)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			continue
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error after iterating rows: %v\n", err)
		return nil
	}

	fmt.Printf("Successfully fetched %d products with category details\n", len(products))
	return products
}

func (repo *ProductRepository) CreateProductRepo(product *models.ProductRequest) error {
	// Query INSERT dengan kolom dan placeholder yang pasti sejajar
	query := "INSERT INTO products (category_id, name, price, stock) VALUES ($1, $2, $3, $4) RETURNING id, created_at"

	err := repo.db.QueryRow(
		query,
		product.CategoryId,
		product.Name,
		product.Price,
		product.Stock,
	).Scan(&product.Id, &product.CreatedAt)

	if err != nil {
		fmt.Printf("Error creating product details: [Columns: 5, Values: 5] Error: %v\n", err)
		return err
	}

	fmt.Printf("Product created successfully with ID: %s\n", product.Id)
	return nil
}

func (repo *ProductRepository) GetProductByCategoryIdRepo(categoryId string) []models.Product {
	var products []models.Product

	// Gunakan SELECT dengan join dan WHERE clause
	query := `
		SELECT 
			p.id, p.created_at, p.category_id, p.name, p.price, p.stock, 
			c.id, c.name, c.description 
		FROM products p 
		JOIN categories c ON c.id = p.category_id
		WHERE p.category_id = $1
	`
	rows, err := repo.db.Query(query, categoryId)
	if err != nil {
		fmt.Printf("Error executing query for category %s: %v\n", categoryId, err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.CreatedAt,
			&product.CategoryId,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Category.Id,
			&product.Category.Name,
			&product.Category.Description,
		)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			continue
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error after iterating rows: %v\n", err)
		return nil
	}

	fmt.Printf("Successfully fetched %d products for category %s with category details\n", len(products), categoryId)
	return products
}

func (repo *ProductRepository) UpdateProductRepo(product *models.ProductRequest) error {
	query := "UPDATE products SET category_id = $1, name = $2, price = $3, stock = $4 WHERE id = $5"
	_, err := repo.db.Exec(query, product.CategoryId, product.Name, product.Price, product.Stock, product.Id)
	if err != nil {
		fmt.Printf("Error updating product details: [Columns: 5, Values: 5] Error: %v\n", err)
		return err
	}
	fmt.Printf("Product updated successfully with ID: %s\n", product.Id)
	return nil
}

func (repo *ProductRepository) DeleteProductRepo(productId string) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := repo.db.Exec(query, productId)
	if err != nil {
		fmt.Printf("Error deleting product details: [Columns: 1, Values: 1] Error: %v\n", err)
		return err
	}
	fmt.Printf("Product deleted successfully with ID: %s\n", productId)
	return nil
}
