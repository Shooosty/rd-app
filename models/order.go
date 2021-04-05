package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Order struct {
	ID          string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
	UserId      string    `json:"user_id"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Owner       string    `json:"owner"`
	Contract    string    `json:"contract"`
	Datetime    string    `json:"datetime"`
}

func (order *Order) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
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
