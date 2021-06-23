package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID           string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updateAt"`
	Name         string    `json:"name" binding:"required"`
	Surname      string    `json:"surname" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Phone        string    `json:"phone" binding:"required"`
	Password     string    `json:"password" sql:"-"`
	PasswordHash string    `sql:"passwordHash"`
	Role         string    `json:"role" binding:"required"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdateUserInput struct {
	Name    *string `json:"name"`
	Surname *string `json:"surname"`
	Phone   *string `json:"phone"`
}

type ChangePasswordInput struct {
	Password     string `json:"password" sql:"-"`
	PasswordHash string `json:"newPassword" sql:"passwordHash"`
}

type ResetPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == nil && i.Surname == nil && i.Phone == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
