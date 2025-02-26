package repository

import (
	"product-service/common"
	"product-service/internal/models"
)

type CategoryRepository interface {
	GetAll(paging *common.Paging) ([]models.Category, error)
	GetByID(id interface{}) (*models.Category, error)
	Create(category *models.AddCategory) error
	Update(category *models.Category) error
	Delete(id interface{}) error
}