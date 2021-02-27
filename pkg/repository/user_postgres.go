package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	rd_app "github.com/shooosty/rd-app"
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

//func (r *UserPostgres) Update(userId, listId int, input rd_app.UpdateListInput) error {
//	setValues := make([]string, 0)
//	args := make([]interface{}, 0)
//	argId := 1
//
//	if input.Title != nil {
//		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
//		args = append(args, *input.Title)
//		argId++
//	}
//
//	if input.Description != nil {
//		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
//		args = append(args, *input.Description)
//		argId++
//	}
//
//	// title=$1
//	// description=$1
//	// title=$1, description=$2
//	setQuery := strings.Join(setValues, ", ")
//
//	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
//		todoListsTable, setQuery, usersListsTable, argId, argId+1)
//	args = append(args, listId, userId)
//
//	logrus.Debugf("updateQuery: %s", query)
//	logrus.Debugf("args: %s", args)
//
//	_, err := r.db.Exec(query, args...)
//	return err
//}
