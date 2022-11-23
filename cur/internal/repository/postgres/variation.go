package postgres

import (
	"database/sql"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/variation"
	"jun2/cur/internal/repository"

	"github.com/jmoiron/sqlx"
)

type variationRepository struct {
	db *sqlx.DB
}

func NewVariationRepository(db *sqlx.DB) repository.Variation {
	return &variationRepository{db}
}

func (r *variationRepository) FindVariationInfoByID(id int) ([]variation.Variation, error) {
	data := []variation.Variation{}
	sqlQuery := `
	SELECT pv.variation_id, pv.product_id, pv.variation_type, pv.unit_type, 

	(select price
	from product$price
	where price_id = (
	SELECT max(price_id)
	FROM product$price
	WHERE variation_id=pv.variation_id
	and active_till = (
	SELECT min(active_till)
	FROM product$price
	WHERE variation_id=pv.variation_id
	and current_timestamp between active_from and active_till))) as price

	FROM product$variations pv
	WHERE pv.product_id=$1`

	err := r.db.Select(&data, sqlQuery, id)

	if err == nil && len(data) == 0 {
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

func (r *variationRepository) CreateVariation(param variation.Params) (variationID int, err error) {
	insertQuery := `
	INSERT INTO product$variations(product_id, variation_type, unit_type)
	VALUES ($1, $2, $3) 
	RETURNING variation_id`

	err = r.db.Get(&variationID, insertQuery, param.ProductID, param.VariationType, param.UnitType)

	return
}

func (r *variationRepository) CheckVariatonExistence(param variation.Params) (int, error) {
	sqlQuery := `
	SELECT pv.variation_id 
		FROM product$variations pv 
	WHERE pv.product_id=$1 
	AND pv.variation_type=$2 AND unit_type=$3`
	var data int
	err := r.db.Get(&data, sqlQuery, param.ProductID, param.VariationType, param.UnitType)

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
