package models

import (
	"errors"
	"time"
)

type Order struct {
	Id        int `json:"-" db:"id"`
	CreatedAt time.Time
	UserId    int    `json:"user_id" db:"user_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type UpdateOrderInput struct {
	Name *string `json:"name"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
