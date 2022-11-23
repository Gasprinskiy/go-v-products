package price

import (
	"time"
)

type Price struct {
	VariationID int       `json:"variation_id" db:"variation_id"`
	Price       float64   `json:"price" db:"price"`
	ActiveFrom  time.Time `json:"active_from" db:"active_from"`
	ActiveTill  time.Time `json:"active_till" db:"active_till"`
}
