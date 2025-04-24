package payment

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func (h *handler) Payment(c fiber.Ctx) error {
	type Req struct {
		MyBalance     int    `json:"my_balance"`
		Method        string `json:"method"`
		CreditBalance int    `json:"credit_balance"`
	}

	var req Req
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Неверный формат данных",
		})
	}

	newBalance, err := h.paymentServise.Payment(req.MyBalance, req.Method, req.CreditBalance)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Ошибка при обработке запроса %v", err),
		})
	}
	return c.JSON(fiber.Map{
		"my_balance": newBalance,
	})
}
