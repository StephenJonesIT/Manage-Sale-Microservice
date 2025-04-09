// @File business.inventory_transaction_service.go
// @Description Implements  transaction CRUD for Inventory Transaction_Date
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package business

import (
	"context"
	"fmt"
	"sync"

	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/repository"
)

type TransactionService interface {
	GetTransactions(ctx context.Context, paging *common.Paging) ([]models.Transaction, error)
	DeleteTransaction(ctx context.Context, transactionID string) error
	CreateGoodsReceipt(ctx context.Context, transaction *models.InventoryTransaction) error
	CreateGoodsIssue(ctx context.Context, transaction *models.InventoryTransaction) error
}

type TransactionServiceImpl struct {
	transactionRepo repository.TransactionRepository
	pwRepo          repository.ProductWarehouseRepo
}

func NewTransactionService(
	transactionRepo repository.TransactionRepository,
	pwRepo repository.ProductWarehouseRepo,
) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		transactionRepo: transactionRepo,
		pwRepo:          pwRepo,
	}
}

func (service *TransactionServiceImpl) GetTransactions(
	ctx context.Context,
	paging *common.Paging,
) ([]models.Transaction, error) {
	paging.Process()
	return service.transactionRepo.GetTransactions(ctx, paging)
}

// Phieu nhap
func (s *TransactionServiceImpl) CreateGoodsReceipt(
	ctx context.Context,
	transaction *models.InventoryTransaction,
) error {
	if transaction.Warehouse_ID == 0 {
		return fmt.Errorf("warehouse id is required")
	}

	if transaction.Quantity < 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}

	if transaction.Product_ID == "" {
		return fmt.Errorf("product id is required")
	}

	_, err := s.pwRepo.GetByID(transaction.Product_ID, transaction.Warehouse_ID)
	if err != nil {
		productWarehouse := models.ProductWarehouses{
			Product_ID:   transaction.Product_ID,
			WareHouse_ID: transaction.Warehouse_ID,
		}

		if errCreate := s.pwRepo.Create(ctx, &productWarehouse); errCreate != nil {
			return fmt.Errorf("faild to create product warehouse")
		}
	}

	tx := s.transactionRepo.(*repository.TransactionRepositoryImpl).DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Sử dụng WaitGroup để đợi các goroutine hoàn thành
	var wg sync.WaitGroup
	errChan := make(chan error, 2) // Buffer cho 2 goroutine
	var finalErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.transactionRepo.GoodsReceipt(ctx, tx, transaction); err != nil {
			errChan <- fmt.Errorf("GoodsReceipt error: %w", err)
		}
	}()

	// Goroutine 2: IncrementQuantity
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.pwRepo.IncrementQuantity(
			ctx,
			tx,
			transaction.Product_ID,
			transaction.Warehouse_ID,
			transaction.Quantity,
		); err != nil {
			errChan <- fmt.Errorf("IncrementQuantity error: %w", err)
		}
	}()

	// Goroutine để đợi và xử lý kết quả
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Kiểm tra lỗi từ các goroutine
	for err := range errChan {
		if err != nil {
			tx.Rollback()
			if finalErr == nil {
				finalErr = err
			} else {
				finalErr = fmt.Errorf("%v; %w", finalErr, err)
			}
		}
	}

	if finalErr != nil {
		return finalErr
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit error: %w", err)
	}

	return nil
}

func (s *TransactionServiceImpl) CreateGoodsIssue(ctx context.Context, transaction *models.InventoryTransaction) error {
	if transaction.Warehouse_ID == 0 {
		return fmt.Errorf("warehouse id is required")
	}

	if transaction.Quantity < 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}

	if transaction.Product_ID == "" {
		return fmt.Errorf("product id is required")
	}

	productWarehouse, err := s.pwRepo.GetByID(transaction.Product_ID, transaction.Warehouse_ID)
	if err != nil {
		return fmt.Errorf("warehouse or product invalid")
	}

	tempQuantity := productWarehouse.Quantity - transaction.Quantity
	if tempQuantity < 0 {
		return fmt.Errorf("transaction failed: quantity requested (%d) exceeds available stock (%d) in warehouse", transaction.Quantity, productWarehouse.Quantity)
	}

	tx := s.transactionRepo.(*repository.TransactionRepositoryImpl).DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	// Sử dụng WaitGroup để đợi các goroutine hoàn thành
	var wg sync.WaitGroup
	errChan := make(chan error, 2) // Buffer cho 2 goroutine
	var finalErr error

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.transactionRepo.GoodsReceipt(ctx, tx, transaction); err != nil {
			errChan <- fmt.Errorf("GoodsReceipt error: %w", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.pwRepo.DecrementQuantity(ctx, tx, transaction.Product_ID, transaction.Warehouse_ID, transaction.Quantity); err != nil {
			errChan <- fmt.Errorf("IncrementQuantity error: %w", err)
		}
	}()

	// Goroutine để đợi và xử lý kết quả
	go func() {
		wg.Wait()
		close(errChan)
	}()

	// Kiểm tra lỗi từ các goroutine
	for err := range errChan {
		if err != nil {
			tx.Rollback()
			if finalErr == nil {
				finalErr = err
			} else {
				finalErr = fmt.Errorf("%v; %w", finalErr, err)
			}
		}
	}

	if finalErr != nil {
		return finalErr
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("commit error: %w", err)
	}

	return nil
}

func (s *TransactionServiceImpl) DeleteTransaction(ctx context.Context, transactionID string) error {
	if transactionID == "" {
		return fmt.Errorf("transaction id is required")
	}

	return s.transactionRepo.Delete(ctx, transactionID)
}