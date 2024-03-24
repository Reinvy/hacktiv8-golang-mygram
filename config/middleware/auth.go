package middleware

import (
	"mygram/domain/dto"
	"mygram/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	claims, err := util.GetJWTClaims(bearerToken)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	ctx.Set("claims", claims)

	ctx.Next()
}
