package handlers

import (
	"strconv"

	"github.com/dhavisiregar/masak-apa/database"
	"github.com/dhavisiregar/masak-apa/models"
	"github.com/gin-gonic/gin"
)

func GetIngredients(c *gin.Context) {
    var ingredients []models.Ingredient
    database.DB.Order("name asc").Find(&ingredients)
    c.JSON(200, ingredients)
}

type MatchRequest struct {
	Ingredients    []string `json:"ingredients"`
	NormalizedName string   `json:"normalized_name"`
}

func MatchRecipes(c *gin.Context) {
	var req MatchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if len(req.Ingredients) == 0 {
		c.JSON(400, gin.H{"error": "ingredients required"})
		return
	}

	// Query Params
	minMatch := 0.0
	if v := c.Query("min_match"); v != "" {
		if parsed, err := strconv.ParseFloat(v, 64); err == nil {
			minMatch = parsed
		}
	}

	exact := c.Query("exact") == "true"

	page := 1
	limit := 10

	if v := c.Query("page"); v != "" {
		if p, err := strconv.Atoi(v); err == nil && p > 0 {
			page = p
		}
	}

	if v := c.Query("limit"); v != "" {
		if l, err := strconv.Atoi(v); err == nil && l > 0 {
			limit = l
		}
	}

	// Get Ingredient IDs
	var ingredients []models.Ingredient
	database.DB.Where("normalized_name IN ?", req.Ingredients).Find(&ingredients)

	if len(ingredients) == 0 {
		c.JSON(200, gin.H{"data": []gin.H{}, "page": page, "limit": limit, "total": 0})
		return
	}

	var inputIDs []uint
	for _, ing := range ingredients {
		inputIDs = append(inputIDs, ing.ID)
	}

	// Use raw SQL to avoid HAVING alias issue in MySQL
	type Result struct {
		ID           uint
		Title        string
		TotalCount   int
		MatchedCount int
	}

	var results []Result

	database.DB.Raw(`
		SELECT
			r.id,
			r.title,
			COUNT(DISTINCT CASE WHEN ri.is_optional = false THEN ri.id END) as total_count,
			COUNT(DISTINCT CASE WHEN ri.is_optional = false AND ri.ingredient_id IN ? THEN ri.id END) as matched_count
		FROM recipes r
		JOIN recipe_ingredients ri ON ri.recipe_id = r.id
		WHERE r.deleted_at IS NULL AND ri.deleted_at IS NULL
		GROUP BY r.id, r.title
		HAVING COUNT(DISTINCT CASE WHEN ri.is_optional = false AND ri.ingredient_id IN ? THEN ri.id END) > 0
		ORDER BY matched_count DESC
	`, inputIDs, inputIDs).Scan(&results)

	var response []gin.H

	for _, r := range results {
		if r.TotalCount == 0 {
			continue
		}

		score := float64(r.MatchedCount) / float64(r.TotalCount) * 100

		if exact && r.MatchedCount != r.TotalCount {
			continue
		}

		if !exact && score < minMatch {
			continue
		}

		// Get missing ingredients
		var missing []string
		database.DB.Raw(`
			SELECT i.name
			FROM recipe_ingredients ri
			JOIN ingredients i ON i.id = ri.ingredient_id
			WHERE ri.recipe_id = ?
			AND ri.is_optional = false
			AND ri.ingredient_id NOT IN ?
			AND ri.deleted_at IS NULL
		`, r.ID, inputIDs).Scan(&missing)

		if missing == nil {
			missing = []string{}
		}

		response = append(response, gin.H{
			"id":                  r.ID,
			"title":               r.Title,
			"match_percentage":    score,
			"missing_ingredients": missing,
		})
	}

	if response == nil {
		response = []gin.H{}
	}

	totalResults := len(response)
	start := (page - 1) * limit
	end := start + limit

	if start > totalResults {
		start = totalResults
	}
	if end > totalResults {
		end = totalResults
	}

	c.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
		"total": totalResults,
		"data":  response[start:end],
	})
}