package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAll() ([]models.User, error) {
	var users []models.User
	err := db.Table(usersTable).Find(&users).Error

	return users, err
}

func (r *UserPostgres) GetById(userId string) (models.User, error) {
	var user models.User
	err := db.Table(usersTable).Where("id = ?", userId).Find(&user).Error

	return user, err
}

func (r *UserPostgres) Delete(userId string) error {
	users := make([]*Users, 0)
	err := db.Table(usersTable).Where("id = ?", userId).Delete(&users).Error

	return err
}

func (r *UserPostgres) Update(userId string, input models.UpdateUserInput) error {
	err := db.Table(usersTable).Where("id = ?", userId).Updates(&input).Error

	return err
}

func (r *UserPostgres) ChangePassword(userId string, input models.ChangePasswordInput) error {
	var user models.User
	err := db.Table(usersTable).Where("id = ? AND password_hash = ?", userId, input.Password).Find(&user).Error

	if err == nil {
		db.Table(usersTable).Where("id = ? AND password_hash = ?", userId, input.Password).Updates(&input)
	}

	return err
}
