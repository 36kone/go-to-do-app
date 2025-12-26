package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	repo := repositories.NewInMemoryTodoRepository()
	service := services.NewTodoService(repo)
	controller := controllers.NewTodoController(service)

	routes.RegisterTodoRoutes(r, controller)

	r.Run(":8080")
}
