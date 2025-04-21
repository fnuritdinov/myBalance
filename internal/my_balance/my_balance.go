package my_balance

import "myBalance/cmd"

func (s *service) MyBalance() (balance int, err error) {
	return cmd.NewBalance, nil
}
