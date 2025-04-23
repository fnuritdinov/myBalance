package main

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/cmd/handler/credit"
	"myBalance/cmd/handler/currency"
	"myBalance/cmd/handler/deposit"
	"myBalance/cmd/handler/my_balance"
	creditSrv "myBalance/internal/credit"
	currencySrv "myBalance/internal/currency"
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

	currencyService := currencySrv.New()
	currencyHandler := currency.New(currencyService)
	app.Get("/currency_rate/:currency_rate", currencyHandler.Currency)

	//app.Get("/convert/30", (c fiber.Ctx) error

	app.Listen(":3000")
}
