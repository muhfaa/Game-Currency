package currency

import (
	"errors"
	"strings"

	currencyBusiness "Game-currency/business/currency"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type MySQLDBRepository struct {
	db *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) *MySQLDBRepository {
	return &MySQLDBRepository{
		db,
	}
}

func (repo *MySQLDBRepository) InsertCurrency(name string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}

	insertQuery := `INSERT INTO currency (
		name
	)
	VALUES
	(?)`

	_, err = tx.Exec(insertQuery, name)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		if strings.Contains(err.Error(), "Duplicate entry") {
			err = errors.New("error duplicate entry data")
			return err
		}

		err = errors.New("resource error")
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}

	return nil
}

func (repo *MySQLDBRepository) GetCurrency() ([]currencyBusiness.Currency, error) {
	var (
		allCurrency []currencyBusiness.Currency
		currency    currencyBusiness.Currency
	)

	selectQuery := `SELECT * FROM currency`

	row, err := repo.db.Queryx(selectQuery)
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return nil, err
	}

	for row.Next() {
		err = row.StructScan(&currency)
		if err != nil {
			log.Error(err)
			err = errors.New("resource error")
			return nil, err
		}

		allCurrency = append(allCurrency, currency)
	}

	return allCurrency, nil
}
