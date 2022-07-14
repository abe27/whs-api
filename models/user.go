package models

import "time"

type User struct {
	ID         string    `gorm:"size:36" json:"id"`
	UserName   string    `gorm:"index:unique,size:10" json:"username"`
	Password   string    `gorm:"size:255"`
	IsVerified bool      `json:"is_verified" default:"false"`
	CreatedAt  time.Time `json:"created_at" default:"now"`
	UpdatedAt  time.Time `json:"updated_at" default:"now"`
}
