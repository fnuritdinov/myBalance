package payment

type Service interface {
	Payment(my_balance int, method string, creditbalance int) (newBalance int, err error)
}

type service struct{}

func New() Service {
	return &service{}
}
