package routes

import (
	"final-project/controller"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(app *gin.RouterGroup) {
	route := app
	product := route.Group("/products")
	{
		product.GET("/", controller.GetAllProduct)
		product.GET("/:uuid", controller.GetProductByUUID)

		product.Use(middlewares.Authentication())
		product.POST("/", controller.CreateProduct)
		product.DELETE("/:uuid", middlewares.ProductAuthorization(), controller.DeleteProductByUUID)
		product.PUT("/:uuid", middlewares.ProductAuthorization(), controller.UpdateProductbyUUID)
	}
}
