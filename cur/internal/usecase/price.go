package usecase

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/price"
	"jun2/cur/rimport"
)

type PriceUsecase struct {
	rimport.RepositoryImports
}

func NewPriceUsecase(ri rimport.RepositoryImports) *PriceUsecase {
	return &PriceUsecase{
		RepositoryImports: ri,
	}
}

// CreatePrice создает цену с временным промежутком
func (u *PriceUsecase) CreatePrice(param price.Price) (id int, err error) {
	id, err = u.Repository.Price.CreatePrice(param)
	if err != nil {
		err = global.ErrInternalError
		return
	}
	return
}
