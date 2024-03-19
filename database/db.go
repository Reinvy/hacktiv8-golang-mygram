package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DBNAME")

	dbPort := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(dbPort)
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, port, dbUser, dbPassword, dbName)
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
