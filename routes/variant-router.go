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
		// variant.Use(middlewares.ProductAuthorization())
		// variant.POST("/", middlewares.VariantAuthorization(), controller.CreateVariant)
		variant.POST("/",  controller.CreateVariant)
		variant.PUT("/:uuid", middlewares.VariantAuthorization(), controller.UpdateVariantByUuid)
		variant.DELETE("/:uuid", controller.DeleteVariantByUUID)
	}
}
