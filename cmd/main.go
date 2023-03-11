package main

import (
	"github.com/fazilnbr/learn-docker/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fazil muhammed kp!")
	})

	app.Listen(":3000")
}
