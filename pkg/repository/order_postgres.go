package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shooosty/rd-app/models"
	"github.com/sirupsen/logrus"
	"strings"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) GetAll() ([]models.Order, error) {
	var orders []models.Order

	query := fmt.Sprintf("SELECT id, name FROM %s", ordersTable)
	err := r.db.Select(&orders, query)

	return orders, err
}

func (r *OrderPostgres) Create(order models.Order) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var orderId int
	createOrderQuery := fmt.Sprintf("INSERT INTO %s (name, user_id) values ($1, $2) RETURNING id", ordersTable)

	row := tx.QueryRow(createOrderQuery, order.Name, order.UserId)
	err = row.Scan(&orderId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return orderId, tx.Commit()
}

func (r *OrderPostgres) GetAllForUser(userId int) ([]models.Order, error) {
	var orders []models.Order

	query := fmt.Sprintf("SELECT id, name FROM %s WHERE user_id = $1", ordersTable)
	err := r.db.Select(&orders, query, userId)

	return orders, err
}

func (r *OrderPostgres) GetById(orderId int) (models.Order, error) {
	var order models.Order

	query := fmt.Sprintf("SELECT id, name, role, phone, email FROM %s WHERE id = $1", ordersTable)
	err := r.db.Get(&order, query, orderId)

	return order, err
}

func (r *OrderPostgres) Delete(orderId int) error {
	query := fmt.Sprintf("DELETE id, name, role, phone, email FROM %s WHERE id = $1", ordersTable)
	_, err := r.db.Exec(query, orderId)

	return err
}

func (r *OrderPostgres) Update(orderId int, input models.UpdateOrderInput) error {
	setValues := make([]string, 0)

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$1"))
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = %d",
		ordersTable, setQuery)

	logrus.Debugf("updateQuery: %s", query)

	_, err := r.db.Exec(query, orderId)
	return err
}
