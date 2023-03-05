package entities

import (
	guuid "github.com/google/uuid"
)

type User struct {
	ID        guuid.UUID `gorm:"primaryKey" json:"-"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"-" `
	UpdatedAt int64      `gorm:"autoUpdateTime:milli" json:"-"`
}
