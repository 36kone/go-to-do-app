type TodoService struct {
	repo repositories.TodoRepository
}

func (s *TodoService) GetTodos() []models.Todo {
	return s.repo.GetAll()
}

func (s *TodoService) CreateTodo(title string) models.Todo {
	todo := models.Todo{
		Title:     title,
		Completed: false,
	}
	return s.repo.Create(todo)
}
