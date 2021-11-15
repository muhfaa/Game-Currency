package currency

import (
	"Game-currency/api/common"
	"Game-currency/api/v1/currency/request"
	"Game-currency/api/v1/currency/response"
	"Game-currency/business/currency"
	"Game-currency/util/validator"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	currencyService currency.Service
}

func NewController(currencyService currency.Service) *Controller {
	return &Controller{
		currencyService,
	}
}

func (controller Controller) InsertCurrency(c echo.Context) error {
	currencyRequest := new(request.InsertCurrencyRequest)

	if err := c.Bind(currencyRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
	}

	if err := validator.GetValidator().Struct(currencyRequest); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	if err := controller.currencyService.InsertCurrency(currencyRequest.Name); err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

func (controller Controller) GetCurrency(c echo.Context) error {
	currency, err := controller.currencyService.GetCurrency()
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	successResponse := common.NewSuccessResponse(response.NewResponse(currency))
	return c.JSON(http.StatusOK, successResponse)
}
