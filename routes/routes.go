package routes

import (
	"mygram/config/middleware"
	"mygram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(g *gin.Engine, db *gorm.DB) {
	userSerrvice := service.NewUserService(db)
	photoService := service.NewPhotoService(db)
	commentService := service.NewCommentService(db)

	api := g.Group("api/v1")
	api.POST("/login", userSerrvice.Login)
	api.POST("/register", userSerrvice.Register)

	userApi := api.Group("users", middleware.AuthMiddleware)
	userApi.PUT("/:id", userSerrvice.Update)
	userApi.DELETE("/:id", userSerrvice.Delete)

	photoApi := api.Group("photos", middleware.AuthMiddleware)
	photoApi.GET("/", photoService.GetAll)
	photoApi.GET("/:id", photoService.GetPhotoByID)
	photoApi.POST("/", photoService.Create)
	photoApi.PUT("/:id", photoService.Update)
	photoApi.DELETE("/:id", photoService.Delete)

	commentApi := api.Group("comments", middleware.AuthMiddleware)
	// commentApi.GET("/", commentService.GetAll)
	commentApi.POST("/", commentService.Create)
	// commentApi.PUT("/:id", commentService.Update)
	// commentApi.DELETE("/:id", commentService.Delete)

}
