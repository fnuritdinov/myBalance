package currency

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/cmd"
	"strconv"
)

func (s *service) Currency(rateStr string) (float64, error) {

	rate, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		return 0, fiber.ErrBadRequest
	}

	return float64(cmd.NewBalance) * rate, nil
}
