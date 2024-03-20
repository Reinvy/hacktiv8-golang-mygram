package routes

import (
	"mygram/controller"
	"mygram/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(g *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userController := controller.NewUserController(userRepository)

	api := g.Group("api/v1")
	api.POST("/login", userController.Login)
	api.POST("/register", userController.Register)

	userApi := api.Group("users")
	userApi.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "passed pong",
		})
	})

}
