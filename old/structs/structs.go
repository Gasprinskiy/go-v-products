package structs

import "time"

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type ProductWithVariations struct {
	Name          string  `json:"product_name" db:"product_name"`
	Description   string  `json:"description" db:"description"`
	Tags          string  `json:"tags" db:"tags"`
	VariationType float64 `json:"variation_type" db:"variation_type"`
	UnitType      string  `json:"unit_type" db:"unit_type"`
}

type Product struct {
	Id          int         `json:"product_id" db:"product_id"`
	Name        string      `json:"product_name" db:"product_name"`
	Description string      `json:"description" db:"description"`
	Tags        string      `json:"tags" db:"tags"`
	Variations  []Variation `json:"variations"`
}

type Variation struct {
	VariationId       int            `json:"variation_id" db:"variation_id"`
	ProductId         int            `json:"product_id" db:"product_id"`
	VariationType     float64        `json:"variation_type" db:"variation_type"`
	UnitType          string         `json:"unit_type" db:"unit_type"`
	Price             float64        `json:"price" db:"price"`
	StockAvailability []StockProduct `json:"stock_availability"`
}

type Price struct {
	PriceId     int     `json:"price_id" db:"price_id"`
	VariationId int     `json:"variation_id" db:"variation_id"`
	Price       float64 `json:"price" db:"price"`
	ActiveFrom  string  `json:"active_from" db:"active_from"`
	ActiveTill  string  `json:"active_till" db:"active_till"`
}

type StockProduct struct {
	AccountingId int    `json:"accounting_id" db:"accounting_id"`
	StockId      int    `json:"stock_id" db:"stock_id"`
	ProductId    int    `json:"product_id" db:"product_id"`
	VariationId  int    `json:"variation_id" db:"variation_id"`
	Amount       int    `json:"amount" db:"amount"`
	StockName    string `json:"stock_name" db:"stock_name"`
	Location     string `json:"location" db:"location"`
}

type StockInfo struct {
	StockId    int    `json:"stock_id" db:"stock_id"`
	StockName  string `json:"stock_name" db:"stock_name"`
	Location   string `json:"location" db:"location"`
	TotalAmout int    `json:"total_amout" db:"total_amount"`
}

type Sales struct {
	SaleId      int    `json:"sale_id" db:"sale_id"`
	ProductId   int    `json:"product_id" db:"product_id"`
	VariationId int    `json:"variation_id" db:"variation_id"`
	StockId     int    `json:"stock_id" db:"stock_id"`
	SoldDate    string `json:"sold_date" db:"sold_date"`
	Amount      int    `json:"amount" db:"amount"`
}

type SalesReportBody struct {
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Limit       int    `json:"limit"`
	ProductName string `json:"product_name"`
	StorageId   int    `json:"storage_id"`
}

type SalesReport struct {
	ProductId     int       `json:"product_id" db:"product_id"`
	ProductName   string    `json:"product_name" db:"product_name"`
	VariationType float64   `json:"variation_type" db:"variation_type"`
	UnitType      string    `json:"unit_type" db:"unit_type"`
	StockName     string    `json:"stock_name" db:"stock_name"`
	SoldDate      time.Time `json:"sold_date" db:"sold_date"`
	Amount        int       `json:"amount" db:"amount"`
}

func NewProduct(source ProductWithVariations) Product {
	return Product{
		Name:        source.Name,
		Description: source.Description,
		Tags:        source.Tags,
	}
}

func NewVariation(source ProductWithVariations, productId int) Variation {
	return Variation{
		ProductId:     productId,
		VariationType: source.VariationType,
		UnitType:      source.UnitType,
	}
}
