package uimport

import (
	"jun2/cur/internal/usecase"
	"jun2/cur/rimport"
)

type UsecaseImports struct {
	Usecase Usecase
}

func NewUsecaseImports(ri rimport.RepositoryImports) UsecaseImports {
	return UsecaseImports{
		Usecase: Usecase{
			Product:   usecase.NewProductUsecase(ri),
			Variation: usecase.NewVariationUsecase(ri),
			Stock:     usecase.NewStockUsecase(ri),
			Price:     usecase.NewPriceUsecase(ri),
			Sales:     usecase.NewSalesUsecase(ri),
		},
	}
}
