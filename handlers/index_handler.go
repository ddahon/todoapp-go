package handlers

import (
	"database/sql"

	"github.com/ddahon/go-todoapp/views"
	"github.com/ddahon/go-todoapp/views/home"
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx, db *sql.DB) error {
	return views.Render(c, home.Home(getData(db)))
}
