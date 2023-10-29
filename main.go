package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func WelcomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func main() {
	app := fiber.New()
	rtr := app.Group("/api/v1")

	rtr.Get("/", WelcomeHandler)

	log.Fatal(app.Listen("0.0.0.0:8080"))
}
