package my_balance

type Service interface {
	MyBalance() (balance int, err error)
}

type service struct{}

func New() Service {
	return &service{}
}
