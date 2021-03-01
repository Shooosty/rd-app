package models

import "errors"

type Order struct {
	Id     int    `json:"-" db:"id"`
	UserId int    `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
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
