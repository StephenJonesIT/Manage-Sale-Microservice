/*
* @File: business.product_service.go
* @Description: Implements category CRUD functions for MySQL
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

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl{
	return &CategoryRepositoryImpl{DB: db}
}

func(repo *CategoryRepositoryImpl) GetAll(
	paging *common.Paging, 
) ([]models.Category, error){
	var result []models.Category

	sql := repo.DB

	if err := sql.Table(models.Category{}.TableName()).Count(&paging.Total).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(
				http.StatusBadRequest,
				err,
				"Table not found",
				"TABLE_NOT_FOUND",
				"The table category does not exits",
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

	if err := sql.Order("category_id desc").
		Offset((paging.Page-1)*paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
			return nil, common.NewAppError(
				http.StatusInternalServerError,
				err,
				"Error retrieving category",
				"RETRIEVING_ERROR",
				"An unexpected error occurred while retrieving the category",
			)
	}

	return result, nil
}

func(repo *CategoryRepositoryImpl) GetByID(
	id interface{},
) (*models.Category, error){
	return nil, nil
}

func(repo *CategoryRepositoryImpl) Create(item *models.AddCategory) error{
	return nil
}

func(repo *CategoryRepositoryImpl) Update(item *models.Category) error {
	return nil
}
	
func(repo *CategoryRepositoryImpl) Delete(id interface{}) error {
	return nil
}