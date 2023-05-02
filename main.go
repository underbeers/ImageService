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
	app.Post("/filePet", controllers.FileUploadPet)
	app.Post("/fileUser", controllers.FileUploadUser)

	app.Listen("0.0.0.0:" + config.EnvPort())
}
