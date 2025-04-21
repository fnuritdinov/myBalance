package deposit

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/cmd"
	"strconv"
)

func (s *service) Deposit(balanceStr string) (newBalance int, err error) {
	balance, err := strconv.Atoi(balanceStr)
	if err != nil {
		return 0, fiber.ErrBadRequest
	}
	return cmd.NewBalance - balance, nil
}
