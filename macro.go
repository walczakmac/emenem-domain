package emenem_domain

import (
	"gorm.io/gorm"
	"time"
)

type Macro struct {
	ID     uint `gorm:"primarykey"`
	MealID uint

	Proteins      float64
	Fats          float64
	Carbohydrates float64
	Fiber         float64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
