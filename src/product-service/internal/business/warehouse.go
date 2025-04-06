/*
* @File: business.warehouse_service.go
* @Description: Implements Warehouse CRUD functions for ProductService
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
 */
package business

import (
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/repository"
)

type WarehouseService interface{
	GetAllWarehouses(paging *common.Paging) ([]models.Warehouses, error)
	GetWarehouseByID(id interface{}) (*models.Warehouses, error)
	CreateWarehouse(item *models.Warehouses) error
	UpdateWarehouse(item *models.Warehouses) error
	DeleteWarehouse(id interface{}) error
}

type WarehouseServiceImpl struct {
	warehouseRepo repository.WarehouseRepository
}

func NewWarehouseService(repo repository.WarehouseRepository) *WarehouseServiceImpl{
	return &WarehouseServiceImpl{warehouseRepo: repo}
}

func(service *WarehouseServiceImpl)GetAllWarehouses(paging *common.Paging) ([]models.Warehouses, error) {
	return service.warehouseRepo.GetAll(paging)
}

func(service *WarehouseServiceImpl)	GetWarehouseByID(id interface{}) (*models.Warehouses, error) {
	return service.warehouseRepo.GetByID(id)
}

func(service *WarehouseServiceImpl)	CreateWarehouse(item *models.Warehouses) error {
	return service.warehouseRepo.Create(item)
}

func(service *WarehouseServiceImpl) UpdateWarehouse(item *models.Warehouses) error {
	return service.warehouseRepo.Update(item)
}
	
func(service *WarehouseServiceImpl) DeleteWarehouse(id interface{}) error{
	return service.warehouseRepo.Delete(id)
}