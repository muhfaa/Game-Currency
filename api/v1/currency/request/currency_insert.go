package request

type InsertCurrencyRequest struct {
	Name string `json:"name" validate:"required"`
}
