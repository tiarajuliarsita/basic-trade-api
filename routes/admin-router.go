package routes

import (
	"final-project/controller"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(app *gin.RouterGroup) {
	route := app
	admin := route.Group("/auth")
	{
		admin.POST("/register", controller.AdminRegister)
		admin.POST("/login", controller.AdminLogin)
	}
}
