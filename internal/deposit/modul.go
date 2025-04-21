package deposit

type Service interface {
	Deposit(balance string) (newBalance int, err error)
}

type service struct{}

func New() Service {
	return &service{}
}
