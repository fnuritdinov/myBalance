package main

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/cmd/handler/credit"
	"myBalance/cmd/handler/deposit"
	"myBalance/cmd/handler/my_balance"
	creditSrv "myBalance/internal/credit"
	depositSrv "myBalance/internal/deposit"
	myBalanceSrv "myBalance/internal/my_balance"
)

func main() {
	app := fiber.New()

	balanceService := myBalanceSrv.New()
	balanceHandler := my_balance.New(balanceService)

	app.Get("/my_balance", balanceHandler.MyBalance)

	depositService := depositSrv.New()
	depositHandler := deposit.New(depositService)

	app.Get("/my_balance/deposit/:balance", depositHandler.Deposit)

	creditService := creditSrv.New()
	creditHandler := credit.New(creditService)
	app.Get("/myBalance/credit/:balance", creditHandler.Credit)

	app.Listen(":3000")
}
