package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type SuggestRequest struct {
	Ingredients []string `json:"ingredients"`
}

type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GroqRequest struct {
	Model    string        `json:"model"`
	Messages []GroqMessage `json:"messages"`
	MaxTokens int          `json:"max_tokens"`
}

type GroqChoice struct {
	Message GroqMessage `json:"message"`
}

type GroqResponse struct {
	Choices []GroqChoice `json:"choices"`
}

func SuggestDishes(c *gin.Context) {
	var req SuggestRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Ingredients) == 0 {
		c.JSON(400, gin.H{"error": "ingredients required"})
		return
	}

	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		c.JSON(500, gin.H{"error": "GROQ_API_KEY not set"})
		return
	}

	ingredientList := strings.Join(req.Ingredients, ", ")

	payload := GroqRequest{
		Model:     "llama-3.1-8b-instant",
		MaxTokens: 100,
		Messages: []GroqMessage{
			{
				Role:    "system",
				Content: "Kamu adalah asisten masak Indonesia. Berikan 2-3 saran nama masakan singkat berdasarkan bahan yang disebutkan. Format: hanya nama masakan dipisahkan koma, tanpa penjelasan tambahan. Contoh: Nasi Goreng, Ayam Goreng, Telur Dadar",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Bahan yang saya punya: %s. Kira-kira bisa masak apa?", ingredientList),
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to build request"})
		return
	}

	httpReq, err := http.NewRequest("POST", "https://api.groq.com/openai/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create request"})
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to call Groq API"})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to read response"})
		return
	}

	var groqResp GroqResponse
	if err := json.Unmarshal(respBody, &groqResp); err != nil || len(groqResp.Choices) == 0 {
		c.JSON(500, gin.H{"error": "failed to parse response", "raw": string(respBody)})
		return
	}

	c.JSON(200, gin.H{
		"suggestion": groqResp.Choices[0].Message.Content,
	})
}