package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
