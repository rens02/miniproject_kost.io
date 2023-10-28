package controller

import (
	"app/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"io/ioutil"
	"net/http"
	"os"
)

func GetRecommendation(c echo.Context) error {
	//ganti jadi AI buat admin JANGAN LUPA ***********
	OpenAI_Key := os.Getenv("API_OPENAI")

	var reqData models.RoomRequest

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if err := json.Unmarshal(body, &reqData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	client := openai.NewClient(OpenAI_Key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Anda merupakan asisten yang dapat membantu untuk memberikan rekomendasi kamar.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Rekomendasi kamar untuk %s orang dengan fasilitas %s .", reqData.Person, reqData.Facility),
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}
	recommendation := resp.Choices[0].Message.Content

	response := models.AIResponse{
		Status: "success",
		Data:   recommendation,
	}

	return c.JSON(http.StatusOK, response)
}
