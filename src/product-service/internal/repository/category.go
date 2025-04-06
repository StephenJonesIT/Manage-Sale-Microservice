/*
* @File: repository.category_repository.go
* @Description: Implements category CRUD functions for MySQL
* @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
 */
package repository

import (
	"errors"
	"net/http"

	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll(paging *common.Paging) ([]models.Category, error)
	GetByID(id interface{}) (*models.Category, error)
	Create(category *models.AddCategory) error
	Update(category *models.Category) error
	Delete(id interface{}) error
}

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
	var result models.Category

	if err :=  repo.DB.First(&result, id).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(
				http.StatusNotFound,
				err,
				"Category not found",
				"CATEGORY_NOT_FOUND",
				"The category with the given ID doesn't exists",
			)
		}
		return nil, common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Error retrieve category",
			"Database error while retrieving category",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return &result, nil
}

func(repo *CategoryRepositoryImpl) Create(item *models.AddCategory) error{
	if err := repo.DB.Create(item).Error; err != nil {
		return common.NewAppError(
			http.StatusInternalServerError, 
			err,
			"Failed to create product",
            "Database error while creating product",
            "INTERNAL_SERVER_ERROR", 
		)
	}
	return nil
}

func(repo *CategoryRepositoryImpl) Update(item *models.Category) error {
	if err := repo.DB.Where("category_id = ?", item.Category_ID).Save(item).Error; err != nil{
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Fail to update category",
			"Database error while updating category",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return nil
}
	
func(repo *CategoryRepositoryImpl) Delete(id interface{}) error {
	if err := repo.DB.Delete(models.Category{}, id).Error; err != nil{
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Fail to delete category",
			"Database error while deleting category",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return nil
}