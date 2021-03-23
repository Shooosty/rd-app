package models

import (
	"errors"
	"time"
)

type Order struct {
	ID        string    `sql:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
	UserId    string    `json:"user_id" db:"user_id" binding:"required"`
	Name      string    `json:"name" binding:"required"`
}

//func (order *Order) BeforeCreate(scope *gorm.Scope) error {
//	_ = scope.SetColumn("ID", uuid.NewV4().String())
//	return nil
//}

type UpdateOrderInput struct {
	Name *string `json:"name"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
