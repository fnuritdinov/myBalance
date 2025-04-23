package credit

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/cmd"
	"strconv"
)

func (s *service) Credit(balanceStr string) (newBalance int, err error) {
	balance, err := strconv.Atoi(balanceStr)
	if err != nil {
		return 0, fiber.ErrBadRequest
	}
	cmd.NewBalance = cmd.NewBalance - balance
	return cmd.NewBalance, nil
}
