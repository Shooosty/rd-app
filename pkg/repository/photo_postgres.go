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

func (r *PhotoPostgres) GetAll() ([]models.Photo, error) {
	var photos []models.Photo
	err := db.Table(photosTable).Find(&photos).Error

	return photos, err
}

func (r *PhotoPostgres) GetById(photoId string) (models.Photo, error) {
	var photo models.Photo
	err := db.Table(photosTable).Where("id = ?", photoId).Find(&photo).Error

	return photo, err
}

func (r *PhotoPostgres) Create(photo models.Photo) (string, error) {
	err := db.Table(photosTable).Create(&photo).Error

	return photo.ID, err
}

func (r *PhotoPostgres) Delete(photoId string) error {
	photo := make([]*Orders, 0)
	err := db.Table(photosTable).Where("id = ?", photoId).Delete(&photo).Error

	return err
}
