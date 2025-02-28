/*
* @File: repository.warehouse_repository.go
* @Description: Implements warehouse CRUD functions for MySQL
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
*/
package repository

import (
	"errors"
	"net/http"
	"product-service/common"
	"product-service/internal/models"

	"gorm.io/gorm"
)

type WarehouseRepository interface {
	GetAll(paging *common.Paging) ([]models.Warehouses, error)
	GetByID(id interface{}) (*models.Warehouses, error)
	Create(warehouse *models.Warehouses) error
	Update(warehouse *models.Warehouses) error
	Delete(id interface{}) error
}

type WarehouseRepositoryImpl struct {
	DB *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) *WarehouseRepositoryImpl{
	return &WarehouseRepositoryImpl{DB: db}
}

func(repo *WarehouseRepositoryImpl) GetAll(
	paging *common.Paging, 
) ([]models.Warehouses, error){
	var result []models.Warehouses

	sql := repo.DB

	if err := sql.Table(models.Warehouses{}.TableName()).Count(&paging.Total).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(
				http.StatusBadRequest,
				err,
				"Table not found",
				"TABLE_NOT_FOUND",
				"The table warehouse does not exits",
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

	if err := sql.Order("warehouse_id desc").
		Offset((paging.Page-1)*paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
			return nil, common.NewAppError(
				http.StatusInternalServerError,
				err,
				"Error retrieving warehouse",
				"RETRIEVING_ERROR",
				"An unexpected error occurred while retrieving the warehouse",
			)
	}

	return result, nil
}

func(repo *WarehouseRepositoryImpl) GetByID(
	id interface{},
) (*models.Warehouses, error){
	var result models.Warehouses

	if err :=  repo.DB.First(&result, id).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(
				http.StatusNotFound,
				err,
				"warehouse not found",
				"warehouse_NOT_FOUND",
				"The warehouse with the given ID doesn't exists",
			)
		}
		return nil, common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Error retrieve warehouse",
			"Database error while retrieving warehouse",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return &result, nil
}

func(repo *WarehouseRepositoryImpl) Create(item *models.Warehouses) error{
	if err := repo.DB.Create(item).Error; err != nil {
		return common.NewAppError(
			http.StatusInternalServerError, 
			err,
			"Failed to create warehouse",
            "Database error while creating warehouse",
            "INTERNAL_SERVER_ERROR", 
		)
	}
	return nil
}

func(repo *WarehouseRepositoryImpl) Update(item *models.Warehouses) error {
	if err := repo.DB.Where("warehouse_id = ?", item.Warehouse_ID).Save(item).Error; err != nil{
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Fail to update warehouse",
			"Database error while updating warehouse",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return nil
}
	
func(repo *WarehouseRepositoryImpl) Delete(id interface{}) error {
	if err := repo.DB.Delete(models.Warehouses{}, id).Error; err != nil{
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Fail to delete warehouse",
			"Database error while deleting warehouse",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return nil
}