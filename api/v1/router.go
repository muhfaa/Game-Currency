package router

import (
	"Game-currency/api/v1/conversion"
	"Game-currency/api/v1/currency"

	"github.com/labstack/echo/v4"
)

func RegisterPath(
	e *echo.Echo,
	currencyController *currency.Controller,
	conversionController *conversion.Controller,
) {

	// currency
	currencyV1 := e.Group("v1/currency")
	currencyV1.POST("/add", currencyController.InsertCurrency)
	currencyV1.GET("", currencyController.GetCurrency)

	// conversion
	conversionV1 := e.Group("v1/conversion")
	conversionV1.POST("/add", conversionController.InsertConversionRate)
	conversionV1.GET("", conversionController.GetConversionRate)

}
