package utils

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) (multipart.File, *multipart.FileHeader, error) {

	fileHeader, err := c.FormFile("image")
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewBadRequestError("Failed to get image from form")
	}

	file, err := fileHeader.Open()
	if err != nil {
		logs.Error(err)
		return nil, nil, errs.NewUnexpectedError()
	}
	defer file.Close()
	return file, fileHeader, nil
}
