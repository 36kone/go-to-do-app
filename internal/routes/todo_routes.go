func RegisterTodoRoutes(r *gin.Engine, controller *TodoController) {
	todos := r.Group("/todos")
	{
		todos.GET("/", controller.GetTodos)
		todos.POST("/", controller.CreateTodo)
	}
}
