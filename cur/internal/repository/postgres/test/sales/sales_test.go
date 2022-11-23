package sales_test

import (
	"jun2/cur/config"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/sales"
	"jun2/cur/internal/repository/postgres"
	"jun2/cur/tools/sqlnull"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestFindSalesList(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewSalesRepository(db)

	t.Run("создание тестовых данных", func(t *testing.T) {
		testSalesData := sales.Sales{
			StockId:     3,
			ProductId:   23,
			VariationId: 56,
			Amount:      3,
		}

		saleID, err := repo.AddProductSale(testSalesData)
		r.NoError(err)
		r.NotZero(saleID)

		t.Run("получение и проверка тестовых данных", func(t *testing.T) {
			data := sales.Sales{}
			err := db.Get(&data, `
			SELECT ps.product_id, ps.variation_id, ps.stock_id, ps.amount
			FROM product$sales ps
			WHERE sale_id = $1`, saleID)

			r.NoError(err)
			r.Equal(data, testSalesData)

			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err := db.Exec("DELETE FROM product$sales WHERE sale_id=$1", saleID)
				r.NoError(err)
			})
		})
	})
}

func TestLoadSalesListCount(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewSalesRepository(db)

	t.Run("получение тестовых данных для сравнения", func(t *testing.T) {
		testIDCount := []int64{}
		r.NoError(db.Select(&testIDCount, `SELECT sale_id FROM product$sales`))
		t.Run("сравнение", func(t *testing.T) {
			count, err := repo.LoadSalesListCount()
			r.NoError(err)
			r.Equal(count, len(testIDCount))
		})
	})
}

func TestAddProductSale(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewSalesRepository(db)

	loc, _ := time.LoadLocation("Asia/Tashkent")
	timeZone := time.Now().In(loc).Location()
	activeTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2015-01-01 00:00:00", timeZone)
	r.NoError(err)

	var (
		variationID = 56
		stockID     = 3
	)

	t.Run("создание тестовых данных", func(t *testing.T) {
		salesList := []sales.SalesReport{
			{
				StockID:       stockID,
				ProductId:     23,
				VariationType: 5.0,
				UnitType:      "kg",
				Amount:        3,
				SoldDate:      activeTime,
				ProductName:   "tes1t",
				StockName:     "Склад 3",
			},
			{
				StockID:       3,
				ProductId:     23,
				VariationType: 5.0,
				UnitType:      "kg",
				Amount:        3,
				SoldDate:      activeTime.Add(time.Hour),
				ProductName:   "tes1t",
				StockName:     "Склад 3",
			},
			{
				StockID:       3,
				ProductId:     23,
				VariationType: 5.0,
				UnitType:      "kg",
				Amount:        3,
				SoldDate:      activeTime.Add(2 * time.Hour),
				ProductName:   "tes1t",
				StockName:     "Склад 3",
			},
		}
		saleIdList := []int64{0, 0, 0}
		for i, sale := range salesList {
			r.NoError(db.Get(&saleIdList[i], `INSERT INTO product$sales(product_id, variation_id, stock_id, sold_date, amount)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING sale_id`, sale.ProductId, variationID, sale.StockID, sale.SoldDate, sale.Amount))
		}
		t.Run("получение и проверка тестовых данных", func(t *testing.T) {
			data, err := repo.FindSalesList(global.DefaultLimit, global.DefaultOffset)
			r.NoError(err)
			for _, sale := range salesList {
				r.Contains(data, sale)
			}
			t.Run("удаление тестовых данных", func(t *testing.T) {
				for _, id := range saleIdList {
					_, err = db.Exec("delete from product$sales where sale_id = $1", id)
					r.NoError(err)
				}
			})
		})
	})
}

