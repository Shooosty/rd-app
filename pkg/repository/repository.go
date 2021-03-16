package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Users interface {
	GetAll() ([]models.User, error)
	GetById(userId int) (models.User, error)
	Delete(userId int) error
	Update(userId int, input models.UpdateUserInput) error
}

type Orders interface {
	GetAll() ([]models.Order, error)
	Create(order models.Order) (int, error)
	GetAllForUser(userId int) ([]models.Order, error)
	GetById(orderId int) (models.Order, error)
	Delete(orderId int) error
	Update(orderId int, input models.UpdateOrderInput) error
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
