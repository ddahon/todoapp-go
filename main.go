package main

import (
	"database/sql"
	"log"

	"github.com/ddahon/go-todoapp/handlers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://postgres:password@localhost:5432/todoapp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return handlers.IndexHandler(c, db)
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		return handlers.AddHandler(c, db)
	})

	app.Post("/delete", func(c *fiber.Ctx) error {
		return handlers.DeleteHandler(c, db)
	})

	log.Fatal(app.Listen(":3000"))
}
