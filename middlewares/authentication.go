package middlewares

import (
	"final-project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenVerify, err := helpers.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unathenticated",
				"message": err.Error(),
			})
			return
		}
		ctx.Set("adminData", tokenVerify)
		ctx.Next()
	}

}
