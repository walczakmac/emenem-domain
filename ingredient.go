package emenem_domain

import (
	"gorm.io/gorm"
	"time"
)

type Ingredient struct {
	ID        uint `gorm:"primarykey"`
	MealID    uint `gorm:"default:null"`
	SnackID   uint `gorm:"default:null"`
	ProductID uint

	Weight float64
	Amount float64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Product Product
}
