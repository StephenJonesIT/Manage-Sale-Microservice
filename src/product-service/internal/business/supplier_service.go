/*
* @File: business.product_service.go
* @Description: Implements Supplier CRUD functions for SupplierService
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/
package business

import (
	"product-service/common"
	"product-service/internal/models"
	"product-service/internal/repository"
)

type SupplierService interface{
	GetAllSuppliers(paging *common.Paging) ([]models.Supplier, error)
	GetSupplierByID(id interface{}) (*models.Supplier, error)
	CreateSupplier(item *models.AddSupplier) error
	UpdateSupplier(item *models.Supplier) error
	DeleteSupplier(id interface{}) error
}

type SupplierServiceImpl struct {
	supplierRepo repository.SupplierRepo
}

func NewSupplierService(repo repository.SupplierRepo) *SupplierServiceImpl{
	return &SupplierServiceImpl{supplierRepo: repo}
}

func (service *SupplierServiceImpl) GetAllSuppliers(paging *common.Paging) ([]models.Supplier, error){
	return service.supplierRepo.GetAll(paging)
}

func (service *SupplierServiceImpl) GetSupplierByID(SupplierID interface{}) (*models.Supplier, error){
	return service.supplierRepo.GetById(SupplierID)
}

func (service *SupplierServiceImpl) CreateSupplier(addSupplier *models.AddSupplier) error{
	return service.supplierRepo.Create(addSupplier)
}

func (service *SupplierServiceImpl) UpdateSupplier(updateSupplier *models.Supplier) error{
	return service.supplierRepo.Update(updateSupplier)
}

func (service *SupplierServiceImpl) DeleteSupplier(supplierID interface{}) error{
	return service.supplierRepo.Delete(supplierID)
}