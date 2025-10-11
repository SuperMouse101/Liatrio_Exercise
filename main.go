package main

import (
	"log"
	"time"
	"github.com/gofiber/fiber/v3"
)

type Responce struct {
    Message string `json:"message"`
    Timestamp int64 `json:"timestamp"`
}

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a json response to the client
        r := Responce {
            Message: "My name is Liam Lassonde",
            Timestamp: time.Now().Unix(),
        }
        return c.JSON(r)
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}