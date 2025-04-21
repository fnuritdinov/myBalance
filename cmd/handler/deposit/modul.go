package deposit

import "github.com/gofiber/fiber/v3"

type Handler interface {
	Deposit(c fiber.Ctx) error
}

type handler struct{}

func New() Handler {
	return &handler{}
}
