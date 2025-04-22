package currency

import (
	"errors"
	"github.com/gofiber/fiber/v3"
)

func (h *handler) Currency(c fiber.Ctx) error {
	rateStr := c.Params("currency_rate")

	my_balance, err := h.currencyService.Currency(rateStr)
	if err != nil {
		if errors.Is(err, fiber.ErrBadRequest) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"balance_tjs": 10,
		"balance_rub": my_balance,
	})
}
