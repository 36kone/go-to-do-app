package repositories

import (
	"go-to-do-app/internal/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll() []models.TodoModel
	FindByID(id int) (models.TodoModel, bool)
	Create(todo models.TodoModel) models.TodoModel
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) FindAll() []models.TodoModel {
	var todos []models.TodoModel
	r.db.Find(&todos)
	return todos
}

func (r *todoRepository) FindByID(id int) (models.TodoModel, bool) {
	var todo models.TodoModel
	result := r.db.First(&todo, id)
	if result.Error != nil {
		return models.TodoModel{}, false
	}
	return todo, true
}

func (r *todoRepository) Create(todo models.TodoModel) models.TodoModel {
	r.db.Create(&todo)
	return todo
}
