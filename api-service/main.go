package main

import (
    "log"
	"time"
	"fmt"
	"os"

    "github.com/gofiber/fiber/v2"
)

func main() {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "devops-api"
	}

    app := fiber.New()

    app.Get("/health", func (c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf(`{"status":"ok","app":"%s"}`, appName))
    })

	app.Get("/time", func (c *fiber.Ctx) error {
		t := time.Now()
		return c.SendString(fmt.Sprintln("Current time: ", t))
	})

    log.Fatal(app.Listen(":8080"))
}
