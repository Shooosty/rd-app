package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
)

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CreateEmployer(user models.User) (string, error)
	GetUser(username, password string) (models.User, error)
}

type Users interface {
	GetAll() ([]models.User, error)
	GetById(userId string) (models.User, error)
	Delete(userId string) error
	Update(userId string, input models.UpdateUserInput) error
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

type Repository struct {
	Authorization
	Users
	Orders
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users:         NewUserPostgres(db),
		Orders:        NewOrderPostgres(db),
	}
}
