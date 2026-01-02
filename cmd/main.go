package main

import (
	"github.com/gin-gonic/gin"

	"go-to-do-app/internal/controllers"
	"go-to-do-app/internal/database"
	"go-to-do-app/internal/repositories"
	"go-to-do-app/internal/routes"
	"go-to-do-app/internal/services"
)

func main() {
	database.DataBaseConnection()

	repos := repositories.NewRepositories(database.DB)

	svcs := services.NewServices(repos)

	todoController := controllers.NewTodoController(svcs.Todo)
	userController := controllers.NewUserController(svcs.User)

	r := gin.Default()
	routes.RegisterRoutes(r, routes.Controllers{
		Todo: todoController,
		User: userController,
	})

	r.Run(":8080")
}
