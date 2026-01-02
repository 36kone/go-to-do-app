package routes

import (
	"go-to-do-app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.Engine, controller *controllers.TodoController) {
	todos := r.Group("/todos")
	{
		todos.GET("/", controller.GetTodos)
		todos.GET("/:id", controller.GetTodoById)
		todos.POST("/", controller.Create)
	}
}
