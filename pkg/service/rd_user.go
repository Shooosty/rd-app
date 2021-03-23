package service

import (
	"github.com/shooosty/rd-app/models"
	"github.com/shooosty/rd-app/pkg/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(userId string) (models.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) Delete(userId string) error {
	return s.repo.Delete(userId)
}

func (s *UserService) Update(userId string, input models.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, input)
}
