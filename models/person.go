package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Person struct {
	ID          string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrderId     string    `json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updateAt"`
	Description string    `json:"description"`
	Name        string    `json:"name"`
	WillBuy     string    `json:"willBuy"`
	Role        string    `json:"role"`
	Surname     string    `json:"surname"`
	Type        string    `json:"type"`
	MiddleName  string    `json:"middleName"`
	PhotosCount int       `json:"photosCount"`
}

func (person *Person) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdatePersonInput struct {
	Name        *string `json:"name"`
	Surname     *string `json:"surname"`
	MiddleName  *string `json:"middleName"`
	WillBuy     *string `json:"willBuy"`
	Role        *string `json:"role"`
	Type        *string `json:"type"`
	Description *string `json:"description"`
	PhotosCount *int    `json:"photosCount"`
}

func (i UpdatePersonInput) Validate() error {
	if i.Name == nil && i.Description == nil &&
		i.Role == nil && i.WillBuy == nil &&
		i.MiddleName == nil && i.Surname == nil &&
		i.Type == nil && i.PhotosCount == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
