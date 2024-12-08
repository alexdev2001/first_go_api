package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"unique"`
}

func main() {
	// load data from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	// get the values from the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// prepare the connection string
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s  password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	// connect to the postgres database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database", err)
	}

	// automatically migrate the database
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate the database")
	}

}
