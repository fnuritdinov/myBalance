package deposit

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/internal/deposit"
)

type Handler interface {
	Deposit(c fiber.Ctx) error
}

type handler struct {
	depositService deposit.Service
}

func New(depositService deposit.Service) Handler {
	return &handler{depositService: depositService}
}
