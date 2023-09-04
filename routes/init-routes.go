package routes

import (
	"final-project/controller"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	app := gin.Default()

	adminRoute := app.Group("/auth")
	{
		adminRoute.POST("/register", controller.AdminRegister)
		adminRoute.POST("/login", controller.AdminLogin)
	}
	productRoute := app.Group("/products")
	{
		productRoute.GET("/", controller.GetAllProduct)
		productRoute.GET("/:uuid", controller.GetProductByUUID)

		productRoute.Use(middlewares.Authentication())
		productRoute.POST("/", controller.CreateProduct)
		productRoute.DELETE("/:uuid", controller.DeleteProductByUUID)
		productRoute.PUT("/:uuid", controller.UpdateProductbyUUID)
	}

	variantsRoute := app.Group("/products/variants")
	{
		variantsRoute.Use(middlewares.Authentication())
		variantsRoute.POST("/",  controller.CreateVariant)
	}
	return app
}
