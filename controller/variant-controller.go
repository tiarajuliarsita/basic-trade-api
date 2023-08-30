package controller

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateVariant(ctx *gin.Context) {
	db := database.GetDb()
	contentType := helpers.GetContentType(ctx)
	newVariant := models.Variant{}

	if contentType == appJson {
		ctx.ShouldBindJSON(&newVariant)
	} else {
		ctx.ShouldBind(&newVariant)
	}

	err := db.Create(&newVariant).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"variant": newVariant,
	})
}
