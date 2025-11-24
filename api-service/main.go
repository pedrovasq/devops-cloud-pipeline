package main

import (
    "log"
	"time"
	"fmt"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/health", func (c *fiber.Ctx) error {
        return c.SendStatus(200)
    })

	app.Get("/time", func (c *fiber.Ctx) error {
		t := time.Now()
		return c.SendString(fmt.Sprintln("Current time: ", t))
	})

    log.Fatal(app.Listen(":8080"))
}
