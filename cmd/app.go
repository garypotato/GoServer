package main

import (
	"GoServer/internal/store"

	"github.com/gofiber/fiber/v2"
)

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	port string
	env  string
	db   dbConfig
}

type application struct {
	config config
	store  *store.Store
}

func (app *application) Run() {
	fiberApp := fiber.New()

	fiberApp.Get("/users", func(c *fiber.Ctx) error {
		user, err := app.store.Users.GetAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error fetching users")
		}
		return c.JSON(user)

	})

	fiberApp.Listen(app.config.port)
}
