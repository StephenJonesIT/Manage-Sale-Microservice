package repository

import (
	"product-service/common"
	"product-service/internal/models"
)

type ProductRepostitory interface {
	GetAll(filter *common.Filter, paging *common.Paging)([]models.Product, error)
	GetByID(productID string)(*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(productID string) error
}