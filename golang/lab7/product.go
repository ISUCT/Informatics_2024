package Lab7

import (
	"errors"
)

type Product interface {
	applyDiscount(Discount float64) error
	getPrice() float64
	getProductInfo() string
}

func validateDiscount(discount float64) error {
	if discount < 0 {
		return errors.New("скидка не может быть меньше 0")
	}
	if discount > 100 {
		return errors.New("скидка не может быть больше 100%")
	}
	return nil
}
