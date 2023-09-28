package route

import (
	"github.com/gin-gonic/gin"
	"github.com/iyunrozikin26/bank_sampah.git/src/domains/controllers"
	"github.com/iyunrozikin26/bank_sampah.git/src/domains/repositories"
	services "github.com/iyunrozikin26/bank_sampah.git/src/domains/services/product"
	"gorm.io/gorm"
)

func addProduRoutes(ctx *gin.Context, rg *gin.RouterGroup, db *gorm.DB) {
	repository := repositories.NewProductRepository(db)
	service := services.NewProductService(repository)
	controller := controllers.NewProductController(service, ctx)

	product := rg.Group("/products")
	{
		product.GET("/", controller.IndexProduct)
		// product.GET("/:id", controller.GetByID)
		// product.POST("/", controller.Create)
		// product.PATCH("/:id", controller.Update)
		// product.DELETE("/:id", controller.Delete)
	}
}
