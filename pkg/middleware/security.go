package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matiasRaspa/Final-BackEnd-E3-API-GO/pkg/web"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewNotFoundApiErrorSecurityEmpty("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewNotFoundApiErrorSecurityInvalid("invalid token"))
			return
		}
		ctx.Next()
	}
}
