package variation

import (
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/tools/sqlnull"
)

type Variation struct {
	VariationID       int                  `json:"variation_id" db:"variation_id"`
	ProductID         int                  `json:"product_id" db:"product_id"`
	VariationType     float64              `json:"variation_type" db:"variation_type"`
	UnitType          string               `json:"unit_type" db:"unit_type"`
	Price             sqlnull.NullFloat64  `json:"price" db:"price"`
	StockAvailability []stock.ProductStock `json:"stock_availability"`
}

type Params struct {
	ProductID     int     `json:"product_id" db:"product_id"`
	VariationType float64 `json:"variation_type" db:"variation_type"`
	UnitType      string  `json:"unit_type" db:"unit_type"`
}

func NewVariation(pID int, vType float64, uType string) Params {
	return Params{
		ProductID:     pID,
		VariationType: vType,
		UnitType:      uType,
	}
}

type ProductWithVariationList struct {
	product.Product
	VariationList []Variation `json:"variation_list"`
}
