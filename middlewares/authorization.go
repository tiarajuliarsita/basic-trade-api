package middlewares

import (
	"final-project/database"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDb()
		productUUID := ctx.Param("productUUID")
		adminData := ctx.MustGet("userData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))

		var getProduct models.Product
		err := db.Select("admin_id").Where("uuid = ?", productUUID).First(&getProduct).Error
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data Not Found",
			})
			return
		}
		if getProduct.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()

	}
}
