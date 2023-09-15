package helpers

import "github.com/gin-gonic/gin"

func GetContentType(ctx *gin.Context) string {
	return ctx.Request.Header.Get("Content-Type")
}

func Binding(ctx *gin.Context, models any, typeHeader string) error {

	contentType := GetContentType(ctx)

	var err error
	if contentType == typeHeader {
		err = ctx.ShouldBindJSON(models)
	} else {
		err = ctx.ShouldBind(models)
	}

	if err != nil {
		return err
	}

	return nil
}
