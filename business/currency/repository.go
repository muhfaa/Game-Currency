package currency

type Repository interface {
	InsertCurrency(name string) error
	GetCurrency() ([]Currency, error)
}
