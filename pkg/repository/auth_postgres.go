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

func (r *AuthPostgres) CreateUser(user models.User) (string, error) {
	type Result struct {
		ID    string
		Name  string
		Email string
	}

	var result Result

	newUser := models.User{Name: user.Name, Email: user.Email, PasswordHash: user.Password, Role: user.Role, Phone: user.Phone}
	err := db.Table(usersTable).Create(&newUser).Scan(&result).Error

	return result.ID, err
}

func (r *AuthPostgres) CreateEmployee(user models.User, generatedPassword string) (string, error) {
	type Result struct {
		ID    string
		Name  string
		Email string
	}

	var result Result

	newUser := models.User{Name: user.Name, Email: user.Email, PasswordHash: user.Password, Role: user.Role, Phone: user.Phone}
	err := db.Table(usersTable).Create(&newUser).Scan(&result).Error

	if err == nil {
		SendPasswordToEmployee(generatedPassword, user.Name, user.Email)
	}

	return result.ID, err
}

func (r *AuthPostgres) ResetPassword(input models.ResetPasswordInput, newPassword string, generatedPassword string) error {
	updatedUser := models.User{PasswordHash: newPassword}

	err := db.Table(usersTable).Where("email = ?", input.Email).Updates(&updatedUser).Error

	if err == nil {
		SendRestoredPassword(generatedPassword, input.Email)
	}

	return err
}

func (r *AuthPostgres) ChangePassword(input models.ChangePasswordInput) error {
	updatedUser := models.User{PasswordHash: input.NewPassword}

	err := db.Table(usersTable).Where("email = ? AND password_hash = ?", input.Email, input.Password).Updates(&updatedUser).Error

	return err
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	err := db.Table(usersTable).Where("email = ? AND password_hash = ?", email, password).Find(&user).Error

	return user, err
}
