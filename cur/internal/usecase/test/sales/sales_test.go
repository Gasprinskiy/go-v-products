package sales_test

import (
	"errors"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/sales"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/rimport"
	"jun2/cur/tools/sqlnull"
	"jun2/cur/uimport"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFindSalesList(t *testing.T) {
	r := assert.New(t)
	type fields struct {
		ri rimport.TestRepositoryImports
	}
	type args struct {
		limit  int
		offset int
	}

	const (
		limit  = 3
		offset = 0
	)

	result := sales.SalesListWithTotalCount{}

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData sales.SalesListWithTotalCount
		err          error
	}{
		{
			name: "успешный результат",
			args: args{
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Sales.EXPECT().FindSalesList(limit, offset).Return(result.SalesList, nil),
					f.ri.Repository.Sales.EXPECT().LoadSalesListCount().Return(result.TotalCount, nil),
				)
			},
			expectedData: result,
			err:          nil,
		},
		{
			name: "неуспешный результат (получнение данных)",
			args: args{
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Sales.EXPECT().FindSalesList(limit, offset).Return(result.SalesList, global.ErrInternalError),
				)
			},
			expectedData: result,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (получнение количестова записей)",
			args: args{
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Sales.EXPECT().FindSalesList(limit, offset).Return(result.SalesList, nil),
					f.ri.Repository.Sales.EXPECT().LoadSalesListCount().Return(result.TotalCount, global.ErrInternalError),
				)
			},
			expectedData: result,
			err:          global.ErrInternalError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				ri: rimport.NewTestRepositoryImports(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ui := uimport.NewUsecaseImports(f.ri.RepositoryImports())

			data, err := ui.Usecase.Sales.FindSalesList(tt.args.limit, tt.args.offset)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}

func TestAddProductSale(t *testing.T) {
	r := assert.New(t)
	type fields struct {
		ri rimport.TestRepositoryImports
	}
	type args struct {
		param sales.Sales
	}

	const (
		accountingID = 0
		product_id   = 31
		variation_id = 58
		stock_id     = 1
		amount       = 67
		exist_amount = 80
		expect_id    = 0
	)

	param := sales.Sales{
		ProductId:   product_id,
		VariationId: variation_id,
		StockId:     stock_id,
		Amount:      amount,
	}

	stockCheckData := stock.ProductStock{
		ProductID:   product_id,
		VariationID: variation_id,
		StockID:     stock_id,
	}

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData int
		err          error
	}{
		{
			name: "успешный результат",
			args: args{
				param: param,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(stockCheckData).Return(accountingID, nil),
					f.ri.Repository.Stock.EXPECT().FindProductAmount(param).Return(exist_amount, nil),
					f.ri.Repository.Stock.EXPECT().DecreaseAccountingAmount(amount, accountingID).Return(nil),
					f.ri.Repository.Sales.EXPECT().AddProductSale(param).Return(expect_id, nil),
				)
			},
			expectedData: expect_id,
			err:          nil,
		},
		{
			name: "неуспешный результат (нет продукта на складе)",
			args: args{
				param: param,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(stockCheckData).Return(expect_id, stock.ErrNoProductInStock),
				)
			},
			expectedData: expect_id,
			err:          stock.ErrNoProductInStock,
		},
		{
			name: "неуспешный результат (получение количества)",
			args: args{
				param: param,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(stockCheckData).Return(0, nil),
					f.ri.Repository.Stock.EXPECT().FindProductAmount(param).Return(expect_id, global.ErrInternalError),
				)
			},
			expectedData: expect_id,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (не достаточно количестово товара)",
			args: args{
				param: param,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(stockCheckData).Return(expect_id, nil),
					f.ri.Repository.Stock.EXPECT().FindProductAmount(param).Return(expect_id, stock.ErrNotEnough),
				)
			},
			expectedData: expect_id,
			err:          stock.ErrNotEnough,
		},
		{
			name: "неуспешный результат (уменьшение количества товара на складе)",
			args: args{
				param: sales.Sales{
					ProductId:   product_id,
					VariationId: variation_id,
					StockId:     stock_id,
					Amount:      0,
				},
			},
			prepare: func(f *fields) {
				param.Amount = 0
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(stockCheckData).Return(0, nil),
					f.ri.Repository.Stock.EXPECT().FindProductAmount(param).Return(expect_id, nil),
					f.ri.Repository.Stock.EXPECT().DecreaseAccountingAmount(param.Amount, accountingID).Return(global.ErrInternalError),
				)
			},
			expectedData: expect_id,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (добавление записи в таблицу продаж)",
			args: args{
				param: param,
			},
			prepare: func(f *fields) {
				param := sales.Sales{
					ProductId:   product_id,
					VariationId: variation_id,
					StockId:     stock_id,
					Amount:      amount,
				}
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(stockCheckData).Return(accountingID, nil),
					f.ri.Repository.Stock.EXPECT().FindProductAmount(param).Return(exist_amount, nil),
					f.ri.Repository.Stock.EXPECT().DecreaseAccountingAmount(param.Amount, accountingID).Return(nil),
					f.ri.Repository.Sales.EXPECT().AddProductSale(param).Return(expect_id, global.ErrInternalError),
				)
			},
			expectedData: expect_id,
			err:          global.ErrInternalError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				ri: rimport.NewTestRepositoryImports(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ui := uimport.NewUsecaseImports(f.ri.RepositoryImports())

			data, err := ui.Usecase.Sales.AddProductSale(tt.args.param)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}

