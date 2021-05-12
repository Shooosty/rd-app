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

func (r *OrderPostgres) Create(order models.Order) (string, error) {
	err := db.Table(ordersTable).Create(&order).Error
	if err != nil {
		var userEmail string
		_ = db.Table(usersTable).Select("email").Where("id = ?", order.UserId).Find(&userEmail)
		SendNewOrderCreated(userEmail)

		var photographerEmail string
		_ = db.Table(usersTable).Select("email").Where("id = ?", order.PhotographerId).Find(&photographerEmail)
		SendNewOrderCreated(photographerEmail)

		var designerEmail string
		_ = db.Table(usersTable).Select("email").Where("id = ?", order.DesignerId).Find(&designerEmail)
		SendNewOrderCreated(designerEmail)
	}

	return order.ID, err
}

func (r *OrderPostgres) GetAllForUser(userId string) ([]models.Order, error) {
	var orders []models.Order
	err := db.Table(ordersTable).Where("user_id = ?", userId).Find(&orders).Error

	return orders, err
}

func (r *OrderPostgres) GetAllForPhotographer(photographerId string) ([]models.Order, error) {
	var orders []models.Order
	err := db.Table(ordersTable).Where("photographer_id = ?", photographerId).Find(&orders).Error

	return orders, err
}

func (r *OrderPostgres) GetAllForDesigner(designerId string) ([]models.Order, error) {
	var orders []models.Order
	err := db.Table(ordersTable).Where("designer_id = ?", designerId).Find(&orders).Error

	return orders, err
}

func (r *OrderPostgres) GetById(orderId string) (models.Order, error) {
	var order models.Order
	err := db.Table(ordersTable).Where("id = ?", orderId).Find(&order).Error

	return order, err
}

func (r *OrderPostgres) Delete(orderId string) error {
	orders := make([]*Orders, 0)
	err := db.Table(ordersTable).Where("id = ?", orderId).Delete(&orders).Error

	return err
}

func (r *OrderPostgres) Update(orderId string, input models.UpdateOrderInput) error {
	err := db.Table(ordersTable).Where("id = ?", orderId).Updates(&input).Error

	return err
}
