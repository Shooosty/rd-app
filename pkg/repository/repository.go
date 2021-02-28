package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/shooosty/rd-app"
)

type Authorization interface {
	CreateUser(user rd_app.User) (int, error)
	GetUser(username, password string) (rd_app.User, error)
}

type Users interface {
	GetAll() ([]rd_app.User, error)
	GetById(userId int) (rd_app.User, error)
	Delete(userId int) error
	Update(userId int, input rd_app.UpdateUserInput) error
}

type Repository struct {
	Authorization
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users:         NewUserPostgres(db),
	}
}
