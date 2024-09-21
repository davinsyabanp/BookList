package main

import (
	"books/database"
	"books/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.Connect()
	defer database.DB.Close()

	// Create router
	router := gin.Default()

	// Define routes
	routes.BookRoutes(router)
	routes.CategoryRoutes(router)

	// Run server
	router.Run(":8080")
}
