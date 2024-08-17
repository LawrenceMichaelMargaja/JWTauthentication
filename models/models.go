package models

import "gorm.io/gorm"

type User struct {
	ID       int
	Username string `gorm:"uniqueIndex;not null;size:50;" json:"username"`
	Email    string `gorm:"uniqueIndex;not null;size:255;" json:"email"`
	Password string `gorm:"not null;" json:"-"`
}

type Report struct {
	NumberUsers int
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
	SKU   string
}
