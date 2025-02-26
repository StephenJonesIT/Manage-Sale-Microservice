package business

import (
	"product-service/common"
	"product-service/internal/models"
	"product-service/internal/repository"
)

type CategoryServiceImpl struct {
	CategoryRepo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryServiceImpl{
	return &CategoryServiceImpl{CategoryRepo: repo}
}

func(service *CategoryServiceImpl) GetAllCategories(paging *common.Paging) ([]models.Category, error){
	return service.CategoryRepo.GetAll(paging)
}

func(service *CategoryServiceImpl) GetCategoryByID(id interface{}) (*models.Category, error){
	return service.CategoryRepo.GetByID(id)
}

func(service *CategoryServiceImpl) CreateCategory(item *models.AddCategory) error{
	return service.CategoryRepo.Create(item)
}

func(service *CategoryServiceImpl) UpdateCategory(item *models.Category) error{
	return service.CategoryRepo.Update(item)
}

func(service *CategoryServiceImpl) DeleteCategory(id interface{}) error{
	return service.CategoryRepo.Delete(id)
}

