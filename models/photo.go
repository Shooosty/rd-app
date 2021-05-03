package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Photo struct {
	gorm.Model
	ID         string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updateAt"`
	Url        string    `json:"url"`
	Name       string    `json:"name"`
	SureName   string    `json:"sureName"`
	MiddleName string    `json:"middleName"`
}

func (photo *Photo) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdatePhotoInput struct {
	Url *string `json:"url"`
}

func (i UpdatePhotoInput) Validate() error {
	if i.Url == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
