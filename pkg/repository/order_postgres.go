package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
)

type OrderPostgres struct {
	db *gorm.DB
}

func NewOrderPostgres(db *gorm.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := db.Table(ordersTable).Find(&orders).Error
	return orders, err
}

func (r *OrderPostgres) Create(order models.Order) (int, error) {
	err := db.Table(ordersTable).Create(&order).Error
	return order.Id, err
}

func (r *OrderPostgres) GetAllForUser(userId int) ([]models.Order, error) {
	var orders []models.Order
	err := db.Table(ordersTable).Where("user_id = ?", userId).Find(&orders).Error
	return orders, err
}

func (r *OrderPostgres) GetById(orderId int) (models.Order, error) {
	var order models.Order
	err := db.Table(ordersTable).Where("id = ?", orderId).Find(&order).Error
	return order, err
}

func (r *OrderPostgres) Delete(orderId int) error {
	orders := make([]*Orders, 0)
	err := db.Table(ordersTable).Where("id = ?", orderId).Delete(&orders).Error
	return err
}

func (r *OrderPostgres) Update(orderId int, input models.UpdateOrderInput) error {
	err := db.Table(ordersTable).Where("id = ?", orderId).Update(&input).Error
	return err
}
