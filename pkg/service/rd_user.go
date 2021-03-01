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

func (s *UserService) GetById(userId int) (models.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserService) Update(userId int, input models.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, input)
}
