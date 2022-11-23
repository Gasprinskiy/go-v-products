package stock_test

import (
	"jun2/cur/config"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/sales"
	"jun2/cur/internal/entity/stock"
	"jun2/cur/internal/repository/postgres"
	"jun2/cur/tools/sqlnull"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateStock(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	t.Run("создание тестовых данных", func(t *testing.T) {
		data := stock.StockInfo{
			StockName: "Test stock",
			Location:  "Somewhere",
		}
		data.StockId, err = repo.CreateStock(data)
		r.NoError(err)
		t.Run("получение и проверка тестовых данных", func(t *testing.T) {
			testData := stock.StockInfo{}

			r.NoError(db.Get(&testData, `
			SELECT * FROM stocks
			WHERE stock_id = $1
			`, data.StockId))

			r.Equal(data, testData)
			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err := db.Exec("delete from stocks where stock_id = $1", data.StockId)
				r.NoError(err)
			})
		})
	})
}

func TestFindVariationAvailability(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)
	t.Run("создание тестовых данных", func(t *testing.T) {
		variationID := 56
		testStockData := []stock.ProductStock{
			{
				StockID:     1,
				ProductID:   23,
				VariationID: variationID,
				Amount:      1,
				Location:    "Юнусабад 8",
				StockName:   "Склад 1",
			},
			{
				StockID:     2,
				ProductID:   23,
				VariationID: variationID,
				Amount:      1,
				Location:    "Сергели 7",
				StockName:   "Склад 2",
			},
		}
		for i, value := range testStockData {
			err := db.Get(&testStockData[i], `
			INSERT INTO product$stocks (stock_id, product_id, variation_id, amount)
			VALUES ($1, $2, $3, $4) 
			RETURNING *`, value.StockID, value.ProductID, value.VariationID, value.Amount)

			r.NoError(err)
		}
		t.Run("получение и проверка тестовых данных", func(t *testing.T) {
			data, err := repo.FindVariationAvailability(variationID)
			r.NoError(err)
			r.Equal(testStockData, data)
			for _, value := range data {
				_, err := db.Exec("DELETE FROM product$stocks where accounting_id=$1", value.AccountingID)
				r.NoError(err)
			}
		})
	})
}

func TestAddProductToStock(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	t.Run("создание тесовых данных", func(t *testing.T) {
		testStockData := stock.ProductStock{
			StockID:     1,
			ProductID:   23,
			VariationID: 56,
			Amount:      1,
		}
		testStockData.AccountingID, err = repo.AddProductToStock(testStockData)
		r.NoError(err)
		r.NotZero(testStockData.AccountingID)

		t.Run("получение и проверка тестовых данных", func(t *testing.T) {
			data := stock.ProductStock{}
			err := db.Get(&data, `
			SELECT ps.accounting_id, ps.stock_id, ps.product_id, ps.variation_id, ps.amount
			FROM product$stocks ps
			WHERE ps.accounting_id=$1`, testStockData.AccountingID)

			r.NoError(err)
			r.Equal(data, testStockData)

			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err = db.Exec("DELETE FROM product$stocks where accounting_id=$1", testStockData.AccountingID)
				r.NoError(err)
			})
		})
	})
}

func TestCheckProductExistence(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	t.Run("создание тесовых данных", func(t *testing.T) {
		testStockData := stock.ProductStock{
			StockID:     1,
			ProductID:   23,
			VariationID: 56,
			Amount:      1,
		}
		err := db.Get(&testStockData.AccountingID, `
		INSERT INTO product$stocks(stock_id, product_id, variation_id, amount)
		VALUES ($1, $2, $3, $4)
		RETURNING accounting_id  
		`, testStockData.StockID, testStockData.ProductID, testStockData.VariationID, testStockData.Amount)
		r.NoError(err)
		r.NotZero(testStockData.AccountingID)

		t.Run("получение и проверка тестовых данных", func(t *testing.T) {
			data, err := repo.CheckProductExistence(testStockData)
			r.Equal(testStockData.AccountingID, data)
			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err = db.Exec("DELETE FROM product$stocks where accounting_id=$1", testStockData.AccountingID)
				r.NoError(err)
			})
		})
	})
}

