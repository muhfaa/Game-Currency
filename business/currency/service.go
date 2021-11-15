package currency

type Service interface {
	InsertCurrency(name string) error
	GetCurrency() ([]Currency, error)
}

type service struct {
	currencyRepository Repository
}

func NewService(currencyRepository Repository) Service {
	return &service{
		currencyRepository,
	}
}

func (s *service) InsertCurrency(name string) error {
	return s.currencyRepository.InsertCurrency(name)
}

func (s *service) GetCurrency() ([]Currency, error) {
	return s.currencyRepository.GetCurrency()
}
