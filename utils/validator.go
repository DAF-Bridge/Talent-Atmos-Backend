package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/types"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(data interface{}) []types.ValidationError {
	var validationErrors []types.ValidationError

	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, types.ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
				Error: true,
			})
		}
	}

	return validationErrors
}

func ParseJSONAndValidate(c *fiber.Ctx, body interface{}) error {
	if c.Body() == nil || len(c.Body()) == 0 {
		return errs.NewBadRequestError("Request Body is Empty")
	}
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if body == nil {
		return errs.NewBadRequestError("Empty JSON body")
	}

	validationErrors := ValidateStruct(body)
	if len(validationErrors) > 0 {
		// Convert validation errors to a slice of strings
		var errorMessages []string
		for _, ve := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation: %s (value: %v)\n", ve.Field, ve.Tag, ve.Value))
		}

		// Join error messages and return as a single error string
		return errs.NewBadRequestError(strings.Join(errorMessages, "; "))

	}

	return nil
}

func UnmarshalAndValidateJSON(jsonStr string, dest interface{}) error {
	if jsonStr == "" {
		return errs.NewBadRequestError("Empty JSON body")
	}

	if err := json.Unmarshal([]byte(jsonStr), dest); err != nil {
		var syntaxErr *json.SyntaxError
		var typeErr *json.UnmarshalTypeError

		if errors.As(err, &syntaxErr) {
			logs.Error(fmt.Sprintf("Syntax error at offset %d", syntaxErr.Offset))
			return errs.NewBadRequestError(fmt.Sprintf("Syntax error at offset %d", syntaxErr.Offset))
		} else if errors.As(err, &typeErr) {
			logs.Error(fmt.Sprintf("Type error: expected %s but got %v (field: %s)", typeErr.Type, typeErr.Value, typeErr.Field))
			return errs.NewBadRequestError(fmt.Sprintf("Type error: expected %s but got %v (field: %s)", typeErr.Type, typeErr.Value, typeErr.Field))
		}

		logs.Error(err)
		return errs.NewBadRequestError(fmt.Sprintf("Invalid JSON format: %s", err.Error()))
	}

	validationErrors := ValidateStruct(dest)
	if len(validationErrors) > 0 {
		var errorMessages []string
		for _, ve := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation: %s (value: %v)\n", ve.Field, ve.Tag, ve.Value))
		}

		return errs.NewBadRequestError(strings.Join(errorMessages, "; "))
	}

	return nil
}
