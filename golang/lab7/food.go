package lab7

import (
	"errors"
)

type Food struct {
	price    float64
	name     string
	brand    string
	category string
}

func (f *Food) Sale(discount float64) error {
	if (0 > discount) || (100 < discount) {
		return errors.New("Скидка должна быть в диапозоне от 0 до 100 процентов")
	}

	f.price -= discount * f.price
	return nil
}

func (f *Food) Change_price(new_price float64) {
	f.price = new_price
}

func (f *Food) Change_specifications(new_name, new_category, new_brand string) {
	f.category = new_category
	f.name = new_name
	f.brand = new_brand
}

func (f *Food) Get_price() float64 {
	return f.price
}

func (f *Food) Get_info() [3]string {
	arr := [3]string{f.name, f.brand, f.category}
	return arr
}

