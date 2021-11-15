package main

import (
	"Game-currency/business/conversion"
	"Game-currency/business/currency"
	"Game-currency/config"
	"Game-currency/modules/mysql"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	conversionModule "Game-currency/modules/conversion"
	currencyModule "Game-currency/modules/currency"

	conversionController "Game-currency/api/v1/conversion"
	currencyController "Game-currency/api/v1/currency"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	GameCurrencyAPP "Game-currency/api/v1"
)

func newDatabaseConnection() *sqlx.DB {

	uri := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.GetConfig().Mysql.User,
		config.GetConfig().Mysql.Password,
		config.GetConfig().Mysql.Host,
		config.GetConfig().Mysql.Port,
		config.GetConfig().Mysql.Name)

	db := mysql.NewDatabaseConnection(uri)

	fmt.Println("uri:", uri)
	return db

}

func newCurrencyService(db *sqlx.DB) currency.Service {
	currencyRepository := currencyModule.NewMySQLRepository(db)

	return currency.NewService(currencyRepository)
}

func newConversionService(db *sqlx.DB) conversion.Service {
	conversionRepository := conversionModule.NewMySQLRepository(db)

	return conversion.NewService(conversionRepository)
}

func main() {

	// database init
	db := newDatabaseConnection()

	// service init
	currencyService := newCurrencyService(db)

	conversionService := newConversionService(db)

	// controller init
	currencyV1 := currencyController.NewController(currencyService)

	conversionV1 := conversionController.NewController(conversionService)

	e := echo.New()

	GameCurrencyAPP.RegisterPath(
		e,
		currencyV1,
		conversionV1,
	)

	// run server
	go func() {
		address := fmt.Sprintf("0.0.0.0:%d", config.GetConfig().BackendPort)

		if err := e.Start(address); err != nil {
			log.Error("failed to start server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Wait for interrupt signal to gracefully shutdown the server with
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error("failed to grafefully shutting down server", err)
	}

}
