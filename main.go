package main

import (
	"log"

	"github.com/ddahon/go-todoapp/views"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, views.Home())
	})

	log.Fatal(app.Listen(":3000"))
}
