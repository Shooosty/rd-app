package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Order struct {
	ID             string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updateAt"`
	UserId         string    `json:"userId"`
	Type           string    `json:"type"`
	Address        string    `json:"address"`
	Description    string    `json:"description"`
	Owner          string    `json:"owner"`
	DesignerId     string    `json:"designerId"`
	Status         string    `json:"status"`
	PhotographerId string    `json:"photographerId"`
	Contract       string    `json:"contract"`
	Datetime       string    `json:"datetime"`
}

func (order *Order) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdateOrderInput struct {
	Address        *string         `json:"address"`
	Status         *string         `json:"status"`
	Owner          *string         `json:"owner"`
	UserId         *string         `json:"userId"`
	DesignerId     *string         `json:"designerId"`
	PhotographerId *string         `json:"photographerId"`
	PeopleIds      *pq.StringArray `sql:"type:text[]" json:"people"`
	Contract       *string         `json:"contract"`
	Datetime       *string         `json:"datetime"`
	Description    *string         `json:"description"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Address == nil && i.Contract == nil && i.Status == nil &&
		i.Owner == nil && i.UserId == nil && i.DesignerId == nil &&
		i.PhotographerId == nil &&
		i.Description == nil && i.Datetime == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
