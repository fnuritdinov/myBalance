package credit

type Service interface {
	Credit(balance string) (newBalance int, err error)
}

type service struct {
}

func New() Service {
	return &service{}
}
