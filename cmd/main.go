package main

import (
	"GoServer/internal/env"
	"GoServer/internal/store"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}
	// initialise configuration
	config := config{
		port: env.GetString("PORT"),
		env:  env.GetString("ENV"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS"),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS"),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME"),
		},
	}
	application := &application{
		config: config,
		store:  store.NewStore(),
	}

	application.Run()
}
