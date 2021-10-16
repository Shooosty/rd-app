package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Order struct {
	ID                      string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Number                  string         `json:"number"`
	Sections                pq.StringArray `sql:"type: text[]" json:"sections"`
	LayoutClientDescription pq.StringArray `sql:"type: text[]" json:"layoutClientDescription"`
	DesignerDescription     pq.StringArray `sql:"type: text[]" json:"designerDescription"`
	DateTimes               pq.StringArray `sql:"type: text[]" json:"dateTimes"`
	Tz                      pq.StringArray `sql:"type: text[]" json:"tz"`
	StatusHistory           pq.StringArray `sql:"type: text[]" json:"statusHistory"`
	CreatedAt               time.Time      `json:"createdAt"`
	UpdatedAt               time.Time      `json:"updateAt"`
	UserId                  string         `json:"userId"`
	Address                 string         `json:"address"`
	Description             string         `json:"description"`
	PupilsMin               int            `json:"pupilsMin"`
	PupilsMax               int            `json:"pupilsMax"`
	TeachersMin             int            `json:"teachersMin"`
	TeachersMax             int            `json:"teachersMax"`
	InitialDescription      string         `json:"initialDescription"`
	DesignerId              string         `json:"designerId"`
	Design                  string         `json:"design"`
	PhotographerId          string         `json:"photographerId"`
	ManagerId               string         `json:"managerId"`
	Status                  string         `json:"status"`
	Contract                string         `json:"contract"`
	PhotoContract           string         `json:"photoContract"`
	AdditionalContract      string         `json:"additionalContract"`
	AttachmentContract      string         `json:"attachmentContract"`
	Layout                  string         `json:"layout"`
	LayoutCover             string         `json:"layoutCover"`
	YandexDisc              string         `json:"yandexDisc"`
	PreFormDate             string         `json:"preFormDate"`
	FormDate                string         `json:"formDate"`
	LayoutDate              string         `json:"layoutDate"`
	PreProdDate             string         `json:"PreProdDate"`
	LayoutFormDate          string         `json:"layoutFormDate"`
}

func (order *Order) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

type UpdateOrderInput struct {
	Status                  *string         `json:"status"`
	Address                 *string         `json:"address"`
	Number                  *string         `json:"number"`
	UserId                  *string         `json:"userId"`
	DesignerId              *string         `json:"designerId"`
	Design                  *string         `json:"design"`
	DateTimes               *pq.StringArray `json:"dateTimes"`
	LayoutClientDescription *pq.StringArray `json:"layoutClientDescription"`
	Sections                *pq.StringArray `json:"sections"`
	Tz                      *pq.StringArray `json:"tz"`
	StatusHistory           *pq.StringArray `json:"statusHistory"`
	DesignerDescription     *pq.StringArray `json:"designerDescription"`
	AdditionalContract      *string         `json:"additionalContract"`
	ManagerId               *string         `json:"managerId"`
	PhotographerId          *string         `json:"photographerId"`
	PupilsMin               *int            `json:"pupilsMin"`
	PupilsMax               *int            `json:"pupilsMax"`
	TeachersMin             *int            `json:"teachersMin"`
	TeachersMax             *int            `json:"teachersMax"`
	LayoutCover             *string         `json:"layoutCover"`
	YandexDisc              *string         `json:"yandexDisc"`
	PreFormDate             *string         `json:"preFormDate"`
	PhotoContract           *string         `json:"photoContract"`
	FormDate                *string         `json:"formDate"`
	Contract                *string         `json:"contract"`
	AttachmentContract      *string         `json:"attachmentContract"`
	Layout                  *string         `json:"layout"`
	LayoutDate              *string         `json:"layoutDate"`
	PreProdDate             *string         `json:"PreProdDate"`
	LayoutFormDate          *string         `json:"layoutFormDate"`
	InitialDescription      *string         `json:"initialDescription"`
	Description             *string         `json:"description"`
}

func (i UpdateOrderInput) Validate() error {
	if i.Address == nil && i.Status == nil && i.AdditionalContract == nil && i.PhotoContract == nil &&
		i.UserId == nil && i.DesignerId == nil && i.LayoutClientDescription == nil &&
		i.ManagerId == nil && i.Layout == nil && i.AttachmentContract == nil && i.Sections == nil &&
		i.PhotographerId == nil && i.Number == nil && i.InitialDescription == nil && i.LayoutDate == nil && i.PreProdDate == nil &&
		i.DesignerDescription == nil && i.Design == nil && i.DateTimes == nil && i.FormDate == nil && i.StatusHistory == nil &&
		i.PreFormDate == nil && i.Description == nil && i.LayoutFormDate == nil && i.Tz == nil &&
		i.LayoutCover == nil && i.YandexDisc == nil && i.PupilsMin == nil && i.PupilsMax == nil &&
		i.TeachersMin == nil && i.TeachersMax == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
