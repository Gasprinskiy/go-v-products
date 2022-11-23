package postgres

import (
	"database/sql"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/sales"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/internal/repository"

	"github.com/jmoiron/sqlx"
)

type stockRepository struct {
	db *sqlx.DB
}

func NewStockRepository(db *sqlx.DB) repository.Stock {
	return &stockRepository{db}
}

func (r *stockRepository) CreateStock(param stock.StockInfo) (data int, err error) {
	sqlQuerry := `
	INSERT INTO stocks(stock_name, location)
	VALUES ($1, $2)
	RETURNING stock_id
	`

	err = r.db.Get(&data, sqlQuerry, param.StockName, param.Location)

	if err != nil {
		return
	}

	return
}

func (r *stockRepository) FindVariationAvailability(id int) (data []stock.ProductStock, err error) {
	sqlQuery := `
	SELECT ps.accounting_id, ps.stock_id, ps.product_id, ps.amount, ps.variation_id, st.stock_name, st.location
	FROM product$stocks ps
        JOIN stocks st ON (ps.stock_id = st.stock_id)
	WHERE (ps.variation_id=$1)`

	err = r.db.Select(&data, sqlQuery, id)

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

func (r *stockRepository) AddProductToStock(param stock.ProductStock) (id int, err error) {
	sqlQuery := `
	INSERT INTO product$stocks(stock_id, product_id, amount, variation_id)
	VALUES ($1, $2, $3, $4)
	RETURNING accounting_id 
	`
	err = r.db.Get(&id, sqlQuery, param.StockID, param.ProductID, param.Amount, param.VariationID)

	if err != nil {
		return
	}

	return
}

func (r *stockRepository) CheckProductExistence(param stock.ProductStock) (data int, err error) {
	sqlQuery := `
	SELECT ps.accounting_id 
	FROM product$stocks ps 
	WHERE ps.stock_id=$1 
	AND ps.variation_id=$2 
	AND ps.product_id=$3
	`

	err = r.db.Get(&data, sqlQuery, param.StockID, param.VariationID, param.ProductID)

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

func (r *stockRepository) IncreaseAccountingAmount(amount, acID int) (err error) {
	sqlQuery := `
	UPDATE product$stocks 
	SET amount=amount+$1 
	WHERE accounting_id=$2
	`

	_, err = r.db.Exec(sqlQuery, amount, acID)

	return
}

func (r *stockRepository) DecreaseAccountingAmount(amount, acID int) (err error) {
	sqlQuery := `
	UPDATE product$stocks 
	SET amount=amount-$1 
	WHERE accounting_id=$2
	`

	_, err = r.db.Exec(sqlQuery, amount, acID)

	return
}

func (r *stockRepository) FindStockInfo(limit, offset int) (data []stock.ProductStockInfo, err error) {
	sqlQuery := `
	SELECT st.stock_id, st.stock_name, st.location,

	(SELECT SUM(ps.amount)
	FROM product$stocks ps
	WHERE (ps.stock_id = st.stock_id)
	) AS total_amount

	FROM stocks st 
	GROUP BY (st.stock_id, st.stock_name, st.location)
	ORDER BY st.stock_id
	DESC
	LIMIT $1
	OFFSET $2
	`

	err = r.db.Select(&data, sqlQuery, limit, offset)

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

func (r *stockRepository) LoadStockCount() (data int, err error) {
	sqlQuery := `
	SELECT count(stock_id) 
	FROM stocks
	`

	err = r.db.Get(&data, sqlQuery)

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

func (r *stockRepository) FindProductStockInfo(productID, limit, offset int) (data []stock.ProductStockInfo, err error) {
	sqlQuery := `
	SELECT ps.stock_id, st.stock_name, st.location, 
	SUM(ps.amount) AS total_amount
	FROM product$stocks ps
	JOIN stocks st 
	ON (ps.stock_id = st.stock_id)
	WHERE (ps.product_id=$1)
	GROUP BY (ps.stock_id, st.stock_name, st.location)
	ORDER BY ps.stock_id
	DESC
	LIMIT $2
	OFFSET $3
	`

	err = r.db.Select(&data, sqlQuery, productID, limit, offset)

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

func (r *stockRepository) FindProductStockCount(productID int) (data int, err error) {
	sqlQuery := `
	SELECT count(DISTINCT st.stock_id)
	FROM stocks st
	JOIN product$stocks ps 
	ON (ps.stock_id = st.stock_id)
	WHERE (ps.product_id=$1)
	`
	err = r.db.Get(&data, sqlQuery, productID)

	if err == sql.ErrNoRows {
		err = global.ErrNoData
	}

	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		return data, global.ErrInternalError
	}
}

func (r *stockRepository) LoadStockList() (data []stock.StockInfo, err error) {
	sqlQuery := `
	SELECT stock_id, stock_name, location 
	FROM stocks 
	`

	err = r.db.Select(&data, sqlQuery)

	if err == nil && len(data) == 0 {
		err = global.ErrNoData
	}

	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		return data, global.ErrInternalError
	}
}

func (r *stockRepository) FindProductAmount(param sales.Sales) (data int, err error) {
	sqlQuery := `
	SELECT ps.amount
	FROM product$stocks ps
	WHERE 
	(ps.product_id=$1 
	and ps.variation_id=$2 
	and ps.stock_id=$3)`

	err = r.db.Get(&data, sqlQuery, param.ProductId, param.VariationId, param.StockId)

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
