package middleware

import (
	"fmt"
	"log"
	"martpedia-backend/internal/app/db"
	"martpedia-backend/internal/app/model/domain"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func DatabaseMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Initialize database
		dbInstance := db.NewDB()

		// Set database instance to the context Fiber
		c.Locals("db", dbInstance)

		// Continue to the next middleware or handler
		return c.Next()
	}
}

func RequiredAuth(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Authorization header is missing",
		})
	}

	// Check if the token is prefixed with "Bearer "
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:] // Delete 'Bearer ' from the token
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token format",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Token has expired",
			})
		}

		// Check if the sub is available in the claims
		sub, ok := claims["sub"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Invalid token claims: sub not found",
			})
		}

		// Check if the database connection is available in locals
		db, ok := c.Locals("db").(*gorm.DB)
		if !ok || db == nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  "error",
				"message": "Database connection not found",
			})
		}

		// Find the user with token sub
		var user domain.User
		if result := db.First(&user, sub); result.Error != nil || user.ID == 0 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "User not found",
			})
		}

		// Attach user to context
		c.Locals("user", user)

		// Continue to the next middleware or handler
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token claims",
		})
	}
}
