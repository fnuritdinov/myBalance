package payment

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/internal/payment"
)

type Handler interface {
	Payment(c fiber.Ctx) (err error)
}

type handler struct {
	paymentServise payment.Service
}

func New(paymentService payment.Service) Handler {
	return &handler{paymentServise: paymentService}
}
