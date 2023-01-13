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
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   60,
		Prompt:      "tell me an interesting fact about capybaras, about their weight, or their size, or their habitat or their predators, or their reputation, or their food, in the style of a reporter, a pirate, a politician, a memer or a programmer or an archaelogist, or an alien, or an ancient egyptian, or an underwater mermaid, or a crocodile (or anybody else)",
		Temperature: 1,
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
