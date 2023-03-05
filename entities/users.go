package entities

import (
	guuid "github.com/google/uuid"
)

type User struct {
	UserId    guuid.UUID `gorm:"primaryKey" json:"userId"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Password  string     `json:"-"`
	CreatedAt int64      `gorm:"autoCreateTime" json:"-" `
	UpdatedAt int64      `gorm:"autoUpdateTime:milli" json:"-"`
}
