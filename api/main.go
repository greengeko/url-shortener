package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"url-shortener/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
