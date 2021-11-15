package conversion

type Conversion struct {
	ID             int `json:"_" db:"id"`
	CurrencyIDFrom int `json:"currencyIDFrom" db:"currencyIDFrom"`
	CurrencyIDTo   int `json:"currencyIDTo" db:"currencyIDTo"`
	Rate           int `json:"rate" db:"rate"`
}
