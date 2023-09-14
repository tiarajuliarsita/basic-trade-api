package pagnation

import (
	"final-project/database"
	"final-project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagnation(ctx *gin.Context) (int, int, int, int, int64) {
	db := database.GetDb()
	offset := ctx.Query("offset")
	limit := ctx.Query("limit")
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = "10"
	}
	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)

	var total int64

	db.Model(&models.Product{}).Count(&total)
	page := offsetInt/limitInt + 1
	lastPage := int(total)/limitInt + 1

	return lastPage, limitInt, offsetInt, page, total
}
