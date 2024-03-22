package middleware

import (
	"mygram/model/dto"
	"mygram/util"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	authorizationValue := ctx.GetHeader("Authorization")
	splittedValue := strings.Split(authorizationValue, "Bearer ")
	if len(splittedValue) <= 1 {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}
	jwtToken := splittedValue[1]

	claims, err := util.GetJWTClaims(jwtToken)
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

func AdminMiddleware(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	mapClaims, ok := claims.(map[string]any)
	if !ok {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	isAdmin, ok := mapClaims["admin"]
	if !ok {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	boolIsAdmin, ok := isAdmin.(bool)
	if !ok {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	if !boolIsAdmin {
		var r dto.Response = dto.Response{
			Status:  http.StatusText(http.StatusUnauthorized),
			Message: "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	ctx.Next()
}
