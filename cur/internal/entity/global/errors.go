package global

import (
	"errors"
	"fmt"
)

var (
	// ErrNoData данные не найдены
	ErrNoData = errors.New("данные не найдены")
	// ErrInvalidParam не валидный параметр
	ErrInvalidParam = func(paramname string) error {
		errstr := fmt.Sprintf("не валидный параметр '%s'", paramname)
		return errors.New(errstr)
	}
	// ErrInternalError сервер временно недоступен
	ErrInternalError = errors.New("сервер временно недоступен")
)
