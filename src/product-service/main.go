package main

import (
	"fmt"
	"io"
	"os"
	"product-service/common"
	_ "product-service/docs"
	"product-service/internal/business"
	"product-service/internal/handlers"
	"product-service/internal/repository"
	"product-service/pkg/database"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	log.Info("Loading configuration...")
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	log.Info("Connecting to database...")
	err = database.ConnectDatabase()
	if err != nil {
		return err
	}

	log.Info("Setting up logging...")
	if common.Config.EnableGinFileLog {
		f, err := os.Create("logs/gin.log")
		if err != nil {
			log.Error("Unable to create log file:", err)
			return err
		}
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
	m.router.Use(gin.Logger(), gin.Recovery())

	log.Info("Server initialization complete")
	return nil
}

func(m *Main) InitializeProductComponents(database *gorm.DB) (*handlers.ProductHandler, error) {
    log.Info("Creating repositories and services...")

    productRepo := repository.NewProductRepository(database)
    if productRepo == nil {
        log.Fatal("Failed to create product repository")
        return nil, fmt.Errorf("failed to create product repository")
    }

    productService := business.NewProductService(productRepo)
    if productService == nil {
        log.Fatal("Failed to create product service")
        return nil, fmt.Errorf("failed to create product service")
    }

    productHandler := handlers.NewProductHandler(productService)
    if productHandler == nil {
        log.Fatal("Failed to create product handler")
        return nil, fmt.Errorf("failed to create product handler")
    }

    return productHandler, nil
}


func main() {
	log.Info("Starting application...")
	m := Main{}
	err := m.initServer()
	if err != nil {
		log.Fatal("Failed to initialize server: ", err)
	}

	if database.DB == nil {
		log.Fatal("Database connection is nil")
	}

	productHandler,_ := m.InitializeProductComponents(database.DB)

	categoryRepo := repository.NewCategoryRepository(database.DB)
	categoryService := business.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	log.Info("Setting up API routes...")

	m.router.GET("/swagger/*any", func(c *gin.Context) {
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})
	
	api := m.router.Group("/api")
	{
		api.GET("/products", productHandler.GetAllProducts)
		api.GET("/product/:id", productHandler.GetProduct)
		api.POST("/product", productHandler.CreateProduct)
		api.PUT("/product", productHandler.UpdateProduct)
		api.DELETE("/product/:id", productHandler.DeleteProduct)
		api.GET("/categories", categoryHandler.GetAllCategories)
	}

	err = m.router.Run(common.Config.Port)
	if err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
