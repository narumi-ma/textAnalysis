package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/narumi-ma/textAnalysis/database"
	"github.com/narumi-ma/textAnalysis/model"
)

// Request parameters (POST)
type Request struct {
	ID      int    `json:"id"`
	JSONRPC string `json:"jsonrpc"` // The value should be "2.0"
	Method  string `json:"method"`  // The value should be "jlp.furiganaservice.furigana"
	Params  Params `json:"params"`
}

type Params struct {
	Q     string `json:"q"`
	Grade int    `json:"grade"`
}

// To do:
// DeleteHistory()
// FindHistory() with id

func GetQuery(c *fiber.Ctx) error {
	query := []model.Furigana{}
	database.Database.Db.Find(&query)

	return c.Status(200).JSON(query)
}

func Furigana(c *fiber.Ctx) error {
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

	var furigana model.Furigana
	furigana.Q = body.Query
	if err := c.BodyParser(&furigana); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&furigana)
	text := Request{
		ID:      furigana.ID,
		JSONRPC: "2.0",
		Method:  "jlp.furiganaservice.furigana",
		Params: Params{
			Q:     body.Query,
			Grade: 8,
		},
	}

	// Call the Post function
	response, err := CreateRequest(text)
	if err != nil {
		log.Printf("Error making POST request: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to process the query",
		})
	}

	// Return the response
	return c.SendString(response)
}

func CreateRequest(text Request) (string, error) {
	// Convert text to JSON
	jsonData, err := json.Marshal(text)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", URL_F, bytes.NewBuffer(jsonData))
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
