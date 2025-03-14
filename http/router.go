package http

import (
	"github.com/gofiber/fiber/v2"
	"golang-transaction-experiment/controller"
)

type RouteConfig struct {
	App             *fiber.App
	OrderController controller.OrderController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/v1/order", c.OrderController.Create)
}
