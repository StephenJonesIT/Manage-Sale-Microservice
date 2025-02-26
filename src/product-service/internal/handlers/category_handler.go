package handlers

import (
	"net/http"
	"product-service/common"
	"product-service/internal/business"
     log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service business.CategoryService
}

func NewCategoryHandler(service business.CategoryService) *CategoryHandler{
	return &CategoryHandler{service: service}
}

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