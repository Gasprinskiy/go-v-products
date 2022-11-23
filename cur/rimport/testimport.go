package rimport

import (
	"jun2/cur/internal/repository"

	"github.com/golang/mock/gomock"
)

type TestRepositoryImports struct {
	Repository MockRepository
}

func NewTestRepositoryImports(ctrl *gomock.Controller) TestRepositoryImports {
	return TestRepositoryImports{
		Repository: MockRepository{
			Product:   repository.NewMockProduct(ctrl),
			Variation: repository.NewMockVariation(ctrl),
			Stock:     repository.NewMockStock(ctrl),
			Price:     repository.NewMockPrice(ctrl),
			Sales:     repository.NewMockSales(ctrl),
		},
	}
}

func (t *TestRepositoryImports) RepositoryImports() RepositoryImports {
	return RepositoryImports{
		Repository: Repository{
			Product:   t.Repository.Product,
			Variation: t.Repository.Variation,
			Stock:     t.Repository.Stock,
			Sales:     t.Repository.Sales,
			Price:     t.Repository.Price,
		},
	}
}
