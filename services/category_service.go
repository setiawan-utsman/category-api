package services

import (
	"backend-api/models"
	"backend-api/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (cs *CategoryService) GetAllCategoryService() []models.Category {
	return cs.repo.GetAllCategoryRepo()
}

func (cs *CategoryService) CreateCategoryService(category *models.Category) error {
	return cs.repo.CreateCategoryRepo(category)
}

func (cs *CategoryService) UpdateCategoryService(category *models.Category) error {
	return cs.repo.UpdateCategoryRepo(category)
}

func (cs *CategoryService) DeleteCategoryService(categoryId string) error {
	return cs.repo.DeleteCategoryRepo(categoryId)
}
