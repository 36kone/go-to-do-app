package database

import (
	"log"

	"go-to-do-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DataBaseConnection() {
	connectinString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectinString))
	if err != nil {
		log.Panic("Failed to connect to database")
	}
	err = DB.AutoMigrate(&models.UserModel{}, &models.TodoModel{})
	if err != nil {
		log.Panic("Failed to migrate database:", err)
	}
}
