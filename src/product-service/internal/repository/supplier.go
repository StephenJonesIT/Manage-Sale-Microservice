// @File repository.supplier_repository.go
// @Description Implements Supplier CRUD functions for MySQL
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package repository

import (
	"errors"
	"net/http"
	"product-service/common"
	"product-service/internal/models"

	"gorm.io/gorm"
)

type SupplierRepo interface {
	GetAll(paging *common.Paging) ([]models.Supplier, error)
	GetById(id interface{}) (*models.Supplier, error) 
	Create(item *models.AddSupplier) error
	Delete(id interface{}) error
	Update(item *models.Supplier) error
}

type SupplierRepoImpl struct {
	DB *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepoImpl {
	return &SupplierRepoImpl{
		DB: db,
	}
}

func(repo *SupplierRepoImpl) GetAll(paging *common.Paging) ([]models.Supplier, error){
	var result []models.Supplier
	
	sql := repo.DB

	if err := sql.Table(models.Supplier{}.TableName()).Count(&paging.Total).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(
				http.StatusBadRequest,
				err,
				"Table not found",
				"TABLE_NOT_FOUND",
				"The table supplier does not exits",
			)
		}

		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, common.NewAppError(
				http.StatusBadRequest,
				err,
				"Invalid data",
				"INVALID_DATA",
				"The data in the table is invalid",
			)
		}

		return nil, common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Fail to count records",
			"COUNT_ERR",
			"An unexpected error occurred while counting the record.",
		)
	}

	if err := sql.Order("created_at DESC").
		Offset((paging.Page-1)*paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
			return nil, common.NewAppError(
				http.StatusInternalServerError,
				err,
				"Error retrieving suppliers",
				"RETRIEVE_ERROR",
				"An unexpected error occurred while retrieving the suppliers.",
			)	
	}

	return result, nil
}

func(repo *SupplierRepoImpl) Create(item *models.AddSupplier) error{
	if err := repo.DB.Create(item).Error; err != nil {
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Fail to create supplier",
			"Database error while createing supplier",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return nil
}

func(repo *SupplierRepoImpl) Update(item *models.Supplier) error{
	if err := repo.DB.Where("supplier_id = ?", item.Supplier_ID).Updates(item).Error; err != nil {
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Fail to update supplier",
			"Database error while updating supplier",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return nil
}

func(repo *SupplierRepoImpl) Delete(id interface{}) error{
	if err := repo.DB.Delete(models.Supplier{}, id).Error; err != nil{
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Fail to delete supplier",
			"Database error while deleting supplier",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return  nil
}

func(repo *SupplierRepoImpl) GetById(id interface{}) (*models.Supplier, error){
	var result models.Supplier
	if err :=  repo.DB.First(&result, id).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(
				http.StatusNotFound,
				err,
				"Supplier not found",
				"SUPPLIER_NOT_FOUND",
				"The supplier with the given ID doesn't exists",
			)
		}
	}
	return &result, nil
}