package controllers

import (
	"github.com/abe27/bugtracker/api/models"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	var r models.Response
	r.Success = true
	r.Message = "Hello, world!"
	return c.Status(fiber.StatusOK).JSON(r)
}

func ErrorHandler(c *fiber.Ctx) error {
	var r models.Response
	r.Success = false
	r.Message = "ขอ อภัย ในความไม่สะดวก\nขณะนี้ระบบกำลังอยู่ในช่วงการพัฒนา"
	return c.Status(fiber.StatusInternalServerError).JSON(r)
}
