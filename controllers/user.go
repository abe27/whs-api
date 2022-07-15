package controllers

import (
	"fmt"

	"github.com/abe27/bugtracker/api/models"
	"github.com/abe27/bugtracker/api/services"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	db := services.DBConn
	var r models.Response
	user := new(models.User)
	// Generate a unique ID
	user.ID = services.Nanoid()
	user.UserName = c.FormValue("username")
	user.Email = c.FormValue("email")
	// Get Password from data form
	passwd := user.Password
	hash, _ := services.HashPassword(passwd)
	user.Password = hash

	err := db.Create(&user).Error
	if err != nil {
		r.ID = services.Nanoid()
		r.Success = false
		r.Message = "An error occurred while saving the data.Please contact your system administrator to fix this error."
		r.Data = err
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	r.ID = services.Nanoid()
	r.Success = true
	r.Message = "Registration information saved."
	r.Data = user
	return c.Status(fiber.StatusCreated).JSON(r)
}

func Login(c *fiber.Ctx) error {
	db := services.DBConn
	var r models.Response
	u := models.User{}
	u.SetFormData(c)
	err := db.Where("username", u.UserName).First(&u).Error
	if err != nil {
		r.ID = services.Nanoid()
		r.Message = "Not Found UserName!"
		r.Data = err
		return c.Status(fiber.StatusNotFound).JSON(r)
	}

	passwd := c.FormValue("password")
	fmt.Println(passwd + " != " + u.Password)
	fmt.Println(len(passwd))
	/// check password HashPassword
	isMatch := services.CheckPasswordHash(passwd, u.Password)
	fmt.Println(isMatch)
	if !isMatch {
		r.ID = services.Nanoid()
		r.Success = false
		r.Message = "Please check your password."
		r.Data = nil
		return c.Status(fiber.StatusUnauthorized).JSON(r)
	}
	r.ID = services.Nanoid()
	r.Success = true
	r.Message = "Login"
	r.Data = u
	return c.Status(fiber.StatusOK).JSON(r)
}
