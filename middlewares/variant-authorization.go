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
		adminData := ctx.MustGet("adminData").(jwt.MapClaims)
		adminID := uint(adminData["id"].(float64))
		var getProduct models.Product
		err := db.Where("admin_id = ?", adminID).Find(&getProduct).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   err.Error(),
					"message": "Data Not Found",
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Internal Server Error",
			})
			return
		}

		ctx.Next()
	}
}

// func VariantAuthorization() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		db := database.GetDb()
// 		adminData := ctx.MustGet("adminData").(jwt.MapClaims)
// 		adminID := uint(adminData["id"].(float64))

// 		// Cek apakah ID admin ada dalam tabel produk
// 		var productCount int64
// 		err := db.Model(&models.Product{}).Where("admin_id = ?", adminID).Count(&productCount).Error
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 				"error":   err.Error(),
// 				"message": "Internal Server Error",
// 			})
// 			return
// 		}

// 		if productCount == 0 {
// 			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 				"message": "Data Not Found",
// 			})
// 			return
// 		}

// 		ctx.Next()
// 	}
// }
