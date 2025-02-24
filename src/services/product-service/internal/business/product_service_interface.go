/*
* @File: business.product_service_interface.go
* @Description: Implements Product CRUD functions for ProductService
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/

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