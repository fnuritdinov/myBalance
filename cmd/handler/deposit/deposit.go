package deposit

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"time"
)

func (h *handler) Deposit(c fiber.Ctx) error {

	balanceStr := c.Params("balance")

	newBalance, err := h.depositService.Deposit(balanceStr)
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
		"balance": newBalance,
		"date":    time.Now().Format("2006-01-02 15:04:05"),
	})
}
