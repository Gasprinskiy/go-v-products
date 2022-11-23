package price_test

import (
	"jun2/cur/config"
	"jun2/cur/internal/entity/price"
	"jun2/cur/internal/repository/postgres"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var loc, _ = time.LoadLocation("Asia/Tashkent")
var timeZone = time.Now().In(loc).Location()

func TestCreatePrice(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewPriceRepository(db)

	var testData price.Price
	t.Run("создание тесовых данных", func(t *testing.T) {
		activeTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-02-03 10:29:30", timeZone)
		r.NoError(err)
		price := price.Price{
			VariationID: 56,
			Price:       255.55,
			ActiveFrom:  activeTime,
			ActiveTill:  activeTime.Add(time.Hour),
		}
		priceID, err := repo.CreatePrice(price)
		r.NoError(err)
		r.NotZero(priceID)

		t.Run("получение и проверка тестовых данных", func(t *testing.T) {
			err := db.Get(&testData, `
			SELECT pr.variation_id, pr.price, pr.active_from, pr.active_till
			FROM product$price pr
			WHERE pr.price_id=$1`, priceID)

			r.NoError(err)
			r.Equal(price, testData)

			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err = db.Exec("delete from product$price where price_id = $1", priceID)
				r.NoError(err)
			})
		})
	})
}
