package domain

import (
	"time"
)

type User struct {
	Model
	Name          string     `gorm:"name"`
	Email         string     `gorm:"email"`
	Age           int        `gorm:"age"`
	PhoneNumber   string     `gorm:"phone_number"`
	PassWord      string     `gorm:"pass_word"`
	Role          int        `gorm:"role"`
	Gender        bool       `gorm:"gender"`
	Avatar        string     `gorm:"avatar"`
	RegisterDate  *time.Time `gorm:"register_date"`
	LastLoginDate *time.Time `gorm:"last_login_date"`
	Disable       bool       `gorm:"disable"`
}
