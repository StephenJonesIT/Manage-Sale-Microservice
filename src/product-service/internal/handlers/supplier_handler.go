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

type SupplierHandler struct {
	supplierHandler business.SupplierService
}

func NewSupplierHandler(service business.SupplierService) *SupplierHandler{
	return &SupplierHandler{supplierHandler: service}
}

// GetAllSuppliers godoc
// @Summary Retrieve a list of Suppliers
// @Description Retrieve all products, with optional paging
// @Tags suppliers
// @Accept  json
// @Produce  json
// @Param paging query common.Paging false "Supplier paging"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/suppliers [get]
func(handler *SupplierHandler) GetAllSuppliers(ctx *gin.Context){
	log.Info("GetAllSuppliers endpoint called")
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

	result, err := handler.supplierHandler.GetAllSuppliers(&paging)

	if err != nil {
        log.Error("Failed to retrieve suppliers list: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("suppliers list retrieved successfully")
    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK, 
        "Successfully retrieved the suppliers list",
        result, 
        paging, 
        nil,
    ))
}

// CreateSupplier godoc
// @Summary Create a new Supplier
// @Description Create a new supplier with the provided parameters
// @Tags suppliers
// @Accept  json
// @Produce  json
// @Param Supplier body models.AddSupplier true "Supplier to create"
// @Success 201 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/supplier [post]
func(handler *SupplierHandler) CreateSupplier(ctx *gin.Context){
    log.Info("CreateSupplier endpoint called")

    var addSupplier models.AddSupplier

    if err := ctx.ShouldBindJSON(&addSupplier); err != nil  {
        log.Warn("Invalid Supplier parameters: ", err)
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid supplier parameters",
            "Error binding supplier parameters",
            "BAD_REQUEST",
        ))
        return
    }

    if err := handler.supplierHandler.CreateSupplier(&addSupplier); err != nil {
        log.Error("Failed to create Supplier: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Supplier created successfully")
    ctx.JSON(http.StatusCreated,common.NewResponse(
        http.StatusCreated,
        "Supplier created successfully",
        addSupplier,
        nil,
        nil,
    ))
}

// UpdateSupplier godoc
// @Summary Update an existing Supplier
// @Description Update an existing Supplier with the provided parameters
// @Tags suppliers
// @Accept json
// @Produce json
// @Param Supplier body models.UpdateSupplier true "Supplier to update"
// @Param id path int true "Supplier ID"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/supplier/{id} [put]
func(handler *SupplierHandler) UpdateSupplier(ctx *gin.Context){
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid Supplier ID",
            "Error converting Supplier ID",
            "BAD_REQUEST",
        ))
        return
    }

    var updateSupplier models.UpdateSupplier

    if err := ctx.ShouldBindJSON(&updateSupplier); err != nil{
        ctx.JSON(http.StatusBadRequest,common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid Supplier parameters",
            "Error binding Supplier parameters",
            "BAD_REQUEST",
        ))
        return 
    }

    existingSupplier, err := handler.supplierHandler.GetSupplierByID(id)
	existingSupplier.Supplier_Name 	= updateSupplier.Supplier_Name
	existingSupplier.Phone 			= updateSupplier.Phone
	existingSupplier.Email			= updateSupplier.Email
	existingSupplier.Address		= updateSupplier.Address
	existingSupplier.City 			= updateSupplier.City
	existingSupplier.Country		= updateSupplier.Country

    if err != nil {
        ctx.JSON(http.StatusNotFound, err)
        return
    }

    if  err := handler.supplierHandler.UpdateSupplier(existingSupplier); err != nil {
        ctx.JSON(http.StatusInternalServerError,err)
        return
    }

    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Supplier update successfully",
        existingSupplier,
        nil,
        nil,
    ))
} 

// DeleteSupplier godoc
// @Summary Deleted a Supplier
// @Description Delete a Supplier by id
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Success 200 {object} common.Response
// @Success 500 {object} common.AppError
// @Router /api/supplier/{id} [delete]
func(handler *SupplierHandler) DeleteSupplier(ctx *gin.Context){
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)

    if err != nil {
        ctx.JSON(http.StatusBadRequest,common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid Supplier parameters",
            "Error binding Supplier parameters",
            "BAD_REQUEST",
        ))
        return 
    }
		supplier, err := handler.supplierHandler.GetSupplierByID(id)
    	if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
        return
    	}

    if err := handler.supplierHandler.DeleteSupplier(supplier.Supplier_ID); err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Supplier deleted successfully",
        true,
        nil,
        nil,
    ))
}

// GetSupplier godoc
// @Summary Get a Supplier
// @Description Get a Supplier by id
// @Tags suppliers
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Success 200 {object} common.Response
// @Success 400 {object} common.AppError
// @Router /api/supplier/{id} [get]
func(handler *SupplierHandler) GetSupplier(ctx *gin.Context){
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid Supplier parameters",
            "Error binding Supplier parameters",
            "BAD_REQUEST",
        ))
        return
    }

    Supplier, err := handler.supplierHandler.GetSupplierByID(id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Get a Supplier successfully",
        Supplier,
        nil,
        nil,
    ))
}