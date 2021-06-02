package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Order struct {
	ID             string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Number         string         `json:"number"`
	Sections       pq.StringArray `json:"sections"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updateAt"`
	UserId         string         `json:"userId"`
	Address        string         `json:"address"`
	Description    string         `json:"description"`
	Owner          string         `json:"owner"`
	DesignerId     string         `json:"designerId"`
	PhotographerId string         `json:"photographerId"`
	ManagerId      string         `json:"managerId"`
	Status         string         `json:"status"`
	Contract       string         `json:"contract"`
	Datetime       string         `json:"datetime"`
}

func (order *Order) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdateOrderInput struct {
	Address        *string `json:"address"`
	Status         *string `json:"status"`
	Number         *string `json:"number"`
	Owner          *string `json:"owner"`
	UserId         *string `json:"userId"`
	DesignerId     *string `json:"designerId"`
	ManagerId      *string `json:"managerId"`
	PhotographerId *string `json:"photographerId"`
	Contract       *string `json:"contract"`
	Datetime       *string `json:"datetime"`
	Description    *string `json:"description"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Address == nil && i.Status == nil &&
		i.Owner == nil && i.UserId == nil && i.DesignerId == nil &&
		i.ManagerId == nil &&
		i.PhotographerId == nil && i.Number == nil &&
		i.Description == nil && i.Datetime == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
