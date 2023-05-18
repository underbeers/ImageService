package controllers

import (
	"ImageService/dtos"
	"ImageService/models"
	"ImageService/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func FileUploadPet(c *fiber.Ctx) error {
	//upload
	formHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Select a file to upload"},
			})
	}

	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	originURL, thumbnailURL, err := services.NewMediaUpload().FileUploadPet(models.File{File: formFile})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	return c.Status(http.StatusOK).JSON(
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &fiber.Map{"original": originURL, "thumbnail": thumbnailURL},
		})
}

func FileUploadUser(c *fiber.Ctx) error {
	//upload
	formHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Select a file to upload"},
			})
	}

	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	uploadURL, err := services.NewMediaUpload().FileUploadUser(models.File{File: formFile})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	return c.Status(http.StatusOK).JSON(
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &fiber.Map{"original": uploadURL},
		})
}

func RemoteUpload(c *fiber.Ctx) error {
	var url models.Url
	//validate the request body

	if err := c.BodyParser(&url); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       &fiber.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := services.NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			dtos.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &fiber.Map{"data": "Error uploading file"},
			})
	}

	return c.Status(http.StatusOK).JSON(
		dtos.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &fiber.Map{"original": uploadUrl},
		})
}
