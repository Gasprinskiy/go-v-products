package product_test

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/internal/entity/variation"
	"jun2/cur/rimport"
	"jun2/cur/tools/sqlnull"
	"jun2/cur/uimport"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
	}

	type args struct {
		param product.AddProductParams
	}

	const (
		name           = "Some prod"
		descr          = "some prod descr"
		tags           = "some, prod"
		variation_type = 1.0
		unit_type      = "nnm"
	)

	result := product.ProductCreationResult{
		ProductID:   0,
		VariationID: 0,
	}

	param := product.AddProductParams{
		Name:          name,
		Description:   descr,
		Tags:          tags,
		VariationType: variation_type,
		UnitType:      unit_type,
	}

	newProduct := product.Product{
		Name:        name,
		Description: descr,
		Tags:        tags,
	}
	newVariation := variation.Params{
		ProductID:     result.ProductID,
		VariationType: variation_type,
		UnitType:      unit_type,
	}

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData product.ProductCreationResult
		err          error
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().CreateProduct(newProduct).Return(result.ProductID, nil),
					f.ri.Repository.Variation.EXPECT().CreateVariation(newVariation).Return(result.VariationID, nil),
				)
			},
			args: args{
				param,
			},
			expectedData: result,
			err:          nil,
		},
		{
			name: "неуспешный результат (создание продукта)",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().CreateProduct(newProduct).Return(result.ProductID, global.ErrInternalError),
				)
			},
			args: args{
				param,
			},
			expectedData: result,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (создание вариации)",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().CreateProduct(newProduct).Return(result.ProductID, nil),
					f.ri.Repository.Variation.EXPECT().CreateVariation(newVariation).Return(result.VariationID, global.ErrInternalError),
				)
			},
			args: args{
				param,
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

			data, err := ui.Usecase.Product.CreateProduct(tt.args.param)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}

func TestFindProductByID(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
	}

	type args struct {
		id int
	}

	const (
		id = 1
	)

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData variation.ProductWithVariationList
		err          error
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				data := variation.ProductWithVariationList{
					Product: product.Product{
						ID:          0,
						Name:        "",
						Description: "",
						Tags:        "",
					},
					VariationList: []variation.Variation{
						{
							VariationID:       0,
							ProductID:         0,
							VariationType:     0,
							UnitType:          "",
							Price:             sqlnull.NewFloat64(0),
							StockAvailability: []stock.ProductStock{},
						},
					},
				}
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductByID(id).Return(data.Product, nil),
					f.ri.Repository.Variation.EXPECT().FindVariationInfoByID(id).Return(data.VariationList, nil),
					f.ri.Repository.Stock.EXPECT().FindVariationAvailability(data.VariationList[0].VariationID).Return(data.VariationList[0].StockAvailability, nil),
				)
			},
			args: args{id: id},
			expectedData: variation.ProductWithVariationList{
				Product: product.Product{
					ID:          0,
					Name:        "",
					Description: "",
					Tags:        "",
				},
				VariationList: []variation.Variation{
					{
						VariationID:       0,
						ProductID:         0,
						VariationType:     0,
						UnitType:          "",
						Price:             sqlnull.NewFloat64(0),
						StockAvailability: []stock.ProductStock{},
					},
				},
			},
			err: nil,
		},
		{
			name: "частично успешный результат (первой вариации нет на складе)",
			prepare: func(f *fields) {
				data := variation.ProductWithVariationList{
					Product: product.Product{
						ID:          0,
						Name:        "",
						Description: "",
						Tags:        "",
					},
					VariationList: []variation.Variation{
						{
							VariationID:       0,
							ProductID:         0,
							VariationType:     0,
							UnitType:          "",
							Price:             sqlnull.NewFloat64(0),
							StockAvailability: []stock.ProductStock{},
						},
						{
							VariationID:       0,
							ProductID:         0,
							VariationType:     0,
							UnitType:          "",
							Price:             sqlnull.NewFloat64(0),
							StockAvailability: []stock.ProductStock{},
						},
					},
				}
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductByID(id).Return(data.Product, nil),
					f.ri.Repository.Variation.EXPECT().FindVariationInfoByID(id).Return(data.VariationList, nil),
					f.ri.Repository.Stock.EXPECT().FindVariationAvailability(data.VariationList[0].VariationID).Return(data.VariationList[0].StockAvailability, global.ErrNoData),
					f.ri.Repository.Stock.EXPECT().FindVariationAvailability(data.VariationList[1].VariationID).Return(data.VariationList[1].StockAvailability, nil),
				)
			},
			args: args{id: id},
			expectedData: variation.ProductWithVariationList{
				Product: product.Product{
					ID:          0,
					Name:        "",
					Description: "",
					Tags:        "",
				},
				VariationList: []variation.Variation{
					{
						VariationID:       0,
						ProductID:         0,
						VariationType:     0,
						UnitType:          "",
						Price:             sqlnull.NewFloat64(0),
						StockAvailability: []stock.ProductStock{},
					},
					{
						VariationID:       0,
						ProductID:         0,
						VariationType:     0,
						UnitType:          "",
						Price:             sqlnull.NewFloat64(0),
						StockAvailability: []stock.ProductStock{},
					},
				},
			},
			err: nil,
		},
		{
			name: "неуспешный результат (нет вариаций)",
			prepare: func(f *fields) {
				data := variation.ProductWithVariationList{
					Product: product.Product{
						ID:          0,
						Name:        "",
						Description: "",
						Tags:        "",
					},
				}
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductByID(id).Return(data.Product, global.ErrNoData),
					// f.ri.Repository.Variation.EXPECT().FindVariationInfoByID(id).Return(nil, global.ErrNoData),
				)
			},
			args: args{id: id},
			expectedData: variation.ProductWithVariationList{
				Product: product.Product{
					ID:          0,
					Name:        "",
					Description: "",
					Tags:        "",
				},
			},
			err: global.ErrNoData,
		},
		{
			name: "неуспешный результат",
			prepare: func(f *fields) {
				data := variation.ProductWithVariationList{}
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductByID(id).Return(data.Product, global.ErrNoData),
				)
			},
			args:         args{id: id},
			expectedData: variation.ProductWithVariationList{},
			err:          global.ErrNoData,
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

			data, err := ui.Usecase.Product.FindProductByID(tt.args.id)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}

