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
	if err == nil {
		user, _ := r.GetUserById(order.UserId)
		SendNewOrderCreatedToClient(user.Email)

		photographer, _ := r.GetUserById(order.PhotographerId)
		SendNewOrderCreatedToEmployee(photographer.Email)
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

	order, _ := r.GetById(orderId)
	user, _ := r.GetUserById(order.UserId)
	SendUpdateOrderToClient(user.Email, localizeStatuses(order.Status))
	return err
}

func (r *OrderPostgres) GetUserById(userId string) (models.User, error) {
	var user models.User
	err := db.Table(usersTable).Where("id = ?", userId).Find(&user).Error

	return user, err
}

func localizeStatuses(status string) (localizeStatus string) {
	switch status {
	case "photoDateApproved":
		return "Cъемка назначена"
	case "needAnotherPhotoDate":
		return "Назначить доп съемку"
	case "anotherPhotoDateApproved":
		return "Доп съемка назначена"
	case "photoDateChecked":
		return "Съемка проведена"
	case "onTheFormation":
		return "На формировании"
	case "onDesign":
		return "В отделе дизайна"
	case "onTheClientApprove":
		return "На сверке у клиента"
	case "onEdits":
		return "Заказ на правках"
	case "onProduction":
		return "В производствe"
	case "done":
		return "Заказ готов"
	case "closed":
		return "Заказ закрыт"
	}
	return
}
