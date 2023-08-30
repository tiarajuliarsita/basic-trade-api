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

	contentType := helpers.GetContentType(ctx)
	if contentType == appJson {
		ctx.ShouldBindJSON(&newAdmin)
	} else {
		ctx.ShouldBind(&newAdmin)
	}

	err := db.Create(&newAdmin).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
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

	contentType := helpers.GetContentType(ctx)
	if contentType == appJson {
		ctx.ShouldBindJSON(&admin)
	} else {
		ctx.ShouldBind(&admin)
	}

	password := admin.Password
	err := db.Debug().Where("email= ?", admin.Email).Take(&admin).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Invalid email error disini",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(admin.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   err,
			"message": "Invalid password",
		})
		return
	}

	token := helpers.GenerateToken(admin.ID, admin.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": "token",
		"TOKEN":token,
	})
}
