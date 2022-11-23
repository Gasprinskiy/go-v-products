package usecase

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/entity/variation"
	"jun2/cur/rimport"
)

type VariationUsecase struct {
	rimport.RepositoryImports
}

func NewVariationUsecase(ri rimport.RepositoryImports) *VariationUsecase {
	return &VariationUsecase{
		RepositoryImports: ri,
	}
}

// CreateVariation проверяет существование вариации и создает если ее нет
func (u *VariationUsecase) CreateVariation(param product.AddProductParams) (int, error) {
	newVariation := variation.NewVariation(param.ProductID, param.VariationType, param.UnitType)
	// проверка существует ли вариация у данного товара с данными параметрами
	id, err := u.Repository.Variation.CheckVariatonExistence(newVariation)
	switch err {
	// если вариация с данными параметрами есть
	case nil:
		return id, variation.ErrVariationExists
	// если вариации нет
	case global.ErrNoData:
		id, err = u.Repository.Variation.CreateVariation(newVariation)
		if err != nil {
			return 0, err
		}
		return id, nil

	default:
		return id, err
	}
}
