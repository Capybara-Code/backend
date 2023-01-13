package Controllers

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func GetFact(c *gin.Context) {
	client := gogpt.NewClient(os.Getenv("OPENAPI_KEY"))
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 100,
		Prompt:    "Tell me a fact about capybara",
	}
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get fact",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"fact": resp.Choices[0].Text,
	})

}
