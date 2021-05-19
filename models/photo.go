package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Photo struct {
	ID       string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PersonID string `json:"personId"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Size     int64  `json:"surname"`
	Type     string `json:"type"`
}

func (person *Photo) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}