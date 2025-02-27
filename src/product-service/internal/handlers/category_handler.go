package handlers

import (
	"net/http"
	"product-service/common"
	"product-service/internal/business"
	"product-service/internal/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CategoryHandler struct {
	service business.CategoryService
}

func NewCategoryHandler(service business.CategoryService) *CategoryHandler{
	return &CategoryHandler{service: service}
}

// GetAllCategories godoc
// @Summary Retrieve a list of categories
// @Description Retrieve all products, with optional filtering and paging
// @Tags categories
// @Accept  json
// @Produce  json
// @Param paging query common.Paging false "Category paging"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/categories [get]
func(handler *CategoryHandler) GetAllCategories(ctx *gin.Context){
	log.Info("GetAllCategories endpoint called")
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

	result, err := handler.service.GetAllCategories(&paging)

	if err != nil {
        log.Error("Failed to retrieve categories list: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Categories list retrieved successfully")
    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK, 
        "Successfully retrieved the categories list",
        result, 
        paging, 
        nil,
    ))
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the provided parameters
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body models.AddCategory true "Category to create"
// @Success 201 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/category [post]
func(handler *CategoryHandler) CreateCategory(ctx *gin.Context){
    log.Info("CreateCategory endpoint called")

    var addCategory models.AddCategory

    if err := ctx.ShouldBindJSON(&addCategory); err != nil  {
        log.Warn("Invalid category parameters: ", err)
        ctx.JSON(http.StatusBadRequest, common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid category parameters",
            "Error binding category parameters",
            "BAD_REQUEST",
        ))
        return
    }

    if err := handler.service.CreateCategory(&addCategory); err != nil {
        log.Error("Failed to create category: ", err)
        ctx.JSON(http.StatusInternalServerError, err)
        return
    }

    log.Info("Category created successfully")
    ctx.JSON(http.StatusCreated,common.NewResponse(
        http.StatusCreated,
        "Category created successfully",
        addCategory,
        nil,
        nil,
    ))
}

// UpdateCategory godoc
// @Summary Update an existing category
// @Description Update an existing category with the provided parameters
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category to update"
// @Param id path int true "Category ID"
// @Success 200 {object} common.Response
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /api/category/{id} [put]
func(handler *CategoryHandler) UpdateCategory(ctx *gin.Context){
    id := ctx.Param("id")
    
    var updateCategory models.Category

    if err := ctx.ShouldBindJSON(&updateCategory); err != nil{
        ctx.JSON(http.StatusBadRequest,common.NewAppError(
            http.StatusBadRequest,
            err,
            "Invalid category parameters",
            "Error binding category parameters",
            "BAD_REQUEST",
        ))
        return 
    }

    existingCategory, err := handler.service.GetCategoryByID(id)

    if err != nil {
        ctx.JSON(http.StatusNotFound, err)
        return
    }

    existingCategory.Category_Name = updateCategory.Category_Name
    existingCategory.Description = updateCategory.Description

    if  err := handler.service.UpdateCategory(existingCategory); err != nil {
        ctx.JSON(http.StatusInternalServerError,err)
        return
    }

    ctx.JSON(http.StatusOK, common.NewResponse(
        http.StatusOK,
        "Category update successfully",
        existingCategory,
        nil,
        nil,
    ))
} 