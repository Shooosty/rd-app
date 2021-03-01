package service

import (
	"github.com/shooosty/rd-app/models"
	"github.com/shooosty/rd-app/pkg/repository"
)

type OrderService struct {
	repo repository.Orders
}

func NewOrderService(repo repository.Orders) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	return s.repo.GetAll()
}

func (s *OrderService) GetAllForUser(userId int) ([]models.Order, error) {
	return s.repo.GetAllForUser(userId)
}

func (s *OrderService) Create(order models.Order) (int, error) {
	return s.repo.Create(order)
}

func (s *OrderService) GetById(orderId int) (models.Order, error) {
	return s.repo.GetById(orderId)
}

func (s *OrderService) Delete(orderId int) error {
	return s.repo.Delete(orderId)
}

func (s *OrderService) Update(orderId int, input models.UpdateOrderInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(orderId, input)
}
