package models

import (
	"fmt"
	"log"
	"os"

	// production
	// _ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	databaseName := os.Getenv("DB_NAME")
	// production
	// dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, user, databaseName, password)
	
	// development
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 sslmode=disable password=%s", dbHost, user, databaseName, password)

	// production
	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DriverName: "cloudsqlpostgres",
	// 	DSN:        dsn,
	// }))

	// development
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Medicine{})
}
