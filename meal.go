package emenem_domain

import (
	"gorm.io/gorm"
	"time"
)

var dayNamesMap = map[string]int{
	"poniedzialek": 1,
	"wtorek":       2,
	"środa":        3,
	"czwartek":     4,
	"piatek":       5,
	"sobota":       6,
	"niedziela":    7,
}

type Meal struct {
	ID         uint `gorm:"primarykey"`
	CategoryID uint

	Name        string
	Tag         string `gorm:"index"`
	Kcal        float64
	Description string
	Day         int
	Person      string
	DailyKcal   int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Category    MealCategory
	Ingredients []Ingredient
	Macro       Macro
	Snacks      []Snack
}
