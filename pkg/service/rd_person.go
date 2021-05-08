package service

import (
	"github.com/shooosty/rd-app/models"
	"github.com/shooosty/rd-app/pkg/repository"
)

type PersonService struct {
	repo repository.Persons
}

func NewPersonService(repo repository.Persons) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) GetAll() ([]models.Person, error) {
	return s.repo.GetAll()
}

func (s *PersonService) GetAllByOrderId(orderId string) ([]models.Person, error) {
	return s.repo.GetAllByOrderId(orderId)
}

func (s *PersonService) Create(person models.Person) (string, error) {
	return s.repo.Create(person)
}

func (s *PersonService) Delete(personId string) error {
	return s.repo.Delete(personId)
}

func (s *PersonService) Update(personId string, input models.UpdatePersonInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(personId, input)
}
