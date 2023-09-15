package controller

import (
	"errors"
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"final-project/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func CreateVariant(ctx *gin.Context) {
	db := database.GetDb()
	newVariant := request.Variant{}
	adminData := ctx.MustGet("adminData").(jwt.MapClaims) 
	adminID := uint(adminData["id"].(float64))

	err:=helpers.Binding(ctx, &newVariant, appJson)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var product models.Product

	err = db.Where("uuid = ?", newVariant.UUID).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Data Not Found",
		})
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Product Not Found",
			"error":   err.Error(),
		})
		return
	}

	// Periksa apakah admin memiliki produk dengan UUID yang sesuai
	if product.AdminID != adminID {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "You are not authorized to create a variant for this product",
		})
		return
	}

	variant := models.Variant{
		VariantName: newVariant.VariantName,
		Quantity:    newVariant.Quantity,
		ProductID:   product.ID,
	}

	if err := db.Create(&variant).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create variant",
			"error":   err.Error(),
		})
		return
	}
	// err = db.Where("uuid = ?", newVariant.UUID).Find(&variant).Error
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"message": err.Error(),
	// 	})
	// 	return 
	// }

	ctx.JSON(http.StatusOK, gin.H{
		"variant": variant,
	})
}

func GetAllVariants(ctx *gin.Context) {
	db := database.GetDb()
	variants := []models.Variant{}
	err := db.Model(&models.Variant{}).Preload("Product").Find(&variants).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"variants": variants,
	})
}

func GetVariantByUuid(ctx *gin.Context) {
	db := database.GetDb()
	variantUUID := ctx.Param("uuid")
	variant := models.Variant{}
	err := db.Preload("Product").Where("uuid = ?", variantUUID).First(&variant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Variant not found",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"variants": variant,
		"succes":   true,
	})
}

func UpdateVariantByUuid(ctx *gin.Context) {
	db := database.GetDb()
	variantUUID := ctx.Param("uuid")
	newVariant := models.Variant{}
	err:=helpers.Binding(ctx, &newVariant, appJson)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	variant := models.Variant{}
	
	err = db.Model(&newVariant).Where("uuid = ?", variantUUID).Updates(newVariant).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = db.Where("uuid = ?", variantUUID).Find(&variant).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"variant": variant,
	})
}

func DeleteVariantByUUID(ctx *gin.Context) {
    db := database.GetDb()
    variantUUID := ctx.Param("uuid")
    var variant models.Variant
    if err := db.Where("uuid = ?", variantUUID).Delete(&variant).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Internal Server Error",
            "message": "Failed to delete the variant",
        })
        return
    }

    // Variant berhasil dihapus
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Variant deleted successfully",
    })
}

