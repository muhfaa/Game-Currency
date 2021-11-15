package conversion

import (
	"Game-currency/api/common"
	"Game-currency/api/v1/conversion/request"
	"Game-currency/api/v1/conversion/response"
	"Game-currency/business"
	"Game-currency/business/conversion"
	"Game-currency/util/validator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	conversionService conversion.Service
}

func NewController(conversionService conversion.Service) *Controller {
	return &Controller{
		conversionService,
	}
}

func (controller Controller) InsertConversionRate(c echo.Context) error {
	insertConversionRate := new(request.InsertConversionRateRequest)

	if err := c.Bind(insertConversionRate); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse)
	}

	if err := validator.GetValidator().Struct(insertConversionRate); err != nil {
		return c.JSON(http.StatusBadRequest, common.NewValidationResponse(err.Error()))
	}

	data, _ := controller.conversionService.GetConversionRate(insertConversionRate.CurrencyIDFrom, insertConversionRate.CurrencyIDTo)
	if data != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(business.ErrDuplicate))
	}

	insertData := insertConversionRate.ToUpsertConversionRateSpec()

	err := controller.conversionService.InsertConversionRate(*insertData)
	if err != nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	return c.JSON(http.StatusOK, common.NewSuccessResponseNoData())
}

func (controller Controller) GetConversionRate(c echo.Context) error {
	idFrom := c.QueryParam("currencyIDFrom")
	idFromInt, _ := strconv.Atoi(idFrom)

	idTo := c.QueryParam("currencyIDTo")
	idToInt, _ := strconv.Atoi(idTo)

	amount := c.QueryParam("amount")
	amountInt, _ := strconv.Atoi(amount)

	conversionRate, err := controller.conversionService.GetConversionRate(idFromInt, idToInt)
	if err != nil || conversionRate == nil {
		return c.JSON(common.NewBusinessErrorMappingResponse(err))
	}

	result := amountInt / conversionRate.Rate

	successResponse := common.NewSuccessResponse(response.NewResponse(result))

	return c.JSON(http.StatusOK, successResponse)

}
