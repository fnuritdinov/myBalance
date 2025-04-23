package credit

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/internal/credit"
)

type Handler interface {
	Credit(c fiber.Ctx) error
}

type handler struct {
	creditService credit.Service
}

func New(creditService credit.Service) Handler {
	return &handler{creditService: creditService}
}
