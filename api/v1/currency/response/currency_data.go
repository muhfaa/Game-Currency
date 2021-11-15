package response

import "Game-currency/business/currency"

type CurrencyDataResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewResponse(currencie []currency.Currency) *[]CurrencyDataResponse {
	var (
		currencyDatas []CurrencyDataResponse
		currencyData  CurrencyDataResponse
	)

	for _, currency := range currencie {
		currencyData.ID = currency.ID
		currencyData.Name = currency.Name

		currencyDatas = append(currencyDatas, currencyData)
	}

	return &currencyDatas

}
