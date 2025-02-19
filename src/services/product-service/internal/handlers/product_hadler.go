package handlers

import (
	"net/http"
	"product-service/common"
	"product-service/internal/business"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductService business.ProductServiceInterface
}

func NewProductHandler(service business.ProductServiceInterface) *ProductHandler{
	return &ProductHandler{ProductService: service}
}

func (handler *ProductHandler) GetAllProducts(ctx *gin.Context) {
	var paging common.Paging
	if err := ctx.ShouldBind(&paging); err != nil {
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
		ctx.JSON(http.StatusBadRequest, common.NewAppError(
			http.StatusBadRequest,
			err,
			"Error getting products",
			"Error getting products",
			"BAD_REQUEST",
		))
		return
	}

	ctx.JSON(http.StatusOK, common.NewResponse(
		http.StatusOK, 
		"Successfully retrieved the product list",
		result, 
		paging, 
		filter,
	))
}
