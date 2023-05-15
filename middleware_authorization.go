package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareAuthorization(c *fiber.Ctx) error {

	fmt.Printf("\n********************************** New Request **********************************\n\n")
	fmt.Printf("***** GET Header ***** | %v", string(c.Request().Header.Header()))

	var token string
	authorization := c.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") {
		token = strings.TrimPrefix(authorization, "Bearer ")
	}

	if token == "" {
		WriteLog("Error Unauthorized", "Token Firebase AppCheck in Header empty!")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Request().Header.Set("X-Firebase-AppCheck", token)
	c.Next()

	return nil
}
