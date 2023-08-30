package app

import (
	"final-project/database"
	"final-project/routes"
)

func App() {
	database.ConnectDB()
	app := routes.Routes()
	app.Run(":8080")

}
