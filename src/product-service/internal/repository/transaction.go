// @File repository.inventory_transaction_repository.go
// @Description Implements inventory transaction CRUD functions for MySQL
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package repository

import (
	"context"

	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactions(ctx context.Context, paging *common.Paging) ([]models.Transaction, error)
	GetTransactionByID(ctx context.Context, transactionID string) (*models.InventoryTransaction, error)
	Update(ctx context.Context, tx *gorm.DB, item *models.InventoryTransaction) error
	Delete(ctx context.Context, transactionID string) error
	GoodsReceipt(ctx context.Context, tx *gorm.DB, transaction *models.InventoryTransaction) error
	GoodsIssue(ctx context.Context, tx *gorm.DB, transaction *models.InventoryTransaction) error
}

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{DB: db}
}

func (repo *TransactionRepositoryImpl) GetTransactions(ctx context.Context, paging *common.Paging) ([]models.Transaction, error) {
	var result []models.Transaction

	if err := repo.DB.Table(models.InventoryTransaction{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := repo.DB.WithContext(ctx).
		Table(models.InventoryTransaction{}.TableName()).
		Select("inventory_transactions.*, p.product_name as product_name, w.warehouse_name as warehouses_name").
		Joins("JOIN products as p ON p.product_id = inventory_transactions.product_id").
    	Joins("JOIN warehouses as w ON w.warehouse_id = inventory_transactions.warehouse_id").
		Order("transaction_date desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *TransactionRepositoryImpl) GetTransactionByID(ctx context.Context, transactionID string) (*models.InventoryTransaction, error) {
	var transaction models.InventoryTransaction
	if err := repo.DB.
		WithContext(ctx).
		Table(models.InventoryTransaction{}.TableName()).
		Where("transaction_id = ?", transactionID).
		First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (repo *TransactionRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, item *models.InventoryTransaction) error {
	return tx.WithContext(ctx).
		Table(models.InventoryTransaction{}.TableName()).
		Where("transaction_id = ?", item.Transaction_ID).
		Updates(&item).Error
}

func (repo *TransactionRepositoryImpl) Delete(ctx context.Context, transactionID string) error {
	return repo.DB.WithContext(ctx).
		Table(models.InventoryTransaction{}.TableName()).
		Where("transaction_id = ?", transactionID).
		Delete(models.InventoryTransaction{}).Error
}

func (r *TransactionRepositoryImpl) GoodsReceipt(ctx context.Context, tx *gorm.DB, transaction *models.InventoryTransaction) error {
	return tx.WithContext(ctx).
		Table(models.InventoryTransaction{}.TableName()).
		Create(&transaction).Error
}

func (r *TransactionRepositoryImpl) GoodsIssue(ctx context.Context, tx *gorm.DB, transaction *models.InventoryTransaction) error {
	return tx.WithContext(ctx).
		Table(models.InventoryTransaction{}.TableName()).
		Create(transaction).Error
}
