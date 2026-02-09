package services

import (
	"backend-api/models"
	"backend-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (ps *ProductService) GetAllProductService(name string) []models.Product {
	return ps.repo.GetAllProductRepo(name)
}

func (ps *ProductService) CreateProductService(product *models.ProductRequest) error {
	return ps.repo.CreateProductRepo(product)
}

func (ps *ProductService) GetProductByCategoryIdService(categoryId string) []models.Product {
	return ps.repo.GetProductByCategoryIdRepo(categoryId)
}

func (ps *ProductService) UpdateProductService(product *models.ProductRequest) error {
	return ps.repo.UpdateProductRepo(product)
}

func (ps *ProductService) DeleteProductService(productId string) error {
	return ps.repo.DeleteProductRepo(productId)
}
