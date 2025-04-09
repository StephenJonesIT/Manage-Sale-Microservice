// @File bussiness.product_warehouse_service.go
// @Description Implements CRUD functions for Product Warehouse Service
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package business

import (
	"context"

	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/repository"
)

type ProductWarehouseService interface {
	GetAllProductWarehouses(paging *common.Paging) ([]models.WarehousesList, error)
	GetProductWarehouse(idProduct, idWarehouse interface{}) (*models.ProductWarehouses, error)
	CreateProductWarehouse(ctx context.Context,item *models.ProductWarehouses) error
	UpdateProductWarehouse(ctx context.Context,item *models.ProductWarehouses) error
	DeleteProductWarehouse(idProduct, idWarehouse interface{}) error
}

type ProductWarehouseServiceImpl struct {
	service repository.ProductWarehouseRepo
}

func NewProductWarehouseService(repo repository.ProductWarehouseRepo) *ProductWarehouseServiceImpl {
	return &ProductWarehouseServiceImpl{service: repo}
}

func(biz *ProductWarehouseServiceImpl) GetAllProductWarehouses(paging *common.Paging) ([]models.WarehousesList, error){
	return biz.service.GetAll(paging)
}

func(biz *ProductWarehouseServiceImpl) GetProductWarehouse(idProduct, idWarehouse interface{}) (*models.ProductWarehouses, error){
	return biz.service.GetByID(idProduct,idWarehouse)
}

func(biz *ProductWarehouseServiceImpl) CreateProductWarehouse(ctx context.Context,item *models.ProductWarehouses) error {
	return biz.service.Create(ctx,item)
}

func(biz *ProductWarehouseServiceImpl) UpdateProductWarehouse(ctx context.Context,item *models.ProductWarehouses) error {
	return biz.service.Update(ctx,item)
}

func(biz *ProductWarehouseServiceImpl) DeleteProductWarehouse(idProduct, idWarehouse interface{}) error{
	return biz.service.Delete(idProduct, idWarehouse)
}