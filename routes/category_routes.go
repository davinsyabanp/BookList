package routes

import (
	"books/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	router.POST("/categories", controllers.CreateCategory)
	router.GET("/categories", controllers.ListCategories)
	router.GET("/categories/:id", controllers.GetCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)
}
