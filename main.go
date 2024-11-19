package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/narumi-ma/textAnalysis/database"
	"github.com/narumi-ma/textAnalysis/router"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome")
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	// welcome
	app.Get("/api", welcome)

	// kana2kanji
	app.Post("/api/kana2kanji", router.Kana2kanji)

	// furigana
	app.Post("/api/furigana", router.Furigana)
	app.Get("/api/furigana/history", router.GetQuery)

	log.Fatal(app.Listen(":3000"))
}
