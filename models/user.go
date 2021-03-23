package models

import (
	"errors"
	"time"
)

type User struct {
	ID           string    `sql:"type:uuid;primary_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"update_at"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Phone        string    `json:"phone" binding:"required"`
	Password     string    `json:"password" sql:"-" binding:"required"`
	PasswordHash string    `sql:"password_hash"`
	Role         string    `json:"role" binding:"required"`
}

//func (user *User) BeforeCreate(scope *gorm.Scope) error {
//	_ = scope.SetColumn("ID", uuid.NewV4().String())
//	return nil
//}

type UpdateUserInput struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == nil && i.Phone == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
