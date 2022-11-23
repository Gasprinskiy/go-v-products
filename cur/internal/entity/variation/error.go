package variation

import "errors"

var (
	// ErrVariationExists вариация уже существует
	ErrVariationExists = errors.New("вариация уже существует")
)
