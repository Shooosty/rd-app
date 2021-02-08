package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/shooosty/rd-app"
)

type Authorization interface {
	CreateUser(user rd_app.User) (int, error)
	GetUser(username, password string) (rd_app.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
