package usecase

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/rimport"
)

type StockUsecase struct {
	rimport.RepositoryImports
}

func NewStockUsecase(ri rimport.RepositoryImports) *StockUsecase {
	return &StockUsecase{
		RepositoryImports: ri,
	}
}

// CreateStock создает склад
func (u *StockUsecase) CreateStock(param stock.StockInfo) (data int, err error) {
	data, err = u.Repository.Stock.CreateStock(param)
	if err != nil {
		return
	}
	return
}

// LoadStockList загружает лист продуктов
func (u *StockUsecase) LoadStockList() (data []stock.StockInfo, err error) {
	data, err = u.Repository.Stock.LoadStockList()
	if err != nil {
		return
	}
	return
}

// AddProductToStock добавление товвара на склад
func (u *StockUsecase) AddProductToStock(newAccounting stock.ProductStock) (id int, err error) {
	// проверяет есть ли продукт с такими параметрами
	id, err = u.Repository.Stock.CheckProductExistence(newAccounting)
	switch err {
	// если продукт есть, то добавляет количество к существующему продукту
	case nil:
		err = u.Repository.Stock.IncreaseAccountingAmount(newAccounting.Amount, id)
		if err != nil {
			return
		}
		return
	// если такого продукта нет, то добавляет новую запись
	case global.ErrNoData:
		id, err = u.Repository.Stock.AddProductToStock(newAccounting)
		if err != nil {
			return
		}
		return
	default:
		return
	}
}

// FindStockInfo поиск информации о складах и количества товара
func (u *StockUsecase) FindStockInfo(productID, limit, offset int) (data stock.ProductStockInfoWithTotalCount, err error) {
	// если limit не передан, выставляеться дефолтное значение
	if limit <= 0 {
		limit = global.DefaultLimit
	}
	// поиск количества на складах по ID продукта
	if productID > 0 {
		// поиск количества продукта по складам
		data.ProductStockInfo, err = u.Repository.Stock.FindProductStockInfo(productID, limit, offset)
		if err != nil {
			return
		}
		// поиск общего количества записей
		data.TotalCount, err = u.Repository.Stock.FindProductStockCount(productID)
		if err != nil {
			return
		}
		return
	}
	// поиск общего количества товаров по складам
	data.ProductStockInfo, err = u.Repository.Stock.FindStockInfo(limit, offset)
	if err != nil {
		return
	}
	// поиск общего количества записей
	data.TotalCount, err = u.Repository.Stock.LoadStockCount()
	if err != nil {
		return
	}
	return
}
