package handlers

import (
    "net/http"
    "product-service/common"
    "product-service/internal/business"
    "product-service/internal/models"

    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
)

type ProductHandler struct {
    ProductService business.ProductServiceInterface
}

func NewProductHandler(service business.ProductServiceInterface) *ProductHandler{
    return &ProductHandler{ProductService: service}
}

// GetAllProducts godoc
// @Summary Retrieve a list of products
// @Description Retrieve all products, with optional filtering and paging
// @Tags products
// @Accept  json
// @Produce  json
// @Param filter query common.Filter false "Product filter"
// @Param paging query common.Paging false "Product paging"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/products [get]
func (handler *ProductHandler) GetAllProducts(ctx *gin.Context) {
    log.Info("GetAllProducts endpoint called")

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

    var filter common.Filter
    if err := ctx.ShouldBind(&filter); err != nil {
        log.Warn("Invalid filter parameters: ", err)
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid filter parameters",
            "Error binding filter parameters",
            "BAD_REQUEST",
        ))
        return
    }

    result, err := handler.ProductService.GetAllProducts(&filter, &paging)
    if err != nil {
        log.Error("Failed to retrieve product list: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Product list retrieved successfully")
    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK, 
        "Successfully retrieved the product list",
        result, 
        paging, 
        filter,
    ))
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided parameters
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Product to create"
// @Success 201 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/product [post]
func(handler *ProductHandler) CreateProduct(ctx *gin.Context){
    log.Info("CreateProduct endpoint called")

    var product models.Product

    if err := ctx.ShouldBindJSON(&product); err != nil  {
        log.Warn("Invalid product parameters: ", err)
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid product parameters",
            "Error binding product parameters",
            "BAD_REQUEST",
        ))
        return
    }

    if err := handler.ProductService.CreateProduct(&product); err != nil {
        log.Error("Failed to create product: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Product created successfully")
    ctx.JSON(http.StatusCreated,common.NewResponse(
        http.StatusCreated,
        "Product created successfully",
        product,
        nil,
        nil,
    ))
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Update an existing product with the provided parameters
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product to update"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/product [put]
func (handler *ProductHandler) UpdateProduct(ctx *gin.Context){
    log.Info("UpdateProduct endpoint called")

    var product models.Product

    if err := ctx.ShouldBindJSON(&product); err != nil {
        log.Warn("Invalid product parameters: ", err)
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid product parameters",
            "Error binding product parameters",
            "BAD_REQUEST",
        ))
        return
    }

    if err := handler.ProductService.UpdateProduct(&product); err != nil {
        log.Error("Failed to update product: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Product updated successfully")
    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Product update successfully",
        product,
        nil,
        nil,
    ))
}

// DeleteProduct godoc
// @Summary Deleted a product
// @Description Delete a product by id
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} common.Response
// @Success 500 {object} common.AppError
// @Router /api/product/{id} [delete]
func (handler *ProductHandler) DeleteProduct(ctx *gin.Context){
    log.Info("DeleteProduct endpoint called")

    idProduct := ctx.Param("id")

    err := handler.ProductService.DeleteProduct(idProduct)

    if err != nil {
        log.Error("Failed to delete product: ", err)
        ctx.JSON(http.StatusInternalServerError,err)
        return
    }

    log.Info("Product deleted successfully")
    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Product successfully deleted",
        nil,
        nil,
        nil,
    ))
}

// GetProduct godoc
// @Summary Get a product
// @Description Get a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
//  @Success 200 {object} common.Response
// @Failure 404 {object} common.AppError
// @Router /product/{id} [get]
func (handler *ProductHandler) GetProduct(ctx *gin.Context) {
    log.Info("GetProduct endpoint called")

    // Lấy product ID từ request parameters
    idProduct := ctx.Param("id")

    // Thử lấy thông tin sản phẩm từ database
    product, err := handler.ProductService.GetProductByID(idProduct)

    if err != nil {
        log.Warn("Product not found: ", err)
		ctx.JSON(http.StatusNotFound, common.NewAppError(
			http.StatusNotFound,
			err,
			"Product Not Found",
			"Product Not Found",
			"NOT_FOUND",
		))
		return
	}

	log.Info("Get a product successfully")
	ctx.JSON(http.StatusOK, common.NewResponse(
		http.StatusOK,
		"Get a product successfully",
		product,
		nil,
		nil,
	))
}