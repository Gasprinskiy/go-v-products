package price_test

import (
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/price"
	"jun2/cur/rimport"
	"jun2/cur/uimport"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePrice(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
	}

	type args struct {
		param price.Price
	}

	var (
		pr   = 10.0
		from = time.Now()
		till = time.Now().Add(1 * time.Hour)

		expectedID = 0
	)

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
					f.ri.Repository.Price.EXPECT().CreatePrice(price.Price{
						Price:      pr,
						ActiveFrom: from,
						ActiveTill: till,
					}).Return(expectedID, nil),
				)
			},
			args: args{
				param: price.Price{
					Price:      pr,
					ActiveFrom: from,
					ActiveTill: till,
				},
			},
			expectedData: expectedID,
			err:          nil,
		},
		{
			name: "неуспешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.Repository.Price.EXPECT().CreatePrice(price.Price{
						Price:      pr,
						ActiveFrom: from,
						ActiveTill: till,
					}).Return(expectedID, global.ErrInternalError),
				)
			},
			args: args{
				param: price.Price{
					Price:      pr,
					ActiveFrom: from,
					ActiveTill: till,
				},
			},
			expectedData: expectedID,
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

			data, err := ui.Usecase.Price.CreatePrice(tt.args.param)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)
		})
	}
}
