package conversion

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"

	"Game-currency/business"
	conversionBusiness "Game-currency/business/conversion"
)

type MySQLDBRepository struct {
	db *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) *MySQLDBRepository {
	return &MySQLDBRepository{
		db,
	}
}

func (repo *MySQLDBRepository) InsertConversionRate(conversion conversionBusiness.Conversion) error {
	tx, err := repo.db.Begin()
	if err != nil {
		log.Error(err)
		err = errors.New("resource error")
		return err
	}

	insertQuery := `INSERT INTO conversion (
		currencyIDFrom,
		currencyIDTo,
		rate
	)
	VALUES
	(?,?,?)`

	_, err = tx.Exec(insertQuery, conversion.CurrencyIDFrom, conversion.CurrencyIDTo, conversion.Rate)
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

func (repo *MySQLDBRepository) GetConversionRate(idFrom, idTo int) (*conversionBusiness.Conversion, error) {
	var conversion conversionBusiness.Conversion
	selectQuery := `SELECT * FROM conversion 
					WHERE currencyIDFrom = ? AND currencyIDTo = ? 
					OR currencyIDFrom = ? AND currencyIDTo = ?`

	err := repo.db.QueryRowx(selectQuery, idFrom, idTo, idTo, idFrom).StructScan(&conversion)
	if err != nil {
		log.Error(err)
		if err == sql.ErrNoRows {
			return nil, business.ErrNotFound
		}
		err = errors.New("resource error")
		return nil, err
	}

	return &conversion, err
}