func TestFindProductList(t *testing.T) {
	r := assert.New(t)
	type fields struct {
		ri rimport.TestRepositoryImports
	}
	type args struct {
		tags   string
		limit  int
		offset int
	}

	const (
		tag    = "tag"
		limit  = 3
		offset = 0
	)

	result := product.ProductListWithTotalCount{}

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData product.ProductListWithTotalCount
		err          error
	}{
		{
			name: "успешный результат (поиск по тегам)",
			args: args{
				tags:   tag,
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductListByTags(limit, offset, tag).Return(result.ProductList, nil),
					f.ri.Repository.Product.EXPECT().FindProductCountByTags(tag).Return(result.TotalCount, nil),
				)
			},
			expectedData: result,
			err:          nil,
		},
		{
			name: "успешный результат (поиск по умолчанию)",
			args: args{
				tags:   "",
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductList(limit, offset).Return(result.ProductList, nil),
					f.ri.Repository.Product.EXPECT().LoadProductCount().Return(result.TotalCount, nil),
				)
			},
			expectedData: result,
			err:          nil,
		},
		{
			name: "неуспешный результат (поиск по тегам)",
			args: args{
				tags:   tag,
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductListByTags(limit, offset, tag).Return(result.ProductList, global.ErrInternalError),
				)
			},
			expectedData: result,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (поиск по умолчанию)",
			args: args{
				tags:   "",
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductList(limit, offset).Return(result.ProductList, global.ErrInternalError),
				)
			},
			expectedData: result,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (поиск по тегам - количестов записей)",
			args: args{
				tags:   tag,
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductListByTags(limit, offset, tag).Return(result.ProductList, nil),
					f.ri.Repository.Product.EXPECT().FindProductCountByTags(tag).Return(result.TotalCount, global.ErrInternalError),
				)
			},
			expectedData: result,
			err:          global.ErrInternalError,
		},
		{
			name: "неуспешный результат (поиск по умолчанию - количестов записей)",
			args: args{
				tags:   "",
				limit:  limit,
				offset: offset,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Product.EXPECT().FindProductList(limit, offset).Return(result.ProductList, nil),
					f.ri.Repository.Product.EXPECT().LoadProductCount().Return(result.TotalCount, global.ErrInternalError),
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

			data, err := ui.Usecase.Product.FindProductList(tt.args.tags, tt.args.limit, tt.args.offset)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}
