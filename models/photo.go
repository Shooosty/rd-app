package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Photo struct {
	ID        string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PersonId  string `json:"personId"`
	OrderId   string `json:"orderId"`
	Name      string `json:"name"`
	NameS3    string `json:"nameS3"`
	Url       string `json:"url"`
	UrlResize string `json:"urlResize"`
	Size      int64  `json:"size"`
	Type      string `json:"type"`
}

func (person *Photo) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}
