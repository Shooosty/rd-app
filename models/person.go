package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Person struct {
	ID          string          `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OrderId     string          `json:"orderId"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updateAt"`
	Description string          `json:"description"`
	Name        string          `json:"name"`
	Surname     string          `json:"surname"`
	Type        string          `json:"type"`
	Photos      *pq.StringArray `sql:"type:text[]" json:"photos"`
	MiddleName  string          `json:"middleName"`
}

func (person *Person) BeforeCreate(scope *gorm.Scope) error {
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
