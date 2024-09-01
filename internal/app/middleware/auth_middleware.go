package middleware

import (
	"fmt"
	"log"
	"martpedia-backend/internal/app/model/domain"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func RequiredAuth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Authorization header is missing"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Token has expired"})
		}

		// Find the user with token sub
		var user domain.User
		if result := c.Locals("db").(*gorm.DB).First(&user, claims["sub"]); result.Error != nil || user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User not found"})
		}

		// Attach user to context
		c.Locals("user", user)

		// Continue to the next middleware or handler
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token claims"})
	}
}
