package service

import (
	"github.com/shooosty/rd-app"
	"github.com/shooosty/rd-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user rd_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	GetCurrentUser(username, password string) (rd_app.User, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
