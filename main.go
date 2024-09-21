package main

import (
	"books/database"
	"books/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	defer database.DB.Close()

	router := gin.Default()

	routes.BookRoutes(router)
	routes.CategoryRoutes(router)

	router.Run(":8080")
}
