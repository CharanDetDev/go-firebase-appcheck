package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Post("/authorization/appcheck", MiddlewareAuthorization, VerifyAppCheckToken)
	log.Fatal(app.Listen(":3001"))
}

func WriteLog(prefixMsg, MsgInfo string) {
	fmt.Printf("***** %v ***** | %v \n\n", prefixMsg, MsgInfo)
	fmt.Printf("*********************************************************************************\n\n")
}
