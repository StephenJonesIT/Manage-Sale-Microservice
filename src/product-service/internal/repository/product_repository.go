/*
* @File: business.product_service.go
* @Description: Implements Product CRUD functions for MySQL
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

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

 func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl{
	return &ProductRepositoryImpl{DB: db}
 }

// GetAll retrieves all products from the database with optional filtering and paging
func (repo *ProductRepositoryImpl) GetAll(
    filter *common.Filter, 
    paging *common.Paging,
) ([]models.Product, error) {
    var result []models.Product
    sql := repo.DB

    // Áp dụng bộ lọc nếu có
    if filter != nil {
        if v := filter.Status; v != "" {
            sql = sql.Where("status = ?", v)
        }
    }

    // Đếm tổng số bản ghi
    if err := sql.Table(models.Product{}.TableName()).Count(&paging.Total).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, common.NewAppError(
                http.StatusNotFound,
                err,
                "Table not found",
                "TABLE_NOT_FOUND",
                "The table Product does not exist.",
            )
        }
        if errors.Is(err, gorm.ErrInvalidData) {
            return nil, common.NewAppError(
                http.StatusBadRequest,
                err,
                "Invalid data",
                "INVALID_DATA",
                "The data in the table is invalid.",
            )
        }
        return nil, common.NewAppError(
            http.StatusInternalServerError,
            err,
            "Failed to count records",
            "COUNT_ERROR",
            "An unexpected error occurred while counting the records.",
        )
    }

    // Truy vấn danh sách sản phẩm với phân trang
    if err := sql.Order("created_at desc").
        Offset((paging.Page - 1) * paging.Limit).
        Limit(paging.Limit).
        Find(&result).Error; err != nil {
        return nil, common.NewAppError(
            http.StatusInternalServerError,
            err,
            "Error retrieving products",
            "RETRIEVE_ERROR",
            "An unexpected error occurred while retrieving the products.",
        )
    }

    return result, nil
}


// GetByID retrieves a product by its ID from the database
func (repo *ProductRepositoryImpl) GetByID(productID string) (*models.Product, error) {
    var product models.Product
    if err := repo.DB.Where("product_id = ?", productID).First(&product).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, common.NewAppError(
                http.StatusNotFound,
                err,
                "Product not found",
                "PRODUCT_NOT_FOUND",
                "The product with the given ID does not exist.",
            )
        }
        return nil, common.NewAppError(
            http.StatusInternalServerError,
            err,
            "Error retrieving product",
            "Database error while retrieving product",
            "INTERNAL_SERVER_ERROR",
        )
    }
    return &product, nil
}


func(repo *ProductRepositoryImpl) Create(product *models.Product) error{
	if err := repo.DB.Create(product).Error; err != nil{
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

// Update updates an existing product in the database
func (repo *ProductRepositoryImpl) Update(product *models.Product) error {
    if err := repo.DB.Save(product).Error; err != nil {
        return common.NewAppError(
            http.StatusInternalServerError,
            err,
            "Failed to update product",
            "Database error while updating product",
            "INTERNAL_SERVER_ERROR",
        )
    }
    return nil
}


func (repo *ProductRepositoryImpl) Delete(productID string) error {
    var product models.Product

    // Tìm sản phẩm theo ID
    if err := repo.DB.Where("product_id = ?", productID).First(&product).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return common.NewAppError(
                http.StatusNotFound,
                err,
                "Product not found",
                "PRODUCT_NOT_FOUND",
                "The product with the given ID does not exist.",
            )
        }
        return common.NewAppError(
            http.StatusInternalServerError,
            err,
            "Error finding product",
            "Error finding product in the database",
            "INTERNAL_SERVER_ERROR",
        )
    }

    // Xóa sản phẩm
    if err := repo.DB.Delete(&product).Error; err != nil {
        return common.NewAppError(
            http.StatusInternalServerError,
            err,
            "Error deleting product",
            "Database error while deleting product",
            "INTERNAL_SERVER_ERROR",
        )
    }

    return nil
}
