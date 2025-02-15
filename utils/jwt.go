package utils

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ExtractJWTClaims(c *fiber.Ctx) (*types.JWT, error) {
	userObj := c.Locals("user")
	if userObj == nil {
		return nil, errs.NewUnauthorizedError("unauthorized")
	}

	claims, ok := userObj.(jwt.MapClaims)
	if !ok {
		return nil, errs.NewUnauthorizedError("invalid token format")
	}

	jwtClaims := &types.JWT{
		Email:  claims["email"].(string),
		Exp:    int64(claims["exp"].(float64)),
		UserID: claims["user_id"].(string),
	}

	return jwtClaims, nil
}
