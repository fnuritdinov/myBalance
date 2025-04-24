package main

import (
	"github.com/gofiber/fiber/v3"
	"myBalance/cmd/handler/credit"
	"myBalance/cmd/handler/currency"
	"myBalance/cmd/handler/deposit"
	"myBalance/cmd/handler/my_balance"
	"myBalance/cmd/handler/payment"
	userSrvc "myBalance/cmd/handler/user"
	creditSrv "myBalance/internal/credit"
	currencySrv "myBalance/internal/currency"
	depositSrv "myBalance/internal/deposit"
	myBalanceSrv "myBalance/internal/my_balance"
	paymentSrvc "myBalance/internal/payment"
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

	paymentService := paymentSrvc.New()
	paymentHandler := payment.New(paymentService)
	app.Post("/payment", paymentHandler.Payment)


	app.Listen(":3000")

