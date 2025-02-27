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
	Create(item *models.Supplier) error
	Delete(id interface{}) error
	Update(item *models.Supplier) error
}

type SupplierRepoImpl struct {
	DB *gorm.DB
}

func NewRepoSupplier(db *gorm.DB) *SupplierRepoImpl {
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

