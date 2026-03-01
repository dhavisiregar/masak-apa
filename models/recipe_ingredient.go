package models

import "gorm.io/gorm"

type RecipeIngredient struct {
	gorm.Model
	RecipeID     uint
	IngredientID uint
	IsOptional   bool
	Ingredient   Ingredient
}