package emenem_domain

import (
	"gorm.io/gorm"
	"time"
)

type MealCategory struct {
	ID uint `gorm:"primarykey"`

	Name string
	Tag  string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
