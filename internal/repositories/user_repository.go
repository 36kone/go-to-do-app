package repositories

import (
	"go-to-do-app/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []models.UserModel
	FindByID(id int) (models.UserModel, bool)
	Create(user models.UserModel) models.UserModel
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAll() []models.UserModel {
	var users []models.UserModel
	r.db.Find(&users)
	return users
}

func (r *userRepository) FindByID(id int) (models.UserModel, bool) {
	var user models.UserModel
	result := r.db.First(&user, id)
	if result.Error != nil {
		return models.UserModel{}, false
	}
	return user, true
}

func (r *userRepository) Create(user models.UserModel) models.UserModel {
	r.db.Create(&user)
	return user
}
