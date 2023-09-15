package middlewares

import (
	"errors"
	"final-project/database"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func VariantAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDb()
		variantUUID := ctx.Param("uuid")
		adminData := ctx.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))

		var variant models.Variant
		err := db.Where("uuid = ?", variantUUID).Find(&variant).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Data Not Found",
			})
			return
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Variant not found",
				"error":   err.Error(),
			})
			return
		}

		// Periksa apakah admin memiliki produk dengan UUID yang sesuai
		var product models.Product
		if err := db.Where("id = ?", variant.ProductID).Find(&product).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": "Failed to retrieve product",
			})
			return
		}

		if product.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorized for this data.",
				"error":err.Error(),
			})
			return
		}

		ctx.Next()
	}
}
