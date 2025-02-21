package utils

import (
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) (multipart.File, *multipart.FileHeader, error) {
	var file multipart.File
	var fileHeader *multipart.FileHeader

	fileHeader, err := c.FormFile("image")
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewBadRequestError("Failed to get image from form")
	}

	file, err = fileHeader.Open()
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewUnexpectedError()
	}

	return file, fileHeader, nil
}

func UploadBackgroundImage(c *fiber.Ctx) (multipart.File, *multipart.FileHeader, error) {
	var file multipart.File
	var fileHeader *multipart.FileHeader

	fileHeader, err := c.FormFile("background_image")
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewBadRequestError("Failed to get background image from form")
	}

	file, err = fileHeader.Open()
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewUnexpectedError()
	}

	return file, fileHeader, nil
}
