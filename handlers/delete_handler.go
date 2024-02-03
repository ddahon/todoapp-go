package handlers

import (
	"database/sql"
	"log"

	"github.com/ddahon/go-todoapp/views"
	"github.com/ddahon/go-todoapp/views/components"
	"github.com/gofiber/fiber/v2"
)

func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	payload := struct {
		Id uint `json:"id"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}
	log.Println(payload.Id)
	_, err := db.Exec("DELETE FROM todos WHERE id=$1", payload.Id)
	if err != nil {
		return err
	}
	return views.Render(c, components.List(getData(db)))

}
