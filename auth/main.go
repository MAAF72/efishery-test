package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MAAF72/efishery-test/adapters"
	"github.com/MAAF72/efishery-test/routers"
	"github.com/MAAF72/efishery-test/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := gin.Default()

	adapters := adapters.Init()
	services.Init(adapters)

	routers.RegisterRouters(app)

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	app.Run(fmt.Sprintf("%s:%s", host, port))
}
