package stock_test

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/rimport"
	"jun2/cur/uimport"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateStock(t *testing.T) {
	r := assert.New(t)
	type fields struct {
		ri rimport.TestRepositoryImports
	}
	type args struct {
		param stock.StockInfo
	}

	const (
		StockName = "Some stock"
		Location  = "Somewhere"
		expect_id = 0
	)

	param := stock.StockInfo{
		StockName: StockName,
		Location:  Location,
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
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CreateStock(param).Return(expect_id, nil),
				)
			},
			args: args{
				param: param,
			},
			expectedData: expect_id,
			err:          nil,
		},
		{
			name: "неуспешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CreateStock(param).Return(expect_id, global.ErrInternalError),
				)
			},
			args: args{
				param: param,
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

			data, err := ui.Usecase.Stock.CreateStock(tt.args.param)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}

func TestLoadStockList(t *testing.T) {
	r := assert.New(t)
	type fields struct {
		ri rimport.TestRepositoryImports
	}

	tests := []struct {
		name         string
		prepare      func(f *fields)
		expectedData []stock.StockInfo
		err          error
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				var data []stock.StockInfo
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().LoadStockList().Return(data, nil),
				)
			},
			err: nil,
		},
		{
			name: "неуспешный результат",
			prepare: func(f *fields) {
				var data []stock.StockInfo
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().LoadStockList().Return(data, global.ErrInternalError),
				)
			},
			err: global.ErrInternalError,
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

			data, err := ui.Usecase.Stock.LoadStockList()
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}

func TestAddProductToStock(t *testing.T) {
	r := assert.New(t)
	type fields struct {
		ri rimport.TestRepositoryImports
	}
	type args struct {
		newAccounting stock.ProductStock
	}

	const (
		StockID     = 1
		ProductID   = 23
		VariationID = 53
		Amount      = 0
		expect_id   = 0
	)

	newAccounting := stock.ProductStock{
		StockID:     StockID,
		ProductID:   ProductID,
		VariationID: VariationID,
		Amount:      Amount,
	}

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData int
		err          error
	}{
		{
			name: "успешный результат (новая запись)",
			args: args{
				newAccounting: newAccounting,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(newAccounting).Return(0, global.ErrNoData),
					f.ri.Repository.Stock.EXPECT().AddProductToStock(newAccounting).Return(expect_id, nil),
				)
			},
			expectedData: expect_id,
			err:          nil,
		},
		{
			name: "успешный результат (увеличение количества)",
			args: args{
				newAccounting: newAccounting,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(newAccounting).Return(expect_id, nil),
					f.ri.Repository.Stock.EXPECT().IncreaseAccountingAmount(newAccounting.Amount, expect_id).Return(nil),
				)
			},
			err: nil,
		},
		{
			name: "неуспешный результат (проверка существования продукта)",
			args: args{
				newAccounting: newAccounting,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(newAccounting).Return(expect_id, global.ErrInternalError),
				)
			},
			expectedData: expect_id,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (новая запись)",
			args: args{
				newAccounting: newAccounting,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(newAccounting).Return(expect_id, global.ErrNoData),
					f.ri.Repository.Stock.EXPECT().AddProductToStock(newAccounting).Return(expect_id, global.ErrInternalError),
				)
			},
			expectedData: expect_id,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (увеличение количества)",
			args: args{
				newAccounting: newAccounting,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().CheckProductExistence(newAccounting).Return(expect_id, nil),
					f.ri.Repository.Stock.EXPECT().IncreaseAccountingAmount(newAccounting.Amount, expect_id).Return(global.ErrInternalError),
				)
			},
			err: global.ErrInternalError,
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

			data, err := ui.Usecase.Stock.AddProductToStock(tt.args.newAccounting)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}

