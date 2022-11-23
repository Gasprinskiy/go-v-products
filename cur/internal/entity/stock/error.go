package stock

import "errors"

var (
	// ErrNotEnough не достаточно количества товара
	ErrNotEnough = errors.New("не достаточно количества товара")
	// ErrNoProductInStock данного продукта нет на складе
	ErrNoProductInStock = errors.New("данного продукта нет на складе")
)
