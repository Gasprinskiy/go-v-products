package postgres

import (
	"fmt"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/sales"
	"jun2/cur/internal/repository"

	"github.com/jmoiron/sqlx"
)

type salesRepository struct {
	db *sqlx.DB
}

func NewSalesRepository(db *sqlx.DB) repository.Sales {
	return &salesRepository{db}
}

func (r *salesRepository) FindSalesList(limit, offset int) (data []sales.SalesReport, err error) {
	sqlQuery := `
	SELECT ps.product_id, pr.product_name, pv.variation_type, pv.unit_type, st.stock_id, st.stock_name, ps.sold_date, ps.amount
	FROM product$sales ps
		JOIN product pr ON (ps.product_id = pr.product_id)
		JOIN product$variations pv ON (ps.variation_id = pv.variation_id)
		JOIN stocks st ON (ps.stock_id = st.stock_id)
	ORDER BY ps.sale_id 
	DESC 
	LIMIT $1
	OFFSET $2
	`

	err = r.db.Select(&data, sqlQuery, limit, offset)

	if err == nil && len(data) <= 0 {
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

func (r *salesRepository) LoadSalesListCount() (data int, err error) {
	sqlQuery := `
	SELECT COUNT(ps.sale_id)
	FROM product$sales ps
		JOIN product pr ON (ps.product_id = pr.product_id)
		JOIN product$variations pv ON (ps.variation_id = pv.variation_id)
		JOIN stocks st ON (ps.stock_id = st.stock_id)
	`

	err = r.db.Get(&data, sqlQuery)

	if err == global.ErrNoData {
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

func (r *salesRepository) AddProductSale(param sales.Sales) (id int, err error) {
	sqlQuery := `
	INSERT INTO product$sales(product_id, variation_id, stock_id, amount)
	VALUES ($1, $2, $3, $4)
	RETURNING sale_id`

	err = r.db.Get(&id, sqlQuery, param.ProductId, param.VariationId, param.StockId, param.Amount)
	if err != nil {
		return
	}
	return
}

func (r *salesRepository) FindSalesReport(param sales.SalesReportRequest) (data []sales.SalesReport, err error) {
	sqlQuery := r.buildSalesReportSqlQuery(param)
	err = r.db.Select(&data, sqlQuery, param.StartDate, param.EndDate)

	if err == nil && len(data) <= 0 {
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

func (r *salesRepository) FindSalesReportCount(param sales.SalesReportRequest) (data int, err error) {
	sqlQuery := r.buildSalesReportCountSqlQuery(param)
	err = r.db.Get(&data, sqlQuery, param.StartDate, param.EndDate)

	if err == global.ErrNoData {
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

func (r *salesRepository) buildSalesReportSqlQuery(p sales.SalesReportRequest) string {
	limit := global.DefaultLimit
	sqlQuery := `
	SELECT ps.product_id, pr.product_name, pv.variation_type, pv.unit_type, st.stock_id, st.stock_name, ps.sold_date, ps.amount
	FROM product$sales ps
		JOIN product pr ON (ps.product_id = pr.product_id)
		JOIN product$variations pv ON (ps.variation_id = pv.variation_id)
		JOIN stocks st ON (ps.stock_id = st.stock_id)
	WHERE ps.sold_date BETWEEN $1 AND $2`

	if p.ProductName.Valid {
		nameQuery := fmt.Sprintf(" AND LOWER(pr.product_name) LIKE LOWER('%s')", "%"+p.ProductName.String+"%")
		sqlQuery += nameQuery
	}
	if p.StorageID.Valid {
		stockIdQuery := fmt.Sprintf(" AND ps.stock_id=%d", p.StorageID.GetInt())
		sqlQuery += stockIdQuery
	}
	if p.Limit.Valid {
		limit = p.Limit.GetInt()
	}

	limitQuery := fmt.Sprintf(` 
	ORDER BY ps.sale_id 
	DESC LIMIT %d OFFSET %d`, limit, p.Offset.GetInt())

	sqlQuery += limitQuery

	return sqlQuery
}

func (r *salesRepository) buildSalesReportCountSqlQuery(p sales.SalesReportRequest) string {
	sqlQuery := `
	SELECT  COUNT(ps.sale_id)
	FROM product$sales ps
		JOIN product pr ON (ps.product_id = pr.product_id)
		JOIN product$variations pv ON (ps.variation_id = pv.variation_id)
		JOIN stocks st ON (ps.stock_id = st.stock_id)
	WHERE ps.sold_date BETWEEN $1 AND $2`

	if p.ProductName.Valid {
		nameQuery := fmt.Sprintf(" AND LOWER(pr.product_name) LIKE LOWER('%s')", "%"+p.ProductName.String+"%")
		sqlQuery += nameQuery
	}
	if p.StorageID.Valid {
		stockIdQuery := fmt.Sprintf(" AND ps.stock_id=%d", p.StorageID.GetInt())
		sqlQuery += stockIdQuery
	}

	return sqlQuery
}
