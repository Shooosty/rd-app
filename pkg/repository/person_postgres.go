package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
)

type PersonPostgres struct {
	db *gorm.DB
}

func NewPersonPostgres(db *gorm.DB) *PersonPostgres {
	return &PersonPostgres{db: db}
}

func (r *PersonPostgres) GetAll() ([]models.Person, error) {
	var persons []models.Person
	err := db.Table(personsTable).Find(&persons).Error

	return persons, err
}

func (r *PersonPostgres) GetAllByOrderId(orderId string) ([]models.Person, error) {
	var persons []models.Person
	err := db.Table(personsTable).Where("order_id = ?", orderId).Find(&persons).Error

	return persons, err
}

func (r *PersonPostgres) Create(person models.Person) (string, error) {
	err := db.Table(personsTable).Create(&person).Error

	return person.ID, err
}

func (r *PersonPostgres) Delete(personId string) error {
	person := make([]*Orders, 0)
	err := db.Table(personsTable).Where("id = ?", personId).Delete(&person).Error

	return err
}

func (r *PersonPostgres) Update(personId string, input models.UpdatePersonInput) error {
	err := db.Table(personsTable).Where("id = ?", personId).Updates(&input).Error

	return err
}
