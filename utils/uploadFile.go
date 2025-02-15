package utils

import (
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) (multipart.File, *multipart.FileHeader, error) {
	var file multipart.File
	var fileHeader *multipart.FileHeader
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return nil, nil, errs.NewBadRequestError("Failed to get image from form")
	}

	file, err = fileHeader.Open()
	if err != nil {
		return nil, nil, errs.NewUnexpectedError()
	}
	defer file.Close()

	return file, fileHeader, nil
}
