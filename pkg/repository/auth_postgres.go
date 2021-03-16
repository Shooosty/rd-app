package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int

	err := db.Table(usersTable).Create(&user).Error

	return id, err
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User

	err := db.Table(usersTable).Where("email = ? AND password_hash = ?", email, password).Find(&user).Error

	return user, err
}
