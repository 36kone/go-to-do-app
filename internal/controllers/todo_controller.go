func (c *TodoController) GetTodos(ctx *gin.Context) {
	todos := c.service.GetTodos()
	ctx.JSON(200, todos)
}
