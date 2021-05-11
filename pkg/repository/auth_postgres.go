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

func (r *AuthPostgres) ResetPassword(email, password string) error {
	var user models.User
	user.PasswordHash = password

	err := db.Table(usersTable).Where("email = ?", email).Updates(&user.PasswordHash).Error

	if err == nil {
		SendPasswordToEmployee(user.Password, user.Name, user.Email)
	}

	return err
}

func (r *AuthPostgres) ChangePassword(email, password, newPassword string) error {
	var user models.User
	user.PasswordHash = newPassword

	err := db.Table(usersTable).Where("email = ? AND password_hash = ?", email, password).Updates(&user.PasswordHash).Error

	if err == nil {
		SendPasswordToEmployee(user.Password, user.Name, user.Email)
	}

	return err
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	err := db.Table(usersTable).Where("email = ? AND password_hash = ?", email, password).Find(&user).Error

	return user, err
}

func SendPasswordToEmployee(password string, name string, email string) {
	subject := "Регистрация в личном кабинете"
	text := "Ваш пароль для входа в кабинет: " + password
	html := "<b>" + name + "," + "</b>" + "<p>" + "Ваш пароль для входа в кабинет: " + password + "<p>" + "</br>" +
		"<p> Рекомендуем сменить пароль при первом входе в кабинет! </p>"
	_ = SendMail(subject, text, html, name, email)
}
