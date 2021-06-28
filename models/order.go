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
	DesignerId          string         `json:"designerId"`
	Design              string         `json:"design"`
	PhotographerId      string         `json:"photographerId"`
	ManagerId           string         `json:"managerId"`
	Status              string         `json:"status"`
	Contract            string         `json:"contract"`
	AttachmentContract  string         `json:"attachmentContract"`
	Layout              string         `json:"layout"`
	PreFormDate         string         `json:"preFormDate"`
	FormDate            string         `json:"formDate"`
	LayoutFormDate      string         `json:"layoutFormDate"`
	DateTimes           pq.StringArray `sql:"type: text[]" json:"dateTimes"`
}

func (order *Order) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdateOrderInput struct {
	Status              *string         `json:"status"`
	Address             *string         `json:"address"`
	Number              *string         `json:"number"`
	UserId              *string         `json:"userId"`
	DesignerId          *string         `json:"designerId"`
	Design              *string         `json:"design"`
	DateTimes           *pq.StringArray `json:"dateTimes"`
	ManagerId           *string         `json:"managerId"`
	PhotographerId      *string         `json:"photographerId"`
	PreFormDate         *string         `json:"preFormDate"`
	FormDate            *string         `json:"formDate"`
	Contract            *string         `json:"contract"`
	AttachmentContract  *string         `json:"attachmentContract"`
	Layout              *string         `json:"layout"`
	LayoutFormDate      *string         `json:"layoutFormDate"`
	InitialDescription  *string         `json:"initialDescription"`
	DesignerDescription *string         `json:"designerDescription"`
	Description         *string         `json:"description"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Address == nil && i.Status == nil &&
		i.UserId == nil && i.DesignerId == nil &&
		i.ManagerId == nil && i.Layout == nil && i.AttachmentContract == nil &&
		i.PhotographerId == nil && i.Number == nil && i.InitialDescription == nil &&
		i.DesignerDescription == nil && i.Design == nil && i.DateTimes == nil && i.FormDate == nil &&
		i.PreFormDate == nil && i.Description == nil && i.LayoutFormDate == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
