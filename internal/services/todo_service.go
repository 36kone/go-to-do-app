package services

import (
	"go-to-do-app/internal/models"
	"go-to-do-app/internal/repositories"
)

type TodoService struct {
	repo repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetTodos() []models.TodoModel {
	return s.repo.FindAll()
}

func (s *TodoService) GeTodoByID(id int) (models.TodoModel, bool) {
	return s.repo.FindByID(id)
}

func (s *TodoService) CreateTodo(title string) models.TodoModel {
	todo := models.TodoModel{
		Title:     title,
		Completed: false,
	}
	return s.repo.Create(todo)
}
