/*
* @File: business.product_service.go
* @Description: Implements Product CRUD functions for ProductService
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
 */

package business

import (
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/repository"
)

type ProductServiceInterface interface {
	GetAllProducts(filter *common.Filter, paging *common.Paging) ([]models.Product, error)
	GetProductByID(productID string) (*models.Product, error)
	CreateProduct(product *models.Product) (error)
	UpdateProduct(product *models.Product) (error)
	DeleteProduct(productID string) (error)
}

type ProductServiceImpl struct {
	ProductRepo repository.ProductRepostitory
}

func NewProductService(repo repository.ProductRepostitory) *ProductServiceImpl{
	return &ProductServiceImpl{ProductRepo: repo}
}

func (service *ProductServiceImpl) GetAllProducts(filter *common.Filter, paging *common.Paging) ([]models.Product, error){
	return service.ProductRepo.GetAll(filter, paging)
}

func (service *ProductServiceImpl) GetProductByID(productID string) (*models.Product, error){
	return service.ProductRepo.GetByID(productID)
}

func (service *ProductServiceImpl) CreateProduct(product *models.Product) error{
	return service.ProductRepo.Create(product)
}

func (service *ProductServiceImpl) UpdateProduct(product *models.Product) error{
	return service.ProductRepo.Update(product)
}

func (service *ProductServiceImpl) DeleteProduct(productID string) error{
	return service.ProductRepo.Delete(productID)
}