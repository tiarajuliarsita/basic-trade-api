package pagnation

import "github.com/gin-gonic/gin"

func Search(ctx *gin.Context)string{
	search:= ctx.Query("search")
	return search
}