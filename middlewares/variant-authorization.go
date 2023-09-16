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
		//cari variant dengan uuid sesuai params
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

		//  cari product dengan berdasarkan product id di tabel variant
		var product models.Product
		err = db.Where("id = ?", variant.ProductID).Find(&product).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   err.Error(),
					"message": "Data Not Found",
				})
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "Failed to retrieve product",
			})
			return
		}
		
		if product.AdminID != adminID {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "You are not authorized for this data.",
			})
			
			return 
		}
		

		ctx.Next()
	}
}
