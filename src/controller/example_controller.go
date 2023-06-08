package controller

import (
	"github.com/gofiber/fiber/v2"
)

func exampleController(router fiber.Router) {
	controller := router.Group("/example")
	controller.Get("/special", specialRoute)
	controller.Get("/specialNotFound", specialRouteNotFound)
}

func specialRoute(c *fiber.Ctx) error {
	return c.SendString("Special")
}

func specialRouteNotFound(ctx *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusNotFound, "Special Error")
}
