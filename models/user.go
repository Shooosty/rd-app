package models

import (
	"errors"
	"time"
)

type User struct {
	Id        int    `json:"-" db:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
	CreatedAt time.Time
}

type UpdateUserInput struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == nil && i.Phone == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
