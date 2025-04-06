// @File repository.product_warehouse_repository.go
// @Description Implements CRUD functions for MySQL
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
package repository

import (
	"context"
	"errors"
	"net/http"

	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"gorm.io/gorm"
)

type ProductWarehouseRepo interface {
	GetAll(paging *common.Paging) ([]models.ProductWarehouses, error)
	GetByID(idProduct, idWarehouse interface{}) (*models.ProductWarehouses, error)
	Create(ctx context.Context, item *models.ProductWarehouses) error
	Update(ctx context.Context, item *models.ProductWarehouses) error
	Delete(idProduct, idWarehouse interface{}) error
	IncrementQuantity(ctx context.Context, tx *gorm.DB, productID string, warehouseID, quantity int) error
	DecrementQuantity(ctx context.Context, tx *gorm.DB, productID string, warehouseID, quantity int) error
}

type ProductWarehouseRepoImple struct {
	db *gorm.DB
}

func NewProductWarehouseRepo(db *gorm.DB) *ProductWarehouseRepoImple {
	return &ProductWarehouseRepoImple{db: db}
}

func (repo *ProductWarehouseRepoImple) GetAll(paging *common.Paging) ([]models.ProductWarehouses, error) {
	var result []models.ProductWarehouses

	sql := repo.db

	if err := sql.Table(models.ProductWarehouses{}.TableName()).Count(&paging.Total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(
				http.StatusNotFound,
				err,
				"Table not found",
				"TABLE_NOT_FOUND",
				"The table Product Warehouse does not exist.",
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

	if err := sql.Order("last_updated desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Error retrieving product warehouses",
			"RETRIEVE_ERROR",
			"An unexpected error occurred while retrieving the product warehouses.",
		)
	}
	return result, nil
}

func (repo *ProductWarehouseRepoImple) GetByID(idProduct, idWarehouse interface{}) (*models.ProductWarehouses, error) {
	var result models.ProductWarehouses
	if err := repo.db.Where("product_id = ? AND warehouse_id = ?", idProduct, idWarehouse).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(
				http.StatusNotFound,
				err,
				"Product not found",
				"PRODUCT_NOT_FOUND",
				"The product warehouse with the given ID does not exist.",
			)
		}
		return nil, common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Error retrieving product warehouse",
			"Database error while retrieving product warehouse",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return &result, nil
}

func (repo *ProductWarehouseRepoImple) Create(ctx context.Context, item *models.ProductWarehouses) error {
	if err := repo.db.Create(item).Error; err != nil {
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Failed to create product warehouse",
			"Database error while creating product warehouse",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return nil
}

func (repo *ProductWarehouseRepoImple) Update(ctx context.Context, item *models.ProductWarehouses) error {
	if err := repo.db.Where("product_id = ? AND warehouse_id = ?", item.Product_ID, item.WareHouse_ID).Save(item).Error; err != nil {
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Failed to update product warehouse",
			"Database error while updating product warehouse",
			"INTERNAL_SERVER_ERROR",
		)
	}
	return nil
}

func (repo *ProductWarehouseRepoImple) Delete(idProduct, idWarehouse interface{}) error {
	// Xóa sản phẩm
	if err := repo.db.
		Table(models.ProductWarehouses{}.TableName()).
		Where("product_id = ? AND warehouse_id = ?", idProduct, idWarehouse).
		Delete(models.ProductWarehouses{}).Error; err != nil {
		return common.NewAppError(
			http.StatusInternalServerError,
			err,
			"Error deleting product warehouse",
			"Database error while deleting product warehouse",
			"INTERNAL_SERVER_ERROR",
		)
	}

	return nil
}

func (r *ProductWarehouseRepoImple) IncrementQuantity(ctx context.Context, tx *gorm.DB, productID string, warehouseID, quantity int) error {
	return tx.WithContext(ctx).
		Table(models.ProductWarehouses{}.TableName()).
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).
		UpdateColumn("quantity", gorm.Expr("quantity + ?", quantity)).
		Error
}

func (r *ProductWarehouseRepoImple) DecrementQuantity(ctx context.Context, tx *gorm.DB, productID string, warehouseID, quantity int) error {
	return tx.WithContext(ctx).
		Table(models.ProductWarehouses{}.TableName()).
		Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).
		UpdateColumn("quantity", gorm.Expr("quantity - ?", quantity)).
		Error
}
