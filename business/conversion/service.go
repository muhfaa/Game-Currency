package conversion

type Service interface {
	InsertConversionRate(conversion Conversion) error
	GetConversionRate(idFrom, idTo int) (*Conversion, error)
}

type service struct {
	conversionRepository Repository
}

func NewService(conversionRepository Repository) Service {
	return &service{
		conversionRepository,
	}
}

func (s *service) InsertConversionRate(conversion Conversion) error {

	return s.conversionRepository.InsertConversionRate(conversion)
}

func (s *service) GetConversionRate(idFrom, idTo int) (*Conversion, error) {

	return s.conversionRepository.GetConversionRate(idFrom, idTo)
}
