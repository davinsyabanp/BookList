package routes

import (
	"books/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.ListBooks)
	router.GET("/books/:id", controllers.GetBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
