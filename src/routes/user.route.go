package route

import (
	"github.com/gin-gonic/gin"
	"github.com/iyunrozikin26/bank_sampah.git/src/domains/controllers"
	"github.com/iyunrozikin26/bank_sampah.git/src/domains/repositories"
	services "github.com/iyunrozikin26/bank_sampah.git/src/domains/services/user"
	"gorm.io/gorm"
)

func addUserRoutes(ctx *gin.Context, rg *gin.RouterGroup, db *gorm.DB) {
	repository := repositories.NewUserRepository(db)
	service := services.NewUserService(repository)
	controller := controllers.NewUserController(service, ctx)

	user := rg.Group("/users")
	{
		user.GET("/", controller.Index)
		user.GET("/:id", controller.GetByID)
		user.POST("/", controller.Create)
		user.PATCH("/:id", controller.Update)
		user.DELETE("/:id", controller.Delete)
	}
}
