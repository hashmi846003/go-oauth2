package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashmi846003/go-oauth2/config"
	"github.com/hashmi846003/go-oauth2/controllers"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()

	app.Get("/google_login", controllers.GoogleLogin)
	app.Post("/google_callback", controllers.GoogleCallback)

	app.Listen(":8080")

}
