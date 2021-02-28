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
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
	}

	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argId))
		args = append(args, *input.Phone)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $1",
		setQuery, usersTable)
	args = append(args, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, userId)
	return err
}