func TestFindStockInfo(t *testing.T) {
	r := assert.New(t)
	type fields struct {
		ri rimport.TestRepositoryImports
	}
	type args struct {
		productID int
		limit     int
		offset    int
	}
	const (
		productID = 1
		limit     = 3
		offset    = 0
	)
	result := stock.ProductStockInfoWithTotalCount{}
	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData stock.ProductStockInfoWithTotalCount
		err          error
	}{
		{
			name: "успешный результат (поиск общего количества товаров по складам и количества записей)",
			args: args{
				productID: 0,
				limit:     limit,
				offset:    offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().FindStockInfo(limit, offset).Return(result.ProductStockInfo, nil),
					f.ri.Repository.Stock.EXPECT().LoadStockCount().Return(result.TotalCount, nil),
				)
			},
			err: nil,
		},
		{
			name: "успешный результат (поиск количества продукта и количества по ID продукта)",
			args: args{
				productID: productID,
				limit:     limit,
				offset:    offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().FindProductStockInfo(productID, limit, offset).Return(result.ProductStockInfo, nil),
					f.ri.Repository.Stock.EXPECT().FindProductStockCount(productID).Return(result.TotalCount, nil),
				)
			},
			err: nil,
		},
		{
			name: "неуспешный результат (поиск общего количества товаров по складам)",
			args: args{
				productID: 0,
				limit:     limit,
				offset:    offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().FindStockInfo(limit, offset).Return(result.ProductStockInfo, global.ErrInternalError),
				)
			},
			err: global.ErrInternalError,
		},
		{
			name: "неуспешный результат (поиск количества по ID продукта)",
			args: args{
				productID: productID,
				limit:     limit,
				offset:    offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().FindProductStockInfo(productID, limit, offset).Return(result.ProductStockInfo, global.ErrInternalError),
				)
			},
			err: global.ErrInternalError,
		},
		{
			name: "неуспешный результат (поиск общего количества товаров и записей по складам)",
			args: args{
				productID: 0,
				limit:     limit,
				offset:    offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().FindStockInfo(limit, offset).Return(result.ProductStockInfo, nil),
					f.ri.Repository.Stock.EXPECT().LoadStockCount().Return(result.TotalCount, global.ErrInternalError),
				)
			},
			err: global.ErrInternalError,
		},
		{
			name: "неуспешный результат (поиск количества продукта и количества по ID продукта)",
			args: args{
				productID: productID,
				limit:     limit,
				offset:    offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Stock.EXPECT().FindProductStockInfo(productID, limit, offset).Return(result.ProductStockInfo, nil),
					f.ri.Repository.Stock.EXPECT().FindProductStockCount(productID).Return(result.TotalCount, global.ErrInternalError),
				)
			},
			err: global.ErrInternalError,
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

			data, err := ui.Usecase.Stock.FindStockInfo(tt.args.productID, tt.args.limit, tt.args.offset)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}

// func TestAwesomePublicMethod(t *testing.T) {
// 	r := assert.New(t)

// 	type fields struct {
// 		ri rimport.TestRepositoryImports
// 		bi *bimport.TestBridgeImports
// 		ts *transaction.MockSession
// 	}
// 	type args struct {
// 		id int
// 	}

// 	const (
// 		id = 1
// 	)

// 	tests := []struct {
// 		name         string
// 		prepare      func(f *fields)
// 		args         args
// 		expectedData template.TemplateObject
// 		err          error
// 	}{
// 		{
// 			name: "успешный результат",
// 			prepare: func(f *fields) {
// 				templateData := template.TemplateObject{}

// 				gomock.InOrder(
// 					f.ri.MockRepository.Template.EXPECT().FindTemplateObj(f.ts, id).Return(templateData, nil),
// 				)
// 			},
// 			args: args{id: id},
// 			err:  nil,
// 		},
// 		{
// 			name: "неуспешный результат",
// 			prepare: func(f *fields) {
// 				templateData := template.TemplateObject{}

// 				gomock.InOrder(
// 					f.ri.MockRepository.Template.EXPECT().FindTemplateObj(f.ts, id).Return(templateData, global.ErrNoData),
// 				)
// 			},
// 			args: args{id: id},
// 			err:  global.ErrInternalError,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			f := fields{
// 				ri: rimport.NewTestRepositoryImports(ctrl),
// 				ts: transaction.NewMockSession(ctrl),
// 			}
// 			if tt.prepare != nil {
// 				tt.prepare(&f)
// 			}

// 			sm := transaction.NewMockSessionManager(ctrl)
// 			ui := uimport.NewUsecaseImports(testLogger, testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports(), sm)

// 			data, err := ui.Usecase.Template.AwesomePublicMethod(f.ts, tt.args.id)
// 			r.Equal(tt.err, err)
// 			r.Equal(tt.expectedData, data)

// 		})
// 	}
// }
