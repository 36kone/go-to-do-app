package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-to-do-app/internal/services"
)

type TodoController struct {
	service *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
	return &TodoController{service: service}
}

func (c *TodoController) GetTodos(ctx *gin.Context) {
	todos := c.service.GetTodos()
	ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodoById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	todo, found := c.service.GeTodoByID(id)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) Create(ctx *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	todo := c.service.CreateTodo(input.Title)
	ctx.JSON(http.StatusCreated, todo)
}
