package emenem_domain

import (
	"gorm.io/gorm"
	"time"
)

type Snack struct {
	ID     uint `gorm:"primarykey"`
	MealID uint

	Ingredient Ingredient

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
