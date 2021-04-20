package emenem_domain

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID uint `gorm:"primarykey"`

	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
