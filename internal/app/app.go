package app

import "github.com/gofiber/fiber/v2"

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

type Application struct {
	config config
}

func (app *Application) Run() {
	fiberApp := fiber.New()

	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fiberApp.Listen(":8080")
}
