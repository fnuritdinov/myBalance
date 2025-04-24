package payment

import (
	"errors"
)

func (s *service) Payment(my_balance int, method string, creditbalance int) (newBalance int, err error) {
	if method == "credit" {
		newBalance = creditbalance - my_balance
	} else if method == "debit" {
		newBalance = creditbalance + my_balance
	} else {
		return 0, errors.New("invalid method")
	}

	return newBalance, nil
}
