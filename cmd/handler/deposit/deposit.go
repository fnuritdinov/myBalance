package deposit

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/cmd"
	"strconv"
)

func (h *handler) Deposit(c fiber.Ctx) error {

	balanceStr := c.Params("balance")
	balance, err := strconv.Atoi(balanceStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "error balance",
		})
	}

	cmd.NewBalance -= balance
	return c.JSON(fiber.Map{
		"balance": cmd.NewBalance,
		"date":    "18.04.25",
	})
}
