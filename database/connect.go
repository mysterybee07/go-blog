package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mysterybee07/blogbackend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ")
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to database")
	} else {
		log.Println("Connected to database")
	}
	DB = db
	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
}
