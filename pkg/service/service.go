package service

import (
	"github.com/shooosty/rd-app/models"
	"github.com/shooosty/rd-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user models.User) (string, error)
	CreateEmployee(user models.User) (string, error)
	GenerateToken(username, password string) (string, error)
	GetCurrentUser(username, password string) (models.User, error)
	ParseToken(token string) (string, error)
}

type Users interface {
	GetAll() ([]models.User, error)
	GetById(userId string) (models.User, error)
	Delete(userId string) error
	Update(userId string, input models.UpdateUserInput) error
}

type Orders interface {
	GetAll() ([]models.Order, error)
	GetAllForUser(userId string) ([]models.Order, error)
	GetAllForPhotographer(photographerId string) ([]models.Order, error)
	GetAllForDesigner(designerId string) ([]models.Order, error)
	GetById(orderId string) (models.Order, error)
	Create(order models.Order) (string, error)
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

type Service struct {
	Authorization
	Users
	Orders
	Persons
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Users:         NewUserService(repos.Users),
		Orders:        NewOrderService(repos.Orders),
		Persons:       NewPersonService(repos.Persons),
	}
}
