package product_test

import (
	"jun2/cur/config"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/repository/postgres"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestFindProductByID(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewProductRepository(db)

	t.Run("создание тестовых данных", func(t *testing.T) {
		product := product.Product{
			Name:        "ариель",
			Description: "парашок",
			Tags:        "нет тэгов",
		}

		r.NoError(db.Get(&product.ID, `
		INSERT INTO product (product_name, description, tags)
		VALUES ($1, $2, $3) 
		RETURNING product_id`, product.Name, product.Description, product.Tags))

		t.Run("поиск и проверка тестовых данных", func(t *testing.T) {
			data, err := repo.FindProductByID(product.ID)
			r.NoError(err)
			r.Equal(product, data)

			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err = db.Exec("delete from product where product_id = $1", product.ID)
				r.NoError(err)
			})
		})
	})
}

func TestCreateProduct(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewProductRepository(db)

	var testData product.Product
	t.Run("создание тесовых данных", func(t *testing.T) {
		product := product.Product{
			Name:        "хуета",
			Description: "полная",
			Tags:        "ааааа",
		}
		product.ID, err = repo.CreateProduct(product)
		r.NoError(err)
		r.NotZero(product.ID)

		t.Run("получение и проверка тестовых данных", func(t *testing.T) {
			r.NoError(db.Get(&testData, `
			SELECT p.product_id, p.product_name, p.description, p.tags 
			FROM product p
			WHERE p.product_id=$1`, product.ID))
			r.Equal(product, testData)

			t.Run("удаление тестовых данных", func(t *testing.T) {
				_, err = db.Exec("delete from product where product_id = $1", product.ID)
				r.NoError(err)
			})
		})
	})
}

func TestFindProductList(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewProductRepository(db)

	t.Run("добавление тестовых данных", func(t *testing.T) {
		productList := []product.Product{
			{
				Name:        "test0100100",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
			{
				Name:        "test0200200",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
			{
				Name:        "test0300300",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
		}
		for i, product := range productList {
			r.NoError(db.Get(&productList[i].ID, `
			INSERT INTO product (product_name, description, tags)
			VALUES ($1, $2, $3) 
			RETURNING product_id`, product.Name, product.Description, product.Tags))
		}
		t.Run("получение данных по limit и offset", func(t *testing.T) {
			data, err := repo.FindProductList(global.DefaultLimit, global.DefaultOffset)
			r.NoError(err)
			r.True(len(data) == global.DefaultLimit)
			for _, product := range productList {
				r.Contains(data, product)
			}
			t.Run("удаление тестовых данных", func(t *testing.T) {
				for _, product := range productList {
					_, err = db.Exec("DELETE FROM product WHERE product_id = $1", product.ID)
					r.NoError(err)
				}
			})
		})
	})
}

func TestFindProductListByTags(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewProductRepository(db)

	t.Run("добавление тестовых данных", func(t *testing.T) {
		productList := []product.Product{
			{
				Name:        "test0100100",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
			{
				Name:        "test0200200",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
			{
				Name:        "test0300300",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
		}

		for i, product := range productList {
			r.NoError(db.Get(&productList[i].ID, `
			INSERT INTO product (product_name, description, tags)
			VALUES ($1, $2, $3) 
			RETURNING product_id`, product.Name, product.Description, product.Tags))
		}

		t.Run("получение и проверка тестовых данных по tags и limit", func(t *testing.T) {
			data, err := repo.FindProductListByTags(global.DefaultLimit, global.DefaultOffset, "001100110011")
			r.NoError(err)
			for _, value := range productList {
				r.Contains(data, value)
			}
			t.Run("удаление тестовых данных", func(t *testing.T) {
				for _, product := range productList {
					_, err = db.Exec("delete from product where product_id = $1", product.ID)
					r.NoError(err)
				}
			})
		})
	})
}

func TestLoadProductCount(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewProductRepository(db)

	t.Run("получение данных для сравнения", func(t *testing.T) {
		productIdList := []int64{}
		r.NoError(db.Select(&productIdList, "SELECT product_id FROM product"))
		t.Run("сравнение", func(t *testing.T) {
			count, err := repo.LoadProductCount()
			r.NoError(err)
			r.Equal(count, len(productIdList))
		})
	})
}

func TestFindProductCountByTags(t *testing.T) {
	r := require.New(t)

	conf := config.NewConfig(os.Getenv("CONF_PATH"))
	r.NotNil(conf)

	db, err := sqlx.Open("postgres", conf.DbConn())
	r.NoError(err)
	defer db.Close()

	repo := postgres.NewProductRepository(db)

	t.Run("добавление тестовых данных", func(t *testing.T) {
		productList := []product.Product{
			{
				Name:        "test0100100",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
			{
				Name:        "test0200200",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
			{
				Name:        "test0300300",
				Description: "test descritpion",
				Tags:        "001100110011",
			},
		}
		for i, product := range productList {
			r.NoError(db.Get(&productList[i].ID, `
			INSERT INTO product (product_name, description, tags)
			VALUES ($1, $2, $3) 
			RETURNING product_id`, product.Name, product.Description, product.Tags))
		}
		t.Run("получение данных и сравнение данных", func(t *testing.T) {
			count, err := repo.FindProductCountByTags("001100110011")
			r.NoError(err)
			r.Equal(count, len(productList))
			t.Run("удаление тестовых данных", func(t *testing.T) {
				for _, product := range productList {
					_, err = db.Exec("delete from product where product_id = $1", product.ID)
					r.NoError(err)
				}
			})
		})
	})

}
