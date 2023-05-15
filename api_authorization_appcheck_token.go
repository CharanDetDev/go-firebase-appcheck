package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

func VerifyAppCheckToken(c *fiber.Ctx) error {

	token := c.Get("X-Firebase-AppCheck")
	if token == "" {
		WriteLog("Error Unauthorized", "X-Firebase-AppCheck in Header empty!")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	fmt.Printf("***** GET Token in Header ***** | %v\n\n", token)

	// change file name to -> firebaseServiceAccountConfigMock.json and edit infomation the config firebservice account of you.
	opt := option.WithCredentialsFile("firebaseServiceAccountConfig.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		WriteLog("Error GET Credential file Firebase appcheck", fmt.Sprintf("%+v", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	appCheckService, err := app.AppCheck(context.Background())
	if err != nil {
		WriteLog("Error app.AppCheck()", fmt.Sprintf("%+v", err.Error()))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	decodeAppCheckToken, err := appCheckService.VerifyToken(token)
	if err != nil {
		WriteLog("Error VerifyToken()", fmt.Sprintf("%+v", err.Error()))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	WriteLog("VerifyToken Successfully", fmt.Sprintf("%+v", decodeAppCheckToken))
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": decodeAppCheckToken,
	})

}
