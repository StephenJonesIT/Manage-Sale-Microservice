// @File bussiness.product_warehouse_service.go
// @Description Implements CRUD functions for Product Warehouse Service
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package business

import (
	"product-service/common"
	"product-service/internal/models"
	"product-service/internal/repository"
)

type ProductWarehouseService interface {
	GetAllProductWarehouses(paging *common.Paging) ([]models.ProductWarehouses, error)
	GetProductWarehouse(idProduct, idWarehouse interface{}) (*models.ProductWarehouses, error)
	CreateProductWarehouse(item *models.ProductWarehouses) error
	UpdateProductWarehouse(item *models.ProductWarehouses) error
	DeleteProductWarehouse(item *models.ProductWarehouses) error
}

type ProductWarehouseServiceImpl struct {
	service repository.ProductWarehouseRepo
}

func NewProductWarehouseService(repo repository.ProductWarehouseRepo) *ProductWarehouseServiceImpl {
	return &ProductWarehouseServiceImpl{service: repo}
}

func(biz *ProductWarehouseServiceImpl) GetAllProductWarehouses(paging *common.Paging) ([]models.ProductWarehouses, error){
	return biz.service.GetAll(paging)
}

func(biz *ProductWarehouseServiceImpl) GetProductWarehouse(idProduct, idWarehouse interface{}) (*models.ProductWarehouses, error){
	return biz.service.GetByID(idProduct,idWarehouse)
}

func(biz *ProductWarehouseServiceImpl) CreateProductWarehouse(item *models.ProductWarehouses) error {
	return biz.service.Create(item)
}

func(biz *ProductWarehouseServiceImpl) UpdateProductWarehouse(item *models.ProductWarehouses) error {
	return biz.service.Update(item)
}

func(biz *ProductWarehouseServiceImpl) DeleteProductWarehouse(item *models.ProductWarehouses) error{
	return biz.service.Delete(item)
}