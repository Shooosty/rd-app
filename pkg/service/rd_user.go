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

func (s *UserService) ChangePassword(userId string, input models.ChangePasswordInput) error {
	input.Password = generatePasswordHash(input.Password)
	input.PasswordHash = generatePasswordHash(input.PasswordHash)

	return s.repo.ChangePassword(userId, input)
}
