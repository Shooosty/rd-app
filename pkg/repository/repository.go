package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CreateEmployee(user models.User, generatedPassword string) (string, error)
	GetUser(username, password string) (models.User, error)
	ResetPassword(input models.ResetPasswordInput, newPassword string, generatedPassword string) error
}

type Users interface {
	GetAll() ([]models.User, error)
	GetById(userId string) (models.User, error)
	Delete(userId string) error
	Update(userId string, input models.UpdateUserInput) error
	ChangePassword(userId string, input models.ChangePasswordInput) error
}

type Orders interface {
	GetAll() ([]models.Order, error)
	Create(order models.Order) (string, error)
	GetAllForUser(userId string) ([]models.Order, error)
	GetAllForPhotographer(photographerId string) ([]models.Order, error)
	GetAllForDesigner(designerId string) ([]models.Order, error)
	GetById(orderId string) (models.Order, error)
	Delete(orderId string) error
	Update(orderId string, input models.UpdateOrderInput) error
}

type Persons interface {
	GetAll() ([]models.Person, error)
	GetAllByOrderId(orderId string) ([]models.Person, error)
	Create(person models.Person) (string, error)
	Delete(personId string) error
	Update(personId string, input models.UpdatePersonInput) error
}

type Repository struct {
	Authorization
	Users
	Orders
	Persons
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users:         NewUserPostgres(db),
		Orders:        NewOrderPostgres(db),
		Persons:       NewPersonPostgres(db),
	}
}