func TestFindSalesReport(t *testing.T) {
	r := assert.New(t)
	type fields struct {
		ri rimport.TestRepositoryImports
	}
	type args struct {
		param sales.SalesReportRequest
	}

	var (
		start   = time.Now()
		end     = time.Now().Add(1 * time.Hour)
		limit   = sqlnull.NewInt64(3)
		offset  = sqlnull.NewInt64(0)
		name    = sqlnull.NewString("some")
		stockID = sqlnull.NewInt64(1)
	)

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData sales.SalesListWithTotalCount
		err          error
	}{
		{
			name: "успешный результат",
			args: args{
				param: sales.SalesReportRequest{
					StartDate:   start,
					EndDate:     end,
					Limit:       limit,
					Offset:      offset,
					ProductName: name,
					StorageID:   stockID,
				},
			},
			prepare: func(f *fields) {
				data := sales.SalesListWithTotalCount{}
				param := sales.SalesReportRequest{
					StartDate:   start,
					EndDate:     end,
					Limit:       limit,
					Offset:      offset,
					ProductName: name,
					StorageID:   stockID,
				}
				gomock.InOrder(
					f.ri.Repository.Sales.EXPECT().FindSalesReport(param).Return(data.SalesList, nil),
					f.ri.Repository.Sales.EXPECT().FindSalesReportCount(param).Return(data.TotalCount, nil),
				)
			},
			expectedData: sales.SalesListWithTotalCount{},
			err:          nil,
		},
		{
			name: "неуспешный результат (отсчет по продажам)",
			args: args{
				param: sales.SalesReportRequest{
					StartDate:   start,
					EndDate:     end,
					Limit:       limit,
					Offset:      offset,
					ProductName: name,
					StorageID:   stockID,
				},
			},
			prepare: func(f *fields) {
				data := sales.SalesListWithTotalCount{}
				param := sales.SalesReportRequest{
					StartDate:   start,
					EndDate:     end,
					Limit:       limit,
					Offset:      offset,
					ProductName: name,
					StorageID:   stockID,
				}
				gomock.InOrder(
					f.ri.Repository.Sales.EXPECT().FindSalesReport(param).Return(data.SalesList, errors.New("сервер временно недоступен")),
				)
			},
			expectedData: sales.SalesListWithTotalCount{},
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (количество записей)",
			args: args{
				param: sales.SalesReportRequest{
					StartDate:   start,
					EndDate:     end,
					Limit:       limit,
					Offset:      offset,
					ProductName: name,
					StorageID:   stockID,
				},
			},
			prepare: func(f *fields) {
				data := sales.SalesListWithTotalCount{}
				param := sales.SalesReportRequest{
					StartDate:   start,
					EndDate:     end,
					Limit:       limit,
					Offset:      offset,
					ProductName: name,
					StorageID:   stockID,
				}
				gomock.InOrder(
					f.ri.Repository.Sales.EXPECT().FindSalesReport(param).Return(data.SalesList, nil),
					f.ri.Repository.Sales.EXPECT().FindSalesReportCount(param).Return(data.TotalCount, errors.New("сервер временно недоступен")),
				)
			},
			expectedData: sales.SalesListWithTotalCount{},
			err:          global.ErrInternalError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			f := fields{
				ri: rimport.NewTestRepositoryImports(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ui := uimport.NewUsecaseImports(f.ri.RepositoryImports())

			data, err := ui.Usecase.Sales.FindSalesReport(tt.args.param)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}
