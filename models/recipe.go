package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title       string
	Slug        string `gorm:"unique"`
	Instructions string
	CookingTime int
	Difficulty  string
	Ingredients []RecipeIngredient
}