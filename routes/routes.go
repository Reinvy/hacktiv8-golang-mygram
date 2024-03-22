package routes

import (
	"mygram/controller"
	"mygram/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(g *gin.Engine, db *gorm.DB) {
	userController := controller.NewUserController(db)
	photoController := controller.NewPhotoController(db)
	commentController := controller.NewCommentController(db)

	api := g.Group("api/v1")
	api.POST("/login", userController.Login)
	api.POST("/register", userController.Register)

	userApi := api.Group("users", middleware.AuthMiddleware)
	userApi.PUT("/:id", userController.Update)
	userApi.DELETE("/:id", userController.Delete)

	photoApi := api.Group("photos", middleware.AuthMiddleware)
	photoApi.GET("/", photoController.GetAll)
	photoApi.GET("/:id", photoController.GetPhotoByID)
	photoApi.POST("/", photoController.Create)
	photoApi.PUT("/:id", photoController.Update)
	photoApi.DELETE("/:id", photoController.Delete)

	commentApi := api.Group("comments", middleware.AuthMiddleware)
	// commentApi.GET("/", commentController.GetAll)
	commentApi.POST("/", commentController.Create)
	// commentApi.PUT("/:id", commentController.Update)
	// commentApi.DELETE("/:id", commentController.Delete)

}
