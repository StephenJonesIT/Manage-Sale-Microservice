// @File handlers.product_warehouse_handler.go
// @Description Implements Product Warehouse API logic functions
// @Author Tran Thanh Sang (tranthanhsang.it.la@gmail.com)

package handlers

import (
	"net/http"
	"product-service/common"
	"product-service/internal/business"
	"product-service/internal/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type ProductWarehouseHandler struct {
	service business.ProductWarehouseService
}

func NewProductWarehouseHanlder(service business.ProductWarehouseService) *ProductWarehouseHandler {
	return &ProductWarehouseHandler{service: service}
}

// GetAllProductWarehouse godoc
// @Summary Retrieve a list of product warehouses
// @Description Retrieve all product warehouses, with optional paging
// @Tags product warehouse
// @Accept  json
// @Produce  json
// @Param paging query common.Paging false "Product paging"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/product/warehouses [get]
func (handler *ProductWarehouseHandler) GetAllProductWarehouses(ctx *gin.Context) {
    log.Info("GetAllProductWarehouses endpoint called")

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

    paging.Process()

    result, err := handler.service.GetAllProductWarehouses(&paging)
    if err != nil {
        log.Error("Failed to retrieve product warehouses list: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Product warehouses list retrieved successfully")
    ctx.JSON(http.StatusOK, common.NewDetailResponse(
        http.StatusOK, 
        "Successfully retrieved the product warehouses list",
        result, 
        paging, 
		nil,
    ))
}

// CreateProductWarehouse godoc
// @Summary Create a new product warehouse
// @Description Create a new product warehouse with the provided parameters
// @Tags product warehouse
// @Accept  json
// @Produce  json
// @Param product body models.ProductWarehouses true "Product warehouse to create"
// @Success 201 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/product/warehouse [post]
func(handler *ProductWarehouseHandler) CreateProduct(ctx *gin.Context){
    log.Info("CreateProductWarehouse endpoint called")

    var temp models.ProductWarehouses

    if err := ctx.ShouldBindJSON(&temp); err != nil  {
        log.Warn("Invalid product warehouse parameters: ", err)
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid product warehouse parameters",
            "Error binding product warehouse parameters",
            "BAD_REQUEST",
        ))
        return
    }

    if err := handler.service.CreateProductWarehouse(ctx,&temp); err != nil {
        log.Error("Failed to create product warehouse: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Product warehouse created successfully")
    ctx.JSON(http.StatusCreated,common.NewDetailResponse(
        http.StatusCreated,
        "Product warehouse created successfully",
        temp,
        nil,
        nil,
    ))
}

// UpdateProductWarehouse godoc
// @Summary Update an existing product warehouse
// @Description Update an existing product warehouse with the provided parameters
// @Tags product warehouse
// @Accept json
// @Produce json
// @Param product body models.ProductWarehouses true "Product warehouse to update"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/product/warehouse [put]
func (handler *ProductWarehouseHandler) UpdateProductWarehouse(ctx *gin.Context){
    log.Info("UpdateProductWarehouse endpoint called")

    var productWarehouse models.ProductWarehouses

    if err := ctx.ShouldBindJSON(&productWarehouse); err != nil {
        log.Warn("Invalid product warehouse parameters: ", err)
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid product warehouse parameters",
            "Error binding product warehouse parameters",
            "BAD_REQUEST",
        ))
        return
    }

	productWarehouseExisting, err := handler.service.GetProductWarehouse(productWarehouse.Product_ID, productWarehouse.WareHouse_ID)

    if err != nil {
        log.Warn("Product warehouse not found: ", err)
		ctx.JSON(http.StatusNotFound, common.NewAppError(
			http.StatusNotFound,
			err,
			"Product warehouse Not Found",
			"Product warehouse Not Found",
			"NOT_FOUND",
		))
		return
	}
	
	productWarehouseExisting.Quantity = productWarehouse.Quantity

    if err := handler.service.UpdateProductWarehouse(ctx,productWarehouseExisting); err != nil {
        log.Error("Failed to update product: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Product warehouse updated successfully")
    ctx.JSON(http.StatusOK, common.NewDetailResponse(
        http.StatusOK,
        "Product warehouse update successfully",
        productWarehouseExisting,
        nil,
        nil,
    ))
}

// DeleteProductWarehouse godoc
// @Summary Deleted a product warehouse
// @Description Delete a product warehouse by id
// @Tags product warehouse
// @Accept json
// @Produce json
// @Param idProduct path string true "Product ID"
// @Param idWarehouse path string true "Warehouse ID"
// @Success 200 {object} common.Response
// @Success 500 {object} common.AppError
// @Router /api/product/warehouse/{idProduct}/{idWarehouse} [delete]
func (handler *ProductWarehouseHandler) DeleteProductWarehouse(ctx *gin.Context){
    log.Info("DeleteProductWarehouse endpoint called")

    idProduct := ctx.Param("idProduct")
	idWarehouse := ctx.Param("idWarehouse")

	productWarehouseExisting, err := handler.service.GetProductWarehouse(idProduct, idWarehouse)

    if err != nil {
        log.Warn("Product warehouse not found: ", err)
		ctx.JSON(http.StatusNotFound, common.NewAppError(
			http.StatusNotFound,
			err,
			"Product warehouse Not Found",
			"Product warehouse Not Found",
			"NOT_FOUND",
		))
		return
	}

    errDelete := handler.service.DeleteProductWarehouse(productWarehouseExisting.Product_ID,productWarehouseExisting.WareHouse_ID)

    if errDelete != nil {
        log.Error("Failed to delete product warehouse: ", errDelete)
        ctx.JSON(http.StatusInternalServerError,err)
        return
    }

    log.Info("Product deleted warehouse successfully")
    ctx.JSON(http.StatusOK, common.NewDetailResponse(
        http.StatusOK,
        "Product warehouse successfully deleted",
        nil,
        nil,
        nil,
    ))
}

// GetProductWarehouse godoc
// @Summary Get a product warehouse
// @Description Get a product warehouse
// @Tags product warehouse
// @Accept json
// @Produce json
// @Param idProduct path string true "Product ID"
// @Param idWarehouse path string true "Warehouse ID"
// @Success 200 {object} common.Response
// @Failure 404 {object} common.AppError
// @Router /api/product/warehouse/{idProduct}/{idWarehouse} [get]
func (handler *ProductWarehouseHandler) GetProductWarehouse(ctx *gin.Context) {
    log.Info("GetProductWarehouse endpoint called")

    // Lấy product ID từ request parameters
    idProduct := ctx.Param("idProduct")
	idWarehouse := ctx.Param("idWarehouse")
    // Thử lấy thông tin sản phẩm từ database
    product, err := handler.service.GetProductWarehouse(idProduct, idWarehouse)

    if err != nil {
        log.Warn("Product warehouse not found: ", err)
		ctx.JSON(http.StatusNotFound, common.NewAppError(
			http.StatusNotFound,
			err,
			"Product warehouse Not Found",
			"Product warehouse  Not Found",
			"NOT_FOUND",
		))
		return
	}

	log.Info("Get a product warehouse successfully")
	ctx.JSON(http.StatusOK, common.NewDetailResponse(
		http.StatusOK,
		"Get a product warehouse successfully",
		product,
		nil,
		nil,
	))
}
