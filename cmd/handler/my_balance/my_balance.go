package my_balance

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"time"
)

func (h *handler) MyBalance(c fiber.Ctx) (err error) {
	balance, err := h.myBalanceSvc.MyBalance()
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
		"balance": balance,
		"date":    time.Now().Format("2006-01-02 15:04:05"),
	})

}
