package rimport

import (
	"jun2/cur/internal/repository"
)

type Repository struct {
	Product   repository.Product
	Variation repository.Variation
	Stock     repository.Stock
	Price     repository.Price
	Sales     repository.Sales
}

type MockRepository struct {
	Product   *repository.MockProduct
	Variation *repository.MockVariation
	Stock     *repository.MockStock
	Price     *repository.MockPrice
	Sales     *repository.MockSales
}
