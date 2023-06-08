package server

import (
	"github.com/cubexsone/fiber-backend-template/src/controller"
	"github.com/cubexsone/fiber-backend-template/src/environment"
	"github.com/cubexsone/fiber-backend-template/src/utils/log"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Server() {
	app := fiber.New()
	app.Use(recover.New())
	env := environment.ENV
	log.Info.Println("Server is running on:", env)

	router := app.Group("/")
	router.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	controller.Api(router)

	app.Listen(":3000")
}
