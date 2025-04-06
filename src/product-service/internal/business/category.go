// @File business.category_service.go
// @Description Implements category CRUD functions for CategoryService
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package business

import (
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/repository"
)

type CategoryService interface {
	GetAllCategories(paging *common.Paging) ([]models.Category, error)
	GetCategoryByID(id interface{}) (*models.Category, error)
	CreateCategory(item *models.AddCategory) error
	UpdateCategory(item *models.Category) error
	DeleteCategory(id interface{}) error
}

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

