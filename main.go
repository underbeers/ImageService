package main

import (
	config "ImageService/configs"
	"ImageService/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Cloudinary"})
	})
	app.Post("/file", controllers.FileUpload)
	app.Post("/remote", controllers.RemoteUpload)

	app.Listen("0.0.0.0:" + config.EnvPort())
}
