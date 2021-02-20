package routers

import (
	"github.com/gin-gonic/gin"
	
	"gin-orm-prac/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/todo", controllers.CreateTodo)
		v1.GET("/todos", controllers.FetchAllTodo)
		v1.GET("/todo/:id", controllers.FetchSingleTodo)
	// 	v1.PUT("/todo/:id", UpdateTodo)
		v1.DELETE("/todo/:id", controllers.DeleteTodo)
	}

	return router;
}