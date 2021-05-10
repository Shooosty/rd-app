package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
)

type PhotoPostgres struct {
	db *gorm.DB
}

func NewPhotoPostgres(db *gorm.DB) *PhotoPostgres {
	return &PhotoPostgres{db: db}
}

func (r *PhotoPostgres) Create(photo models.Photo) (string, error) {
	err := db.Table(photosTable).Create(&photo).Error

	return photo.ID, err
}
