package rimport

import (
	"jun2/cur/internal/repository/postgres"

	"github.com/jmoiron/sqlx"
)

type RepositoryImports struct {
	Repository Repository
}

func NewRepositoryImports(db *sqlx.DB) RepositoryImports {
	return RepositoryImports{
		Repository: Repository{
			Product:   postgres.NewProductRepository(db),
			Variation: postgres.NewVariationRepository(db),
			Stock:     postgres.NewStockRepository(db),
			Price:     postgres.NewPriceRepository(db),
			Sales:     postgres.NewSalesRepository(db),
		},
	}
}
