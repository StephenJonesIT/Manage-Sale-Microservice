package main

import (
	"io"
	"os"
	"product-service/common"
	"product-service/internal/business"
	"product-service/internal/handlers"
	"product-service/internal/repository"
	"product-service/pkg/database"

	"github.com/gin-gonic/gin"
)

type Main struct{
	router *gin.Engine
}

func(m *Main) initServer() error{
	var err error
	err = common.LoadConfig()
    if err != nil {
		return err
    }

	err = database.ConnectDatabase()
	if err != nil {
		return err
	}

	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		}else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	}else{
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()
	m.router.Use(gin.Logger(), gin.Recovery())

	return nil

}
func main(){

	m := Main{}
	m.initServer()
	
	productRepo := repository.NewProductRepository(database.DB)
	productService := business.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)


	api := m.router.Group("/api")
	{
		api.GET("/products", productHandler.GetAllProducts)
	}
	
	m.router.Run(common.Config.Port)
}