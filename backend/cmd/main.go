package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/quiz", getQuiz)

	r.Run(":8080")
}

func getQuiz(c *gin.Context) {
	quiz := map[string]interface{}{
		"title": "Pattern Sniper Quiz",
		"questions": []map[string]interface{}{
			{
				"text":         "Will Boeing shares go up or down?",
				"options":      []string{"Up", "Down"},
				"correctIndex": 0,
			},
			{
				"text":         "Buy or Sell?",
				"options":      []string{"Buy", "Sell"},
				"correctIndex": 0,
			},
		},
	}

	c.JSON(http.StatusOK, quiz)
}
