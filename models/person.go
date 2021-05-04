package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Person struct {
	gorm.Model
	ID          string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updateAt"`
	Url         string    `json:"url"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	MiddleName  string    `json:"middleName"`
	Description string    `json:"description"`
}

func (people *Person) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdatePersonInput struct {
	Name *string `json:"name"`
}

func (i UpdatePersonInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
