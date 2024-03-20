package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"mygram/database"
	"mygram/model"
	"mygram/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var validate *validator.Validate

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validate = validator.New(validator.WithRequiredStructEnabled())

	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})
	if err != nil {
		panic(err)
	}

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	port, err := strconv.Atoi(serverPort)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%d", serverHost, port)

	r := gin.Default()
	routes.SetRoutes(r, db)
	r.Run(addr)
}
