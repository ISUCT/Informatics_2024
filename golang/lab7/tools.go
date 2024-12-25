package lab7

import (
	"errors"
)

type Tools struct {
	price    float64
	name     string
	brand    string
	category string
}

func (t *Tools) Sale(discount float64) error {
	if (0 > discount) || (100 < discount) {
		return errors.New("Скидка должна быть в диапозоне от 0 до 100 процентов")
	}

	t.price -= discount * t.price
	return nil
}

func (t *Tools) Change_price(new_price float64) {
	t.price = new_price
}

func (t *Tools) Change_specifications(new_name, new_brand, new_category string) {
	t.brand = new_brand
	t.name = new_name
	t.category = new_category
}

func (t *Tools) Get_price() float64 {
	return t.price
}

func (t *Tools) Get_info() [3]string {
	arr := [3]string{t.name, t.brand, t.category}
	return arr
}