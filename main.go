package main

import (
	"GoServer/internal/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// get env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &app.Application{}

	app.Run()
}
