package business

import (
	"product-service/common"
	"product-service/internal/models"
)

type ProductServiceInterface interface {
	GetAllProducts(filter *common.Filter, paging *common.Paging) ([]models.Product, error)
	GetProductByID(productID string) (*models.Product, error)
	CreateProduct(product *models.Product) (error)
	UpdateProduct(product *models.Product) (error)
	DeleteProduct(productID string) (error)
}