package controller

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"golang-transaction-experiment/config"
	"golang-transaction-experiment/exception"
	"golang-transaction-experiment/service"
)

type OrderController interface {
	Create(ctx *fiber.Ctx) error
}

type OrderControllerImpl struct {
	orderService service.OrderService
}

func NewOrderControllerImpl(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{orderService: orderService}
}

func (c OrderControllerImpl) Create(ctx *fiber.Ctx) error {
	var orderRequest *config.CreateOrderDTO

	if err := ctx.BodyParser(&orderRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, "failed to parse json"))
	}

	err, newOrder := c.orderService.Create(context.Background(), orderRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(exception.NewErrorResponse(fiber.StatusBadRequest, err.Error()))
	}

	response := config.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "new period has been created",
		Data:   newOrder,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
