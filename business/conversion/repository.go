package conversion

type Repository interface {
	InsertConversionRate(conversion Conversion) error
	GetConversionRate(idFrom, idTo int) (*Conversion, error)
}
