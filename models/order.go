package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Order struct {
	gorm.Model
	ID             string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updateAt"`
	UserId         string         `json:"userId"`
	Type           string         `json:"type"`
	Address        string         `json:"address"`
	Description    string         `json:"description"`
	Status         string         `json:"status"`
	Owner          string         `json:"owner"`
	DesignerId     string         `json:"designerId"`
	PeopleIds      pq.StringArray `gorm:"type:string[]" json:"people"`
	PhotographerId string         `json:"photographerId"`
	Contract       string         `json:"contract"`
	Datetime       string         `json:"datetime"`
}

func (order *Order) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdateOrderInput struct {
	Address *string `json:"address"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Address == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
