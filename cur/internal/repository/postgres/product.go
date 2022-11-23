package postgres

import (
	"database/sql"
	"fmt"
	"jun2/cur/internal/entity/global"
	"jun2/cur/internal/entity/product"
	"jun2/cur/internal/repository"

	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) repository.Product {
	return &productRepository{db}
}

func (r *productRepository) FindProductByID(id int) (product.Product, error) {
	data := product.Product{}

	sqlQuery := `
		SELECT p.product_id, p.product_name, p.description, p.tags 
		FROM product p
		WHERE p.product_id=$1
	`

	err := r.db.Get(&data, sqlQuery, id)
	if err == sql.ErrNoRows {
		err = global.ErrNoData
	}

	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		return data, err
	}

}

func (r *productRepository) CreateProduct(newProduct product.Product) (newProductID int, err error) {
	sqlQuery := `
	INSERT INTO product (product_name, description, tags)
	VALUES ($1, $2, $3) 
	RETURNING product_id`

	err = r.db.Get(&newProductID, sqlQuery, newProduct.Name, newProduct.Description, newProduct.Tags)
	if err != nil {
		return
	}
	return
}

func (r *productRepository) FindProductList(limit, offset int) (data []product.Product, err error) {
	sqlQuery := `
	SELECT p.product_id, p.product_name, 
		   p.description, p.tags 
	FROM product p 
	ORDER BY product_id 
	DESC 
	LIMIT $1
	OFFSET $2`

	err = r.db.Select(&data, sqlQuery, limit, offset)
	if err == nil && len(data) == 0 {
		err = global.ErrNoData
	}

	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		return data, err
	}
}

func (r *productRepository) FindProductListByTags(limit, offset int, tags string) (data []product.Product, err error) {
	mainQuery := `
	SELECT p.product_id, p.product_name, 
	p.description, p.tags 
	FROM product p`

	qeuryLastPart := `
	ORDER BY product_id 
	DESC 
	LIMIT $1
	OFFSET $2`

	tagsQuerry := "%" + tags + "%"
	tagsFilter := fmt.Sprintf("WHERE LOWER(p.tags) LIKE LOWER('%s') OR LOWER(p.product_name) LIKE LOWER('%s')", tagsQuerry, tagsQuerry)
	sqlQuery := fmt.Sprintf("%s %s %s", mainQuery, tagsFilter, qeuryLastPart)

	err = r.db.Select(&data, sqlQuery, limit, offset)

	if err == nil && len(data) <= 0 {
		err = global.ErrNoData
	}

	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		return data, err
	}
}

func (r *productRepository) LoadProductCount() (data int, err error) {
	sqlQery := `
	SELECT count(product_id) 
	FROM product
	`
	err = r.db.Get(&data, sqlQery)

	if err == sql.ErrNoRows {
		err = global.ErrNoData
	}

	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		return data, err
	}
}

func (r *productRepository) FindProductCountByTags(tags string) (data int, err error) {
	mainQuery := `
	SELECT count(p.product_id) 
	FROM product p`
	tagsFilter := fmt.Sprintf(" WHERE LOWER(p.tags) LIKE LOWER('%s')", "%"+tags+"%")
	sqlQuery := fmt.Sprintf("%s %s", mainQuery, tagsFilter)

	err = r.db.Get(&data, sqlQuery)

	if err == sql.ErrNoRows {
		err = global.ErrNoData
	}

	switch err {
	case nil:
		return data, nil
	case global.ErrNoData:
		return data, global.ErrNoData
	default:
		return data, err
	}
}
