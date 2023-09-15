package routes

import (
	"final-project/controller"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func VariantsRoutes(app *gin.RouterGroup) {
	route := app
	variant := route.Group("/products/variants")
	{
		variant.GET("/", controller.GetAllVariants)
		variant.GET("/:uuid", controller.GetVariantByUuid)

		variant.Use(middlewares.Authentication())
		variant.POST("/",  controller.CreateVariant)
		variant.PUT("/:uuid", middlewares.VariantAuthorization(), controller.UpdateVariantByUuid)
		variant.DELETE("/:uuid", middlewares.VariantAuthorization(), controller.DeleteVariantByUUID)
	}
}
