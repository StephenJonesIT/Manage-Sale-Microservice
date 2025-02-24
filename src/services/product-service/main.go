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
	log "github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

type Main struct {
    router *gin.Engine
}

func (m *Main) initServer() error {
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
        } else {
            gin.DefaultWriter = io.MultiWriter(f)
        }
    } else {
        if !common.Config.EnableGinConsoleLog {
            gin.DefaultWriter = io.MultiWriter()
        }
    }

    m.router = gin.Default()
	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    m.router.Use(gin.Logger(), gin.Recovery())

    return nil
}

func main() {

    m := Main{}
    err := m.initServer()
    if err != nil {
        log.Fatal("Failed to initialize server: ", err)
    }

    if database.DB == nil {
        log.Fatal("Database connection is nil")
    }

    productRepo := repository.NewProductRepository(database.DB)
    if productRepo == nil {
        log.Fatal("Failed to create product repository")
    }

    productService := business.NewProductService(productRepo)
    if productService == nil {
        log.Fatal("Failed to create product service")
    }

    productHandler := handlers.NewProductHandler(productService)
    if productHandler == nil {
        log.Fatal("Failed to create product handler")
    }

    api := m.router.Group("/api")
    {
        api.GET("/products", productHandler.GetAllProducts)
        api.GET("/product/:id", productHandler.GetProduct)
        api.POST("/product", productHandler.CreateProduct)
        api.PUT("/product", productHandler.UpdateProduct)
        api.DELETE("/product/:id", productHandler.DeleteProduct)
    }

    err = m.router.Run(common.Config.Port)
    if err != nil {
        log.Fatal("Failed to run server: ", err)
    }
}
