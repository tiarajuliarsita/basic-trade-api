package controller

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
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

	contentType := helpers.GetContentType(ctx)
	if contentType == appJson {
		ctx.ShouldBindJSON(&newProduct)
	} else {
		ctx.ShouldBind(&newProduct)
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func GetAllProduct(ctx *gin.Context) {
	db := database.GetDb()
	products := []models.Product{}

	err := db.Preload("Variants").Find(&products).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": products,
	})

}
func GetProductByUUID(ctx *gin.Context) {
	db := database.GetDb()
	productUUID := ctx.Param("uuid")
	products := []models.Product{}

	err := db.Preload("Admin").Preload("Variants").Where("uuid = ?", productUUID).Find(&products).Error
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
	adminData := ctx.MustGet("adminData").(jwt.MapClaims)
	adminID := uint(adminData["id"].(float64))

	product := models.Product{}
	productUUID := ctx.Param("uuid")
	// admin := models.Admin{}
	err := db.Where("admin_id=?", adminID).First(&product).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete the product",
		})
		return
	}
	// Perbaikan 1: Hapus tanda kurung kosong
	err = db.Where("uuid = ?", productUUID).Delete(&product).Error
	// Handle kesalahan jika penghapusan gagal
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": "Failed to delete the product",
		})
		return
	}

	// Produk berhasil dihapus
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
