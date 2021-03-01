package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shooosty/rd-app/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, phone, role, password_hash) values ($1, $2, $3, $4, $5) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Phone, user.Role, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, email, phone, name, role FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
