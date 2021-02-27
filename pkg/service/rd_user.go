package service

import (
	rd_app "github.com/shooosty/rd-app"
	"github.com/shooosty/rd-app/pkg/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]rd_app.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(userId int) (rd_app.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

//func (s *UserService) Update(userId, listId int, input rd_app.UpdateListInput) error {
//	if err := input.Validate(); err != nil {
//		return err
//	}
//
//	return s.repo.Update(userId, listId, input)
//}
