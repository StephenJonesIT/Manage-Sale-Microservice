/*
 * @File: handlers.warehouse_handler.go
 * @Description: Implements Warehouse API logic functions
 * @Author: Tran Thanh Sang (tranthanhsang.it.la@gmail.com)
 */
package handlers

import (
	"net/http"
	"product-service/common"
	"product-service/internal/business"
	"product-service/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type WarehouseHandler struct {
	warehouseService business.WarehouseService
}

func NewWarehouseHandler(service business.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{warehouseService: service}
}

// GetAllWarehouses godoc
// @Summary Retrieve a list of warehouses
// @Description Retrieve all warehouses, with optional paging
// @Tags warehouses
// @Accept  json
// @Produce  json
// @Param paging query common.Paging false "Warehouse paging"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/warehouses [get]
func(handler *WarehouseHandler) GetAllWarehouses(ctx *gin.Context){
	log.Info("GetAllWarehouses endpoint called")
	var paging common.Paging

	if err := ctx.ShouldBind(&paging); err != nil{
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

	paging.Process()

	result, err := handler.warehouseService.GetAllWarehouses(&paging)

	if err != nil {
        log.Error("Failed to retrieve warehouses list: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Warehouses list retrieved successfully")
    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK, 
        "Successfully retrieved the warehouses list",
        result, 
        paging, 
        nil,
    ))
}

// CreateWarehouse godoc
// @Summary Create a new warehouse
// @Description Create a new warehouse with the provided parameters
// @Tags warehouses
// @Accept  json
// @Produce  json
// @Param warehouse body models.Warehouses true "Warehouse to create"
// @Success 201 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/warehouse [post]
func(handler *WarehouseHandler) CreateWarehouses(ctx *gin.Context){
    log.Info("CreateWarehouse endpoint called")

    var addWarehouses models.Warehouses

    if err := ctx.ShouldBindJSON(&addWarehouses); err != nil  {
        log.Warn("Invalid category parameters: ", err)
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid warehouses parameters",
            "Error binding warehouses parameters",
            "BAD_REQUEST",
        ))
        return
    }

    if err := handler.warehouseService.CreateWarehouse(&addWarehouses); err != nil {
        log.Error("Failed to create warehouses: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Warehouses created successfully")
    ctx.JSON(http.StatusCreated,common.NewResponse(
        http.StatusCreated,
        "Warehouses created successfully",
        addWarehouses,
        nil,
        nil,
    ))
}

// UpdateWarehouse godoc
// @Summary Update an existing warehouse
// @Description Update an existing warehouse with the provided parameters
// @Tags warehouses
// @Accept json
// @Produce json
// @Param warehouse body models.Warehouses true "Warehouse to update"
// @Param id path int true "Warehouses ID"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/warehouse/{id} [put]
func(handler *WarehouseHandler) UpdateWarehouses(ctx *gin.Context){
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid warehouse ID",
            "Error converting warehouse ID",
            "BAD_REQUEST",
        ))
        return
    }

    var updateWarehouse models.Warehouses

    if err := ctx.ShouldBindJSON(&updateWarehouse); err != nil{
        ctx.JSON(http.StatusBadRequest,common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid warehouse parameters",
            "Error binding warehouse parameters",
            "BAD_REQUEST",
        ))
        return 
    }

    existingWarehouse, err := handler.warehouseService.GetWarehouseByID(id)

    if err != nil {
        ctx.JSON(http.StatusNotFound, err)
        return
    }

    existingWarehouse.Warehouses_Name = updateWarehouse.Warehouses_Name
    existingWarehouse.Location        = updateWarehouse.Location

    if  err := handler.warehouseService.UpdateWarehouse(existingWarehouse); err != nil {
        ctx.JSON(http.StatusInternalServerError,err)
        return
    }

    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Warehouse update successfully",
        existingWarehouse,
        nil,
        nil,
    ))
} 

// DeleteWarehouse godoc
// @Summary Deleted a warehouse
// @Description Delete a warehouse by id
// @Tags warehouses
// @Accept json
// @Produce json
// @Param id path string true "Warehouse ID"
// @Success 200 {object} common.Response
// @Success 500 {object} common.AppError
// @Router /api/warehouse/{id} [delete]
func(handler *WarehouseHandler) DeleteWarehouse(ctx *gin.Context){
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)

    if err != nil {
        ctx.JSON(http.StatusBadRequest,common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid warehouse parameters",
            "Error binding warehouse parameters",
            "BAD_REQUEST",
        ))
        return 
    }

    if err := handler.warehouseService.DeleteWarehouse(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Warehouse deleted successfully",
        true,
        nil,
        nil,
    ))
}

// GetWarehouse godoc
// @Summary Get a warehouse
// @Description Get a warehouse by id
// @Tags warehouses
// @Accept json
// @Produce json
// @Param id path string true "Warehouse ID"
// @Success 200 {object} common.Response
// @Success 400 {object} common.AppError
// @Router /api/warehouse/{id} [get]
func(handler *WarehouseHandler) GetWarehouse(ctx *gin.Context){
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid warehouse parameters",
            "Error binding warehouse parameters",
            "BAD_REQUEST",
        ))
        return
    }

    category, err := handler.warehouseService.GetWarehouseByID(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Get a warehouse successfully",
        category,
        nil,
        nil,
    ))
}