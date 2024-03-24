package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"mygram/config/database"
	"mygram/domain/entity"
	"mygram/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.User{}, &entity.Photo{}, &entity.Comment{}, &entity.SocialMedia{})
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
