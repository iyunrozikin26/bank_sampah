package route

import (
	"github.com/gin-gonic/gin"
	"github.com/iyunrozikin26/bank_sampah.git/src/modules/user"
	"gorm.io/gorm"
)

func addUserRoutes(ctx *gin.Context, rg *gin.RouterGroup, db *gorm.DB) {
	repository := user.NewUserRepository(db)
	service := user.NewUserService(repository)
	controller := user.NewUserController(service, ctx)

	user := rg.Group("/users")
	{
		user.GET("/", controller.Index)
		user.GET("/:id", controller.GetByID)
		user.POST("/", controller.Create)
		user.PATCH("/:id", controller.Update)
		user.DELETE("/:id", controller.Delete)
	}
}
