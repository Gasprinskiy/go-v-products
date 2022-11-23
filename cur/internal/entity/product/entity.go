package product

type Product struct {
	ID          int    `json:"product_id" db:"product_id"`
	Name        string `json:"product_name" db:"product_name"`
	Description string `json:"description" db:"description"`
	Tags        string `json:"tags" db:"tags"`
}

type ProductListWithTotalCount struct {
	ProductList []Product `json:"product_list"`
	TotalCount  int       `json:"total_count"`
}

type AddProductParams struct {
	ProductID     int     `json:"product_id" db:"product_id"`
	Name          string  `json:"product_name" db:"product_name"`
	Description   string  `json:"description" db:"description"`
	Tags          string  `json:"tags" db:"tags"`
	VariationType float64 `json:"variation_type" db:"variation_type"`
	UnitType      string  `json:"unit_type" db:"unit_type"`
}

type ProductCreationResult struct {
	ProductID   int `json:"product_id"`
	VariationID int `json:"variation_id"`
}

func NewProduct(source AddProductParams) Product {
	return Product{
		Name:        source.Name,
		Description: source.Description,
		Tags:        source.Tags,
	}
}
