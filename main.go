package main

import (
    "log"
    "time"
    "github.com/gofiber/fiber/v3"
)

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        r := time.Now().UTC().GoString();
        r += "\nHello, World 👋!";
        return c.SendString(r)
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}