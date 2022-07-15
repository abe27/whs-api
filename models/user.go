package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID         string    `gorm:"size:36" json:"id"`
	UserName   string    `gorm:"column:username;unique;not null;size:10" json:"username"`
	Password   string    `gorm:"not null;size:60" json:"-"`
	Email      string    `gorm:"default:null;size:25" json:"email"`
	IsVerified bool      `json:"is_verified" default:"false"`
	CreatedAt  time.Time `json:"created_at" default:"now"`
	UpdatedAt  time.Time `json:"updated_at" default:"now"`
}

func (u *User) SetFormData(c *fiber.Ctx) {
	u.UserName = c.FormValue("username")
	u.Password = c.FormValue("password")
	u.Email = c.FormValue("email")
}
