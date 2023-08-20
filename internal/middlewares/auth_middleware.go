package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	jwtKey []byte
}

func NewAuthMiddleware(jwtKey []byte) *AuthMiddleware {
	return &AuthMiddleware{
		jwtKey: jwtKey,
	}
}

func (middleware *AuthMiddleware) VerifyToken(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token bulunamadı"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("geçersiz token yöntemi")
		}
		return middleware.jwtKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Geçersiz token"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("userID", claims["id"].(float64)) // User ID'yi yerel değişkende sakla
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Geçersiz token"})
}
