package main

import (
	"log"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
    Message   string `json:"message"`
    Timestamp int64  `json:"timestamp"`
    Version   string `json:"version"`
}

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	
	version := os.Getenv("APP_VERSION")

	if version == "" {
		version = "dev"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		response := Response{
			Message:   "My name is Sam",
			Timestamp: time.Now().UnixMilli(),
			Version: version,
		}
		return c.JSON(response)
	})

	log.Fatal(app.Listen(":" + port))
}
