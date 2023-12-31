package controller

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"final-project/pagnation"
	"final-project/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateProduct(ctx *gin.Context) {
	newProduct := request.Product{}
	db := database.GetDb()

	adminData := ctx.MustGet("adminData").(jwt.MapClaims)
	adminID := uint(adminData["id"].(float64))

	err := helpers.Binding(ctx, &newProduct, appJson)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	fileName := helpers.RemoveExtention(newProduct.File.Filename)
	uploadResult, err := helpers.UploadFile(newProduct.File, fileName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	product := models.Product{}
	product.AdminID = adminID
	product.ImageURL = uploadResult
	product.Name = newProduct.Name

	err = db.Create(&product).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func GetAllProduct(ctx *gin.Context) {
	db := database.GetDb()
	// search and pagnation
	search := pagnation.Search(ctx)
	lastPage, limitInt, offsetInt, page, total := pagnation.Pagnation(ctx)

	products := []models.Product{}
	err := db.Model(&models.Product{}).Where("name LIKE ?", "%"+search+"%").Offset(offsetInt).Limit(limitInt).Preload("Admin").Preload("Variants").Find(&products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
		"pagination": gin.H{
			"last_page": lastPage,
			"limit":     limitInt,
			"offset":    offsetInt,
			"page":      page,
			"total":     total,
		},
	})
}

func GetProductByUUID(ctx *gin.Context) {
	db := database.GetDb()
	productUUID := ctx.Param("uuid")
	products := []models.Product{}

	err := db.Preload("Variants").Where("uuid = ?", productUUID).Find(&products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"product": products,
	})
}

func DeleteProductByUUID(ctx *gin.Context) {
	db := database.GetDb()
	productUUID := ctx.Param("uuid")

	var product models.Product
	if err := db.Where("uuid = ?", productUUID).Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to delete the product",
		})
		return
	}

	// Produk berhasil dihapus
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
		"data":    "null",
	})
}

func UpdateProductbyUUID(ctx *gin.Context) {
	db := database.GetDb()
	newProduct := request.Product{}
	var product models.Product

	productUUID := ctx.Param("uuid")
	err := helpers.Binding(ctx, &newProduct, appJson)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if newProduct.File != nil {
		fileName := helpers.RemoveExtention(newProduct.File.Filename)
		uploadResult, err := helpers.UploadFile(newProduct.File, fileName)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})

		}
		product.ImageURL = uploadResult
	}

	product.Name = newProduct.Name
	err = db.Model(&product).Where("uuid = ?", productUUID).Updates(product).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = db.Where("uuid = ?", productUUID).Find(&product).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	// succes updated product
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": product,
	})

}
