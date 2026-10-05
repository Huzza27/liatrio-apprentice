package main

import 
(
	"time"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
} 

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
		response := Response{
			Message: "My name is Sam",
			Timestamp: time.Now().Unix(),
		}
        return c.JSON(response)
    })

    app.Listen(":3000")
}