package my_balance

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/internal/my_balance"
)

type Handler interface {
	MyBalance(c fiber.Ctx) (err error)
}

type handler struct {
	myBalanceSvc my_balance.Service
}

func New(service my_balance.Service) Handler {
	return &handler{myBalanceSvc: service}
}
