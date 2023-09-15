package controller

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func AdminRegister(ctx *gin.Context) {
	newAdmin := models.Admin{}
	db := database.GetDb()

	err := helpers.Binding(ctx, &newAdmin, appJson)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	err = db.Create(&newAdmin).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": true,
		"data":    newAdmin,
	})
}

func AdminLogin(ctx *gin.Context) {
	admin := models.Admin{}
	db := database.GetDb()
	
	err:= helpers.Binding(ctx, &admin, appJson)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"error":   "Bad request",
		})
		return
	}
	password := admin.Password
	err = db.Debug().Where("email= ?", admin.Email).Take(&admin).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	comparePass, err := helpers.ComparePass([]byte(admin.Password), []byte(password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "invalid password",
		})
		return
	}
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	token := helpers.GenerateToken(admin.ID, admin.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
