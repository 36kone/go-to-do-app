package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-to-do-app/internal/services"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GeUsers(ctx *gin.Context) {
	users := c.service.GetUsers()
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, found := c.service.GetUserByID(id)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Create(ctx *gin.Context) {
	var input struct {
		Title string `json:"title"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	user := c.service.CreateUser(input.Title)
	ctx.JSON(http.StatusCreated, user)
}
