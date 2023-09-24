package route

import (
	"github.com/gin-gonic/gin"
	userControl "github.com/iyunrozikin26/bank_sampah.git/src/domain/controllers"
	userRepo "github.com/iyunrozikin26/bank_sampah.git/src/domain/repositories"
	userServ "github.com/iyunrozikin26/bank_sampah.git/src/domain/services/user"
	"gorm.io/gorm"
)

var (
	ctx *gin.Context
)

func addUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	userRepository := userRepo.NewRepository(db)
	userService := userServ.NewUserService(userRepository)
	userController := userControl.NewUserController(userService, ctx)

	user := rg.Group("/users")
	{
		user.GET("/", userController.Index)
		user.GET("/:id", userController.GetByID)
		user.POST("/", userController.Create)
		user.PATCH("/:id", userController.Update)
		user.DELETE("/:id", userController.Delete)
	}
}
