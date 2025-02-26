// @File business.category_service_interface.go
// @Description Implements category CRUD functions for CategoryService
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package business

import (
	"product-service/common"
	"product-service/internal/models"
)

type CategoryService interface {
	GetAllCategories(paging *common.Paging) ([]models.Category, error)
	GetCategoryByID(id interface{}) (*models.Category, error)
	CreateCategory(item *models.AddCategory) error
	UpdateCategory(item *models.Category) error
	DeleteCategory(id interface{}) error
}