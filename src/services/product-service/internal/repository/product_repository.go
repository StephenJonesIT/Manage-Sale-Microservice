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

func(repo *ProductRepositoryImpl) GetAll(
	filter *common.Filter, 
	paging *common.Paging,
) ([]models.Product, error){
	var result []models.Product
	sql := repo.DB

	if filter !=nil {
		if v:=filter.Status; v != ""{
			sql = sql.Where("Status = ?", v)
		}
	}

	if err := sql.Table(models.Product{}.TableName()).Count(&paging.Total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewAppError(http.StatusNotFound, err, "Table not found", "TABLE_NOT_FOUND", "The table Product does not exist.")
		}
		if errors.Is(err, gorm.ErrInvalidData) {
			return nil, common.NewAppError(http.StatusBadRequest, err, "Invalid data", "INVALID_DATA", "The data in the table is invalid.")
		}
		return nil, common.NewAppError(http.StatusInternalServerError, err, "Failed to count records", "COUNT_ERROR", "An unexpected error occurred while counting the records.")
	}


	if err := sql.Order("created_at desc").
		Offset((paging.Page-1)*paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil{
			return nil, err
	}

	return result, nil
}

func(repo *ProductRepositoryImpl) GetByID(productID string)(*models.Product, error){
	var product models.Product
	if err := repo.DB.Where("product_id = ?", productID).First(&product).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, common.NewAppError(http.StatusNotFound, err, "Product not found", "PRODUCT_NOT_FOUND", "The product with the given ID does not exist.")
		}
		return nil, err
	}
	return &product, nil
}

func(repo *ProductRepositoryImpl) Create(product *models.Product) error{
	if err := repo.DB.Create(product).Error; err != nil{
		return err
	}
	return nil
}

func(repo *ProductRepositoryImpl) Update(product *models.Product) error{
	if err := repo.DB.Save(product).Error; err != nil{
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) Delete(productID string) error {
	
	var product models.Product
	if err := repo.DB.Where("product_id = ?", productID).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return common.NewAppError(http.StatusNotFound, err, "Product not found", "PRODUCT_NOT_FOUND", "The product with the given ID does not exist.")
		}
		return err
	}

	// Xóa sản phẩm
	if err := repo.DB.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}