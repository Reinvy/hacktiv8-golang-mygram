package routes

import (
	"mygram/config/middleware"
	"mygram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(g *gin.Engine, db *gorm.DB) {
	userService := service.NewUserService(db)
	photoService := service.NewPhotoService(db)
	commentService := service.NewCommentService(db)
	socialMediaService := service.NewSocialMediaService(db)

	api := g.Group("api/v1")
	api.POST("users/login", userService.Login)
	api.POST("users/register", userService.Register)

	userApi := api.Group("users", middleware.AuthMiddleware)
	userApi.PUT("/:id", userService.Update)
	userApi.DELETE("/:id", userService.Delete)

	photoApi := api.Group("photos", middleware.AuthMiddleware)
	photoApi.GET("/", photoService.GetAll)
	photoApi.GET("/:id", photoService.GetPhotoByID)
	photoApi.POST("/", photoService.Create)
	photoApi.PUT("/:id", photoService.Update)
	photoApi.DELETE("/:id", photoService.Delete)

	commentApi := api.Group("comments", middleware.AuthMiddleware)
	commentApi.GET("/", commentService.GetAll)
	commentApi.GET("/:id", commentService.GetCommentByID)
	commentApi.POST("/", commentService.Create)
	commentApi.PUT("/:id", commentService.Update)
	commentApi.DELETE("/:id", commentService.Delete)

	socialMediaAPI := api.Group("socialmedias", middleware.AuthMiddleware)
	socialMediaAPI.GET("/", socialMediaService.GetAll)
	socialMediaAPI.GET("/:id", socialMediaService.GetSocialMediaByID)
	socialMediaAPI.POST("/", socialMediaService.Create)
	socialMediaAPI.PUT("/:id", socialMediaService.Update)
	socialMediaAPI.DELETE("/:id", socialMediaService.Delete)

}
