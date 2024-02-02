package handlers

import (
	"database/sql"

	"github.com/ddahon/go-todoapp/models"
	"github.com/ddahon/go-todoapp/views"
	"github.com/ddahon/go-todoapp/views/components"
	"github.com/gofiber/fiber/v2"
)

func AddHandler(c *fiber.Ctx, db *sql.DB, todos []models.Todo) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	db.Exec("INSERT INTO todos VALUES ($1)", todo.Description)
	todos = append(todos, *todo)
	return views.Render(c, components.List(todos))
}
