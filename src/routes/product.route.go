package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nasrunrozikinmm/bank_sampah.git/src/modules/product"
	"gorm.io/gorm"
)

func addProduRoutes(ctx *gin.Context, rg *gin.RouterGroup, db *gorm.DB) {
	repository := product.NewProductRepository(db)
	service := product.NewProductService(repository)
	controller := product.NewProductController(service, ctx)

	product := rg.Group("/products")
	{
		product.GET("/", controller.IndexProduct)
		product.GET("/:id", controller.GetByID)
		product.POST("/", controller.CreateProduct)
		// product.PATCH("/:id", controller.Update)
		// product.DELETE("/:id", controller.Delete)
	}
}