func TestFindSalesReport(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewSalesRepository(db)
	loc, _ := time.LoadLocation("Asia/Tashkent")
	timeZone := time.Now().In(loc).Location()

	activeTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2015-01-01 00:00:00", timeZone)
	r.NoError(err)

	params := sales.SalesReportRequest{
		StartDate: activeTime,
		EndDate:   activeTime.Add(72 * time.Hour),
	}

	var (
		variationID = 56
		stockID     = 3
		nameQuery   = "tes1t"
		limit       = 2
	)

	testData := []sales.SalesReport{
		{
			StockID:       stockID,
			ProductId:     23,
			VariationType: 5.0,
			UnitType:      "kg",
			Amount:        3,
			SoldDate:      activeTime,
			ProductName:   "tes1t",
			StockName:     "Склад 3",
		},
		{
			StockID:       stockID,
			ProductId:     23,
			VariationType: 5.0,
			UnitType:      "kg",
			Amount:        3,
			SoldDate:      activeTime.Add(time.Hour),
			ProductName:   "tes1t",
			StockName:     "Склад 3",
		},
		{
			StockID:       stockID,
			ProductId:     23,
			VariationType: 5.0,
			UnitType:      "kg",
			Amount:        3,
			SoldDate:      activeTime.Add(2 * time.Hour),
			ProductName:   "tes1t",
			StockName:     "Склад 3",
		},
	}

	t.Run("создание тестовых данных ", func(t *testing.T) {
		for _, value := range testData {
			_, err := db.Exec(`
			INSERT INTO product$sales(product_id, variation_id, stock_id, sold_date, amount)
			VALUES($1, $2, $3, $4, $5)
			RETURNING sale_id
			`, value.ProductId, variationID, stockID, value.SoldDate, value.Amount)

			r.NoError(err)
		}

		t.Run("получение данных с параметрами start_date & end_date (параметры по умолчанию)", func(t *testing.T) {
			data, err := repo.FindSalesReport(params)

			r.NoError(err)
			for _, value := range testData {
				r.Contains(data, value)
			}
		})

		t.Run("получение данных с параметром product_name", func(t *testing.T) {
			params.ProductName = sqlnull.NewString(nameQuery)
			data, err := repo.FindSalesReport(params)

			r.NoError(err)
			for _, value := range data {
				r.Contains(value.ProductName, nameQuery)
			}

			params.ProductName = sqlnull.NullString{}
		})

		t.Run("получение данных с параметром storage_id", func(t *testing.T) {
			params.StorageID = sqlnull.NewInt64(stockID)

			data, err := repo.FindSalesReport(params)

			r.NoError(err)
			for _, value := range data {
				r.Equal(value.StockID, stockID)
			}

			params.StorageID = sqlnull.NullInt64{}
		})

		t.Run("получение данных с ограничителем limit", func(t *testing.T) {
			params.Limit = sqlnull.NewInt64(limit)

			data, err := repo.FindSalesReport(params)

			r.NoError(err)
			r.True(len(data) == limit)

			params.Limit = sqlnull.NullInt64{}
		})

		t.Run("получение данных со всеми параметрами", func(t *testing.T) {
			params.ProductName = sqlnull.NewString(nameQuery)
			params.StorageID = sqlnull.NewInt64(stockID)
			params.Limit = sqlnull.NewInt64(limit)

			data, err := repo.FindSalesReport(params)

			r.NoError(err)
			r.True(len(data) == limit)
			for _, value := range testData {
				r.Equal(value.StockID, stockID)
				r.Contains(value.ProductName, nameQuery)
			}
		})

		t.Run("удаление тестовых данных", func(t *testing.T) {
			_, err := db.Exec("DELETE FROM product$sales WHERE variation_id=$1", variationID)
			r.NoError(err)
		})
	})
}

func TestFindSalesReportCount(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewSalesRepository(db)
	loc, _ := time.LoadLocation("Asia/Tashkent")
	timeZone := time.Now().In(loc).Location()

	activeTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2015-01-01 00:00:00", timeZone)
	r.NoError(err)

	params := sales.SalesReportRequest{
		StartDate: activeTime,
		EndDate:   activeTime.Add(72 * time.Hour),
	}

	var (
		variationID = 56
		stockID     = 3
		nameQuery   = "s1"
	)

	testData := []sales.SalesReport{
		{
			StockID:       stockID,
			ProductId:     23,
			VariationType: 5.0,
			UnitType:      "kg",
			Amount:        3,
			SoldDate:      activeTime,
			ProductName:   "tes1t",
			StockName:     "Склад 3",
		},
		{
			StockID:       stockID,
			ProductId:     23,
			VariationType: 5.0,
			UnitType:      "kg",
			Amount:        3,
			SoldDate:      activeTime.Add(time.Hour),
			ProductName:   "tes1t",
			StockName:     "Склад 3",
		},
		{
			StockID:       stockID,
			ProductId:     23,
			VariationType: 5.0,
			UnitType:      "kg",
			Amount:        3,
			SoldDate:      activeTime.Add(2 * time.Hour),
			ProductName:   "tes1t",
			StockName:     "Склад 3",
		},
	}

	t.Run("создание тестовых данных", func(t *testing.T) {
		for _, value := range testData {
			_, err := db.Exec(`
			INSERT INTO product$sales(product_id, variation_id, stock_id, sold_date, amount)
			VALUES($1, $2, $3, $4, $5)
			RETURNING sale_id
			`, value.ProductId, variationID, stockID, value.SoldDate, value.Amount)

			r.NoError(err)
		}
		t.Run("получение данных с параметрами start_date & end_date (параметры по умолчанию)", func(t *testing.T) {
			count, err := repo.FindSalesReportCount(params)

			r.NoError(err)
			r.Equal(count, len(testData))
		})

		t.Run("получение данных с параметром product_name", func(t *testing.T) {
			params.ProductName = sqlnull.NewString(nameQuery)
			count, err := repo.FindSalesReportCount(params)

			r.NoError(err)
			r.Equal(count, len(testData))

			params.ProductName = sqlnull.NullString{}
		})

		t.Run("получение данных со всеми параметрами", func(t *testing.T) {
			params.ProductName = sqlnull.NewString(nameQuery)
			params.StorageID = sqlnull.NewInt64(stockID)

			count, err := repo.FindSalesReportCount(params)

			r.NoError(err)
			r.Equal(count, len(testData))
		})

		t.Run("удаление тестовых данных", func(t *testing.T) {
			_, err := db.Exec("DELETE FROM product$sales WHERE variation_id=$1", variationID)
			r.NoError(err)
		})
	})
}
