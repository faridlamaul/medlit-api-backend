package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	DbHost := os.Getenv("DB_HOST")
	Dbdriver := os.Getenv("DB_DRIVER")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	
	if err != nil {
		fmt.Println("Cannot connect to database", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("Connected to database", Dbdriver)
	}

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Medicine{})
}