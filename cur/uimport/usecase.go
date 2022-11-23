package uimport

import "jun2/cur/internal/usecase"

type Usecase struct {
	Product   *usecase.ProductUsecase
	Variation *usecase.VariationUsecase
	Stock     *usecase.StockUsecase
	Price     *usecase.PriceUsecase
	Sales     *usecase.SalesUsecase
}
