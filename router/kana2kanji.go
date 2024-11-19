package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/narumi-ma/textAnalysis/model"
)

func Kana2kanji(c *fiber.Ctx) error {
	// Parse the input query
	type RequestBody struct {
		Query string `json:"query"`
	}
	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call the Post function
	response, err := Post(body.Query)
	if err != nil {
		log.Printf("Error making POST request: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process the query",
		})
	}

	// Return the response
	return c.SendString(response)
}

func Post(query string) (string, error) {
	// Create the payload

	payload := model.Request_K{
		ID:      "1234-1",
		JSONRPC: "2.0",
		Method:  "jlp.jimservice.conversion",
		Params: model.Params_K{
			Q:          query,
			Format:     "hiragana",
			Mode:       "kanakanji",
			Option:     []string{"hiragana", "katakana", "alphanumeric", "half_katakana", "half_alphanumeric"},
			Dictionary: []string{"base", "name", "place", "zip", "symbol"},
			Results:    999,
		},
	}

	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", URL_K, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("Yahoo AppID: %s", APPID))

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body := new(bytes.Buffer)
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return body.String(), nil
}
