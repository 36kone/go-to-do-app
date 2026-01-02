package repositories

import "gorm.io/gorm"

type Repositories struct {
	Todo TodoRepository
	User UserRepository
}

func NewRepositories(db *gorm.DB) Repositories {
	return Repositories{
		Todo: NewTodoRepository(db),
		User: NewUserRepository(db),
	}
}
