package app

import (
	"final-project/database"
	"final-project/routes"

	"github.com/gin-gonic/gin"
)

func App() {
	app := gin.Default()
	database.ConnectDB()

	routes.Routes(app)
	app.Run(":8080")

}
