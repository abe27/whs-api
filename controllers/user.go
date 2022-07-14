package controllers

import (
	"github.com/abe27/bugtracker/api/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var r models.Response
	return c.Status(fiber.StatusOK).JSON(r)
}
