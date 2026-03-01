package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name           string `gorm:"unique"`
	Slug           string
	NormalizedName string `gorm:"index"`
}