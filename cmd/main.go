package main

import (
	config "ImageService/configs"
	"ImageService/controllers"
	"ImageService/services"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Cloudinary"})
	})
	app.Post("/filePet", controllers.FileUploadPet)
	app.Post("/fileUser", controllers.FileUploadUser)

	services.GWConnection()

	err := app.Listen("0.0.0.0:" + config.EnvPort())
	if err != nil {
		log.Default().Fatalf("failed to start listen to port %s", config.EnvPort())
	}
}
