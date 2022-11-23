package usecase

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/sales"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/rimport"
)

type SalesUsecase struct {
	rimport.RepositoryImports
}

func NewSalesUsecase(ri rimport.RepositoryImports) *SalesUsecase {
	return &SalesUsecase{
		RepositoryImports: ri,
	}
}

// FindSalesList поиск истории продаж по параметрам limit / offset
func (u *SalesUsecase) FindSalesList(limit, offset int) (data sales.SalesListWithTotalCount, err error) {
	// обработка параметра limit
	if limit <= 0 {
		limit = global.DefaultLimit
	}
	// получение историй продаж по параметрам limit / offset
	data.SalesList, err = u.Repository.Sales.FindSalesList(limit, offset)
	if err != nil {
		err = global.ErrInternalError
		return
	}
	// получение количества записей истории продаж
	data.TotalCount, err = u.Repository.Sales.LoadSalesListCount()
	if err != nil {
		err = global.ErrInternalError
		return
	}
	return
}

// AddProductSale добавление продажи товара со склада
func (u *SalesUsecase) AddProductSale(param sales.Sales) (int, error) {
	// stockCheckData параметры для проверки существования товара на складе
	stockCheckData := stock.NewStockFromSales(param)
	// проверка существования товара на складе
	accountingID, err := u.Repository.Stock.CheckProductExistence(stockCheckData)
	switch err {
	// если това есть на складе
	case nil:
		// получение количестова товара
		stockAmount, err := u.Repository.Stock.FindProductAmount(param)
		if err != nil {
			return 0, err
		}
		// если количетво покуки больше чем есть на складе
		if param.Amount > stockAmount {
			return 0, stock.ErrNotEnough
		}
		// уменьшение количества товара на складе
		err = u.Repository.Stock.DecreaseAccountingAmount(param.Amount, accountingID)
		if err != nil {
			return 0, err
		}
		// добавление записи в таблицу продаж
		id, err := u.Repository.Sales.AddProductSale(param)
		if err != nil {
			return 0, err
		}
		return id, nil

	// если такого товара нет на складе
	case global.ErrNoData:
		return 0, stock.ErrNoProductInStock
	default:
		return 0, err
	}
}

// FindSalesReport отсчет по продажам по данным параметрам
func (u *SalesUsecase) FindSalesReport(param sales.SalesReportRequest) (data sales.SalesListWithTotalCount, err error) {
	// отсчет по продажам
	data.SalesList, err = u.Repository.Sales.FindSalesReport(param)
	if err != nil {
		err = global.ErrInternalError
		return
	}
	// количество записей
	data.TotalCount, err = u.Repository.Sales.FindSalesReportCount(param)
	if err != nil {
		err = global.ErrInternalError
		return
	}
	return data, nil
}
