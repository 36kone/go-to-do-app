package routes

import (
	"go-to-do-app/internal/controllers"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	Todo *controllers.TodoController
	User *controllers.UserController
}

func RegisterRoutes(r *gin.Engine, c Controllers) {
	RegisterTodoRoutes(r, c.Todo)
	RegisterUserRoutes(r, c.User)
}
