package main

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/cmd"
	"myBalance/cmd/handler/deposit"
)

func main() {
	app := fiber.New()

	app.Get("/my_balance", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"date":    "18.04.25",
			"balance": cmd.NewBalance,
		})
	})

	depositHandler := deposit.New()

	app.Get("/my_balance/deposit/:balance", depositHandler.Deposit)

	app.Get("/my_balance/credit/:balance", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"date":    "18.04.25",
			"balance": cmd.NewBalance - 0,
		})
	})

	app.Listen(":3000")
}
