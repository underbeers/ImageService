package dtos

import "github.com/gofiber/fiber/v2"

type MediaDto struct {
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	Data       *fiber.Map `json:"data"`
}
