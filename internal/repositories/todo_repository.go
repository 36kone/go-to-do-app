package repositories

import "go-to-do-app/internal/models"

type TodoRepository interface {
	FindAll() []models.Todo
	FindByID(id int) (models.Todo, bool)
	Create(todo models.Todo) models.Todo
}
