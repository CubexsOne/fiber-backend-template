package controller

import (
	"github.com/gofiber/fiber/v2"
)

func Api(router fiber.Router) {
	apiController := router.Group("/api")
	exampleController(apiController)
}
