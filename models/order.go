package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Order struct {
	ID                  string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Number              string         `json:"number"`
	Sections            pq.StringArray `sql:"type: text[]" json:"sections"`
	CreatedAt           time.Time      `json:"createdAt"`
	UpdatedAt           time.Time      `json:"updateAt"`
	UserId              string         `json:"userId"`
	Address             string         `json:"address"`
	Description         string         `json:"description"`
	DesignerDescription string         `json:"designerDescription"`
	InitialDescription  string         `json:"initialDescription"`
	Owner               string         `json:"owner"`
	DesignerId          string         `json:"designerId"`
	PhotographerId      string         `json:"photographerId"`
	ManagerId           string         `json:"managerId"`
	Status              string         `json:"status"`
	Contract            string         `json:"contract"`
	AttachmentContract  string         `json:"attachmentContract"`
	Layout              string         `json:"layout"`
	Datetime            string         `json:"datetime"`
}

func (order *Order) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdateOrderInput struct {
	Status              *string `json:"status"`
	Address             *string `json:"address"`
	Number              *string `json:"number"`
	Owner               *string `json:"owner"`
	UserId              *string `json:"userId"`
	DesignerId          *string `json:"designerId"`
	ManagerId           *string `json:"managerId"`
	PhotographerId      *string `json:"photographerId"`
	Contract            *string `json:"contract"`
	AttachmentContract  *string `json:"attachmentContract"`
	Layout              *string `json:"layout"`
	Datetime            *string `json:"datetime"`
	InitialDescription  *string `json:"initialDescription"`
	DesignerDescription *string `json:"designerDescription"`
	Description         *string `json:"description"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Address == nil && i.Status == nil &&
		i.Owner == nil && i.UserId == nil && i.DesignerId == nil &&
		i.ManagerId == nil && i.Layout == nil && i.AttachmentContract == nil &&
		i.PhotographerId == nil && i.Number == nil && i.InitialDescription == nil &&
		i.DesignerDescription == nil &&
		i.Description == nil && i.Datetime == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
