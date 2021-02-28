package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	rd_app "github.com/shooosty/rd-app"
	"github.com/sirupsen/logrus"
	"strings"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAll() ([]rd_app.User, error) {
	var users []rd_app.User

	query := fmt.Sprintf("SELECT id, name, role, phone, email FROM %s", usersTable)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserPostgres) GetById(userId int) (rd_app.User, error) {
	var user rd_app.User

	query := fmt.Sprintf("SELECT id, name, role, phone, email FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *UserPostgres) Delete(userId int) error {
	query := fmt.Sprintf("DELETE id, name, role, phone, email FROM %s WHERE id = $1", usersTable)
	_, err := r.db.Exec(query, userId)

	return err
}

func (r *UserPostgres) Update(userId int, input rd_app.UpdateUserInput) error {
	setValues := make([]string, 0)

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$1"))
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$2"))
	}

	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$3"))
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?",
		usersTable, setQuery)

	logrus.Debugf("updateQuery: %s", query)

	_, err := r.db.Exec(query, userId)
	return err
}
