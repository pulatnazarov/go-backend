package constants

import "errors"

var (
	ErrNotEnough      = errors.New("not enough money for this products")
	ErrDivisionByZero = errors.New("can not divide by 0")
)