func TestIncreaseAccountingAmount(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	t.Run("создание тесовых данных", func(t *testing.T) {
		amount := 1
		testStockData := stock.ProductStock{
			StockID:     1,
			ProductID:   23,
			VariationID: 56,
			Amount:      amount,
		}
		err := db.Get(&testStockData.AccountingID, `
		INSERT INTO product$stocks(stock_id, product_id, variation_id, amount)
		VALUES ($1, $2, $3, $4)
		RETURNING accounting_id  
		`, testStockData.StockID, testStockData.ProductID, testStockData.VariationID, testStockData.Amount)
		r.NoError(err)
		r.NotZero(testStockData.AccountingID)

		t.Run("добавление количества к тестовым данным", func(t *testing.T) {
			err := repo.IncreaseAccountingAmount(amount, testStockData.AccountingID)
			r.NoError(err)

			t.Run("получение и проверка тестовых данных", func(t *testing.T) {
				data := stock.ProductStock{}
				err := db.Get(&data, `
				SELECT ps.stock_id, ps.product_id, ps.variation_id, ps.amount
				FROM product$stocks ps
				WHERE accounting_id=$1
				`, testStockData.AccountingID)
				r.NoError(err)
				r.Equal((testStockData.Amount + amount), data.Amount)

				t.Run("удаление тестовых данных", func(t *testing.T) {
					_, err = db.Exec("DELETE FROM product$stocks where accounting_id=$1", testStockData.AccountingID)
					r.NoError(err)
				})
			})
		})
	})
}

func TestDecreaseAccountingAmount(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	t.Run("создание тесовых данных", func(t *testing.T) {
		amount := 1
		testStockData := stock.ProductStock{
			StockID:     1,
			ProductID:   23,
			VariationID: 56,
			Amount:      amount,
		}
		err := db.Get(&testStockData.AccountingID, `
		INSERT INTO product$stocks(stock_id, product_id, variation_id, amount)
		VALUES ($1, $2, $3, $4)
		RETURNING accounting_id  
		`, testStockData.StockID, testStockData.ProductID, testStockData.VariationID, testStockData.Amount)
		r.NoError(err)
		r.NotZero(testStockData.AccountingID)

		t.Run("уменьшение количества тестовых данных", func(t *testing.T) {
			err := repo.DecreaseAccountingAmount(amount, testStockData.AccountingID)
			r.NoError(err)

			t.Run("получение и проверка тестовых данных", func(t *testing.T) {
				data := stock.ProductStock{}
				err := db.Get(&data, `
				SELECT ps.stock_id, ps.product_id, ps.variation_id, ps.amount
				FROM product$stocks ps
				WHERE accounting_id=$1
				`, testStockData.AccountingID)
				r.NoError(err)
				r.Equal((testStockData.Amount - amount), data.Amount)

				t.Run("удаление тестовых данных", func(t *testing.T) {
					_, err = db.Exec("DELETE FROM product$stocks where accounting_id=$1", testStockData.AccountingID)
					r.NoError(err)
				})
			})
		})
	})
}

