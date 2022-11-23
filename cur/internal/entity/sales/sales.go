package sales

import (
	"jun2/cur/tools/sqlnull"
	"time"
)

type Sales struct {
	SaleId      int       `json:"sale_id" db:"sale_id"`
	ProductId   int       `json:"product_id" db:"product_id"`
	VariationId int       `json:"variation_id" db:"variation_id"`
	StockId     int       `json:"stock_id" db:"stock_id"`
	SoldDate    time.Time `json:"sold_date" db:"sold_date"`
	Amount      int       `json:"amount" db:"amount"`
}

type SalesReport struct {
	StockID       int       `json:"stock_id" db:"stock_id"`
	ProductId     int       `json:"product_id" db:"product_id"`
	ProductName   string    `json:"product_name" db:"product_name"`
	VariationType float64   `json:"variation_type" db:"variation_type"`
	UnitType      string    `json:"unit_type" db:"unit_type"`
	StockName     string    `json:"stock_name" db:"stock_name"`
	SoldDate      time.Time `json:"sold_date" db:"sold_date"`
	Amount        int       `json:"amount" db:"amount"`
}

type SalesListWithTotalCount struct {
	SalesList  []SalesReport `json:"sales_list"`
	TotalCount int           `json:"total_count"`
}

type SalesReportRequest struct {
	StartDate   time.Time          `json:"start_date"`
	EndDate     time.Time          `json:"end_date"`
	Limit       sqlnull.NullInt64  `json:"limit"`
	Offset      sqlnull.NullInt64  `json:"offset"`
	ProductName sqlnull.NullString `json:"product_name"`
	StorageID   sqlnull.NullInt64  `json:"storage_id"`
}
