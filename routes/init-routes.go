package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(app *gin.Engine) {
	// app := gin.Default()
	routes := app.Group("")
	AdminRoutes(routes)
	ProductRoutes(routes)
	VariantsRoutes(routes)

	// adminRoute := app.Group("/auth")
	// {
	// 	adminRoute.POST("/register", controller.AdminRegister)
	// 	adminRoute.POST("/login", controller.AdminLogin)
	// }
	// productRoute := app.Group("/products")
	// {
	// 	productRoute.GET("/", controller.GetAllProduct)
	// 	productRoute.GET("/:uuid", controller.GetProductByUUID)

	// 	productRoute.Use(middlewares.Authentication())
	// 	productRoute.POST("/", controller.CreateProduct)
	// 	productRoute.DELETE("/:uuid", middlewares.ProductAuthorization(), controller.DeleteProductByUUID)
	// 	productRoute.PUT("/:uuid", middlewares.ProductAuthorization(), controller.UpdateProductbyUUID)
	// }

	// variantsRoute := app.Group("/products/variants")
	// {
	// 	variantsRoute.GET("/", controller.GetAllVariants)
	// 	variantsRoute.GET("/:uuid", controller.GetVariantByUuid)

	// 	variantsRoute.Use(middlewares.Authentication())
	// 	variantsRoute.POST("/", middlewares.VariantAuthorization(), controller.CreateVariant)
	// 	variantsRoute.PUT("/:uuid", middlewares.VariantAuthorization(), controller.UpdateVariantByUuid)
	// 	variantsRoute.DELETE("/:uuid", middlewares.VariantAuthorization(), controller.DeleteVariantByUUID)
	// }
	// return app
}
