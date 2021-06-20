package service

import (
	"github.com/shooosty/rd-app/models"
	"github.com/shooosty/rd-app/pkg/repository"
)

type PhotoService struct {
	repo repository.Photos
}

func NewPhotoService(repo repository.Photos) *PhotoService {
	return &PhotoService{repo: repo}
}

func (s *PhotoService) GetAll() ([]models.Photo, error) {
	return s.repo.GetAll()
}

func (s *PhotoService) GetById(photoId string) (models.Photo, error) {
	return s.repo.GetById(photoId)
}

func (s *PhotoService) GetAllPhotosByPersonId(personId string) ([]models.Photo, error) {
	return s.repo.GetAllPhotosByPersonId(personId)
}

func (s *PhotoService) GetAllPhotosByOrderId(orderId string) ([]models.Photo, error) {
	return s.repo.GetAllPhotosByOrderId(orderId)
}

func (s *PhotoService) Create(photo models.Photo) (string, error) {
	return s.repo.Create(photo)
}

func (s *PhotoService) Delete(photoId string) error {
	return s.repo.Delete(photoId)
}
