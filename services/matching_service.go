package services

import (
	"github.com/dhavisiregar/masak-apa/database"
	"github.com/dhavisiregar/masak-apa/models"
)

type MatchParams struct {
	Ingredients []string
	MinMatch    float64
	Exact       bool
}

type MatchResult struct {
	RecipeID          uint     `json:"recipe_id"`
	Title             string   `json:"title"`
	MatchPercentage   float64  `json:"match_percentage"`
	MissingIngredients []string `json:"missing_ingredients"`
}

func MatchRecipes(params MatchParams) ([]MatchResult, error) {

	// 1️⃣ Get ingredient IDs
	var inputIngredients []models.Ingredient
	err := database.DB.
		Where("normalized_name IN ?", params.Ingredients).
		Find(&inputIngredients).Error

	if err != nil {
		return nil, err
	}

	if len(inputIngredients) == 0 {
		return []MatchResult{}, nil
	}

	inputMap := make(map[uint]bool)
	var inputIDs []uint

	for _, ing := range inputIngredients {
		inputMap[ing.ID] = true
		inputIDs = append(inputIDs, ing.ID)
	}

	// 2️⃣ Load recipes with ingredients (biar bisa hitung missing)
	var recipes []models.Recipe
	err = database.DB.
		Preload("Ingredients.Ingredient").
		Find(&recipes).Error

	if err != nil {
		return nil, err
	}

	var results []MatchResult

	for _, recipe := range recipes {

		total := 0
		matched := 0
		var missing []string

		for _, ri := range recipe.Ingredients {
			if !ri.IsOptional {
				total++

				if inputMap[ri.IngredientID] {
					matched++
				} else {
					missing = append(missing, ri.Ingredient.Name)
				}
			}
		}

		if total == 0 {
			continue
		}

		score := (float64(matched) / float64(total)) * 100

		if params.Exact {
			if matched == total {
				results = append(results, MatchResult{
					RecipeID:          recipe.ID,
					Title:             recipe.Title,
					MatchPercentage:   100,
					MissingIngredients: []string{},
				})
			}
		} else {
			if score >= params.MinMatch {
				results = append(results, MatchResult{
					RecipeID:          recipe.ID,
					Title:             recipe.Title,
					MatchPercentage:   score,
					MissingIngredients: missing,
				})
			}
		}
	}

	return results, nil
}