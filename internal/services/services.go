package services

import "go-to-do-app/internal/repositories"

type Services struct {
	Todo *TodoService
	User *UserService
}

func NewServices(repos repositories.Repositories) Services {
	return Services{
		Todo: NewTodoService(repos.Todo),
		User: NewUserService(repos.User),
	}
}
