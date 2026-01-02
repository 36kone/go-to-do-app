package services

import (
	"go-to-do-app/internal/models"
	"go-to-do-app/internal/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsers() []models.UserModel {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id int) (models.UserModel, bool) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(name string) models.UserModel {
	user := models.UserModel{
		Name: name,
	}
	return s.repo.Create(user)
}
