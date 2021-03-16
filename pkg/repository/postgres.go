package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/shooosty/rd-app/models"
	"os"
)

var db *gorm.DB

const (
	usersTable  = "users"
	ordersTable = "orders"
)

func init() {
	username, _ := os.LookupEnv("DB_USERNAME")
	host, _ := os.LookupEnv("DB_HOST")
	port, _ := os.LookupEnv("DB_PORT")
	dbname, _ := os.LookupEnv("DB_NAME")
	sslmode, _ := os.LookupEnv("DB_SSL_MODE")
	pass, _ := os.LookupEnv("DB_PASSWORD")

	dbConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, username, dbname, pass, sslmode)

	conn, err := gorm.Open("postgres", dbConfig)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	//Миграция базы данных
	db.Debug().AutoMigrate(&models.User{}, &models.Order{})
}
