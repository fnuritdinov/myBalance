package currency

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/internal/currency"
)

type Handler interface {
	Currency(c fiber.Ctx) error
}

type handler struct {
	currencyService currency.Service
}

func New(currencyService currency.Service) Handler {
	return &handler{currencyService: currencyService}
}
