package stock

import (
	"jun2/cur/internal/entity/sales"
	"jun2/cur/tools/sqlnull"
)

type ProductStock struct {
	AccountingID int    `json:"accounting_id" db:"accounting_id"`
	StockID      int    `json:"stock_id" db:"stock_id"`
	ProductID    int    `json:"product_id" db:"product_id"`
	VariationID  int    `json:"variation_id" db:"variation_id"`
	Amount       int    `json:"amount" db:"amount"`
	StockName    string `json:"stock_name" db:"stock_name"`
	Location     string `json:"location" db:"location"`
}

type ProductStockInfoWithTotalCount struct {
	ProductStockInfo []ProductStockInfo `json:"product_stock_info"`
	TotalCount       int                `json:"total_count"`
}

type ProductStockInfo struct {
	StockId    int               `json:"stock_id" db:"stock_id"`
	StockName  string            `json:"stock_name" db:"stock_name"`
	Location   string            `json:"location" db:"location"`
	TotalAmout sqlnull.NullInt64 `json:"total_amount" db:"total_amount"`
}

type StockInfo struct {
	StockId   int    `json:"stock_id" db:"stock_id"`
	StockName string `json:"stock_name" db:"stock_name"`
	Location  string `json:"location" db:"location"`
}

func NewStockFromSales(s sales.Sales) ProductStock {
	return ProductStock{
		StockID:     s.StockId,
		VariationID: s.VariationId,
		ProductID:   s.ProductId,
	}
}
