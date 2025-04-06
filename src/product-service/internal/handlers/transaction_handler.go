// @File handlers.inventory_transaction_handler.go
// @Description Implements Inventory transaction API logic functions
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package handlers

import (
	"net/http"

	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/common"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/business"
	"github.com/StephenJonesIT/Manage-Sale-Microservice/src/product-service/internal/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type InventoryTransactionHandler struct {
	transactionService business.TransactionService
}

func NewInventoryTransactionHandler(
	transactionService business.TransactionService,
) *InventoryTransactionHandler {
	return &InventoryTransactionHandler{
		transactionService: transactionService,
	}
}

// GetAllTransactions godoc
// @Summary Retrieve a list of inventory 
// @Description Retrieve all inventory , with optional paging
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param paging query common.Paging false "Inventory paging"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.ErrorResponse
// @Router /api/transactions [get]
func (handler *InventoryTransactionHandler) GetAllTransactions(ctx *gin.Context) {
	log.Info("GetAllTransactions endpoint called")
	var paging common.Paging

	if err := ctx.ShouldBind(&paging); err != nil {
		log.Warn("Invalid paging parameters: ", err)
		ctx.JSON(http.StatusBadRequest, common.NewAppError(
			http.StatusBadRequest,
			err,
			"Invalid paging parameters",
			"Error binding paging parameters",
			"BAD_REQUEST",
		))
		return
	}

	result, err := handler.transactionService.GetTransactions(ctx, &paging)

	if err != nil {
		log.Error("Failed to retrieve transactions list: ", err)
		ctx.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error()))
		return
	}

	log.Info("Transactions list retrieved successfully")
	ctx.JSON(http.StatusOK, common.NewDetailResponse(
		http.StatusOK,
		"Successfully retrieved the transactions list",
		result,
		paging,
		nil,
	))
}

// CreateInventoryTransaction godoc
// @Summary Create a new inventory 
// @Description Create a new inventory  with the provided parameters
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param inventoryData body models.InventoryTransaction true "Inventory to create"
// @Success 201 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.ErrorResponse
// @Router /api/transaction [post]
func (handler *InventoryTransactionHandler) CreateTransaction(ctx *gin.Context) {
	log.Info("CreateInventoryTransaction endpoint called")

	var temp models.InventoryTransaction
	if err := ctx.ShouldBindJSON(&temp); err != nil {
		log.Warn("Invalid inventory transaction parameters: ", err)
		ctx.JSON(http.StatusBadRequest, common.NewAppError(http.StatusBadRequest, err, "Invalid inventory transaction parameters", "Error binding inventory transaction parameters", "BAD_REQUEST"))
		return
	}

	temp.Transaction_ID = common.GeneralProductCode("H D")

	// Validate transaction type
	if temp.Transaction_Type.String() != "in" && temp.Transaction_Type.String() != "out" {
		log.Warn("Invalid transaction type")
		ctx.JSON(http.StatusBadRequest, common.NewErrorResponse("Invalid transaction type"))
		return
	}

	switch temp.Transaction_Type.String() {
	case "in":
		if err := handler.transactionService.CreateGoodsReceipt(ctx, &temp); err != nil {
			log.Error("Failed to create goods receipt")
			ctx.JSON(http.StatusBadRequest, common.NewErrorResponse("Failed to create goods receipt"))
			return
		}
	case "out":
		if err := handler.transactionService.CreateGoodsIssue(ctx, &temp); err != nil {
			log.Error("Failed to create goods issue")
			ctx.JSON(http.StatusBadRequest, common.NewErrorResponse("Failed to create goods issue"))
			return
		}
	default:
		log.Warn("Invalid transaction type")
		ctx.JSON(http.StatusBadRequest, common.NewErrorResponse("Invalid transaction type"))
		return
	}

	log.Info("Create transaction successfully")
	ctx.JSON(http.StatusOK, common.NewResponse("Create transaction successfully"))
}

// DeleteTransaction godoc
// @Summary Deleted a inventory transaction
// @Description Delete a inventory transaction by id
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} common.Response
// @Success 500 {object} common.ErrorResponse
// @Router /api/transaction/{id} [delete]
func (handler *InventoryTransactionHandler) DeleteTransaction(ctx *gin.Context) {
	idParam := ctx.Param("id")

	if err := handler.transactionService.DeleteTransaction(ctx, idParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.NewDetailResponse(
		http.StatusOK,
		"Inventory transaction deleted successfully",
		true,
		nil,
		nil,
	))
}
