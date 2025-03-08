package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetStructFromLocals[T any](ctx *fiber.Ctx, key string) (T, error) {
	var zeroValue T // Zero value of the desired type
	value := ctx.Locals(key)
	if value == nil {
		return zeroValue, fmt.Errorf("key %s not found in Locals", key)
	}

	typedValue, ok := value.(T)
	if !ok {
		return zeroValue, fmt.Errorf("value for key %s is not of type %T", key, zeroValue)
	}

	return typedValue, nil
}

func GetJWTFromLocals(c *fiber.Ctx, key string) (jwt.MapClaims, error) {
	claims, err := GetStructFromLocals[jwt.MapClaims](c, key)
	if err != nil {
		return nil, err
	}

	return claims, nil

}