func TestFindStockInfo(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)
	t.Run("создание тестового склада", func(t *testing.T) {
		amount := 4
		testStock := stock.ProductStockInfo{
			Location:   "7й круг ада",
			StockName:  "ад!",
			TotalAmout: sqlnull.NewInt64(amount),
		}

		err := db.Get(&testStock.StockId, `
		INSERT INTO stocks(location, stock_name)
		VALUES ($1, $2)
		RETURNING stock_id
		`, testStock.Location, testStock.StockName)

		r.NoError(err)

		t.Run("создание тестовых данных на складе", func(t *testing.T) {
			testStockData := []stock.ProductStock{
				{
					StockID:     testStock.StockId,
					ProductID:   23,
					VariationID: 56,
					Amount:      amount / 2,
				},
				{
					StockID:     testStock.StockId,
					ProductID:   18,
					VariationID: 61,
					Amount:      amount / 2,
				},
			}

			for i, value := range testStockData {
				err := db.Get(&testStockData[i].AccountingID, `
				INSERT INTO product$stocks(stock_id, product_id, variation_id, amount)
				VALUES ($1, $2, $3, $4)
				RETURNING accounting_id
				`, value.StockID, value.ProductID, value.VariationID, value.Amount)

				r.NoError(err)
				r.NotZero(testStockData[i].AccountingID)
			}

			t.Run("получение и проверка тестовых данных", func(t *testing.T) {
				data, err := repo.FindStockInfo(global.DefaultLimit, global.DefaultOffset)
				r.NoError(err)
				r.Contains(data, testStock)

				t.Run("удаление тестовых данных", func(t *testing.T) {
					for _, value := range testStockData {
						_, err = db.Exec("DELETE FROM product$stocks where accounting_id=$1", value.AccountingID)
						r.NoError(err)
					}
					_, err = db.Exec("DELETE FROM stocks where stock_id=$1", testStock.StockId)
					r.NoError(err)
				})
			})
		})
	})
}

func TestLoadStockCount(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	t.Run("получение данных для сравнения", func(t *testing.T) {
		stockIdList := []int64{}
		r.NoError(db.Select(&stockIdList, "SELECT stock_id FROM stocks"))
		t.Run("сравнение", func(t *testing.T) {
			count, err := repo.LoadStockCount()
			r.NoError(err)
			r.Equal(count, len(stockIdList))
		})
	})
}

func TestFindProductStockInfo(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)
	t.Run("создание тестового склада", func(t *testing.T) {
		amount := 4
		testStockList := []stock.ProductStockInfo{
			{
				Location:   "Somewhere",
				StockName:  "Some stock 1",
				TotalAmout: sqlnull.NewInt64(amount / 2),
			},
			{
				Location:   "I dont know",
				StockName:  "Some stock 2",
				TotalAmout: sqlnull.NewInt64(amount / 2),
			},
		}

		for i, value := range testStockList {
			err := db.Get(&testStockList[i].StockId, `
			INSERT INTO stocks(location, stock_name)
			VALUES ($1, $2)
			RETURNING stock_id
			`, value.Location, value.StockName)

			r.NoError(err)
		}

		t.Run("создание тестовых данных на складе", func(t *testing.T) {
			productID := 23
			testStockData := []stock.ProductStock{
				{
					StockID:     testStockList[0].StockId,
					ProductID:   productID,
					VariationID: 56,
					Amount:      amount / 2,
				},
				{
					StockID:     testStockList[1].StockId,
					ProductID:   productID,
					VariationID: 61,
					Amount:      amount / 2,
				},
			}

			for i, value := range testStockData {
				err := db.Get(&testStockData[i].AccountingID, `
				INSERT INTO product$stocks(stock_id, product_id, variation_id, amount)
				VALUES ($1, $2, $3, $4)
				RETURNING accounting_id
				`, value.StockID, value.ProductID, value.VariationID, value.Amount)

				r.NoError(err)
				r.NotZero(testStockData[i].AccountingID)
			}

			t.Run("получение и проверка тестовых данных", func(t *testing.T) {
				data, err := repo.FindProductStockInfo(productID, global.DefaultLimit, global.DefaultOffset)

				r.NoError(err)
				for _, stock := range testStockList {
					r.Contains(data, stock)
				}

				t.Run("удаление тестовых данных", func(t *testing.T) {
					for _, value := range testStockData {
						_, err = db.Exec("DELETE FROM product$stocks WHERE accounting_id=$1", value.AccountingID)
						r.NoError(err)
					}
					for _, value := range testStockList {
						_, err = db.Exec("DELETE FROM stocks WHERE stock_id=$1", value.StockId)
						r.NoError(err)
					}
				})
			})
		})
	})
}

