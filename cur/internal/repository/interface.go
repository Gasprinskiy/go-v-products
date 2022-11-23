package repository

import (
	"jun2/cur/internal/entity/price"
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/entity/sales"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/internal/entity/variation"
)

type Product interface {
	FindProductByID(id int) (product.Product, error)
	CreateProduct(newProduct product.Product) (int, error)
	FindProductList(limit, offset int) ([]product.Product, error)
	FindProductListByTags(limit, offset int, tags string) ([]product.Product, error)
	LoadProductCount() (int, error)
	FindProductCountByTags(paramtags string) (int, error)
}

type Variation interface {
	FindVariationInfoByID(id int) ([]variation.Variation, error)
	CreateVariation(newVariation variation.Params) (int, error)
	CheckVariatonExistence(param variation.Params) (int, error)
}

type Stock interface {
	CreateStock(param stock.StockInfo) (int, error)
	FindVariationAvailability(id int) ([]stock.ProductStock, error)
	AddProductToStock(newAccounting stock.ProductStock) (int, error)
	CheckProductExistence(param stock.ProductStock) (int, error)
	IncreaseAccountingAmount(amount, acID int) error
	DecreaseAccountingAmount(amount, acID int) error
	FindStockInfo(limit, offset int) ([]stock.ProductStockInfo, error)
	LoadStockCount() (int, error)
	FindProductStockCount(productID int) (int, error)
	LoadStockList() ([]stock.StockInfo, error)
	FindProductStockInfo(productID, limit, offset int) ([]stock.ProductStockInfo, error)
	FindProductAmount(param sales.Sales) (int, error)
}

type Price interface {
	CreatePrice(param price.Price) (int, error)
}

type Sales interface {
	FindSalesList(limit, offset int) ([]sales.SalesReport, error)
	LoadSalesListCount() (int, error)
	AddProductSale(param sales.Sales) (int, error)
	FindSalesReport(param sales.SalesReportRequest) ([]sales.SalesReport, error)
	FindSalesReportCount(param sales.SalesReportRequest) (int, error)
}
