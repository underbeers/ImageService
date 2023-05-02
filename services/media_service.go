package services

import (
	"ImageService/helper"
	"ImageService/models"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUploadPet(file models.File) (string, string, error)
	FileUploadUser(file models.File) (string, error)
	RemoteUpload(url models.Url) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUploadPet(file models.File) (string, string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", "", err
	}

	//upload
	originURL, thumbnailURL, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", "", err
	}
	return originURL, thumbnailURL, nil
}

func (*media) FileUploadUser(file models.File) (string, error) {

	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadURL, err := helper.ImageUploadHelperUser(file.File)
	if err != nil {
		return "", err
	}
	return uploadURL, nil

}

func (*media) RemoteUpload(url models.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	originURL, _, errUrl := helper.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return originURL, nil
}