func TestFindProductStockCount(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	const (
		variationID = 56
		productID   = 23
	)

	t.Run("создание тестовых данных", func(t *testing.T) {
		data := []stock.ProductStock{
			{
				StockID:     3,
				VariationID: variationID,
				ProductID:   productID,
				Amount:      1,
			},
			{
				StockID:     1,
				VariationID: variationID,
				ProductID:   productID,
				Amount:      1,
			},
		}
		for i, value := range data {
			r.NoError(db.Get(&data[i].AccountingID, `
			INSERT INTO product$stocks(stock_id, product_id, variation_id, amount)
			VALUES ($1, $2, $3, $4)
			RETURNING accounting_id
			`, value.StockID, value.ProductID, value.VariationID, value.Amount))
		}
		t.Run("получение тестовых данных", func(t *testing.T) {
			count, err := repo.FindProductStockCount(productID)
			r.NoError(err)
			r.Equal(count, len(data))
			t.Run("удаление тестовых данных", func(t *testing.T) {
				for _, value := range data {
					_, err := db.Exec("DELETE FROM product$stocks WHERE accounting_id=$1", value.AccountingID)
					r.NoError(err)
				}
			})
		})
	})
}

func TestLoadStockList(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	t.Run("создание тестовых данных", func(t *testing.T) {
		data := []stock.StockInfo{
			{
				StockName: "Test stock1",
				Location:  "Somewhere",
			},
			{
				StockName: "Test stock2",
				Location:  "I don't know",
			},
		}
		for i, value := range data {
			r.NoError(db.Get(&data[i].StockId, `
			INSERT INTO stocks(stock_name, location)
			VALUES ($1, $2)
			RETURNING stock_id
			`, value.StockName, value.Location))
		}
		t.Run("получение и проверка тестовых данный", func(t *testing.T) {
			stocks, err := repo.LoadStockList()
			r.NoError(err)
			for _, value := range data {
				r.Contains(stocks, value)
			}
			t.Run("уделаение тестовых данныйх", func(t *testing.T) {
				for _, value := range data {
					_, err := db.Exec("DELETE FROM stocks WHERE stock_id=$1", value.StockId)
					r.NoError(err)
				}
			})
		})
	})
}

func TestFindProductAmount(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewStockRepository(db)

	t.Run("создание тестового склада", func(t *testing.T) {
		amount := 1
		testStock := stock.ProductStockInfo{
			Location:  "7й круг ада",
			StockName: "ад!",
		}

		err := db.Get(&testStock.StockId, `
		INSERT INTO stocks(location, stock_name)
		VALUES ($1, $2)
		RETURNING stock_id
		`, testStock.Location, testStock.StockName)

		r.NoError(err)
		r.NotZero(testStock.StockId)

		t.Run("создание тестовых данных на складе", func(t *testing.T) {
			testProductData := sales.Sales{
				StockId:     testStock.StockId,
				ProductId:   23,
				VariationId: 56,
				Amount:      amount,
			}
			var accountingID int
			err := db.Get(&accountingID, `
				INSERT INTO product$stocks(stock_id, product_id, variation_id, amount)
				VALUES ($1, $2, $3, $4)
				RETURNING accounting_id 
				`,
				testProductData.StockId, testProductData.ProductId, testProductData.VariationId, testProductData.Amount)

			r.NoError(err)
			r.NotZero(accountingID)

			t.Run("получение и проверка тестовых данных", func(t *testing.T) {
				data, err := repo.FindProductAmount(testProductData)
				r.NoError(err)
				r.Equal(data, amount)

				t.Run("удаление тестовых данных", func(t *testing.T) {
					_, err = db.Exec("DELETE FROM product$stocks WHERE accounting_id=$1", accountingID)
					r.NoError(err)
					_, err = db.Exec("DELETE FROM stocks WHERE stock_id=$1", testStock.StockId)
					r.NoError(err)
				})
			})
		})
	})
}
