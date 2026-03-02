package main

import (
	"log"
	"os"

	"github.com/dhavisiregar/masak-apa/database"
	"github.com/dhavisiregar/masak-apa/handlers"
	"github.com/dhavisiregar/masak-apa/models"
	"github.com/dhavisiregar/masak-apa/seed"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	database.Connect()

	database.DB.AutoMigrate(
		&models.Recipe{},
		&models.Ingredient{},
		&models.RecipeIngredient{},
	)

	seed.SeedData()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://masak-apa-fe.vercel.app"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Content-Type"},
	}))

	r.GET("/ingredients", handlers.GetIngredients)
	r.POST("/match-recipes", handlers.MatchRecipes)
	r.POST("/suggest", handlers.SuggestDishes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}