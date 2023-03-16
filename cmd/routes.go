package main

import (
	"github.com/fazilnbr/learn-docker/handlers"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func setupRoutes(app *fiber.App) {
	// Swagger docs

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}
