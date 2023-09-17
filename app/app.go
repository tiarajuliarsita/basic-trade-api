package app

import (
	"final-project/database"
	"final-project/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func App() {
	app := gin.Default()
	database.ConnectDB()

	routes.Routes(app)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	app.Run(":"+ port)

}
