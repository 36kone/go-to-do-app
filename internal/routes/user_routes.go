package routes

import (
	"go-to-do-app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, controller *controllers.UserController) {
	users := r.Group("/users")
	{
		users.POST("/", controller.Create)
		users.GET("/", controller.GeUsers)
		users.GET("/:id", controller.GetUserByID)
	}
}
