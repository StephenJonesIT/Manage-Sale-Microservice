package business

import (
	"product-service/common"
	"product-service/internal/models"
	"product-service/internal/repository"
)

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