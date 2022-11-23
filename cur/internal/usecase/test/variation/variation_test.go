package variation_test

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/entity/variation"
	"jun2/cur/rimport"
	"jun2/cur/uimport"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateVariation(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
	}

	type args struct {
		param product.AddProductParams
	}

	const (
		product_id     = 1
		variation_type = 0.1
		unit_type      = "nnm"
		expect_id      = 0
	)

	productParam := product.AddProductParams{
		ProductID:     product_id,
		VariationType: variation_type,
		UnitType:      unit_type,
	}
	variationParam := variation.Params{
		ProductID:     product_id,
		VariationType: variation_type,
		UnitType:      unit_type,
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
				param: productParam,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Variation.EXPECT().CheckVariatonExistence(variationParam).Return(expect_id, global.ErrNoData),
					f.ri.Repository.Variation.EXPECT().CreateVariation(variationParam).Return(expect_id, nil),
				)
			},
			expectedData: expect_id,
			err:          nil,
		},
		{
			name: "неуспешный результат (вариация существует)",
			args: args{
				param: productParam,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Variation.EXPECT().CheckVariatonExistence(variationParam).Return(expect_id, variation.ErrVariationExists),
				)
			},
			expectedData: expect_id,
			err:          variation.ErrVariationExists,
		},
		{
			name: "неуспешный результат (создание вариации)",
			args: args{
				param: productParam,
			},
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Variation.EXPECT().CheckVariatonExistence(variationParam).Return(expect_id, global.ErrNoData),
					f.ri.Repository.Variation.EXPECT().CreateVariation(variationParam).Return(expect_id, global.ErrInternalError),
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

			data, err := ui.Usecase.Variation.CreateVariation(tt.args.param)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}
