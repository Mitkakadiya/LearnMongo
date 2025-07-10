package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"mongo_db/config"
	"mongo_db/pkg/routers"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	config.ConnectDB()
	app := fiber.New()
	routers.Router(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
