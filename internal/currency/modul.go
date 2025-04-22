package currency

type Service interface {
	Currency(somoni string) (rubl float64, err error)
}

type service struct {
	currency_rate float64
}

func New() Service {
	return &service{}
}
