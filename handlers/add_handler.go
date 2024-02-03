package handlers

import (
	"database/sql"

	"github.com/ddahon/go-todoapp/models"
	"github.com/ddahon/go-todoapp/views"
	"github.com/ddahon/go-todoapp/views/components"
	"github.com/gofiber/fiber/v2"
)

func AddHandler(c *fiber.Ctx, db *sql.DB) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	_, err := db.Exec("INSERT INTO todos VALUES ($1)", todo.Description)
	if err != nil {
		return err
	}
	return views.Render(c, components.List(getData(db)))
}
