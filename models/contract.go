package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Contract struct {
	ID   string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Size string `json:"surname"`
	Type string `json:"type"`
}

func (person *Contract) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}