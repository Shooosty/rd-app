package service

import (
	"github.com/shooosty/rd-app/models"
	"github.com/shooosty/rd-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	GetCurrentUser(username, password string) (models.User, error)
	ParseToken(token string) (int, error)
}

type Users interface {
	GetAll() ([]models.User, error)
	GetById(userId int) (models.User, error)
	Delete(userId int) error
	Update(userId int, input models.UpdateUserInput) error
}

type Orders interface {
	GetAll() ([]models.Order, error)
	GetAllForUser(userId int) ([]models.Order, error)
	GetById(orderId int) (models.Order, error)
	Create(order models.Order) (int, error)
	Delete(orderId int) error
	Update(orderId int, input models.UpdateOrderInput) error
}

type Service struct {
	Authorization
	Users
	Orders
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Users:         NewUserService(repos.Users),
		Orders:        NewOrderService(repos.Orders),
	}
}
