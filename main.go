package main

import (
	"database/sql"
	"log"

	"github.com/ddahon/go-todoapp/handlers"
	"github.com/ddahon/go-todoapp/models"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func getData(db *sql.DB) []models.Todo {
	var todos []models.Todo
	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		t := new(models.Todo)
		rows.Scan(&t.Description)
		todos = append(todos, *t)
	}

	return todos
}

func main() {
	connStr := "postgresql://postgres:password@localhost:5432/todoapp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	todos := getData(db)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		todos = getData(db)
		return handlers.IndexHandler(c, db, todos)
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		todos = getData(db)
		return handlers.AddHandler(c, db, todos)
	})

	log.Fatal(app.Listen(":3000"))
}
