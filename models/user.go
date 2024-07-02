package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"not null;unipue"`
	Password string `gorm:"not null"`
	Products []Product `gorm:"constraint:OnDelete:CASCADE"`
}