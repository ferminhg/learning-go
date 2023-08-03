package main

import (
	"github.com/ferminhg/learning-go/cmd/api/bootstrap"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	host, port := loadSettings()
	if err := bootstrap.Run(host, port); err != nil {
		log.Fatal(err)
	}
}

const defaultPort = 8080
const defaultHost = "localhost"

func loadSettings() (string, uint) {
	var host = defaultHost
	var port = uint(defaultPort)

	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}

	if len(os.Getenv("API_PORT")) > 0 {
		portInt, _ := strconv.Atoi(os.Getenv("API_PORT"))
		port = uint(portInt)
	}

	if len(os.Getenv("API_HOST")) > 0 {
		host = os.Getenv("API_HOST")
	}

	return host, port
}
