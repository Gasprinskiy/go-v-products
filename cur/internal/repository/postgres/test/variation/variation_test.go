package variation_test

import (
	"jun2/cur/config"
	"jun2/cur/internal/entity/variation"
	"jun2/cur/internal/repository/postgres"
	"jun2/cur/tools/sqlnull"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestFindVariationInfoByID(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewVariationRepository(db)

	t.Run("создание тестовых данных", func(t *testing.T) {
		productID := 23
		testData := []variation.Variation{
			{
				VariationID:   52,
				ProductID:     productID,
				VariationType: 5345.11231232,
				Price:         sqlnull.NewFloat64(1.45),
				UnitType:      "smg",
			},
			{
				VariationID:   56,
				ProductID:     productID,
				VariationType: 5.0,
				Price:         sqlnull.NewFloat64(2.45),
				UnitType:      "kg",
			},
		}

		for _, value := range testData {
			_, err := db.Exec(`
			INSERT INTO product$price(variation_id, price, active_till, active_from)
			VALUES($1, $2, $3, $4)
			`, value.VariationID, value.Price, time.Now().Add(24*time.Hour), time.Now())

			r.NoError(err)
		}

		t.Run("поиск тестовых данных", func(t *testing.T) {
			data, err := repo.FindVariationInfoByID(productID)
			r.NoError(err)
			r.Equal(data, testData)

			t.Run("удаление тестовых данных", func(t *testing.T) {
				for _, value := range testData {
					_, err := db.Exec("DELETE FROM product$price WHERE variation_id = $1", value.VariationID)
					r.NoError(err)
				}
			})
		})
	})
}

func TestCreateVariation(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewVariationRepository(db)

	t.Run("создание тестовых данных", func(t *testing.T) {
		testData := variation.Params{
			ProductID:     23,
			VariationType: 0.05,
			UnitType:      "g",
		}

		id, err := repo.CreateVariation(testData)
		r.NoError(err)

		t.Run("получение тестовых данных", func(t *testing.T) {
			data := variation.Params{}

			err := db.Get(&data, `
			SELECT pv.product_id, pv.variation_type, pv.unit_type
			FROM product$variations pv
			WHERE variation_id = $1
			`, id)

			r.NoError(err)
			r.Equal(data, testData)
			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err := db.Exec("DELETE FROM product$variations WHERE variation_id = $1", id)
				r.NoError(err)
			})
		})
	})
}

func TestCheckVariatonExistence(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewVariationRepository(db)

	t.Run("создание тестовых данных", func(t *testing.T) {
		var variationID int
		testData := variation.Params{
			ProductID:     23,
			VariationType: 0.05,
			UnitType:      "g",
		}

		err := db.Get(&variationID, `
		INSERT INTO product$variations(product_id, variation_type, unit_type)
		VALUES ($1, $2, $3)
		RETURNING variation_id
		`, testData.ProductID, testData.VariationType, testData.UnitType)

		r.NoError(err)

		t.Run("проверка тестовых данных", func(t *testing.T) {
			data, err := repo.CheckVariatonExistence(testData)

			r.NoError(err)
			r.Equal(variationID, data)

			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err := db.Exec("DELETE FROM product$variations WHERE variation_id = $1", variationID)
				r.NoError(err)
			})
		})
	})
}
