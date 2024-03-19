package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "passed pong",
		})
	})

}
