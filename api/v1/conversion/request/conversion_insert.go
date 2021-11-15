package request

import "Game-currency/business/conversion"

type InsertConversionRateRequest struct {
	CurrencyIDFrom int `json:"currencyIDFrom" validate:"required"`
	CurrencyIDTo   int `json:"currencyIDTo" validate:"required"`
	Rate           int `json:"rate" validate:"required"`
}

func (req *InsertConversionRateRequest) ToUpsertConversionRateSpec() *conversion.Conversion {
	var conversionSpec conversion.Conversion

	conversionSpec.CurrencyIDFrom = req.CurrencyIDFrom
	conversionSpec.CurrencyIDTo = req.CurrencyIDTo
	conversionSpec.Rate = req.Rate

	return &conversionSpec

}
