package domain

import "time"

type Model struct {
	ID           uint       `gorm:"id"`
	CreatedAt    *time.Time `gorm:"created_at"`
	CreateUserId uint       `gorm:"create_user_id"`
	UpdatedAt    *time.Time `gorm:"updated_at"`
	UpdateUserId uint       `gorm:"update_user_id"`
	Deleted      bool       `gorm:"deleted"`
}
