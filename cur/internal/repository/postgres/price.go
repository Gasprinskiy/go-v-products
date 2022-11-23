package postgres

import (
	"database/sql"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/price"
	"jun2/cur/internal/repository"

	"github.com/jmoiron/sqlx"
)

type priceRepository struct {
	db *sqlx.DB
}

func NewPriceRepository(db *sqlx.DB) repository.Price {
	return &priceRepository{db}
}

func (r *priceRepository) CreatePrice(param price.Price) (data int, err error) {
	sqlQuery := `
	INSERT INTO product$price(variation_id, price, active_till, active_from)
	VALUES ($1, $2, $3, $4)
	RETURNING price_id`

	err = r.db.Get(&data, sqlQuery, param.VariationID, param.Price, param.ActiveTill, param.ActiveFrom)

	if err == sql.ErrNoRows {
		err = global.ErrNoData
	}

	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		return data, err
	}
}
