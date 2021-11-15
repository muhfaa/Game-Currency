package currency

type Currency struct {
	ID   int    `json:"_" db:"id"`
	Name string `json:"name" db:"name"`
}
